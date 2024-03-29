// Code generated by qtc from "cubit.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line cubit.qtpl:1
package templates

//line cubit.qtpl:1
import "fmt"

//line cubit.qtpl:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line cubit.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line cubit.qtpl:2
func StreamGenerateCubit(qw422016 *qt422016.Writer, class *ClassDetails) {
//line cubit.qtpl:2
	qw422016.N().S(`
import 'package:bloc/bloc.dart';
import 'package:meta/meta.dart';
import 'package:equatable/equatable.dart';

part '`)
//line cubit.qtpl:7
	qw422016.E().S(class.Name)
//line cubit.qtpl:7
	qw422016.N().S(`_state.dart';

class `)
//line cubit.qtpl:9
	qw422016.E().S(class.Name)
//line cubit.qtpl:9
	qw422016.N().S(`Cubit extends Cubit<`)
//line cubit.qtpl:9
	qw422016.E().S(class.Name)
//line cubit.qtpl:9
	qw422016.N().S(`State> {
  `)
//line cubit.qtpl:10
	qw422016.E().S(class.Name)
//line cubit.qtpl:10
	qw422016.N().S(`Cubit() : super(`)
//line cubit.qtpl:10
	qw422016.E().S(class.Name)
//line cubit.qtpl:10
	qw422016.N().S(`State());

  `)
//line cubit.qtpl:12
	qw422016.E().S(class.cubitString())
//line cubit.qtpl:12
	qw422016.N().S(`
}

`)
//line cubit.qtpl:15
}

//line cubit.qtpl:15
func WriteGenerateCubit(qq422016 qtio422016.Writer, class *ClassDetails) {
//line cubit.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line cubit.qtpl:15
	StreamGenerateCubit(qw422016, class)
//line cubit.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line cubit.qtpl:15
}

//line cubit.qtpl:15
func GenerateCubit(class *ClassDetails) string {
//line cubit.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line cubit.qtpl:15
	WriteGenerateCubit(qb422016, class)
//line cubit.qtpl:15
	qs422016 := string(qb422016.B)
//line cubit.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line cubit.qtpl:15
	return qs422016
//line cubit.qtpl:15
}

//line cubit.qtpl:21
func (C *ClassDetails) cubitString() string {

	var functionStr = ""

	for _, variable := range C.Variables {

		for _, action := range variable.CubitActions {

			switch variable.Type {
			case "string":
				functionStr += fmt.Sprintf("\tvoid %s(){\n\temit(state.copyWith(%s:%s));\n\t}\n", action.Name, variable.Name, action.NewValue)
			case "int":
				switch action.IntAction {
				case Add:
					functionStr += fmt.Sprintf("\tvoid %s(){\n\temit(state.copyWith(%s:state.%s+%s));\n\t}\n", action.Name, variable.Name, variable.Name, action.NewValue)
				case Subtract:
					functionStr += fmt.Sprintf("\tvoid %s(){\n\temit(state.copyWith(%s:state.%s-%s));\n\t}\n", action.Name, variable.Name, variable.Name, action.NewValue)
				case Multiply:
					functionStr += fmt.Sprintf("\tvoid %s(){\n\temit(state.copyWith(%s:state.%s*%s));\n\t}\n", action.Name, variable.Name, variable.Name, action.NewValue)
				case Divide:
					functionStr += fmt.Sprintf("\tvoid %s(){\n\temit(state.copyWith(%s:state.%s/%s));\n\t}\n", action.Name, variable.Name, variable.Name, action.NewValue)
				case UpdatedInt:
					functionStr += fmt.Sprintf("\tvoid %s(){\n\temit(state.copyWith(%s:%s));\n\t}\n", action.Name, variable.Name, action.NewValue)
				}
			}

		}

	}
	return functionStr
}
