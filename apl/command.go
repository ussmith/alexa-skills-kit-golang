package apl

//CommandBase holds common command properties
type CommandBase struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
	Delay       int    `json:"delay,omitempty"`
	ScreenLock  *bool  `json:"screenLock,omitEmpty"`
	When        *bool  `json:"when,omitempty"`
}

// Easing ..
type Easing string

// RepeatMode ..
type RepeatMode string

const (
	//LinearEasing ..
	LinearEasing Easing = "linear"
	//Ease ..
	Ease Easing = "ease"
	//EaseIn ..
	EaseIn Easing = "ease-in"
	//EaseOut ..
	EaseOut Easing = "ease-out"
	//EaseInOut ..
	EaseInOut Easing = "ease-in-out"

	//Restart ..
	Restart RepeatMode = "repeat"
	//Reverse ..
	Reverse RepeatMode = "reverse"
)

// AnimationProperty ..
type AnimationProperty struct {
	CommandBase
	//From is the starting value of the propery
	From []map[string]int `json:"From"`

	//Property is the name of the property to animate
	Property string `json:"property"`

	//To is the ending value of the property
	To []map[string]int `json:"to"`
}

//AnimateItemCommand Runs a fixed-duration animation sequence on one or more properties of a single component
type AnimateItemCommand struct {
	CommandBase
	ComponentID string              `json:"componentId"`
	Duration    int                 `json:"duration"`
	Easing      Easing              `json:"easing"`
	RepeatCount int                 `json:"repeatCount"`
	RepeatMode  RepeatMode          `json:"repeatMode"`
	Value       []AnimationProperty `json:"animationProperty"`
}

// AutoPage automatically progresses through a seris of pages displayed in a Pager comonent
type AutoPage struct {
	CommandBase
	ComponentID string `json:"componentId"`
	Count       int    `json:"count"`
	Duration    int    `json:"duration"`
}

// OpenURL if succesful opens a web browser or other application on the device
type OpenURL struct {
	CommandBase

	//Source is the URL to open and is required
	Source string `json:"source"`

	//OnFail Command to execute if the URL fails to open
	OnFail []interface{} `json:"onFail"`
}

// ParallelCommand executes a series of commands in parallel
type ParallelCommand struct {
	CommandBase
	Commands []interface{} `json:"command"`
}

// ScrollCommand scrolls a ScrollView or Sequence forward or backward by a set number of pages
type ScrollCommand struct {
	CommandBase
	ComponentID string `json:"componentId"`
	Distance    int    `json:"distance"`
}

// ScrollAlign ..
type ScrollAlign string

const (
	//ScrollAlignFirst ..
	ScrollAlignFirst ScrollAlign = "first"
	//ScrollAlignCenter ..
	ScrollAlignCenter ScrollAlign = "center"
	//ScrollAlignLast ..
	ScrollAlignLast ScrollAlign = "last"
	//ScrollAlignVisible ..
	ScrollAlignVisible ScrollAlign = "visible"
)

// ScrollToCommand scrolls to ensure that a particular component is in view
type ScrollToCommand struct {
	CommandBase
	Align       ScrollAlign `json:"align"`
	ComponentID string      `json:"componentId"`
}

// ScrollToIndexCommand scrolls forward or backward to ensure that a particular child component is in view
type ScrollToIndexCommand struct {
	CommandBase
	Align       ScrollAlign `json:"align"`
	ComponentID string      `json:"componentId"`
	Index       int         `json:"index"`
}

// SendEventCommand generates and sends a command to Alexa
type SendEventCommand struct {
	CommandBase
	Argument  []interface{} `json:"argument"`
	Component []string      `json:"component"`
}

// SequentialCommand executes a series of commands in order
type SequentialCommand struct {
	CommandBase
	//Slice of commands
	Catch []interface{} `json:"catch"`
	//Slice of commands
	Commands    []interface{} `json:"commands"`
	RepeatCount int           `json:"repeatCount"`
	//Slice of commands
	Finally []interface{} `json:"finally"`
}

