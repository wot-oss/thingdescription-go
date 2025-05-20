package thingDescription

type Hreflang struct {
	String      *string
	StringArray []string
}

func (x *Hreflang) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	_, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	return err
}

func (x *Hreflang) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}
