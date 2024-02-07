package thingDescription

type AdditionalResponsesDefinition struct {
	ContentType *string `json:"contentType,omitempty"`
	Schema      *string `json:"schema,omitempty"`
	Success     *bool   `json:"success,omitempty"`
}
