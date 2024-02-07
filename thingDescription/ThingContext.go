package thingDescription

type ThingContext struct {
	AnythingArray []interface{}
	Enum          *ThingContextTdURIV1
}

func (x *ThingContext) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.Enum = nil
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.AnythingArray, false, nil, false, nil, true, &x.Enum, false)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *ThingContext) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.AnythingArray != nil, x.AnythingArray, false, nil, false, nil, x.Enum != nil, x.Enum, false)
}
