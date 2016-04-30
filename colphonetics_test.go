package colphonetics

import (
	"fmt"
	"testing"
)

var fixtures = []struct {
	input, expected string
}{
	{"", ""},
	{"Müller-Lüdenscheidt", "65752682"},
	{"Wikipedia", "3412"},
	{"Breschnew", "17863"},
	{"Bréschnew", "17863"},
	{"1253", ""},
	{"ß", "8"},
}

func TestCode(t *testing.T) {
	for _, f := range fixtures {
		if c := Code(f.input); c != f.expected {
			t.Errorf("Error: expected %v got %v", f.expected, c)
		}
	}
}

func BenchmarkCode_LongWord(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Code("Müller-Lüdenscheidt")
	}
}

func BenchmarkCode_ShortWord(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Code("Bla")
	}
}

func ExampleCode() {
	fmt.Println(Code("Müller-Lüdenscheidt"))
	// Output: 65752682
}
