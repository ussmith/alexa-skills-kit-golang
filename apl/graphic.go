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
	Description string               `json:"description"`
	Height      string               `json:"height"`
	Item        []AlexaVectorGraphic `json:"item"`

	Parameters      []AVGParameter `json:"parameters"`
	ScaleTypeHeight ScaleType      `json:"scaleTypeHeight"`
	ScaleTypeWidth  ScaleType      `json:"scaleTypeWidth"`

	//Value is AVG
	Type string `json:"type"`

	//Value is "1.0"
	Version        string  `json:"version"`
	ViewportHeight float32 `json:"viewportHeight"`
	ViewportWidth  float32 `json:"viewportWidth"`
	Width          string  `json:"width"`
}
