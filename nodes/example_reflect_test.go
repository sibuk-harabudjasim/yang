package nodes_test

import (
	"github.com/c2stack/c2g/node"
	"github.com/c2stack/c2g/nodes"
	"github.com/c2stack/c2g/val"
)

func ExampleReflect_struct() {
	model := `
	    container foo {
			leaf bar {
				type string;
			}
		} 
		list foos {
			key "bar";
			leaf bar {
				type string;
			}	
		}
		leaf-list bleep {
			type string;
		}
		`
	data := struct {
		Foo   *foo
		Foos  map[string]*foo
		Bleep []string
	}{
		Foo: &foo{
			Bar: "x",
		},
		Foos: map[string]*foo{
			"y": &foo{
				Bar: "y",
			},
		},
		Bleep: []string{
			"a", "b",
		},
	}
	sel := exampleSelection(model, nodes.ReflectChild(&data))
	examplePrint(sel)
	// Output:
	// {"foo":{"bar":"x"},"foos":[{"bar":"y"}],"bleep":["a","b"]}
}

func ExampleReflect_extend() {
	model := `
		leaf-list bleep {
			type string;
		}
		leaf len {
			type int32;
			config false;
		}
	`

	data := struct {
		Bleep []string
	}{
		Bleep: []string{
			"a", "b",
		},
	}
	n := &nodes.Extend{
		Base: nodes.ReflectChild(&data),
		OnField: func(parent node.Node, r node.FieldRequest, hnd *node.ValueHandle) error {
			switch r.Meta.Ident() {
			case "len":
				hnd.Val = val.Int32(len(data.Bleep))
			default:
				return parent.Field(r, hnd)
			}
			return nil
		},
	}

	sel := exampleSelection(model, n)

	examplePrint(sel)
	// Output:
	// {"bleep":["a","b"],"len":2}
}

type ExampleApp struct {
	SuggestionBox *ExampleMessage
}

type ExampleMessage struct {
	Id      string
	Message string
}
