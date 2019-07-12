package apl

//Layout ..
type Layout struct {
	Description string      `json:"description"`
	Parameters  []Parameter `json:"parameters"`
}

//Parameter ..
type Parameter struct {
	Default     interface{} `json:"default,omitempty"`
	Description string      `json:"description,omitempty"`
	Name        string      `json:"name"`
	Type        Type        `json:"type,omitempty"`
}

//Type is a string that defines the type of the value expected in a parameter
type Type string

const (
	//AnyType type
	AnyType Type = "any"
	//ArrayType ..
	ArrayType Type = "array"
	//BooleanType ..
	BooleanType Type = "boolean"
	//ColorType ..
	ColorType Type = "color"
	//ComponentType ..
	ComponentType Type = "component"
	//DimensionType ..
	DimensionType Type = "dimension"
	//IntegerType ..
	IntegerType Type = "integer"
	//MapType ..
	MapType Type = "map"
	//NumberType ..
	NumberType Type = "number"
	//ObjectType ..
	ObjectType Type = "object"
	//StringType ..
	StringType Type = "string"
)

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

//Property ..
type Property struct {
	Bindings           []Binding `json:"bind,omitempty"`
	Entities           []Entity  `json:"entity,omitempty"`
	Height             Dimension `json:"height,omitempty"`
	ID                 string    `json:"id,omitemtpy"`
	InheritParentState bool      `json:"inheritParentState,omitempty"`
	MaxHeight          Dimension `json:"maxHeight,omitempty"`
	MaxWidth           Dimension `json:"maxWidth,omitempty"`
	MinHeight          Dimension `json:"minHeight,omitempty"`
	MinWidth           Dimension `json:"minWidth,omitempty"`
	PaddingBotton      Dimension `json:"paddingBotton,omitemtpy"`
	PaddingLeft        Dimension `json:"paddingLeft,omitemtpy"`
	PaddingRight       Dimension `json:"paddingRight,omitemtpy"`
	PaddingTop         Dimension `json:"paddingTop,omitemtpy"`
	Style              Style     `json:"style,omitempty"`
	Type               string    `json:"type"`
	When               bool      `json:"when,omitempty"`
	Width              Dimension `json:"width,omitempty"`
}

//Import ..
type Import struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Source  string `json:"source,omitempty"`
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
	Boolean     map[string]bool   `json:"boolean,omitempty"`
	Colors      map[string]string `json:"colors,omitempty"`
	Description string            `json:"description,omitempty"`
	Dimensions  map[string]string `json:"dimensions,omitempty"`
	Strings     map[string]string `json:"string,omitempty"`
	When        bool              `json:"when,omitempty"`
}

//Style ..
type Style struct {
	Description  string              `json:"description,omitempty"`
	ParentStyles []string            `json:"extends,omitempty"`
	Values       []map[string]string `json:"values,omitempty"`
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
	Commands    map[string]interface{}        `json:"commands"`
	Description string                        `json:"description,omitempty"`
	Graphics    map[string]AlexaVectorGraphic `json:"graphics"`
	Imports     []Import                      `json:"import,omitempty"`
	//start here
	Layouts      map[string]Layout `json:"layouts"`
	MainTemplate Layout            `json:"mainTemplate"`

	//Slice of Commands
	OnMount  []interface{}          `json:"onMount"`
	Resource Resource               `json:"resources"`
	Settings map[string]interface{} `json:"settings"`
	Styles   map[string]Style       `json:"styles"`
	Theme    string                 `json:"theme"`
	Type     string                 `json:"type"`
	Version  string                 `json:"version"`
}

//New gives provdes an empty document with the
//correct type and version configured
func New() Document {
	var d Document
	d.Type = "APL"
	d.Version = "1.1"

	return d
}
