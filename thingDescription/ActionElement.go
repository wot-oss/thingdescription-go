package thingDescription

type ActionElement struct {
	Type         *TypeDeclaration      `json:"@type,omitempty"`
	Description  *string               `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	Forms        []FormElementAction   `json:"forms"`
	Idempotent   *bool                 `json:"idempotent,omitempty"`
	Input        *DataSchema           `json:"input,omitempty"`
	Output       *DataSchema           `json:"output,omitempty"`
	Safe         *bool                 `json:"safe,omitempty"`
	Synchronous  *bool                 `json:"synchronous,omitempty"`
	Title        *string               `json:"title,omitempty"`
	Titles       map[string]string     `json:"titles,omitempty"`
	URIVariables map[string]DataSchema `json:"uriVariables,omitempty"`
}
