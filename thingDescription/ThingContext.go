package thingDescription

type ThingContext struct {
	AnythingArray []interface{}
	Enum          *ThingContextTdURIV1
}

func (x *ThingContext) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.Enum = nil
	_, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.AnythingArray, false, nil, false, nil, true, &x.Enum, false)
	return err
}

func (x *ThingContext) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.AnythingArray != nil, x.AnythingArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}
