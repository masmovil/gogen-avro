// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     example.avsc
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

type DemoSchema struct {
	IntField int32 `json:"IntField"`

	DoubleField float64 `json:"DoubleField"`

	StringField string `json:"StringField"`

	BoolField bool `json:"BoolField"`

	BytesField Bytes `json:"BytesField"`
}

const DemoSchemaAvroCRC64Fingerprint = "\xc4V\xa9\x04ʛf\xad"

func NewDemoSchema() DemoSchema {
	r := DemoSchema{}
	return r
}

func DeserializeDemoSchema(r io.Reader) (DemoSchema, error) {
	t := NewDemoSchema()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeDemoSchemaFromSchema(r io.Reader, schema string) (DemoSchema, error) {
	t := NewDemoSchema()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeDemoSchema(r DemoSchema, w io.Writer) error {
	var err error
	err = vm.WriteInt(r.IntField, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.DoubleField, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.StringField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.BoolField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.BytesField, w)
	if err != nil {
		return err
	}
	return err
}

func (r DemoSchema) Serialize(w io.Writer) error {
	return writeDemoSchema(r, w)
}

func (r DemoSchema) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"DoubleField\",\"type\":\"double\"},{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"DemoSchema\",\"type\":\"record\"}"
}

func (r DemoSchema) SchemaName() string {
	return "DemoSchema"
}

func (_ DemoSchema) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ DemoSchema) SetInt(v int32)       { panic("Unsupported operation") }
func (_ DemoSchema) SetLong(v int64)      { panic("Unsupported operation") }
func (_ DemoSchema) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ DemoSchema) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ DemoSchema) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ DemoSchema) SetString(v string)   { panic("Unsupported operation") }
func (_ DemoSchema) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *DemoSchema) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.Int{Target: &r.IntField}
	case 1:
		return &types.Double{Target: &r.DoubleField}
	case 2:
		return &types.String{Target: &r.StringField}
	case 3:
		return &types.Boolean{Target: &r.BoolField}
	case 4:
		return &BytesWrapper{Target: &r.BytesField}
	}
	panic("Unknown field index")
}

func (r *DemoSchema) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *DemoSchema) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ DemoSchema) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ DemoSchema) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ DemoSchema) HintSize(int)                     { panic("Unsupported operation") }
func (_ DemoSchema) Finalize()                        {}

func (_ DemoSchema) AvroCRC64Fingerprint() []byte {
	return []byte(DemoSchemaAvroCRC64Fingerprint)
}

func (r DemoSchema) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["IntField"], err = json.Marshal(r.IntField)
	if err != nil {
		return nil, err
	}
	output["DoubleField"], err = json.Marshal(r.DoubleField)
	if err != nil {
		return nil, err
	}
	output["StringField"], err = json.Marshal(r.StringField)
	if err != nil {
		return nil, err
	}
	output["BoolField"], err = json.Marshal(r.BoolField)
	if err != nil {
		return nil, err
	}
	output["BytesField"], err = json.Marshal(r.BytesField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *DemoSchema) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["IntField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.IntField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for IntField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["DoubleField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.DoubleField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for DoubleField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["StringField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StringField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for StringField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BoolField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BoolField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BoolField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["BytesField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.BytesField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for BytesField")
	}
	return nil
}
