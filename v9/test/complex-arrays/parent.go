// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

var _ = fmt.Printf

type Parent struct {
	Children []Child `json:"Children"`
}

const ParentAvroCRC64Fingerprint = "p\x9fn\xf8\x11i\x98\x9b"

func NewParent() Parent {
	r := Parent{}
	r.Children = make([]Child, 0)

	r.Children = make([]Child, 2)
	r.Children[0] = NewChild()
	r.Children[0].Name = "record1"

	r.Children[1] = NewChild()
	r.Children[1].Name = "record2"

	return r
}

func DeserializeParent(r io.Reader) (Parent, error) {
	t := NewParent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeParentFromSchema(r io.Reader, schema string) (Parent, error) {
	t := NewParent()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeParent(r Parent, w io.Writer) error {
	var err error
	err = writeArrayChild(r.Children, w)
	if err != nil {
		return err
	}
	return err
}

func (r Parent) Serialize(w io.Writer) error {
	return writeParent(r, w)
}

func (r Parent) Schema() string {
	return "{\"fields\":[{\"default\":[{\"name\":\"record1\"},{\"name\":\"record2\"}],\"name\":\"Children\",\"type\":{\"items\":{\"fields\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Child\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Parent\",\"type\":\"record\"}"
}

func (r Parent) SchemaName() string {
	return "Parent"
}

func (_ Parent) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ Parent) SetInt(v int32)       { panic("Unsupported operation") }
func (_ Parent) SetLong(v int64)      { panic("Unsupported operation") }
func (_ Parent) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ Parent) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ Parent) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ Parent) SetString(v string)   { panic("Unsupported operation") }
func (_ Parent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Parent) Get(i int) types.Field {
	switch i {
	case 0:
		r.Children = make([]Child, 0)

		return &ArrayChildWrapper{Target: &r.Children}
	}
	panic("Unknown field index")
}

func (r *Parent) SetDefault(i int) {
	switch i {
	case 0:
		r.Children = make([]Child, 2)
		r.Children[0] = NewChild()
		r.Children[0].Name = "record1"

		r.Children[1] = NewChild()
		r.Children[1].Name = "record2"

		return
	}
	panic("Unknown field index")
}

func (r *Parent) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ Parent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ Parent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ Parent) HintSize(int)                     { panic("Unsupported operation") }
func (_ Parent) Finalize()                        {}

func (_ Parent) AvroCRC64Fingerprint() []byte {
	return []byte(ParentAvroCRC64Fingerprint)
}

func (r Parent) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Children"], err = json.Marshal(r.Children)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *Parent) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Children"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Children); err != nil {
			return err
		}
	} else {
		r.Children = make([]Child, 0)

		r.Children = make([]Child, 2)
		r.Children[0] = NewChild()
		r.Children[0].Name = "record1"

		r.Children[1] = NewChild()
		r.Children[1].Name = "record2"

	}
	return nil
}
