package tomlbuilder_test

import (
	"fmt"

	"github.com/kharland/tomlbuilder"
)

func ExampleAddNewLine() {
	builder := tomlbuilder.New()
	builder.AddNewLine()
	fmt.Println(builder.String())
	// Output:
	//
}

func ExampleAddComment() {
	builder := tomlbuilder.New()
	builder.AddComment("This is a comment")
	fmt.Println(builder.String())
	// Output:
	// # This is a comment
}

func ExampleAddString() {
	builder := tomlbuilder.New()
	builder.AddString("string", "hello")
	fmt.Println(builder.String())
	// Output:
	// string = "hello"
}

func ExampleAddInt() {
	builder := tomlbuilder.New()
	builder.AddInt("int", 23)
	fmt.Println(builder.String())
	// Output:
	// int = 23
}

func ExampleAddBool() {
	builder := tomlbuilder.New()
	builder.AddBool("boolA", true)
	builder.AddBool("boolB", false)
	fmt.Println(builder.String())
	// Output:
	// boolA = true
	// boolB = false
}

func ExampleAddFloat() {
	builder := tomlbuilder.New()
	builder.AddFloat("floatA", 23.45)
	builder.AddFloat("floatB", 8)
	builder.AddFloat("floatC", 0)
	builder.AddFloat("floatD", .004)
	fmt.Println(builder.String())
	// Output:
	// floatA = 23.45
	// floatB = 8.0
	// floatC = 0.0
	// floatD = 0.004
}

func ExampleAddStringArray() {
	builder := tomlbuilder.New()
	builder.AddStringArray("array", "a", "b", "c")
	fmt.Println(builder.String())
	// Output:
	// array = [
	//   "a",
	//   "b",
	//   "c",
	// ]
	//
}

func ExampleAddIntArray() {
	builder := tomlbuilder.New()
	builder.AddIntArray("array", 1, 2, 3)
	fmt.Println(builder.String())
	// Output:
	// array = [
	//   1,
	//   2,
	//   3,
	// ]
	//
}

func ExampleAddFloatArray() {
	builder := tomlbuilder.New()
	builder.AddFloatArray("array", 0, 1.2, .003, 4)
	fmt.Println(builder.String())
	// Output:
	// array = [
	//   0.0,
	//   1.2,
	//   0.003,
	//   4.0,
	// ]
}

func ExampleAddBoolArray() {
	builder := tomlbuilder.New()
	builder.AddBoolArray("array", true, false)
	fmt.Println(builder.String())
	// Output:
	// array = [
	//   true,
	//   false,
	// ]
}

func ExampleAddTable() {
	builder := tomlbuilder.New()
	builder.AddTable("table", func(b *tomlbuilder.TomlBuilder) {
		b.AddString("foo", "bar")
		b.AddFloat("baz", 123.456)
	})
	fmt.Println(builder.String())
	// Output:
	// [table]
	// foo = "bar"
	// baz = 123.456
}

func ExampleAddArrayOfTables() {
	builder := tomlbuilder.New()
	builder.AddArrayOfTables("tablesArray", func(b *tomlbuilder.TomlBuilder) {
		b.AddString("foo", "bar")
		b.AddFloat("baz", 123.456)
	})
	fmt.Println(builder.String())
	// Output:
	// [[tablesArray]]
	// foo = "bar"
	// baz = 123.456
}
