package alexa

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

const sdkVersion = "1.0"
const launchRequestName = "LaunchRequest"
const intentRequestName = "IntentRequest"
const sessionEndedRequestName = "SessionEndedRequest"

var timestampTolerance = 150

// ErrRequestEnvelopeNil reports that the request envelope was nil
// there might be edge case which causes panic if for whatever reason this object is empty
var ErrRequestEnvelopeNil = errors.New("request envelope was nil")

// Alexa defines the primary interface to use to create an Alexa request handler.
type Alexa struct {
	ApplicationID       string
	RequestHandler      RequestHandler
	IgnoreApplicationID bool
	IgnoreTimestamp     bool
}

// RequestHandler defines the interface that must be implemented to handle
// Alexa Requests
type RequestHandler interface {
	OnSessionStarted(context.Context, *Request, *Session, *Context, *Response) error
	OnLaunch(context.Context, *Request, *Session, *Context, *Response) error
	OnIntent(context.Context, *Request, *Session, *Context, *Response) error
	OnSessionEnded(context.Context, *Request, *Session, *Context, *Response) error
}

// RequestEnvelope contains the data passed from Alexa to the request handler.
type RequestEnvelope struct {
	Version string   `json:"version"`
	Session *Session `json:"session"`
	Request *Request `json:"request"`
	Context *Context `json:"context"`
}

// Session contains the session data from the Alexa request.
type Session struct {
	New        bool   `json:"new"`
	SessionID  string `json:"sessionId"`
	Attributes struct {
		String map[string]string `json:"string"`
	} `json:"attributes"`
	User struct {
		UserID      string            `json:"userId"`
		AccessToken string            `json:"accessToken"`
		Permissions map[string]string `json:"permissions,omitempty"`
	} `json:"user"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
}

// Context contains the context data from the Alexa Request.
type Context struct {
	System struct {
		Device struct {
			DeviceID            string `json:"deviceId"`
			SupportedInterfaces struct {
				AudioPlayer struct{}  `json:"AudioPlayer"`
				Display     *struct{} `json:"Display"`
				VideoApp    *struct{} `json:"VideoApp"`
			} `json:"supportedInterfaces"`
		} `json:"device"`
		Application struct {
			ApplicationID string `json:"applicationId"`
		} `json:"application"`
		User struct {
			UserID      string `json:"userId"`
			AccessToken string `json:"accessToken"`
			Permissions struct {
				ConsentToken string `json:"consentToken"`
			} `json:"permissions"`
		} `json:"user"`
		APIEndpoint    string `json:"apiEndpoint"`
		APIAccessToken string `json:"apiAccessToken"`
	} `json:"System"`
	AudioPlayer struct {
		PlayerActivity       string `json:"playerActivity"`
		Token                string `json:"token"`
		OffsetInMilliseconds int    `json:"offsetInMilliseconds"`
	} `json:"AudioPlayer"`
}

// Request contains the data in the request within the main request.
type Request struct {
	Type        string `json:"type"`
	Timestamp   string `json:"timestamp"`
	RequestID   string `json:"requestId"`
	Locale      string `json:"locale"`
	DialogState string `json:"dialogState"`
	Intent      Intent `json:"intent"`
	Name        string `json:"name"`
}

// Intent contains the data about the Alexa Intent requested.
type Intent struct {
	Name               string                `json:"name"`
	ConfirmationStatus string                `json:"confirmationStatus,omitempty"`
	Slots              map[string]IntentSlot `json:"slots"`
}

// IntentSlot contains the data for one Slot
type IntentSlot struct {
	Name               string       `json:"name"`
	ConfirmationStatus string       `json:"confirmationStatus,omitempty"`
	Value              string       `json:"value"`
	Resolutions        *Resolutions `json:"resolutions,omitempty"`
}

// Resolutions contain the (optional) ID of a slot
type Resolutions struct {
	ResolutionsPerAuthority []struct {
		Authority string `json:"authority"`
		Status    struct {
			Code string `json:"code"`
		} `json:"status"`
		Values []struct {
			Value struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}

// ResponseEnvelope contains the Response and additional attributes.
type ResponseEnvelope struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Response          *Response              `json:"response"`
}

// Response contains the body of the response.
type Response struct {
	OutputSpeech     *OutputSpeech `json:"outputSpeech,omitempty"`
	Card             *Card         `json:"card,omitempty"`
	Reprompt         *Reprompt     `json:"reprompt,omitempty"`
	Directives       []interface{} `json:"directives,omitempty"`
	ShouldSessionEnd bool          `json:"shouldEndSession"`
}

// OutputSpeech contains the data the defines what Alexa should say to the user.
type OutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

// Card contains the data displayed to the user by the Alexa app.
type Card struct {
	Type    string `json:"type"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Text    string `json:"text,omitempty"`
	Image   *Image `json:"image,omitempty"`
}

