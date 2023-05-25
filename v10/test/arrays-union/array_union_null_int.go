// Code generated by github.com/masmovil/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/masmovil/gogen-avro/v10/vm"
	"github.com/masmovil/gogen-avro/v10/vm/types"
)

func writeArrayUnionNullInt(r []*UnionNullInt, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeUnionNullInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type ArrayUnionNullIntWrapper struct {
	Target *[]*UnionNullInt
}

func (_ ArrayUnionNullIntWrapper) SetBoolean(v bool)                { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) SetInt(v int32)                   { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) SetLong(v int64)                  { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) SetFloat(v float32)               { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) SetDouble(v float64)              { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) SetBytes(v []byte)                { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ ArrayUnionNullIntWrapper) Finalize()                        {}
func (_ ArrayUnionNullIntWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
func (r ArrayUnionNullIntWrapper) HintSize(s int) {
	if len(*r.Target) == 0 {
		*r.Target = make([]*UnionNullInt, 0, s)
	}
}
func (r ArrayUnionNullIntWrapper) NullField(i int) {
	(*r.Target)[len(*r.Target)-1] = nil
}

func (r ArrayUnionNullIntWrapper) AppendArray() types.Field {
	var v *UnionNullInt
	v = NewUnionNullInt()

	*r.Target = append(*r.Target, v)

	return (*r.Target)[len(*r.Target)-1]
}
