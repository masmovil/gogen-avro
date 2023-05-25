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

type UnionNullDatatypeUUIDTypeEnum int

const (
	UnionNullDatatypeUUIDTypeEnumDatatypeUUID UnionNullDatatypeUUIDTypeEnum = 1
)

type UnionNullDatatypeUUID struct {
	Null         *types.NullVal
	DatatypeUUID DatatypeUUID
	UnionType    UnionNullDatatypeUUIDTypeEnum
}

func writeUnionNullDatatypeUUID(r *UnionNullDatatypeUUID, w io.Writer) error {

	if r == nil {
		err := vm.WriteLong(0, w)
		return err
	}

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionNullDatatypeUUIDTypeEnumDatatypeUUID:
		return writeDatatypeUUID(r.DatatypeUUID, w)
	}
	return fmt.Errorf("invalid value for *UnionNullDatatypeUUID")
}

func NewUnionNullDatatypeUUID() *UnionNullDatatypeUUID {
	return &UnionNullDatatypeUUID{}
}

func (r *UnionNullDatatypeUUID) Serialize(w io.Writer) error {
	return writeUnionNullDatatypeUUID(r, w)
}

func DeserializeUnionNullDatatypeUUID(r io.Reader) (*UnionNullDatatypeUUID, error) {
	t := NewUnionNullDatatypeUUID()
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

func DeserializeUnionNullDatatypeUUIDFromSchema(r io.Reader, schema string) (*UnionNullDatatypeUUID, error) {
	t := NewUnionNullDatatypeUUID()
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

func (r *UnionNullDatatypeUUID) Schema() string {
	return "[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"bodyworks.datatype\",\"type\":\"record\"}]"
}

func (_ *UnionNullDatatypeUUID) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetInt(v int32)      { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionNullDatatypeUUID) SetLong(v int64) {

	r.UnionType = (UnionNullDatatypeUUIDTypeEnum)(v)
}

func (r *UnionNullDatatypeUUID) Get(i int) types.Field {

	switch i {
	case 0:
		return r.Null
	case 1:
		r.DatatypeUUID = NewDatatypeUUID()
		return &types.Record{Target: (&r.DatatypeUUID)}
	}
	panic("Unknown field index")
}
func (_ *UnionNullDatatypeUUID) NullField(i int)                  { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) HintSize(i int)                   { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *UnionNullDatatypeUUID) Finalize()                        {}

func (r *UnionNullDatatypeUUID) MarshalJSON() ([]byte, error) {

	if r == nil {
		return []byte("null"), nil
	}

	switch r.UnionType {
	case UnionNullDatatypeUUIDTypeEnumDatatypeUUID:
		return json.Marshal(map[string]interface{}{"bodyworks.datatype.UUID": r.DatatypeUUID})
	}
	return nil, fmt.Errorf("invalid value for *UnionNullDatatypeUUID")
}

func (r *UnionNullDatatypeUUID) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["bodyworks.datatype.UUID"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.DatatypeUUID)
	}
	return fmt.Errorf("invalid value for *UnionNullDatatypeUUID")
}
