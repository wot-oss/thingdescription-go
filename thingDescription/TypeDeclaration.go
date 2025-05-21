package thingDescription

type TypeDeclaration struct {
	String      *string
	StringArray []string
}

func (x *TypeDeclaration) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	_, err := unmarshalUnion(data, nil, nil, nil, &x.String, true, &x.StringArray, false, nil, false, nil, false, nil, false)
	return err
}

func (x *TypeDeclaration) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, x.String, x.StringArray != nil, x.StringArray, false, nil, false, nil, false, nil, false)
}
