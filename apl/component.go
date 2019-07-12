package apl

// DisplayState ..
type DisplayState string

const (

	//Invisible ..
	Invisible DisplayState = "invisible"
	//None ..
	None DisplayState = "none"
	//Normal ..
	Normal DisplayState = "normal"
)

// Transform ..
type Transform struct {
	Rotate     *float32 `json:"rotate,omitempty"`
	Scale      *float32 `json:"scale,omitempty"`
	ScaleX     *float32 `json:"scaleX,omitempty"`
	ScaleY     *float32 `json:"scaleY,omitempty"`
	SkewX      *float32 `json:"skewX,omitempty"`
	SkewY      *float32 `json:"skewY,omitempty"`
	TranslateX *float32 `json:"translateX,omitempty"`
	TranslateY *float32 `json:"translateY,omitempty"`
}

//Component ..
type Component struct {
	AccessibilityLabel string        `json:"accessibilityLabel,omitempty"`
	Bind               []Binding     `json:"bind,omitempty"`
	Description        string        `json:"description,omitempty"`
	Checked            bool          `json:"checked,omitempty"`
	Disabled           bool          `json:"disabled,omitempty"`
	Display            DisplayState  `json:"display,omitempty"`
	Entity             []interface{} `json:"entity,omitempty"`
	Height             string        `json:"height,omitempty"`
	ID                 string        `json:"id,omitempty"`
	InheritParentState bool          `json:"inheritParentState,omitempty"`
	MaxHeight          string        `json:"maxHeight,omitEmpty"`
	MinHeight          string        `json:"minHeight,omitEmpty"`
	MaxWidth           string        `json:"maxWidth,omitEmpty"`
	MinWidth           string        `json:"minWidth,omitEmpty"`
	//Slice of command
	OnMount       []interface{} `json:"onMount,omitempty"`
	Opacity       float32       `json:"opacity,omitempty"`
	PaddingBottom string        `json:"paddingBottom,omitempty"`
	PaddingLeft   string        `json:"paddingLeft,omitempty"`
	PaddingRight  string        `json:"paddingRight,omitempty"`
	PaddingTop    string        `json:"paddingTop,omitempty"`
	Speech        interface{}   `json:"opaque,omitempty"`
	Style         Style         `json:"style,omitempty"`
	Transform     []Transform   `json:"transform,omitempty"`
	Type          string        `json:"type"`
	When          string        `json:"when,omitempty"`
	Width         string        `json:"width,omitempty"`
}

// AlignType ..
type AlignType string

//Direction ..
type Direction string

//JustifyContent ..
type JustifyContent string

const (
	//StretchAlignType ..
	StretchAlignType AlignType = "stretch"
	//CenterAlignType ..
	CenterAlignType AlignType = "center"
	//StartAlignType ..
	StartAlignType AlignType = "start"
	//EndAlignType ..
	EndAlignType AlignType = "end"
	//BaselineAlignType ..
	BaselineAlignType AlignType = "baseline"

	//Column ..
	Column Direction = "column"
	//Row ..
	Row Direction = "row"

	//StartJustify ..
	StartJustify JustifyContent = "start"
	//EndJustify ..
	EndJustify JustifyContent = "end"
	//CenterJustify ..
	CenterJustify JustifyContent = "center"
	//SpaceBetweenJustify ..
	SpaceBetweenJustify JustifyContent = "spaceBetween"
	//SpaceAround ..
	SpaceAround JustifyContent = "spaceAround"
)

//Container ..
type Container struct {
	Component
	AlignItems     AlignType      `json:"alignItems"`
	Data           []interface{}  `json:"data"`
	Direction      Direction      `json:"direction"`
	FirstItem      []Component    `json:"component"`
	Items          []Component    `json:"item"`
	JustifyContent JustifyContent `json:"justifyContent"`
	LastItem       []Component    `json:"lastItem"`
	Numbered       bool           `json:"numbered"`
}

//Frame ..
type Frame struct {
	Component
	BackgroundColor         string      `json:"backgroundColor"`
	BorderBottomLeftRadius  string      `json:"borderBottomLeftRadius"`
	BorderBottomRightRadius string      `json:"borderBottomRightRadius"`
	BorderColor             string      `json:"borderColor"`
	BorderRadius            string      `json:"borderRadius"`
	BorderTopLeftRadius     string      `json:"borderTopLeftRadius"`
	BorderTopRightRadius    string      `json:"borderTopRightRadius"`
	BorderWidth             float32     `json:"borderWidth"`
	Item                    []Component `json:"item"`
}

