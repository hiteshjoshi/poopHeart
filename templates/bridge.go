package templates

type ClassDetails struct {
	Name string
	Variables []Variable
}

type Variable struct{
	Name string
	Type string
	InitialValue string
	CubitActions []CubitAction
}

type CubitAction struct{
	Name string
	IntAction IntActions
	NewValue string
	StringAction StringActions
}

//This defines the actions integers are supposed to do
//@TODO: Action for HTTP
type IntActions int
const (
	Add IntActions = iota
	Subtract
	Multiply
	Divide
	UpdatedInt
)

type StringActions int
const (
	Update StringActions = iota
)

