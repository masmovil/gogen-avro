package avro

import (
	"io"
	"testing"

	"github.com/actgardner/gogen-avro/v9/container"
	"github.com/actgardner/gogen-avro/v9/test"
)

import (
	"github.com/actgardner/gogen-avro/v9/compiler"
)

func init() {
	compiler.LoggingEnabled = true
}

func TestRoundTrip(t *testing.T) {
	test.RoundTrip(t,
		func() container.AvroRecord { return &MapTestRecord{} },
		func(r io.Reader) (container.AvroRecord, error) {
			record, err := DeserializeMapTestRecord(r)
			return &record, err
		})
}
