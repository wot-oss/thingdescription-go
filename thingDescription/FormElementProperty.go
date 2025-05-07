package thingDescription

import (
	"net/url"

	"github.com/go-json-experiment/json"
)

type FormElementProperty struct {
	Op                  *FormElementPropertyOp          `json:"op,omitempty"`
	AdditionalResponses []AdditionalResponsesDefinition `json:"additionalResponses,omitempty"`
	ContentCoding       *string                         `json:"contentCoding,omitempty"`
	ContentType         *string                         `json:"contentType,omitempty"`
	Href                url.URL                         `json:"href"`
	Response            *ExpectedResponse               `json:"response,omitempty"`
	Scopes              *TypeDeclaration                `json:"scopes,omitempty"`
	Security            *TypeDeclaration                `json:"security,omitempty"`
	Subprotocol         *string                         `json:"subprotocol,omitempty"`
	AdditionalFields    map[string]interface{}          `json:",unknown"`
}

func (r *FormElementProperty) UnmarshalJSON(data []byte) error {
	type FormElementPropertyRaw FormElementProperty

	tmp := struct {
		Href string `json:"href"`
		*FormElementPropertyRaw
	}{
		FormElementPropertyRaw: (*FormElementPropertyRaw)(r),
	}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	href, err := url.Parse(tmp.Href)
	if err != nil {
		return err
	}
	r.Href = *href
	return nil
}

func (r FormElementProperty) MarshalJSON() ([]byte, error) {
	type FormElementPropertyRaw FormElementProperty
	return json.Marshal(&struct {
		Href string `json:"href"`
		*FormElementPropertyRaw
	}{
		Href:                   r.Href.String(),
		FormElementPropertyRaw: (*FormElementPropertyRaw)(&r),
	})
}
