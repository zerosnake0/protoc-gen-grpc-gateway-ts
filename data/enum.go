package data

type EnumValue struct {
	Name   string
	Number int32
}

// Enum is the data out to render Enums in a file
// Enums that nested inside messages will be pulled out to the top level
// Because the way it works in typescript
type Enum struct {
	// Name will be the package level unique name
	// Nested names will concat with their parent messages so that it will remain unique
	// This also means nested type might be a bit ugly in type script but whatever
	Name string
	// Due to the fact that Protos allows alias fields which is not a feature
	// in Typescript, it's better to use string representation of it.
	// So Values here will basically be the name of the field.
	Values []EnumValue
	// Values are integers
	EnumAsInts bool
}

// NewEnum creates an enum instance.
func NewEnum() *Enum {
	return &Enum{
		Name:   "",
		Values: make([]EnumValue, 0),
	}
}
