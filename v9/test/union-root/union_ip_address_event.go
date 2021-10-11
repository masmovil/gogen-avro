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

type UnionIp_addressEventTypeEnum int

const (
	UnionIp_addressEventTypeEnumIp_address UnionIp_addressEventTypeEnum = 0

	UnionIp_addressEventTypeEnumEvent UnionIp_addressEventTypeEnum = 1
)

type UnionIp_addressEvent struct {
	Ip_address Ip_address
	Event      Event
	UnionType  UnionIp_addressEventTypeEnum
}

func writeUnionIp_addressEvent(r UnionIp_addressEvent, w io.Writer) error {

	err := vm.WriteLong(int64(r.UnionType), w)
	if err != nil {
		return err
	}
	switch r.UnionType {
	case UnionIp_addressEventTypeEnumIp_address:
		return writeIp_address(r.Ip_address, w)
	case UnionIp_addressEventTypeEnumEvent:
		return writeEvent(r.Event, w)
	}
	return fmt.Errorf("invalid value for UnionIp_addressEvent")
}

func NewUnionIp_addressEvent() UnionIp_addressEvent {
	return UnionIp_addressEvent{}
}

func (r UnionIp_addressEvent) Serialize(w io.Writer) error {
	return writeUnionIp_addressEvent(r, w)
}

func DeserializeUnionIp_addressEvent(r io.Reader) (UnionIp_addressEvent, error) {
	t := NewUnionIp_addressEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func DeserializeUnionIp_addressEventFromSchema(r io.Reader, schema string) (UnionIp_addressEvent, error) {
	t := NewUnionIp_addressEvent()
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)

	if err != nil {
		return t, err
	}
	return t, err
}

func (r UnionIp_addressEvent) Schema() string {
	return "[{\"key\":\"value\",\"metadata\":{\"a\":\"b\",\"c\":123},\"name\":\"ip_address\",\"size\":16,\"type\":\"fixed\"},{\"fields\":[{\"doc\":\"Unique ID for this event.\",\"metadata\":{\"key\":\"field1\"},\"name\":\"id\",\"type\":\"string\"},{\"doc\":\"Start IP of this observation's IP range.\",\"metadata\":{\"key\":\"field2\"},\"name\":\"start_ip\",\"type\":\"ip_address\"},{\"doc\":\"End IP of this observation's IP range.\",\"metadata\":{\"key\":\"field3\"},\"name\":\"end_ip\",\"type\":\"ip_address\"}],\"metadata\":{\"key\":\"value\"},\"name\":\"event\",\"subject\":\"event\",\"type\":\"record\"}]"
}

func (_ UnionIp_addressEvent) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) SetInt(v int32)      { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) SetDouble(v float64) { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) SetBytes(v []byte)   { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) SetString(v string)  { panic("Unsupported operation") }

func (r *UnionIp_addressEvent) SetLong(v int64) {

	r.UnionType = (UnionIp_addressEventTypeEnum)(v)
}

func (r *UnionIp_addressEvent) Get(i int) types.Field {

	switch i {
	case 0:
		return &Ip_addressWrapper{Target: (&r.Ip_address)}
	case 1:
		r.Event = NewEvent()
		return &types.Record{Target: (&r.Event)}
	}
	panic("Unknown field index")
}
func (_ UnionIp_addressEvent) NullField(i int)                  { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) HintSize(i int)                   { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) SetDefault(i int)                 { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ UnionIp_addressEvent) Finalize()                        {}

func (r UnionIp_addressEvent) MarshalJSON() ([]byte, error) {

	switch r.UnionType {
	case UnionIp_addressEventTypeEnumIp_address:
		return json.Marshal(map[string]interface{}{"ip_address": r.Ip_address})
	case UnionIp_addressEventTypeEnumEvent:
		return json.Marshal(map[string]interface{}{"event": r.Event})
	}
	return nil, fmt.Errorf("invalid value for UnionIp_addressEvent")
}

func (r *UnionIp_addressEvent) UnmarshalJSON(data []byte) error {

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	if len(fields) > 1 {
		return fmt.Errorf("more than one type supplied for union")
	}
	if value, ok := fields["ip_address"]; ok {
		r.UnionType = 0
		return json.Unmarshal([]byte(value), &r.Ip_address)
	}
	if value, ok := fields["event"]; ok {
		r.UnionType = 1
		return json.Unmarshal([]byte(value), &r.Event)
	}
	return fmt.Errorf("invalid value for UnionIp_addressEvent")
}
