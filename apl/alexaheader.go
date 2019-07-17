package apl

//AlexaHeader ..
type AlexaHeader struct {
	HeaderAttributionImage             string      `json:"headerAttributionImage"`
	HeaderAttributionPrimacy           bool        `json:"headerAttributionPrimacy"`
	HeaderAttributionText              string      `json:"headerAttributionText"`
	HeaderBackButtonAccessibilityLabel string      `json:"headerBackButtonAccessibilityLabel"`
	HeaderBackButtonCommand            interface{} `json:"headerBackButtonCommand"`
	HeaderBackButton                   bool        `json:"headerBackButton"`
	HeaderBackgroundColor              string      `json:"headerBackgroundColor"`
	HeaderDivider                      bool        `json:"headerDivider"`
	HeaderSubtitle                     string      `json:"headerSubtitle"`
	HeaderTitle                        string      `json:"headerTitle"`
	Theme                              string      `json:"theme"`
	Type                               string      `json:"type"`
}
