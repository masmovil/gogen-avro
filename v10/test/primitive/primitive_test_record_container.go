// Code generated by github.com/masmovil/gogen-avro/v10. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/masmovil/gogen-avro/v10/compiler"
	"github.com/masmovil/gogen-avro/v10/container"
	"github.com/masmovil/gogen-avro/v10/vm"
)

func NewPrimitiveTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewPrimitiveTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type PrimitiveTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewPrimitiveTestRecordReader(r io.Reader) (*PrimitiveTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewPrimitiveTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &PrimitiveTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r PrimitiveTestRecordReader) Read() (PrimitiveTestRecord, error) {
	t := NewPrimitiveTestRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
