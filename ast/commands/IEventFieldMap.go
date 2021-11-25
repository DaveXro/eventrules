package commands

type IEventFieldMap interface {
	StringField(name string) (string, error)
	IntegerField(name string) (int, error)
	FloatField(name string) (float32, error)
	Field(name string) (interface{}, error)
}
