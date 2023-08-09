package avro

import (
	"testing"

	"github.com/masmovil/gogen-avro/v10/compiler"

	"github.com/stretchr/testify/assert"
)

func TestLaxNames(t *testing.T) {
	recordFoo := NewRecordFoo()
	recordBar := NewRecordBar()

	// Should throw an error by default
	_, err := compiler.CompileSchemaBytes([]byte(recordFoo.Schema()), []byte(recordBar.Schema()))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Incompatible types by name")

	// With AllowLaxNames() option it should compile schemas successfully
	program, err := compiler.CompileSchemaBytes(
		[]byte(recordFoo.Schema()),
		[]byte(recordBar.Schema()),
		compiler.AllowLaxNames(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, program)
}
