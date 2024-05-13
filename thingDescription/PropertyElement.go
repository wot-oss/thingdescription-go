package thingDescription

type PropertyElement struct {
	Type                *TypeDeclaration       `json:"@type,omitempty"`
	Const               interface{}            `json:"const,omitempty"`
	Default             interface{}            `json:"default,omitempty"`
	Description         *string                `json:"description,omitempty"`
	Descriptions        map[string]string      `json:"descriptions,omitempty"`
	Enum                []interface{}          `json:"enum,omitempty"`
	ExclusiveMaximum    *float64               `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum    *float64               `json:"exclusiveMinimum,omitempty"`
	Format              *string                `json:"format,omitempty"`
	Forms               []FormElementProperty  `json:"forms"`
	Items               *Items                 `json:"items,omitempty"`
	Maximum             *float64               `json:"maximum,omitempty"`
	MaxItems            *int64                 `json:"maxItems,omitempty"`
	MaxLength           *int64                 `json:"maxLength,omitempty"`
	Minimum             *float64               `json:"minimum,omitempty"`
	MinItems            *int64                 `json:"minItems,omitempty"`
	MinLength           *int64                 `json:"minLength,omitempty"`
	MultipleOf          *MultipleOfDefinition  `json:"multipleOf,omitempty"`
	Observable          *bool                  `json:"observable,omitempty"`
	OneOf               []DataSchema           `json:"oneOf,omitempty"`
	Properties          *Properties            `json:"properties,omitempty"`
	ReadOnly            *bool                  `json:"readOnly,omitempty"`
	Required            []string               `json:"required,omitempty"`
	Title               *string                `json:"title,omitempty"`
	Titles              map[string]string      `json:"titles,omitempty"`
	PropertyElementType *DataSchemaType        `json:"type,omitempty"`
	Unit                *string                `json:"unit,omitempty"`
	URIVariables        map[string]DataSchema  `json:"uriVariables,omitempty"`
	WriteOnly           *bool                  `json:"writeOnly,omitempty"`
	AdditionalFields    map[string]interface{} `json:",unknown"`
}

func (pe *PropertyElement) ToDataSchema() DataSchema {
	return DataSchema{
		Type:             pe.Type,
		ReadOnly:         pe.ReadOnly,
		WriteOnly:        pe.WriteOnly,
		Const:            pe.Const,
		Default:          pe.Default,
		Enum:             pe.Enum,
		Description:      pe.Description,
		Descriptions:     pe.Descriptions,
		ExclusiveMaximum: pe.ExclusiveMaximum,
		ExclusiveMinimum: pe.ExclusiveMinimum,
		Format:           pe.Format,
		Items:            pe.Items,
		Maximum:          pe.Maximum,
		Minimum:          pe.Minimum,
		MaxItems:         pe.MaxItems,
		MinItems:         pe.MinItems,
		MaxLength:        pe.MaxLength,
		MinLength:        pe.MinLength,
		MultipleOf:       pe.MultipleOf,
		OneOf:            pe.OneOf,
		Properties:       pe.Properties,
		Required:         pe.Required,
		Title:            pe.Title,
		Titles:           pe.Titles,
		Unit:             pe.Unit,
		DataSchemaType:   pe.PropertyElementType,
	}
}
