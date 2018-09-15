package generate

import (
	"testing"
)

func TestGenerateSingle(t *testing.T) {
	templ := " {{ HELLO }} "
	substitutions := make(map[string]string)
	substitutions["HELLO"] = "hello test"
	gen, err := Generate(templ, substitutions)
	if err != nil {
		t.Fatal(err)
	}
	expectedOut := " hello test "
	if gen != expectedOut {
		t.Fatalf("expected %v, but got %v", expectedOut, gen)
	}
}

func TestGenerateComplex(t *testing.T) {
	templ := `VER: 1.0
	name: {{ NAME }}
	type: person
	password: {{ PASSWORD }}
	notes: none`
	substitutions := make(map[string]string)
	substitutions["NAME"] = "mike"
	substitutions["PASSWORD"] = "badpassword"
	gen, err := Generate(templ, substitutions)
	if err != nil {
		t.Fatal(err)
	}
	expectedOut := `VER: 1.0
	name: mike
	type: person
	password: badpassword
	notes: none`
	if gen != expectedOut {
		t.Fatalf("expected \n%v\n, but got \n%v\n", expectedOut, gen)
	}
}

func TestGenerateComplexRandomSpacing(t *testing.T) {
	templ := `VER: 1.0
	name: {{ NAME}}
	type: person
	password: {{    PASSWORD }}
	notes: none`
	substitutions := make(map[string]string)
	substitutions["NAME"] = "mike"
	substitutions["PASSWORD"] = "badpassword"
	gen, err := Generate(templ, substitutions)
	if err != nil {
		t.Fatal(err)
	}
	expectedOut := `VER: 1.0
	name: mike
	type: person
	password: badpassword
	notes: none`
	if gen != expectedOut {
		t.Fatalf("expected \n%v\n got: \n%v\n", expectedOut, gen)
	}
}

func TestGenerateComplex_NotEnoughSubstitutionsFailure(t *testing.T) {
	templ := `VER: 1.0
	name: {{ NAME }}
	type: person
	password: {{ PASSWORD }}
	notes: none`
	substitutions := make(map[string]string)
	substitutions["NAME"] = "mike"
	gen, err := Generate(templ, substitutions)
	if err == nil {
		t.Fatal("expected error from not enough substitutions, but got no error")
	}
	if gen != "" {
		t.Fatalf("expected failure with empty string, but got \n%v\n", gen)
	}
}
