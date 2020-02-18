package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestCodegenComment(t *testing.T) {
	tt := []struct {
		in             []string
		sourcesComment bool
		out            string
	}{
		{[]string{"file_1.avsc", "file_2.avsc"}, true, `// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     file_1.avsc
 *     file_2.avsc
 */`},
		{[]string{"file_1.avsc"}, true, `// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     file_1.avsc
 */`},
		{[]string{"file_1.avsc", "file_2.avsc"}, false, `// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.`},
	}

	for _, tc := range tt {
		assert.Equal(t, tc.out, codegenComment(config{files: tc.in, sourcesComment: tc.sourcesComment}))
	}
}
