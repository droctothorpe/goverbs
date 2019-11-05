package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

func main() {
	i := 15
	f := 123.456
	s := "cafe"
	n := 65

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Verb", "Description", "Input", "Output"})
	t.AppendRows([]table.Row{
		{"%v", "Default format", "15", fmt.Sprintf("%v", i)},
		{"%#v", "Go-Syntax format", "15", fmt.Sprintf("%#v", i)},
		{"%T", "The type of value", "15", fmt.Sprintf("%T", i)},
		{"%d", "Base 10", "15", fmt.Sprintf("%d", i)},
		{"%+d", "Always show sign", "15", fmt.Sprintf("%+d", i)},
		{"%4d", "Pad with spaces (width 4, right justified)", "15", "␣␣15"},
		{"%-4d", "Pad with spaces (width 4, left justified)", "15", "15␣␣"},
		{"%04d", "Pad with zeroes (width 4)", "15", fmt.Sprintf("%04d", i)},
		{"%b", "Base 2", "15", fmt.Sprintf("%b", i)},
		{"%o", "Base 8", "15", fmt.Sprintf("%o", i)},
		{"%x", "Base 16, lowercase", "15", fmt.Sprintf("%x", i)},
		{"%X", "Base 16, uppercase", "15", fmt.Sprintf("%X", i)},
		{"%#x", "Base 16, with leading 0x", "15", fmt.Sprintf("%#x", i)},
		{"%c", "Character", "65", fmt.Sprintf("%c", n)},
		{"%q", "Quoted Character", "65", fmt.Sprintf("%q", n)},
		{"%U", "Unicode", "65", fmt.Sprintf("%U", n)},
		{"%#U", "Unicode with character", "65", fmt.Sprintf("%#U", n)},
		{"%t", "Boolean", "true", fmt.Sprintf("%t", true)},
		{"%e", "Scientific notation", fmt.Sprintf("%v", f), fmt.Sprintf("%e", f)},
		{"%f", "Decimal point, no exponent", fmt.Sprintf("%v", f), fmt.Sprintf("%f", f)},
		{"%.2f", "Default width, precision 2", fmt.Sprintf("%v", f), fmt.Sprintf("%.2f", f)},
		{"%8.2f", "Width 8, precision 2", fmt.Sprintf("%v", f), fmt.Sprintf("%8.2f", f)},
		{"%s", "Plain string", fmt.Sprintf("%v", s), fmt.Sprintf("%s", s)},
		{"%6s", "Width 6, right justify", fmt.Sprintf("%v", s), fmt.Sprintf("%6s", s)},
		{"%-6s", "Width 6, left justify", fmt.Sprintf("%v", s), fmt.Sprintf("%-6s", s)},
		{"%p", "Pointer", "", fmt.Sprintf("%p", &i)},
		// {"\\a", "U+0007 alert or bell", "", ""},
		// {"\\b", "U+0008 backspace", "", ""},
		// {"\\\\", "U+005c backslash", "", ""},
		// {"\\t", "U+0009 horizontal tab", "", ""},
		// {"\\n", "U+000A line feed or newline", "", ""},
		// {"\\f", "U+000C form feed", "", ""},
		// {"\\r", "U+000D carriage return", "", ""},
		// {"\\v", "U+000b vertical tab", "", ""},
	})

	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.Render()
}
