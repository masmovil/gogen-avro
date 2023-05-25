package avro

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/masmovil/gogen-avro/v10/container"
	"github.com/masmovil/gogen-avro/v10/test"
)

func TestRoundTrip(t *testing.T) {
	test.RoundTripExactBytes(t,
		func() container.AvroRecord { return &EnumTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeEnumTestRecord(r)
			return &record, err
		})
}

func TestInvalidStringConversion(t *testing.T) {
	enumified, err := NewTestEnumTypeValue("bogus")
	assert.Error(t, err)
	assert.Equal(t, TestEnumType(-1), enumified)
}
