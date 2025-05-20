package thingDescription

type FormElementPropertyOp struct {
	Enum        *StickyDescription
	StringArray []string
}

func (x *FormElementPropertyOp) UnmarshalJSON(data []byte) error {
	x.StringArray = nil
	x.Enum = nil
	_, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.StringArray, false, nil, false, nil, true, &x.Enum, false)
	return err
}

func (x *FormElementPropertyOp) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.StringArray != nil, x.StringArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}
