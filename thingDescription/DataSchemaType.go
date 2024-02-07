package thingDescription

type DataSchemaType string

const (
	Array   DataSchemaType = "array"
	Boolean DataSchemaType = "boolean"
	Integer DataSchemaType = "integer"
	Null    DataSchemaType = "null"
	Number  DataSchemaType = "number"
	Object  DataSchemaType = "object"
	String  DataSchemaType = "string"
)
