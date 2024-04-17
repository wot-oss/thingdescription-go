package thingDescription

type EventElement struct {
	Type         *TypeDeclaration      `json:"@type,omitempty"`
	Cancellation *DataSchema           `json:"cancellation,omitempty"`
	Data         *DataSchema           `json:"data,omitempty"`
	DataResponse *DataSchema           `json:"dataResponse,omitempty"`
	Description  *string               `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	Forms        []FormElementEvent    `json:"forms"`
	Subscription *DataSchema           `json:"subscription,omitempty"`
	Title        *string               `json:"title,omitempty"`
	Titles       map[string]string     `json:"titles,omitempty"`
	URIVariables map[string]DataSchema `json:"uriVariables,omitempty"`
}
