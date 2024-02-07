package thingDescription

type Items struct {
	DataSchema      *DataSchema
	DataSchemaArray []DataSchema
}

func (x *Items) UnmarshalJSON(data []byte) error {
	x.DataSchemaArray = nil
	x.DataSchema = nil
	var c DataSchema
	object, err := unmarshalUnion(data, nil, nil, nil, nil, true, &x.DataSchemaArray, true, &c, false, nil, false, nil, false)
	if err != nil {
		return err
	}
	if object {
		x.DataSchema = &c
	}
	return nil
}

func (x *Items) MarshalJSON() ([]byte, error) {
	return marshalUnion(nil, nil, nil, nil, x.DataSchemaArray != nil, x.DataSchemaArray, x.DataSchema != nil, x.DataSchema, false, nil, false, nil, false)
}
