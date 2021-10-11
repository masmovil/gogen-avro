// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
	"io"
)

func writeMapDouble(r map[string]float64, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for k, e := range r {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = vm.WriteDouble(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapDoubleWrapper struct {
	Target *map[string]float64
	keys   []string
	values []float64
}

func (_ *MapDoubleWrapper) SetBoolean(v bool)     { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) SetInt(v int32)        { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) SetLong(v int64)       { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) SetFloat(v float32)    { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) SetDouble(v float64)   { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) SetBytes(v []byte)     { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) SetString(v string)    { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) SetUnionElem(v int64)  { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapDoubleWrapper) SetDefault(i int)      { panic("Unsupported operation") }

func (r *MapDoubleWrapper) HintSize(s int) {
	if r.keys == nil {
		r.keys = make([]string, 0, s)
		r.values = make([]float64, 0, s)
	}
}

func (r *MapDoubleWrapper) NullField(_ int) {
	panic("Unsupported operation")
}

func (r *MapDoubleWrapper) Finalize() {
	for i := range r.keys {
		(*r.Target)[r.keys[i]] = r.values[i]
	}
}

func (r *MapDoubleWrapper) AppendMap(key string) types.Field {
	r.keys = append(r.keys, key)
	var v float64
	r.values = append(r.values, v)
	return &types.Double{Target: &r.values[len(r.values)-1]}
}

func (_ *MapDoubleWrapper) AppendArray() types.Field { panic("Unsupported operation") }
