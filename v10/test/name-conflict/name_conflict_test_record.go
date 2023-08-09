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

type NameConflictTestRecord struct {
	Field_Schema bool `json:"Schema"`

	Field_Serialize bool `json:"Serialize"`

	Field_SchemaName bool `json:"SchemaName"`

	Field_MarshalJSON bool `json:"MarshalJSON"`

	Field_UnmarshalJSON bool `json:"UnmarshalJSON"`

	Field_AvroCRC64Fingerprint bool `json:"AvroCRC64Fingerprint"`

	Field_SetBoolean bool `json:"SetBoolean"`

	Field_SetInt bool `json:"SetInt"`

	Field_SetLong bool `json:"SetLong"`

	Field_SetFloat bool `json:"SetFloat"`

	Field_SetDouble bool `json:"SetDouble"`

	Field_SetBytes bool `json:"SetBytes"`

	Field_SetString bool `json:"SetString"`

	Field_Get bool `json:"Get"`

	Field_SetDefault bool `json:"SetDefault"`

	Field_AppendMap bool `json:"AppendMap"`

	Field_AppendArray bool `json:"AppendArray"`

	Field_NullField bool `json:"NullField"`

	Field_Finalize bool `json:"Finalize"`
}

const NameConflictTestRecordAvroCRC64Fingerprint = "\xcesIh9\x9f&\b"

func NewNameConflictTestRecord() NameConflictTestRecord {
	r := NameConflictTestRecord{}
	return r
}