// Align ..
type Align string

const (
	//BottomAlign ..
	BottomAlign Align = "bottom"
	//BottomLeftAlign ..
	BottomLeftAlign Align = "bottom-left"
	//BottomRightAlign ..
	BottomRightAlign Align = "bottom-right"
	//CenterAlign ..
	CenterAlign Align = "center"
	//LeftAlign ..
	LeftAlign Align = "left"
	//RightAlign ..
	RightAlign Align = "right"
	//TopAlign ..
	TopAlign Align = "top"
	//TopLeftAlign ..
	TopLeftAlign Align = "top-left"
	//TopRigthAlign ..
	TopRigthAlign Align = "top-right"
)

//type Filter struct {
//   Type string `json:"filter"`
//}

// Scale ..
type Scale string

const (
	//NoneScale ..
	NoneScale Scale = "none"
	//FillScale ..
	FillScale Scale = "fill"
	//BestFillScale ..
	BestFillScale Scale = "best-fill"
	//BestFitScale ..
	BestFitScale Scale = "best-fit"
	//BestFitDownScale ..
	BestFitDownScale Scale = "best-fit-down"
)

// Image ..
type Image struct {
	Component
	Alignment    Align  `json:"align"`
	BorderRadius string `json:"borderRadius"`
	//Filter []Filter `json:"filter"`
	OverlayColor    string `json:"overlayColor"`
	OverlayGradient string `json:"overlayGradient"`
	Scale           Scale  `json:"scale"`
	Source          string `json:"source"`
}

//Navigation ..
type Navigation string

const (
	//NormalNav ..
	NormalNav Navigation = "normal"
	//NoNav ..
	NoNav Navigation = "none"
	//WrapNav ..
	WrapNav Navigation = "wrap"
	//ForwardOnly ..
	ForwardOnly Navigation = "forward-only"
)

//Pager ..
type Pager struct {
	Component
	Data        interface{} `json:"data"`
	FirstItem   []Component `json:"firstItem"`
	InitialPage int         `json:"initialPage"`
	Item        []Component `json:"item"`
	LastItem    []Component `json:"lastItem"`
	Navigation  Navigation  `json:"navigation"`
	//Slice of command
	OnPageChange []interface{} `json:"onPageChange"`
}

//ScrollView ..
type ScrollView struct {
	Item []Component `json:"component"`
	//Slice of command
	OnScroll []interface{} `json:"onScroll"`
}

// ScrollDirection ..
type ScrollDirection string

const (
	//HorizontalScroll ..
	HorizontalScroll ScrollDirection = "horizontal"
	//VerticalScroll ..
	VerticalScroll ScrollDirection = "vertical"
)

//Sequence uses a data set to inflate a repeating set of components and display them in a long scrolling list.
type Sequence struct {
	Component
	Data            []interface{}   `json:"data"`
	ScrollDirection ScrollDirection `json:"scrollDirection"`
	FirstItem       []Component     `json:"firstItem"`
	Item            []Component     `json:"item"`
	LastItem        []Component     `json:"lastItem"`
	Numbered        bool            `json:"numbered"`
	//Slice of command
	OnScroll []interface{} `json:"onScroll"`
}

//FontStyle to display
type FontStyle string

// FontWeight ..
type FontWeight string

// TextAlign ..
type TextAlign string

// TextAlignVertical ..
type TextAlignVertical string

const (
	//NormalFont ..
	NormalFont FontStyle = "normal"
	//ItalicFont ..
	ItalicFont FontStyle = "italic"

	//NormalWeight ..
	NormalWeight FontWeight = "normal"
	//BoldWeight ..
	BoldWeight FontWeight = "bold"
	//Weight100 ..
	Weight100 FontWeight = "100"
	//Weight200 ..
	Weight200 FontWeight = "200"
	//Weight300 ..
	Weight300 FontWeight = "300"
	//Weight400 ..
	Weight400 FontWeight = "400"
	//Weight500 ..
	Weight500 FontWeight = "500"
	//Weight600 ..
	Weight600 FontWeight = "600"
	//Weight700 ..
	Weight700 FontWeight = "700"
	//Weight800 ..
	Weight800 FontWeight = "800"
	//Weight900 ..
	Weight900 FontWeight = "900"

	//AlignAuto ..
	AlignAuto TextAlign = "auto"
	//AlignLeft ..
	AlignLeft TextAlign = "left"
	//AlignRight ..
	AlignRight TextAlign = "right"
	//AlignCenter ..
	AlignCenter TextAlign = "center"

	//VertAlignAuto ..
	VertAlignAuto TextAlignVertical = "auto"
	//VertAlignTop ..
	VertAlignTop TextAlignVertical = "top"
	//VertAlignBottom ..
	VertAlignBottom TextAlignVertical = "bottom"
	//VertAlignCenter ..
	VertAlignCenter TextAlignVertical = "center"
)

