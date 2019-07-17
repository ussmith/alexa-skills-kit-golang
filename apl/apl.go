package apl

//Layout ..
type Layout struct {
	Description string        `json:"description"`
	Parameters  []Parameter   `json:"parameters"`
	Item        []interface{} `json:"item"`
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

//Import ..
type Import struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Source  string `json:"source,omitempty"`
}

//Resource ..
type Resource struct {
	Boolean     map[string]bool   `json:"boolean,omitempty"`
	Colors      map[string]string `json:"colors,omitempty"`
	Description string            `json:"description,omitempty"`
	Dimensions  map[string]int    `json:"dimensions,omitempty"`
	Strings     map[string]string `json:"string,omitempty"`
	When        string            `json:"when,omitempty"`
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
	Type     string        `json:"type"`
	Commands []interface{} `json:"commands,omitemtpy"`
	Finally  string        `json:"finally,omitemtpy"`
}

// Setting holds a map of key-value pairs that define document-wide properties.
type Setting struct {
	Key   string
	Value interface{}
}

//Document ..
type Document struct {
	Commands     map[string]interface{}        `json:"commands,omitempty"`
	Description  string                        `json:"description,omitempty"`
	Graphics     map[string]AlexaVectorGraphic `json:"graphics,omitempty"`
	Imports      []Import                      `json:"import,omitempty"`
	Layouts      map[string]Layout             `json:"layouts,omitempty"`
	MainTemplate Layout                        `json:"mainTemplate"`

	//Slice of Commands
	OnMount   OnMount                `json:"onMount,omitempty"`
	Resources []Resource             `json:"resources,omitempty"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
	Styles    map[string]Style       `json:"styles,omitempty"`
	Theme     string                 `json:"theme,omitempty"`
	Type      string                 `json:"type,omitempty"`
	Version   string                 `json:"version,omitempty"`
}

// DocumentWrapper ..
type DocumentWrapper struct {
	Document `json:"document"`
}

//NewDocument gives provdes an empty document with the
//correct type and version configured
func NewDocument() DocumentWrapper {
	var d DocumentWrapper
	d.Type = "APL"
	d.Version = "1.1"

	return d
}

//WithSettings adds settings to the docuemnt
func (d *DocumentWrapper) WithSettings(setting ...Setting) *DocumentWrapper {
	if d.Settings == nil {
		d.Settings = make(map[string]interface{})
	}

	for _, v := range setting {
		d.Settings[v.Key] = v.Value
	}

	return d
}

//WithDescription ..
func (d *DocumentWrapper) WithDescription(description string) *DocumentWrapper {
	d.Description = description
	return d
}

//WithTheme ..
func (d *DocumentWrapper) WithTheme(theme string) *DocumentWrapper {
	d.Theme = theme
	return d
}

//WithImports ..
func (d *DocumentWrapper) WithImports(imports ...Import) *DocumentWrapper {

	for _, i := range imports {
		d.Imports = append(d.Imports, i)
	}

	return d
}

//WithResources ..
func (d *DocumentWrapper) WithResources(resources ...Resource) *DocumentWrapper {
	for _, r := range resources {
		d.Resources = append(d.Resources, r)
	}
	return d
}

//WithStyles ..
func (d *DocumentWrapper) WithStyles(styles map[string]Style) *DocumentWrapper {
	d.Styles = styles
	return d
}

//AddStyle ..
func (d *DocumentWrapper) AddStyle(name string, style Style) *DocumentWrapper {
	if d.Styles == nil {
		d.Styles = make(map[string]Style)
	}

	d.Styles[name] = style
	return d
}

//WithOnMount ..
func (d *DocumentWrapper) WithOnMount(onMount OnMount) *DocumentWrapper {
	d.OnMount = onMount
	return d
}

//WithGraphics ..
func (d *DocumentWrapper) WithGraphics(graphics map[string]AlexaVectorGraphic) *DocumentWrapper {
	d.Graphics = graphics
	return d
}

//AddGraphic ..
func (d *DocumentWrapper) AddGraphic(name string, graphic AlexaVectorGraphic) *DocumentWrapper {
	if d.Graphics == nil {
		d.Graphics = make(map[string]AlexaVectorGraphic)
	}

	d.Graphics[name] = graphic
	return d

}
