package libs

import (
	"fmt"
	"testing"
)

func TestDSNParser(t *testing.T) {
	// map of dsn (string) expecting error (bool)
	testCase := map[string]bool{
		"user:pass@localhost:21": false,
		"user:pass@localhost":    false,
		"user:pass@:21":          false,
		"user@localhost:21":      false,
		"user@localhost":         false,
		"user@:21":               false,
		"localhost:21":           false,
		":21":                    false,
		"":                       false,
		"user:password":          true,
		":password@localhost":    true,
	}

	for v, er := range testCase {
		user, pass, host, port, err := parseDsn(v)

		switch {
		case er && (err == nil):
			t.Errorf("Expecting error for [%v]\n", v)

		case !er && (err != nil):
			t.Errorf("Unexpected error [%v] => [%v][%v][%v][%v]: %v\n", v, user, pass, host, port, err)

		case !er && reconstructDSN(user, pass, host, port) != v:
			t.Errorf("Unexpected result for [%v] => [%v][%v][%v][%v]\n", v, user, pass, host, port)

		case er:
			fmt.Printf("Expected error for [%v]: %v\n", v, err)
		}
	}
}