// Image provides URL(s) to the image to display in resposne to the request.
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// Reprompt contains data about whether Alexa should prompt the user for more data.
type Reprompt struct {
	OutputSpeech *OutputSpeech `json:"outputSpeech,omitempty"`
}

// PlayBehavior Describes playback behavior
type PlayBehavior string

const (
	// ReplaceAll ...
	ReplaceAll PlayBehavior = "REPLACE_ALL"
	// Enqueue ...
	Enqueue PlayBehavior = "ENQUEUE"
	// ReplaceEnqueued ...
	ReplaceEnqueued PlayBehavior = "REPLACE_ENQUEUED"
)

// ClearBehavior Describes the clear behavior
type ClearBehavior string

const (
	// ClearEnqueued ...
	ClearEnqueued ClearBehavior = "CLEAR_ENQUEUED"
	// ClearAll ...
	ClearAll ClearBehavior = "CLEAR_ALL"
)

// PlayAudioDirective Sends Alexa a command to stream the audio file identified by the specified audioItem.
type PlayAudioDirective struct {
	Type         string       `json:"type"`
	PlayBehavior PlayBehavior `json:"playBehavior"`
	AudioItem    AudioItem    `json:"audioItem"`
}

// StopAudioDirective Stops the current audio playback.
type StopAudioDirective struct {
	Type string `json:"type"`
}

// ClearQueueDirective Clears the audio playback queue.
type ClearQueueDirective struct {
	Type          string        `json:"type"`
	ClearBehavior ClearBehavior `json:"clearBehavior"`
}

// AudioItem contains an audio Stream definition for playback.
type AudioItem struct {
	Stream   Stream         `json:"stream"`
	Metadata *AudioMetadata `json:"metadata,omitempty"`
}

// Stream contains instructions on playing an audio stream.
type Stream struct {
	Token                 string `json:"token"`
	URL                   string `json:"url"`
	ExpectedPreviousToken string `json:"expectedPreviousToken,omitempty"`
	OffsetInMilliseconds  int    `json:"offsetInMilliseconds"`
}

// AudioMetadata contains instructions for providing audio metadata
type AudioMetadata struct {
	Title           string      `json:"title,omitempty"`
	Subtitle        string      `json:"subtitle,omitempty"`
	Art             *AudioImage `json:"art,omitempty"`
	BackgroundImage *AudioImage `json:"backgroundImage,omitempty"`
}

// AudioImage contains the data for images used within Audio Directives
type AudioImage struct {
	ContentDescription string             `json:"contentDescription,omitempty"`
	Sources            []AudioImageSource `json:"sources"`
}

// AudioImageSource contains the data for an image
type AudioImageSource struct {
	URL          string  `json:"url"`
	Size         *string `json:"size,omitempty"`
	WidthPixels  *int    `json:"widthPixels,omitempty"`
	HeightPixels *int    `json:"heightPixels,omitempty"`
}

// DialogDirective contains directives for use in Dialog prompts.
type DialogDirective struct {
	Type          string  `json:"type"`
	SlotToElicit  string  `json:"slotToElicit,omitempty"`
	SlotToConfirm string  `json:"slotToConfirm,omitempty"`
	UpdatedIntent *Intent `json:"updatedIntent,omitempty"`
}

// RenderDocumentDirective contain directives for use with Amazon APL documents
type RenderDocumentDirective struct {
	Type        string       `json:"type"`
	Token       string       `json:"token"`
	Document    interface{}  `json:"document"`
	DataSources *interface{} `json:"datasources,omitempty"`
}

