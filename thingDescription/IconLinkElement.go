package thingDescription

type IconLinkElement struct {
	Anchor   *string   `json:"anchor,omitempty"`
	Href     string    `json:"href"`
	Hreflang *Hreflang `json:"hreflang"`
	Rel      *string   `json:"rel,omitempty"`
	Type     *string   `json:"type,omitempty"`
	Sizes    *string   `json:"sizes,omitempty"`
}
