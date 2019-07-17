package apl

// ScaleType ..
type ScaleType string

const (
	//ScaleTypeNone ..
	ScaleTypeNone ScaleType = "none"
	//ScaleTypeGrow ..
	ScaleTypeGrow ScaleType = "grow"
	//ScaleTypeShrink ..
	ScaleTypeShrink ScaleType = "shrink"
	//ScaleTypeStretch ..
	ScaleTypeStretch ScaleType = "stretch"
)

//AVGParamType ..
type AVGParamType string

const (
	//AnyParamType ..
	AnyParamType AVGParamType = "any"
	//StringParamType ..
	StringParamType AVGParamType = "string"
	//NumberParamType ..
	NumberParamType AVGParamType = "number"
	//ColorParamType  ..
	ColorParamType AVGParamType = "color"
)

// AVGParameter ..
type AVGParameter struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Type        AVGParamType `json:"type"`
	Default     interface{}  `json:"default"`
}

// AlexaVectorGraphic is a format for defining vector graphics
type AlexaVectorGraphic struct {
	Description string        `json:"description"`
	Height      int           `json:"height"`
	Width       int           `json:"width"`
	Item        []interface{} `json:"item"`

	Parameters      []AVGParameter `json:"parameters"`
	ScaleTypeHeight ScaleType      `json:"scaleTypeHeight"`
	ScaleTypeWidth  ScaleType      `json:"scaleTypeWidth"`

	//Value is AVG at the top level
	Type string `json:"type"`

	//Value is "1.0" at the top level
	Version        string  `json:"version"`
	ViewportHeight float32 `json:"viewportHeight"`
	ViewportWidth  float32 `json:"viewportWidth"`
}

//NewGraphic creates a new graphic with the prefilled type and version
func NewGraphic() AlexaVectorGraphic {
	return AlexaVectorGraphic{Type: "AVG", Version: "1.0"}
}

//AVGPath item defines one or more rendered shapes with common fill and strokes
type AVGPath struct {
	Type        string  `json:"path"`
	FillOpacity float32 `json:"fillOpacity"`

	//This is a color
	Fill          string  `json:"fill"`
	PathData      string  `json:"pathData"`
	StrokeOpacity float32 `json:"strokeOpacity"`

	//This is a color
	Stroke      string  `json:"stroke"`
	StrokeWidth float32 `json:"strokeWidth"`
}

//NewAVGPath returns a new path object with the type preset
func NewAVGPath() AVGPath {
	return AVGPath{Type: "path"}
}

// AVGGroup is used to apply transformations to the children
type AVGGroup struct {
	Type       string        `json:"type"`
	Opacity    float32       `json:"opacity"`
	Rotation   float32       `json:"rotation"`
	PivotX     float32       `json:"pivotX"`
	PivotY     float32       `json:"pivotY"`
	ScaleX     float32       `json:"scaleX"`
	ScaleY     float32       `json:"scaleY"`
	TranslateX float32       `json:"translateX"`
	TranslateY float32       `json:"translateY"`
	Items      []interface{} `json:"item"`
}

//NewAVGGroup ..
func NewAVGGroup() AVGGroup {
	return AVGGroup{Type: "group"}
}
