package apl

//Layout ..
type Layout struct {
	Description string          `json:"description"`
	Parameters  []APLParameters `json:"parameters"`
}

//Parameter ..
type Parameter struct {
	Default     interface{} `json:"default,omitempty"`
	Description string      `json:"description,omitempty"`
	Name        string      `json:"name"`
	Type        string      `json:"type,omitempty"`
}

//BindType ..
type BindType string

//Binding ..
type Binding struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Type  BindType    `json:"type"`
}

//Entity ..
type Entity interface{}

//Dimension ..
type Dimension string

//Style ..
type Style string

//Property ..
type Property struct {
	Bindings           []APLBinding `json:"bind,omitempty"`
	Entities           []APLEntity  `json:"entity,omitempty"`
	Height             Dimension    `json:"height,omitempty"`
	ID                 string       `json:"id,omitemtpy"`
	InheritParentState bool         `json:"inheritParentState,omitempty"`
	MaxHeight          Dimension    `json:"maxHeight,omitempty"`
	MaxWidth           Dimension    `json:"maxWidth,omitempty"`
	MinHeight          Dimension    `json:"minHeight,omitempty"`
	MinWidth           Dimension    `json:"minWidth,omitempty"`
	PaddingBotton      Dimension    `json:"paddingBotton,omitemtpy"`
	PaddingLeft        Dimension    `json:"paddingLeft,omitemtpy"`
	PaddingRight       Dimension    `json:"paddingRight,omitemtpy"`
	PaddingTop         Dimension    `json:"paddingTop,omitemtpy"`
	Style              Style        `json:"style,omitempty"`
	Type               string       `json:"type"`
	When               bool         `json:"when,omitempty"`
	Width              Dimension    `json:"width,omitempty"`
}

//Settings ..
type Settings struct {
	IdleTimeout int `json:"idleTimeout,omitempty"`
}

//Import ..
type Import struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

//ColorVals ..
type ColorVals struct {
	ColorTextPrimary string `json:"colorTextPrimary,omitempty"`
	//TODO .. What is missing?
}

//DimensionVals ..
type DimensionVals struct {
	TextSizeBody          int `json:"textSizeBody"`
	TextSizePrimary       int `json:"textSizePrimary"`
	TextSizeSecondary     int `json:"textSizeSecondary"`
	TextSizeSecondaryHint int `json:"textSizeSecondaryHint"`
	SpacingThin           int `json:"spacingThin,omitempty"`
	SpacingSmall          int `json:"spacingSmall,omitempty"`
	SpacingMedium         int `json:"spacingMedium,omitempty"`
	SpacingLarge          int `json:"spacingLarge,omitempty"`
	SpacingExtraLarge     int `json:"spacingExtraLarge,omitempty"`
	MarginTop             int `json:"marginTop,omitempty"`
	MarginLeft            int `json:"marginLeft,omitempty"`
	MarginRight           int `json:"marginRight,omitempty"`
	MarginBottom          int `json:"marginBottom,omitempty"`
}

//Resource ..
type Resource struct {
	Description string        `json:"description,omitempty"`
	Colors      ColorVals     `json:"colors,omitempty"`
	When        string        `json:"when,omitempty"`
	Dimensions  DimensionVals `json:"dimensions,omitempty"`
}

//TextStyleValue ..
type TextStyleValue struct {
	Color      string `json:"color,omitempty"`
	FontWeight string `json:"fontWeight,omitempty"`
	//TODO - What's missing
}

//TextStyle ..
type TextStyle struct {
	Description string           `json:"description"`
	Values      []TextStyleValue `json:"values"`
}

//MixinValues ..
type MixinValues struct {
	FontWeight string `json:"fontWeight"`
	//TODO what's missing
}

//Mixin ..
type Mixin struct {
	Values MixinValues `json:"values,omitemtpy"`
}

//TextStyle ..
type TextStyle struct {
	Extends []string `json:"extend"`
}

//TextHint ..
type TextHint struct {
	FontFamily string `json:fontFamily,omitempty"`
	FontStyle  string `json:fontStyle,omitempty"`
	FontSize   string `json:fontSize,omitempty"`
	Color      string `json:color,omitempty"`
}

//TextStyleHint ..
type TextStyleHint struct {
	TextHint TextHint `json:"values"`
}

//Styles ..
type Styles struct {
	TextStyleBase          TextStyle     `json:"TextStyleBase,omitempty"`
	TextStyleBase0         TextStyle     `json:"TextStyleBase0,omitempty"`
	TextStyleBase1         TextStyle     `json:"TextStyleBase1,omitempty"`
	MixinBody              Mixin         `json:"MixinBody,omitempty"`
	MixinPrimary           Mixin         `json:"MixinPrimary,omitempty"`
	MixinSecondary         Mixin         `json:"MixinPrimary,omitempty"`
	TextStylePrimary       TextStyle     `json:"textStylePrimary,omitempty"`
	TextStyleSecondary     TextStyle     `json:"textStyleSecondary,omitempty"`
	TextStyleBody          TextStyle     `json:"textStyleBody,omitempty"`
	TextStyleSecondaryHint TextStyleHint `json:"textStyleSecondaryHint"`
}

//MountCommand ..
type MountCommand struct {
	Type     string `json:"type,omitempty"`
	Commands string `json:"commands,omitempty"`
}

//OnMount ..
type OnMount struct {
	Type     string         `json:"type"`
	Commands []MountCommand `json:"commands,omitemtpy"`
	Finally  string         `json:"finally,omitemtpy"`
}

//Document ..
type Document struct {
	Type        string            `json:"type"`
	Version     string            `json:"version"`
	Settings    Settings          `json:"settings"`
	Theme       string            `json:"theme"`
	Description string            `json:"description,omitempty"`
	Imports     []Import          `json:"import,omitempty"`
	Resources   []Resource        `json:"resources"`
	Styles      Styles            `json:"styles"`
	OnMount     []MountInsruction `json:"onMount"`
	Layouts     map[string]Layout `json:"layouts"`
}
