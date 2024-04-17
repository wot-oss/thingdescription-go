package thingDescription

type DataSchema struct {
	Type             *TypeDeclaration      `json:"@type,omitempty"`
	Const            interface{}           `json:"const,omitempty"`
	ContentEncoding  *string               `json:"contentEncoding,omitempty"`
	ContentMediaType *string               `json:"contentMediaType,omitempty"`
	Default          interface{}           `json:"default,omitempty"`
	Description      *string               `json:"description,omitempty"`
	Descriptions     map[string]string     `json:"descriptions,omitempty"`
	Enum             []interface{}         `json:"enum,omitempty"`
	ExclusiveMaximum *float64              `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum *float64              `json:"exclusiveMinimum,omitempty"`
	Format           *string               `json:"format,omitempty"`
	Items            *Items                `json:"items,omitempty"`
	Maximum          *float64              `json:"maximum,omitempty"`
	MaxItems         *int64                `json:"maxItems,omitempty"`
	MaxLength        *int64                `json:"maxLength,omitempty"`
	Minimum          *float64              `json:"minimum,omitempty"`
	MinItems         *int64                `json:"minItems,omitempty"`
	MinLength        *int64                `json:"minLength,omitempty"`
	MultipleOf       *MultipleOfDefinition `json:"multipleOf,omitempty"`
	OneOf            []DataSchema          `json:"oneOf,omitempty"`
	Properties       *Properties           `json:"properties,omitempty"`
	ReadOnly         *bool                 `json:"readOnly,omitempty"`
	Required         []string              `json:"required,omitempty"`
	Title            *string               `json:"title,omitempty"`
	Titles           map[string]string     `json:"titles,omitempty"`
	DataSchemaType   *DataSchemaType       `json:"type,omitempty"`
	Unit             *string               `json:"unit,omitempty"`
	WriteOnly        *bool                 `json:"writeOnly,omitempty"`
}
