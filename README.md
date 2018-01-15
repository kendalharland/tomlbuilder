# tomlbuilder
A go package for programmatically building TOML files

## Usage
A TomlBuilder object is used to buffer the contents of a TOML file. Values
written to the builder will appear from top to bottom in the generated TOML 
file.

Example

```go
builder := tomlbuilder.New()
builder.AddString("title", "shopping-list")
builder.AddString("day", "sunday")
builder.AddFloat("budget", 50)
builder.AddTable("groceries", func(b *tomlbuilder.TomlBuilder) {
    b.AddString("candy", "m&ms")
    b.AddInt("apples", 3)
})
fmt.Println(builder.String())
    
// Prints:
// title = "shopping-list"
// day = "sunday"
// budget = 50.0
// 
// [groceries]
// candy = "m&ms"
// apples = 3
```