func DeserializeNameConflictTestRecord(r io.Reader) (NameConflictTestRecord, error) {
	t := NewNameConflictTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeNameConflictTestRecordFromSchema(r io.Reader, schema string) (NameConflictTestRecord, error) {
	t := NewNameConflictTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeNameConflictTestRecord(r NameConflictTestRecord, w io.Writer) error {
	var err error
	err = vm.WriteBool(r.Field_Schema, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_Serialize, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SchemaName, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_MarshalJSON, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_UnmarshalJSON, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_AvroCRC64Fingerprint, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetBoolean, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetInt, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetLong, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetFloat, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetDouble, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetBytes, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetString, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_Get, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_SetDefault, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_AppendMap, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_AppendArray, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_NullField, w)
	if err != nil {
		return err
	}
	err = vm.WriteBool(r.Field_Finalize, w)
	if err != nil {
		return err
	}
	return err
}

func (r NameConflictTestRecord) Serialize(w io.Writer) error {
	return writeNameConflictTestRecord(r, w)
}

func (r NameConflictTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"Schema\",\"type\":\"boolean\"},{\"name\":\"Serialize\",\"type\":\"boolean\"},{\"name\":\"SchemaName\",\"type\":\"boolean\"},{\"name\":\"MarshalJSON\",\"type\":\"boolean\"},{\"name\":\"UnmarshalJSON\",\"type\":\"boolean\"},{\"name\":\"AvroCRC64Fingerprint\",\"type\":\"boolean\"},{\"name\":\"SetBoolean\",\"type\":\"boolean\"},{\"name\":\"SetInt\",\"type\":\"boolean\"},{\"name\":\"SetLong\",\"type\":\"boolean\"},{\"name\":\"SetFloat\",\"type\":\"boolean\"},{\"name\":\"SetDouble\",\"type\":\"boolean\"},{\"name\":\"SetBytes\",\"type\":\"boolean\"},{\"name\":\"SetString\",\"type\":\"boolean\"},{\"name\":\"Get\",\"type\":\"boolean\"},{\"name\":\"SetDefault\",\"type\":\"boolean\"},{\"name\":\"AppendMap\",\"type\":\"boolean\"},{\"name\":\"AppendArray\",\"type\":\"boolean\"},{\"name\":\"NullField\",\"type\":\"boolean\"},{\"name\":\"Finalize\",\"type\":\"boolean\"}],\"name\":\"NameConflictTestRecord\",\"type\":\"record\"}"
}

func (r NameConflictTestRecord) SchemaName() string {
	return "NameConflictTestRecord"
}

func (_ NameConflictTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ NameConflictTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ NameConflictTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ NameConflictTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ NameConflictTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ NameConflictTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ NameConflictTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ NameConflictTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NameConflictTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		w := types.Boolean{Target: &r.Field_Schema}

		return w

	case 1:
		w := types.Boolean{Target: &r.Field_Serialize}

		return w

	case 2:
		w := types.Boolean{Target: &r.Field_SchemaName}

		return w

	case 3:
		w := types.Boolean{Target: &r.Field_MarshalJSON}

		return w

	case 4:
		w := types.Boolean{Target: &r.Field_UnmarshalJSON}

		return w

	case 5:
		w := types.Boolean{Target: &r.Field_AvroCRC64Fingerprint}

		return w

	case 6:
		w := types.Boolean{Target: &r.Field_SetBoolean}

		return w

	case 7:
		w := types.Boolean{Target: &r.Field_SetInt}

		return w

	case 8:
		w := types.Boolean{Target: &r.Field_SetLong}

		return w

	case 9:
		w := types.Boolean{Target: &r.Field_SetFloat}

		return w

	case 10:
		w := types.Boolean{Target: &r.Field_SetDouble}

		return w

	case 11:
		w := types.Boolean{Target: &r.Field_SetBytes}

		return w

	case 12:
		w := types.Boolean{Target: &r.Field_SetString}

		return w

	case 13:
		w := types.Boolean{Target: &r.Field_Get}

		return w

	case 14:
		w := types.Boolean{Target: &r.Field_SetDefault}

		return w

	case 15:
		w := types.Boolean{Target: &r.Field_AppendMap}

		return w

	case 16:
		w := types.Boolean{Target: &r.Field_AppendArray}

		return w

	case 17:
		w := types.Boolean{Target: &r.Field_NullField}

		return w

	case 18:
		w := types.Boolean{Target: &r.Field_Finalize}

		return w

	}
	panic("Unknown field index")
}

func (r *NameConflictTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *NameConflictTestRecord) NullField(i int) {
	switch i {
	}
	panic("Not a nullable field index")
}

func (_ NameConflictTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ NameConflictTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ NameConflictTestRecord) HintSize(int)                     { panic("Unsupported operation") }
func (_ NameConflictTestRecord) Finalize()                        {}

func (_ NameConflictTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(NameConflictTestRecordAvroCRC64Fingerprint)
}

func (r NameConflictTestRecord) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["Schema"], err = json.Marshal(r.Field_Schema)
	if err != nil {
		return nil, err
	}
	output["Serialize"], err = json.Marshal(r.Field_Serialize)
	if err != nil {
		return nil, err
	}
	output["SchemaName"], err = json.Marshal(r.Field_SchemaName)
	if err != nil {
		return nil, err
	}
	output["MarshalJSON"], err = json.Marshal(r.Field_MarshalJSON)
	if err != nil {
		return nil, err
	}
	output["UnmarshalJSON"], err = json.Marshal(r.Field_UnmarshalJSON)
	if err != nil {
		return nil, err
	}
	output["AvroCRC64Fingerprint"], err = json.Marshal(r.Field_AvroCRC64Fingerprint)
	if err != nil {
		return nil, err
	}
	output["SetBoolean"], err = json.Marshal(r.Field_SetBoolean)
	if err != nil {
		return nil, err
	}
	output["SetInt"], err = json.Marshal(r.Field_SetInt)
	if err != nil {
		return nil, err
	}
	output["SetLong"], err = json.Marshal(r.Field_SetLong)
	if err != nil {
		return nil, err
	}
	output["SetFloat"], err = json.Marshal(r.Field_SetFloat)
	if err != nil {
		return nil, err
	}
	output["SetDouble"], err = json.Marshal(r.Field_SetDouble)
	if err != nil {
		return nil, err
	}
	output["SetBytes"], err = json.Marshal(r.Field_SetBytes)
	if err != nil {
		return nil, err
	}
	output["SetString"], err = json.Marshal(r.Field_SetString)
	if err != nil {
		return nil, err
	}
	output["Get"], err = json.Marshal(r.Field_Get)
	if err != nil {
		return nil, err
	}
	output["SetDefault"], err = json.Marshal(r.Field_SetDefault)
	if err != nil {
		return nil, err
	}
	output["AppendMap"], err = json.Marshal(r.Field_AppendMap)
	if err != nil {
		return nil, err
	}
	output["AppendArray"], err = json.Marshal(r.Field_AppendArray)
	if err != nil {
		return nil, err
	}
	output["NullField"], err = json.Marshal(r.Field_NullField)
	if err != nil {
		return nil, err
	}
	output["Finalize"], err = json.Marshal(r.Field_Finalize)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *NameConflictTestRecord) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["Schema"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_Schema); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Schema")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Serialize"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_Serialize); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Serialize")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SchemaName"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SchemaName); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SchemaName")
	}
	val = func() json.RawMessage {
		if v, ok := fields["MarshalJSON"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_MarshalJSON); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for MarshalJSON")
	}
	val = func() json.RawMessage {
		if v, ok := fields["UnmarshalJSON"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_UnmarshalJSON); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for UnmarshalJSON")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AvroCRC64Fingerprint"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_AvroCRC64Fingerprint); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AvroCRC64Fingerprint")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SetBoolean"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SetBoolean); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SetBoolean")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SetInt"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SetInt); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SetInt")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SetLong"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SetLong); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SetLong")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SetFloat"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SetFloat); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SetFloat")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SetDouble"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SetDouble); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SetDouble")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SetBytes"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SetBytes); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SetBytes")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SetString"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SetString); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SetString")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Get"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_Get); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Get")
	}
	val = func() json.RawMessage {
		if v, ok := fields["SetDefault"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_SetDefault); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for SetDefault")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AppendMap"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_AppendMap); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AppendMap")
	}
	val = func() json.RawMessage {
		if v, ok := fields["AppendArray"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_AppendArray); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for AppendArray")
	}
	val = func() json.RawMessage {
		if v, ok := fields["NullField"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_NullField); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for NullField")
	}
	val = func() json.RawMessage {
		if v, ok := fields["Finalize"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Field_Finalize); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for Finalize")
	}
	return nil
}
