// Code generated by github.com/masmovil/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
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

type PrimitiveTestRecord struct {
	NewLongField int64 `json:"NewLongField"`

	NewStringField string `json:"NewStringField"`

	NewFloatField float32 `json:"NewFloatField"`

	NewBytesField Bytes `json:"NewBytesField"`

	NewDoubleField float64 `json:"NewDoubleField"`

	NewIntField int32 `json:"NewIntField"`

	NewBoolField bool `json:"NewBoolField"`
}

const PrimitiveTestRecordAvroCRC64Fingerprint = "d\xa9\x9e\xb8\x86&\xa6\x17"

func NewPrimitiveTestRecord() PrimitiveTestRecord {
	r := PrimitiveTestRecord{}
	r.NewLongField = 1234
	r.NewStringField = "testdefault"
	r.NewFloatField = 12.34
	r.NewBytesField = []byte("\x00\xff\x01\xfe")
	r.NewDoubleField = 56.78
	r.NewIntField = 5678
	r.NewBoolField = true
	return r
}

func DeserializePrimitiveTestRecord(r io.Reader) (PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializePrimitiveTestRecordFromSchema(r io.Reader, schema string) (PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writePrimitiveTestRecord(r PrimitiveTestRecord, w io.Writer) error {
	var err error
	err = vm.WriteLong(r.NewLongField, w)
	if err != nil {
		return err
	}
	err = vm.WriteString(r.NewStringField, w)
	if err != nil {
		return err
	}
	err = vm.WriteFloat(r.NewFloatField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBytes(r.NewBytesField, w)
	if err != nil {
		return err
	}
	err = vm.WriteDouble(r.NewDoubleField, w)
	if err != nil {
		return err
	}
	err = vm.WriteInt(r.NewIntField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.NewBoolField, w)
	if err != nil {
		return err
	}
	return err
}

func (r PrimitiveTestRecord) Serialize(w io.Writer) error {
	return writePrimitiveTestRecord(r, w)
}

func (r PrimitiveTestRecord) Schema() string {
	return "{\"fields\":[{\"default\":1234,\"name\":\"NewLongField\",\"type\":\"long\"},{\"default\":\"testdefault\",\"name\":\"NewStringField\",\"type\":\"string\"},{\"default\":12.34,\"name\":\"NewFloatField\",\"type\":\"float\"},{\"default\":\"\\u0000ÿ\\u0001þ\",\"name\":\"NewBytesField\",\"type\":\"bytes\"},{\"default\":56.78,\"name\":\"NewDoubleField\",\"type\":\"double\"},{\"default\":5678,\"name\":\"NewIntField\",\"type\":\"int\"},{\"default\":true,\"name\":\"NewBoolField\",\"type\":\"boolean\"}],\"name\":\"PrimitiveTestRecord\",\"type\":\"record\"}"
}

func (r PrimitiveTestRecord) SchemaName() string {
	return "PrimitiveTestRecord"
}

func (_ PrimitiveTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PrimitiveTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Long{Target: &r.NewLongField}

		return w

	case 1:
		w := types.String{Target: &r.NewStringField}

		return w

	case 2:
		w := types.Float{Target: &r.NewFloatField}

		return w

	case 3:
		w := BytesWrapper{Target: &r.NewBytesField}

		return w

	case 4:
		w := types.Double{Target: &r.NewDoubleField}

		return w

	case 5:
		w := types.Int{Target: &r.NewIntField}

		return w

	case 6:
		w := types.Boolean{Target: &r.NewBoolField}

		return w

	}
	panic("Unknown field index")
}

func (r *PrimitiveTestRecord) SetDefault(i int) {
	switch i {
	case 0:
		r.NewLongField = 1234
		return
	case 1:
		r.NewStringField = "testdefault"
		return
	case 2:
		r.NewFloatField = 12.34
		return
	case 3:
		r.NewBytesField = []byte("\x00\xff\x01\xfe")
		return
	case 4:
		r.NewDoubleField = 56.78
		return
	case 5:
		r.NewIntField = 5678
		return
	case 6:
		r.NewBoolField = true
		return
	}
	panic("Unknown field index")
}

func (r *PrimitiveTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ PrimitiveTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ PrimitiveTestRecord) Finalize()                        {}

func (_ PrimitiveTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(PrimitiveTestRecordAvroCRC64Fingerprint)
}

func (r PrimitiveTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["NewLongField"], err = json.Marshal(r.NewLongField)
	if err != nil {
		return nil, err
	}
	output["NewStringField"], err = json.Marshal(r.NewStringField)
	if err != nil {
		return nil, err
	}
	output["NewFloatField"], err = json.Marshal(r.NewFloatField)
	if err != nil {
		return nil, err
	}
	output["NewBytesField"], err = json.Marshal(r.NewBytesField)
	if err != nil {
		return nil, err
	}
	output["NewDoubleField"], err = json.Marshal(r.NewDoubleField)
	if err != nil {
		return nil, err
	}
	output["NewIntField"], err = json.Marshal(r.NewIntField)
	if err != nil {
		return nil, err
	}
	output["NewBoolField"], err = json.Marshal(r.NewBoolField)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *PrimitiveTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["NewLongField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewLongField); err != nil {
			return err
		}
	} else {
		r.NewLongField = 1234
	}
	val = func() json.RawMessage {
		if v, ok := fields["NewStringField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewStringField); err != nil {
			return err
		}
	} else {
		r.NewStringField = "testdefault"
	}
	val = func() json.RawMessage {
		if v, ok := fields["NewFloatField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewFloatField); err != nil {
			return err
		}
	} else {
		r.NewFloatField = 12.34
	}
	val = func() json.RawMessage {
		if v, ok := fields["NewBytesField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewBytesField); err != nil {
			return err
		}
	} else {
		r.NewBytesField = []byte("\x00\xff\x01\xfe")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NewDoubleField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewDoubleField); err != nil {
			return err
		}
	} else {
		r.NewDoubleField = 56.78
	}
	val = func() json.RawMessage {
		if v, ok := fields["NewIntField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewIntField); err != nil {
			return err
		}
	} else {
		r.NewIntField = 5678
	}
	val = func() json.RawMessage {
		if v, ok := fields["NewBoolField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.NewBoolField); err != nil {
			return err
		}
	} else {
		r.NewBoolField = true
	}
	return nil
}