// SupportsVideo returns true if this skill was initiated from
// a device that supports video, false otherwise
func SupportsVideo(ctx *Context) bool {
	return (ctx.System.Device.SupportedInterfaces.VideoApp != nil)
}

// SupportsDisplay returns true if this skill was initiated from
// a device that supports video, false otherwise
func SupportsDisplay(ctx *Context) bool {
	return (ctx.System.Device.SupportedInterfaces.Display != nil)
}

// ProcessRequest handles a request passed from Alexa
func (alexa *Alexa) ProcessRequest(ctx context.Context, requestEnv *RequestEnvelope) (*ResponseEnvelope, error) {
	if requestEnv == nil {
		return nil, ErrRequestEnvelopeNil
	}

	if !alexa.IgnoreApplicationID {
		err := alexa.verifyApplicationID(requestEnv)
		if err != nil {
			return nil, err
		}
	}
	if !alexa.IgnoreTimestamp {
		err := alexa.verifyTimestamp(requestEnv)
		if err != nil {
			return nil, err
		}
	} else {
		log.Println("Ignoring timestamp verification.")
	}

	request := requestEnv.Request
	session := requestEnv.Session
	if session.Attributes.String == nil {
		session.Attributes.String = make(map[string]string)
	}
	context := requestEnv.Context

	responseEnv := &ResponseEnvelope{}
	responseEnv.Version = sdkVersion
	responseEnv.Response = &Response{}
	responseEnv.Response.ShouldSessionEnd = true // Set default value.

	response := responseEnv.Response

	// If it is a new session, invoke onSessionStarted
	if session.New {
		err := alexa.RequestHandler.OnSessionStarted(ctx, request, session, context, response)
		if err != nil {
			log.Println("Error handling OnSessionStarted.", err.Error())
			return nil, err
		}
	}

	switch requestEnv.Request.Type {
	case launchRequestName:
		err := alexa.RequestHandler.OnLaunch(ctx, request, session, context, response)
		if err != nil {
			log.Println("Error handling OnLaunch.", err.Error())
			return nil, err
		}
	case intentRequestName:
		err := alexa.RequestHandler.OnIntent(ctx, request, session, context, response)
		if err != nil {
			log.Println("Error handling OnIntent.", err.Error())
			return nil, err
		}
	case sessionEndedRequestName:
		err := alexa.RequestHandler.OnSessionEnded(ctx, request, session, context, response)
		if err != nil {
			log.Println("Error handling OnSessionEnded.", err.Error())
			return nil, err
		}
	}

	// Copy Session Attributes into ResponseEnvelope
	responseEnv.SessionAttributes = make(map[string]interface{})
	for n, v := range session.Attributes.String {
		fmt.Println("Setting ", n, "to", v)
		responseEnv.SessionAttributes[n] = v
	}

	return responseEnv, nil
}

// SetTimestampTolerance sets the maximum number of seconds to allow between
// the current time and the request Timestamp.  Default value is 150 seconds.
func (alexa *Alexa) SetTimestampTolerance(seconds int) {
	timestampTolerance = seconds
}

// SetSimpleCard creates a new simple card with the specified content.
func (r *Response) SetSimpleCard(title string, content string) {
	r.Card = &Card{Type: "Simple", Title: title, Content: content}
}

// SetStandardCard creates a new standard card with the specified content.
func (r *Response) SetStandardCard(title string, text string, smallImageURL string, largeImageURL string) {
	r.Card = &Card{Type: "Standard", Title: title, Text: text}
	r.Card.Image = &Image{SmallImageURL: smallImageURL, LargeImageURL: largeImageURL}
}

// SetLinkAccountCard creates a new LinkAccount card.
func (r *Response) SetLinkAccountCard() {
	r.Card = &Card{Type: "LinkAccount"}
}

// SetOutputText sets the OutputSpeech type to text and sets the value specified.
func (r *Response) SetOutputText(text string) {
	r.OutputSpeech = &OutputSpeech{Type: "PlainText", Text: text}
}

// SetOutputSSML sets the OutputSpeech type to ssml and sets the value specified.
func (r *Response) SetOutputSSML(ssml string) {
	r.OutputSpeech = &OutputSpeech{Type: "SSML", SSML: ssml}
}

