// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

func writeArrayChild(r []Child, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeChild(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayChildWrapper struct {
	Target *[]Child
}

func (_ *ArrayChildWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayChildWrapper) Finalize()                        {}
func (_ *ArrayChildWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r *ArrayChildWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]Child, 0, s)
	}
}
func (r *ArrayChildWrapper) NullField(i int) {
	panic("Unsupported operation")
}

func (r *ArrayChildWrapper) AppendArray() types.Field {
	var v Child
	v = NewChild()

	*r.Target = append(*r.Target, v)
	return &types.Record{Target: &(*r.Target)[len(*r.Target)-1]}
}
