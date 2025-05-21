package thingDescription

type MultipleOfDefinition struct {
	Double  *float64
	Integer *int64
}

func (x *MultipleOfDefinition) UnmarshalJSON(data []byte) error {
	_, err := unmarshalUnion(data, &x.Integer, &x.Double, nil, nil, false, nil, false, nil, false, nil, false, nil, false)
	return err
}

func (x *MultipleOfDefinition) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, nil, nil, false, nil, false, nil, false, nil, false, nil, false)
}
