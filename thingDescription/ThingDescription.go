package thingDescription

import (
	"bytes"
	"io"
	"net/url"

	"github.com/fredbi/uri"
	"github.com/go-json-experiment/json"
)

func UnmarshalThingDescription(data []byte) (ThingDescription, error) {
	var r ThingDescription
	err := json.Unmarshal(data, &r)
	return r, err
}

func (c *ThingDescription) HandleJSONData(data map[string]interface{}) error {
	c.AdditionalFields = data
	return nil
}

func (r *ThingDescription) UnmarshalJSONRead(input io.Reader) error {
	type ThingDescriptionRaw ThingDescription
	tmp := struct {
		Base string `json:"base"`
		ID   string `json:"id"`
		*ThingDescriptionRaw
	}{
		ThingDescriptionRaw: (*ThingDescriptionRaw)(r),
	}
	err := json.UnmarshalRead(input, &tmp)
	if err != nil {
		return err
	}
	base, err := url.Parse(tmp.Base)
	if err != nil {
		return err
	}
	r.Base = *base
	id, err := uri.Parse(tmp.ID)
	if err != nil {
		return err
	}
	r.ID = id
	return nil
}

func (r ThingDescription) MarshalJSONWrite(output io.Writer) error {
	type ThingDescriptionRaw ThingDescription
	return json.MarshalWrite(output, &struct {
		Base string `json:"base,omitempty"`
		ID   string `json:"id"`
		*ThingDescriptionRaw
	}{
		Base:                r.Base.String(),
		ID:                  r.ID.String(),
		ThingDescriptionRaw: (*ThingDescriptionRaw)(&r),
	})
}

func (r *ThingDescription) UnmarshalJSON(data []byte) error {
	return r.UnmarshalJSONRead(bytes.NewReader(data))
}

func (r ThingDescription) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	err := r.MarshalJSONWrite(buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// JSON Schema for validating TD instances against the TD information model. TD instances
// can be with or without terms that have default values
type ThingDescription struct {
	Context             *ThingContext              `json:"@context"`
	Type                *TypeDeclaration           `json:"@type,omitempty"`
	Actions             map[string]ActionElement   `json:"actions,omitempty"`
	Base                url.URL                    `json:"base,omitempty"`
	Created             string                     `json:"created,omitempty"`
	Description         string                     `json:"description,omitempty"`
	Descriptions        map[string]string          `json:"descriptions,omitempty"`
	Events              map[string]EventElement    `json:"events,omitempty"`
	Forms               []FormElementRoot          `json:"forms,omitempty"`
	ID                  uri.URI                    `json:"id,omitempty"`
	Links               []IconLinkElement          `json:"links,omitempty"`
	Modified            *string                    `json:"modified,omitempty"`
	Profile             *TypeDeclaration           `json:"profile,omitempty"`
	Properties          map[string]PropertyElement `json:"properties,omitempty"`
	SchemaDefinitions   map[string]DataSchema      `json:"schemaDefinitions,omitempty"`
	Security            *TypeDeclaration           `json:"security"`
	SecurityDefinitions map[string]SecurityScheme  `json:"securityDefinitions"`
	Support             string                     `json:"support,omitempty"`
	Title               string                     `json:"title"`
	Titles              map[string]string          `json:"titles,omitempty"`
	URIVariables        map[string]DataSchema      `json:"uriVariables,omitempty"`
	Version             *Version                   `json:"version,omitempty"`
	AdditionalFields    map[string]interface{}     `json:",unknown"`
}