//Text ..
type Text struct {
	Component
	Color             string            `json:"color"`
	FontFamily        string            `json:"fontFamily"`
	FontSize          string            `json:"fontSize"`
	FontStyle         FontStyle         `json:"fontStyle"`
	FontWeight        FontWeight        `json:"fontWeight"`
	LetterSpacing     string            `json:"letterSpacing"`
	LineHeight        string            `json:"lineHeight"`
	MaxLines          int               `json:"maxLines"`
	Text              string            `json:"text"`
	TextAlign         TextAlign         `json:"textAlign"`
	TextAlignVertical TextAlignVertical `json:"textAlignVertical"`
}

// TouchWrapper ..
type TouchWrapper struct {
	Component
	Item []Component `json:"item"`
	//Slice of command
	OnPress []interface{} `json:"onPress"`
}

// VectorGraphicAlign ..
type VectorGraphicAlign string

// VectorGraphicScale ..
type VectorGraphicScale string

const (
	//VGAlignBotton ..
	VGAlignBotton VectorGraphicAlign = "bottom"
	//VGAlignBottonLeft ..
	VGAlignBottonLeft VectorGraphicAlign = "bottom-left"
	//VGAlignBottonRight ..
	VGAlignBottonRight VectorGraphicAlign = "bottom-right"
	//VGAlignCenter ..
	VGAlignCenter VectorGraphicAlign = "center"
	//VGAlignLeft ..
	VGAlignLeft VectorGraphicAlign = "left"
	//VGAlignRight ..
	VGAlignRight VectorGraphicAlign = "right"
	//VGAlignTop ..
	VGAlignTop VectorGraphicAlign = "top"
	//VGAlignTopLeft ..
	VGAlignTopLeft VectorGraphicAlign = "top-left"
	//VGAlignTopRight ..
	VGAlignTopRight VectorGraphicAlign = "top-right"

	//VGScaleNone ..
	VGScaleNone VectorGraphicScale = "none"
	//VGScaleFill ..
	VGScaleFill VectorGraphicScale = "fill"
	//VGScaleBestFill ..
	VGScaleBestFill VectorGraphicScale = "best-fill"
	//VGScaleBestFit ..
	VGScaleBestFit VectorGraphicScale = "best-fit"
)

// VectorGraphic ..
type VectorGraphic struct {
	Alignment VectorGraphicAlign `json:"align"`
	Scale     VectorGraphicScale `json:"scale"`
	Source    string             `json:"source"`
}

// AudioTrackType ..
type AudioTrackType string

// AudioScale ..
type AudioScale string

const (
	//ATForeground ..
	ATForeground AudioTrackType = "foreground"
	//ATBackground ..
	ATBackground AudioTrackType = "background"
	//ATNone ..
	ATNone AudioTrackType = "none"

	//ATScaleBestFill ..
	ATScaleBestFill AudioScale = "best-fill"
	//ATScaleBestFit ..
	ATScaleBestFit AudioScale = "best-fit"
)

// MediaSource ..
type MediaSource struct {
	Description string `json:"description"`
	Offset      int    `json:"offset"`
	URL         string `json:"url"`
	RepeatCount int8   `json:"repeatCount"`
}

// Video ..
type Video struct {
	AudioTrack AudioTrackType `json:"audioTrack"`
	Autoplay   bool           `json:"autoplay"`
	//Slice of command
	OnEnd []interface{} `json:"onEnd"`
	//Slice of command
	OnPause []interface{} `json:"onPause"`
	//Slice of command
	OnPlay []interface{} `json:"onPlay"`
	//Slice of command
	OnTimeUpdate []interface{} `json:"onTimeUpdate"`
	//Slice of command
	OnTrackUpdate []interface{} `json:"onTrackUpdate"`
	Scale         AudioScale    `json:"scale"`
	Source        []MediaSource `json:"source"`
}