// SetRepromptText created a Reprompt if needed and sets the OutputSpeech type to text and sets the value specified.
func (r *Response) SetRepromptText(text string) {
	if r.Reprompt == nil {
		r.Reprompt = &Reprompt{}
	}
	r.Reprompt.OutputSpeech = &OutputSpeech{Type: "PlainText", Text: text}
}

// SetRepromptSSML created a Reprompt if needed and sets the OutputSpeech type to ssml and sets the value specified.
func (r *Response) SetRepromptSSML(ssml string) {
	if r.Reprompt == nil {
		r.Reprompt = &Reprompt{}
	}
	r.Reprompt.OutputSpeech = &OutputSpeech{Type: "SSML", SSML: ssml}
}

// AddPlayAudioPlayerDirective adds an AudioPlayer directive to the Response.
func (r *Response) AddPlayAudioPlayerDirective(playBehavior PlayBehavior, streamToken, url string, offsetInMilliseconds int, metadata *AudioMetadata) {
	d := PlayAudioDirective{
		Type:         "AudioPlayer.Play",
		PlayBehavior: playBehavior,
		AudioItem: AudioItem{
			Stream: Stream{
				Token:                streamToken,
				URL:                  url,
				OffsetInMilliseconds: offsetInMilliseconds,
			},
			Metadata: metadata,
		},
	}
	r.Directives = append(r.Directives, d)
}

// AddStopAudioPlayerDirective adds an AudioPlayer directive to stop the Audio.
func (r *Response) AddStopAudioPlayerDirective() {
	d := StopAudioDirective{
		Type: "AudioPlayer.Stop",
	}
	r.Directives = append(r.Directives, d)
}

// AddClearQueueAudioPlayerDirective adds an AudioPlayer directive to clear the audio queue.
func (r *Response) AddClearQueueAudioPlayerDirective(clearBehavior ClearBehavior) {
	d := ClearQueueDirective{
		Type:          "AudioPlayer.ClearQueue",
		ClearBehavior: clearBehavior,
	}
	r.Directives = append(r.Directives, d)
}

// AddDialogDirective adds a Dialog directive to the Response.
func (r *Response) AddDialogDirective(dialogType, slotToElicit, slotToConfirm string, intent *Intent) {
	d := DialogDirective{
		Type:          dialogType,
		SlotToElicit:  slotToElicit,
		SlotToConfirm: slotToConfirm,
		UpdatedIntent: intent,
	}
	r.Directives = append(r.Directives, d)
}

// AddRenderDocumentDirective adds a RenderDocument directive to the Response.
func (r *Response) AddRenderDocumentDirective(token string, document, datasources interface{}) {
	d := RenderDocumentDirective{
		Type:        "Alexa.Presentation.APL.RenderDocument",
		Token:       token,
		Document:    document,
		DataSources: &datasources,
	}
	r.Directives = append(r.Directives, d)
}

// verifyApplicationId verifies that the ApplicationID sent in the request
// matches the one configured for this skill.
func (alexa *Alexa) verifyApplicationID(request *RequestEnvelope) error {
	if request == nil {
		return ErrRequestEnvelopeNil
	}

	appID := alexa.ApplicationID
	requestAppID := request.Session.Application.ApplicationID
	if appID == "" {
		return errors.New("application ID was set to an empty string")
	}
	if requestAppID == "" {
		return errors.New("request Application ID was set to an empty string")
	}
	if appID != requestAppID {
		return errors.New("request Application ID does not match expected ApplicationId")
	}

	return nil
}

// verifyTimestamp compares the request timestamp to the current timestamp
// and returns an error if they are too far apart.
func (alexa *Alexa) verifyTimestamp(request *RequestEnvelope) error {
	if request == nil {
		return ErrRequestEnvelopeNil
	}

	timestamp, err := time.Parse(time.RFC3339, request.Request.Timestamp)
	if err != nil {
		return errors.New("unable to parse request timestamp.  Err: " + err.Error())
	}
	now := time.Now()
	delta := now.Sub(timestamp)
	deltaSecsAbs := math.Abs(delta.Seconds())
	if deltaSecsAbs > float64(timestampTolerance) {
		return errors.New("invalid Timestamp. The request timestamp " + timestamp.String() + " was off the current time " + now.String() + " by more than " + strconv.FormatInt(int64(timestampTolerance), 10) + " seconds.")
	}

	return nil
}
