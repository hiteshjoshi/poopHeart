package main

import (
	"bytes"
	"fmt"
	"github.com/hiteshjoshi/poopTemplates/templates"
)

func main() {

	 variables := []templates.Variable{
		{
			Name: "counter",
			InitialValue: "0",
			Type: "int",
			CubitActions: []templates.CubitAction{
				{
					Name: "increment",
					IntAction: templates.Add,
					NewValue: "1",
				},
				{
					Name: "decrement",
					IntAction: templates.Subtract,
					NewValue: "1",
				},
			},
		},
	}
	// qtc creates Write* function for each template function.
	// Such functions accept io.Writer as first parameter:
	classDetails := &templates.ClassDetails{
		Name: "Counter",
		Variables: variables,
	}
	var buf bytes.Buffer
	templates.WriteGenerateState(&buf, classDetails)
	fmt.Printf("buf=\n%s", buf.Bytes())

	templates.WriteGenerateCubit(&buf, classDetails)
	fmt.Printf("buf=\n%s", buf.Bytes())


}