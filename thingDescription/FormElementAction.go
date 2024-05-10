package thingDescription

type FormElementAction struct {
	Op                  *FormElementActionOp            `json:"op"`
	AdditionalResponses []AdditionalResponsesDefinition `json:"additionalResponses,omitempty"`
	ContentCoding       *string                         `json:"contentCoding,omitempty"`
	ContentType         *string                         `json:"contentType,omitempty"`
	Href                string                          `json:"href"`
	Response            *ExpectedResponse               `json:"response,omitempty"`
	Scopes              *TypeDeclaration                `json:"scopes"`
	Security            *TypeDeclaration                `json:"security"`
	Subprotocol         *string                         `json:"subprotocol,omitempty"`
	AdditionalFields    map[string]interface{}          `json:",unknown"`
}
