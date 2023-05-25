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

func NewDatatypeUUIDWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewDatatypeUUID()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type DatatypeUUIDReader struct {
	r io.Reader
	p *vm.Program
}

func NewDatatypeUUIDReader(r io.Reader) (*DatatypeUUIDReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewDatatypeUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &DatatypeUUIDReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r DatatypeUUIDReader) Read() (DatatypeUUID, error) {
	t := NewDatatypeUUID()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}
