package thingDescription

type Properties struct {
	AnythingArray []interface{}
	Bool          *bool
	DataSchemaMap map[string]DataSchema
	Double        *float64
	Integer       *int64
	String        *string
}

func (x *Properties) UnmarshalJSON(data []byte) error {
	x.AnythingArray = nil
	x.DataSchemaMap = nil
	object, err := unmarshalUnion(data, &x.Integer, &x.Double, &x.Bool, &x.String, true, &x.AnythingArray, false, nil, true, &x.DataSchemaMap, false, nil, true)
	if err != nil {
		return err
	}
	if object {
	}
	return nil
}

func (x *Properties) MarshalJSON() ([]byte, error) {
	return marshalUnion(x.Integer, x.Double, x.Bool, x.String, x.AnythingArray != nil, x.AnythingArray, false, nil, x.DataSchemaMap != nil, x.DataSchemaMap, false, nil, true)
}