// SetFocusCommand Changes the actionable component that is in focus
type SetFocusCommand struct {
	CommandBase
	ComponentID string `json:"componentId"`
}

// SetPagePosition ..
type SetPagePosition string

const (
	//SetPageRelative ..
	SetPageRelative SetPagePosition = "relative"
	//SetPageAbsolute ..
	SetPageAbsolute SetPagePosition = "absolute"
)

// SetPageCommand changes the page displayed in a Pager component
type SetPageCommand struct {
	CommandBase
	ComponentID string          `json:"componentId"`
	Position    SetPagePosition `json:"position"`
	Value       int             `json:"value"`
}

// SetValueCommand changes a dynamic property of a componenet
type SetValueCommand struct {
	CommandBase
	ComponentID string `json:"componentId"`
	Property    string `json:"property"`
	Value       string `json:"value"`
}

// SpeakAlign ..
type SpeakAlign string

// SpeakItemHighlightMode ..
type SpeakItemHighlightMode string

const (
	//SpeakAlignFirst ..
	SpeakAlignFirst SpeakAlign = "first"
	//SpeakAlignLast ..
	SpeakAlignLast SpeakAlign = "last"
	//SpeakAlignCenter ..
	SpeakAlignCenter SpeakAlign = "center"
	//SpeakAlignVisible ..
	SpeakAlignVisible SpeakAlign = "visible"

	//HighlightModeLine ..
	HighlightModeLine SpeakItemHighlightMode = "line"
	//HighlightModeBlock ..
	HighlightModeBlock SpeakItemHighlightMode = "block"
)

// SpeakItemCommand reads teh contents of a single item on the screen
type SpeakItemCommand struct {
	CommandBase
	Align            SpeakAlign             `json:"align"`
	ComponentID      string                 `json:"componentId"`
	HighlightMode    SpeakItemHighlightMode `json:"highlightMode"`
	MinimumDwellTime int                    `json:"minimumDwellTime"`
}

// SpeakListCommand read the contents of a range of items inside a common container.
type SpeakListCommand struct {
	CommandBase
	Align            SpeakAlign `json:"align"`
	ComponentID      string     `json:"componentId"`
	Count            int        `json:"count"`
	MinimumDwellTime int        `json:"minimumDwellTime"`
	Start            int        `json:"start"`
}

/////////////////////////////////// Media Commands /////////////////////////////////

// AudioTrackMode ..
type AudioTrackMode string

const (
	//AudioTrackBackground ..
	AudioTrackBackground AudioTrackMode = "background"
	//AudioTrackForeground ..
	AudioTrackForeground AudioTrackMode = "foreground"
	//AudioTrackNone       ..
	AudioTrackNone AudioTrackMode = "none"
)

// PlayMediaCommand plays media on a video player.
type PlayMediaCommand struct {
	CommandBase
	AudioTrack  AudioTrackMode `json:"audioTrack"`
	ComponentID string         `json:"componentId"`
	Source      string         `json:"source"`
	//Source      []MediaSource       `json:"source"`
}

// MediaCommand ..
type MediaCommand string

const (
	//PlayMedia ..
	PlayMedia MediaCommand = "play"
	//PauseMedia ..
	PauseMedia MediaCommand = "pause"
	//RewindMedia ..
	RewindMedia MediaCommand = "rewind"
	//Next ..
	Next MediaCommand = "next"
	//Previous ..
	Previous MediaCommand = "previous"
	//Seek ..
	Seek MediaCommand = "seek"
	//SetTrack ..
	SetTrack MediaCommand = "setTrack"
)

// ControlMediaCommand controls a media player to play, pause, change tracks, or perform some other common action.
type ControlMediaCommand struct {
	CommandBase
	Command     MediaCommand `json:"command"`
	ComponentID string       `json:"componentId"`
	Value       int          `json:"value"`
}
