// Package tomlbuilder is a package for programmatically building TOML files
package tomlbuilder

import (
	"bytes"
	"fmt"
	"strconv"
)

// TomlBuilder is used to create TOML files.
//
// The TomlBuilder supports creating all of the TOML data types:
// * String
// * Integer
// * Float
// * Boolean
// [TODO] Offset Date-time
// [TODO] Local Date-time
// [TODO] Local Date
// [TODO] Local Time
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
//       # This table conflicts with the previous table array
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

// AddNewLine adds a new line to the builder.
func (w *TomlBuilder) AddNewLine() {
	w.write("\n")
}

// AddComment adds a comment to the builder.
func (w *TomlBuilder) AddComment(msg string) {
	w.write("# %v", msg)
}

// AddString adds a string key-value pair to the builder.
func (w *TomlBuilder) AddString(key string, value string) {
	w.write("%v = \"%v\"\n", key, value)
}

// AddInt adds an integer key-value pair to the builder.
func (w *TomlBuilder) AddInt(key string, value int) {
	w.write("%v = %v\n", key, value)
}

// AddFloat adds a float key-value pair to the builder.
func (w *TomlBuilder) AddFloat(key string, value float64) {
	w.write("%v = %v\n", key, formatFloat(value))
}

func (w *TomlBuilder) AddBool(key string, value bool) {
	w.write("%v = %v\n", key, strconv.FormatBool(value))
}

// AddStringArray adds an array of strings to the builder.
func (w *TomlBuilder) AddStringArray(key string, array ...string) {
	vals := make([]string, len(array))
	for i, val := range array {
		vals[i] = fmt.Sprintf("\"%v\"", string(val))
	}
	w.addArray(key, vals)
}

// AddIntArray adds an array of ints to the builder.
func (w *TomlBuilder) AddIntArray(key string, array ...int) {
	vals := make([]string, len(array))
	for i, val := range array {
		vals[i] = fmt.Sprintf("%d", val)
	}
	w.addArray(key, vals)
}

// AddFloatArray adds an array of floats to the builder.
func (w *TomlBuilder) AddFloatArray(key string, array ...float64) {
	vals := make([]string, len(array))
	for i, val := range array {
		vals[i] = formatFloat(val)
	}
	w.addArray(key, vals)
}

func (w *TomlBuilder) AddBoolArray(key string, array ...bool) {
	vals := make([]string, len(array))
	for i, val := range array {
		vals[i] = strconv.FormatBool(val)
	}
	w.addArray(key, vals)
}

func (w *TomlBuilder) addArray(key string, array []string) {
	w.write("%v = [\n", key)
	w.indent()
	for _, val := range array {
		w.write("%v,\n", val)
	}
	w.unindent()
	w.write("]\n")
}

// AddFloatArray(key string, array ...float64)

// AddTable adds a table to the builder.
func (w *TomlBuilder) AddTable(name string, write func(*TomlBuilder)) {
	w.write("[%v]\n", name)
	write(&TomlBuilder{
		IndentSize: w.IndentSize,

		indentation: w.indentation,
		buf:         w.buf,
	})
}

// AddArrayOfTables adds an array of tables to the builder.  name is the name of
// the array.  write is the callback used to build the contents of the array.
func (w *TomlBuilder) AddArrayOfTables(name string, write func(*TomlBuilder)) {
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

func formatFloat(val float64) string {
	if val == float64(int64(val)) {
		return fmt.Sprintf("%v.0", val)
	}
	return fmt.Sprintf("%v", strconv.FormatFloat(val, 'f', -1, 64))
}
