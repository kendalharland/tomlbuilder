// Package tomlbuilder is a package for programmatically constructing TOML files
package tomlbuilder

import (
	"bytes"
	"fmt"
)

// TomlBuilder is used to create TOML files.
//
// The TomlBuilder supports creating all of the TOML data types:
// * String
// * Integer
// * Float
// * Boolean
// X Offset Date-time
// X Local Date-time
// X Local Date
// X Local Time
// * Array
// * Table
// * Inline Table
// * Array of Tables
// * Filename Extension
//
// TomlBuilder will additionally produce formatted output for nested data
// types such as array and array of tables.  The default indentation is 2
// spaces, but can be overriden by setting TomlBuilder.IndentSize.
//
// TomlBuilder will not perform validation on your TOML file, which will be
// caught by any correctly implemented TOML parser. For example, attempting to
// define a normal table with the same name as an already established array will
// not cause an error:
//
//     Exmaple invalid TOML:
//
//       [[fruit]]
//       name = "apple"
//
//       [[fruit.variety]]
//       name = "red delicious"
//
//       # This table conflicts with the previous table
//       [fruit.variety]
//       name = "granny smith"
//
type TomlBuilder struct {
	IndentSize int

	indentation string
	buf         *bytes.Buffer
}

// New creates a new TomlBuilder.
func New() *TomlBuilder {
	return &TomlBuilder{
		IndentSize: 2,

		indentation: "",
		buf:         new(bytes.Buffer),
	}
}

// WriteIntArray(key string, array ...int)
// WriteFloatArray(key string, array ...float64)

// WriteNewLine adds a new line to the builder.
func (w *TomlBuilder) WriteNewLine() {
	w.write("\n")
}

// WriteComment adds a comment to the builder.
func (w *TomlBuilder) WriteComment(msg string) {
	w.write("# %v", msg)
}

// WriteString adds a string key-value pair to the builder
func (w *TomlBuilder) WriteString(key string, value string) {
	w.write("%v = \"%v\"\n", key, value)
}

// WriteInt adds an integer key-value pair to the builder
func (w *TomlBuilder) WriteInt(key string, value int) {
	w.write("%v = %v\n", key, value)
}

// WriteFloat adds a float key-value pair to the builder
func (w *TomlBuilder) WriteFloat(key string, value float64) {
	w.write("%v = %v\n", key, value)
}

// WriteStringArray adds a string array key-value pair to the builder
func (w *TomlBuilder) WriteStringArray(key string, array ...string) {
	w.write("%v = [\n", key)
	w.indent()
	for _, val := range array {
		w.write("\"%v\",\n", val)
	}
	w.unindent()
	w.write("]\n")
}

// WriteArrayOfTables adds an array of tables to the builder.  name is the name
// of the array, while the write callback is used to build the contents of the
// array.
//
// Example:
//
//   builder.WriteArrayOfTables("foo", func(b *TomlBuilder) {
//       b.WriteString("bar", "baz")
//       b.WriteInt("bang", 1)
//   })
//   println(builder.String())
//
// Outputs:
//
//   [[foo]]
//   bar = "baz"
//   bang = 1
//
// FIXME: Move this to an actual example.
func (w *TomlBuilder) WriteArrayOfTables(name string,
	write func(*TomlBuilder)) {
	w.write("[[%v]]\n", name)
	write(&TomlBuilder{
		IndentSize: w.IndentSize,

		indentation: w.indentation,
		buf:         w.buf,
	})
}

// String converts the builder's buffer into a string of TOML file contents.
func (w *TomlBuilder) String() string {
	return w.buf.String()
}

func (w *TomlBuilder) write(format string, args ...interface{}) {
	w.buf.Write([]byte(w.indentation + fmt.Sprintf(format, args...)))
}

func (w *TomlBuilder) indent() {
	w.indentation += "  "
}

func (w *TomlBuilder) unindent() {
	if len(w.indentation) == 0 {
		return
	}
	if len(w.indentation) <= w.IndentSize {
		w.indentation = ""
		return
	}

	w.indentation = w.indentation[:len(w.indentation)-w.IndentSize]
}
