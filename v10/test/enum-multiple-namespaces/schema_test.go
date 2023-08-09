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
		func() container.AvroRecord { return &ComCompanyTeamSomeType{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeComCompanyTeamSomeType(r)
			return &record, err
		})
}

func TestInvalidStringConversion_TeamSomeEnum(t *testing.T) {
	enumified, err := NewComCompanyTeamSomeEnumValue("bogus")
	assert.Error(t, err)
	assert.Equal(t, ComCompanyTeamSomeEnum(-1), enumified)
}

func TestInvalidStringConversion_SharedSomeEnum(t *testing.T) {
	enumified, err := NewComCompanySharedSomeEnumValue("bogus")
	assert.Error(t, err)
	assert.Equal(t, ComCompanySharedSomeEnum(-1), enumified)
}
