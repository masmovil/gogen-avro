// Code generated by github.com/masmovil/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     evolution.avsc
 */
package avro

import (
	"encoding/json"
	"io"

	"github.com/masmovil/gogen-avro/v10/util"
	"github.com/masmovil/gogen-avro/v10/vm/types"
)

func writeIPAddress(r IPAddress, w io.Writer) error {
	_, err := w.Write(r[:])
	return err
}

type IPAddressWrapper struct {
	Target *IPAddress
}

type IPAddress [16]byte

func (b *IPAddress) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	codepoints := util.DecodeByteString(s)
	copy((*b)[:], codepoints)
	return nil
}

func (b IPAddress) MarshalJSON() ([]byte, error) {
	return []byte(util.EncodeByteString(b[:])), nil
}

func (_ IPAddressWrapper) SetBoolean(v bool)   { panic("Unsupported operation") }
func (_ IPAddressWrapper) SetInt(v int32)      { panic("Unsupported operation") }
func (_ IPAddressWrapper) SetLong(v int64)     { panic("Unsupported operation") }
func (_ IPAddressWrapper) SetFloat(v float32)  { panic("Unsupported operation") }
func (_ IPAddressWrapper) SetDouble(v float64) { panic("Unsupported operation") }
func (r IPAddressWrapper) SetBytes(v []byte) {
	copy((*r.Target)[:], v)
}
func (_ IPAddressWrapper) SetString(v string)               { panic("Unsupported operation") }
func (_ IPAddressWrapper) SetUnionElem(v int64)             { panic("Unsupported operation") }
func (_ IPAddressWrapper) Get(i int) types.Field            { panic("Unsupported operation") }
func (_ IPAddressWrapper) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ IPAddressWrapper) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ IPAddressWrapper) NullField(int)                    { panic("Unsupported operation") }
func (_ IPAddressWrapper) HintSize(int)                     { panic("Unsupported operation") }
func (_ IPAddressWrapper) Finalize()                        {}
func (_ IPAddressWrapper) SetDefault(i int)                 { panic("Unsupported operation") }
