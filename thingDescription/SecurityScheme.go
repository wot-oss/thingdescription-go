package thingDescription

// Applies to additional SecuritySchemes not defined in the WoT TD specification.
type SecurityScheme struct {
	Type          *TypeDeclaration  `json:"@type"`
	Description   *string           `json:"description,omitempty"`
	Descriptions  map[string]string `json:"descriptions,omitempty"`
	Proxy         *string           `json:"proxy,omitempty"`
	Scheme        string            `json:"scheme"`
	OneOf         []string          `json:"oneOf,omitempty"`
	AllOf         []string          `json:"allOf,omitempty"`
	In            *In               `json:"in,omitempty"`
	Name          *string           `json:"name,omitempty"`
	Qop           *Qop              `json:"qop,omitempty"`
	Alg           *string           `json:"alg,omitempty"`
	Authorization *string           `json:"authorization,omitempty"`
	Format        *string           `json:"format,omitempty"`
	Identity      *string           `json:"identity,omitempty"`
	Flow          *string           `json:"flow,omitempty"`
	Refresh       *string           `json:"refresh,omitempty"`
	Scopes        *TypeDeclaration  `json:"scopes"`
	Token         *string           `json:"token,omitempty"`
}
