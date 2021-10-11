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

type UnionNullBodyworksTraceTypeEnum int

const (
	UnionNullBodyworksTraceTypeEnumBodyworksTrace UnionNullBodyworksTraceTypeEnum = 1
)

type UnionNullBodyworksTrace struct {
	Null           *types.NullVal
	BodyworksTrace BodyworksTrace
	UnionType      UnionNullBodyworksTraceTypeEnum
}

func writeUnionNullBodyworksTrace(r *UnionNullBodyworksTrace, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullBodyworksTraceTypeEnumBodyworksTrace:
		return writeBodyworksTrace(r.BodyworksTrace, w)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksTrace")
}

func NewUnionNullBodyworksTrace() *UnionNullBodyworksTrace {
	return &UnionNullBodyworksTrace{}
}

func (r *UnionNullBodyworksTrace) Serialize(w io.Writer) error {
	return writeUnionNullBodyworksTrace(r, w)
}

func DeserializeUnionNullBodyworksTrace(r io.Reader) (*UnionNullBodyworksTrace, error) {
	t := NewUnionNullBodyworksTrace()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionNullBodyworksTraceFromSchema(r io.Reader, schema string) (*UnionNullBodyworksTrace, error) {
	t := NewUnionNullBodyworksTrace()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r *UnionNullBodyworksTrace) Schema() string {
	return "[\"null\",{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"Trace\",\"type\":\"record\"}]"
}

func (_ *UnionNullBodyworksTrace) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullBodyworksTrace) SetLong(v int64) {

	r.UnionType = (UnionNullBodyworksTraceTypeEnum)(v)
}

func (r *UnionNullBodyworksTrace) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.BodyworksTrace = NewBodyworksTrace()
		return &types.Record{Target: (&r.BodyworksTrace)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullBodyworksTrace) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullBodyworksTrace) Finalize()                        {}

func (r *UnionNullBodyworksTrace) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullBodyworksTraceTypeEnumBodyworksTrace:
		return json.Marshal(map[string]interface{}{"bodyworks.Trace": r.BodyworksTrace})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullBodyworksTrace")
}

func (r *UnionNullBodyworksTrace) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["bodyworks.Trace"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.BodyworksTrace)
	}
	return fmt.Errorf("invalid value for *UnionNullBodyworksTrace")
}
