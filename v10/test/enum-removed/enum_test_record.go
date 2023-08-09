// Code generated by github.com/masmovil/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/masmovil/gogen-avro/v10/compiler"
	"github.com/masmovil/gogen-avro/v10/vm"
	"github.com/masmovil/gogen-avro/v10/vm/types"
)

var _ = fmt.Printf

type EnumTestRecord struct {
	StringField string `json:"stringField"`

	EnumField TestEnumType `json:"EnumField"`
}

const EnumTestRecordAvroCRC64Fingerprint = "\xb9\x836\xc5\x18\xd3\xf2\xb3"

func NewEnumTestRecord() EnumTestRecord {
	r := EnumTestRecord{}
	r.EnumField = TestEnumTypeTestSymbol3
	return r
}

func DeserializeEnumTestRecord(r io.Reader) (EnumTestRecord, error) {
	t := NewEnumTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeEnumTestRecordFromSchema(r io.Reader, schema string) (EnumTestRecord, error) {
	t := NewEnumTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeEnumTestRecord(r EnumTestRecord, w io.Writer) error {
	var err error
	err = vm.WriteString(r.StringField, w)
	if err != nil {
		return err
	}
	err = writeTestEnumType(r.EnumField, w)
	if err != nil {
		return err
	}
	return err
}

func (r EnumTestRecord) Serialize(w io.Writer) error {
	return writeEnumTestRecord(r, w)
}

func (r EnumTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"stringField\",\"type\":\"string\"},{\"default\":\"testSymbol3\",\"name\":\"EnumField\",\"type\":{\"doc\":\"Test \\nenum\",\"name\":\"TestEnumType\",\"symbols\":[\"TestSymbol1\",\"testSymbol2\",\"testSymbol3\"],\"type\":\"enum\"}}],\"name\":\"EnumTestRecord\",\"type\":\"record\"}"
}

func (r EnumTestRecord) SchemaName() string {
	return "EnumTestRecord"
}

func (_ EnumTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ EnumTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ EnumTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ EnumTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ EnumTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ EnumTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ EnumTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ EnumTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *EnumTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.String{Target: &r.StringField}

		return w

	case 1:
		w := TestEnumTypeWrapper{Target: &r.EnumField}

		return w

	}
	panic("Unknown field index")
}

func (r *EnumTestRecord) SetDefault(i int) {
	switch i {
	case 1:
		r.EnumField = TestEnumTypeTestSymbol3
		return
	}
	panic("Unknown field index")
}

func (r *EnumTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ EnumTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ EnumTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ EnumTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ EnumTestRecord) Finalize()                        {}

func (_ EnumTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(EnumTestRecordAvroCRC64Fingerprint)
}

func (r EnumTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["stringField"], err = json.Marshal(r.StringField)
	if err != nil {
		return nil, err
	}
	output["EnumField"], err = json.Marshal(r.EnumField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *EnumTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["stringField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.StringField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for stringField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["EnumField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EnumField); err != nil {
			return err
		}
	} else {
		r.EnumField = TestEnumTypeTestSymbol3
	}
	return nil
}
