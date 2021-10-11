// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/container"
	"github.com/actgardner/gogen-avro/v9/vm"
)

func NewMapTestRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewMapTestRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type MapTestRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewMapTestRecordReader(r io.Reader) (*MapTestRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewMapTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &MapTestRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r MapTestRecordReader) Read() (MapTestRecord, error) {
	t := NewMapTestRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
