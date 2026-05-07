package internal

import (
	"testing"
)

func TestParseInput(t *testing.T) {

	tests := []struct {
		input    string
		expected Command
	}{
		{"ls -la", Command{Name: "ls", Args: []string{"-la"}}},
		{"echo \"Hello World\"", Command{Name: "echo", Args: []string{"Hello World"}}},
		{"cd /home/user", Command{Name: "cd", Args: []string{"/home/user"}}},
		{"   ", Command{}},
	}

	for _, test := range tests {
		result := ParseInput(test.input)
		if result.Name != test.expected.Name || len(result.Args) != len(test.expected.Args) {
			t.Errorf("ParseInput(%q) = %v, expected %v", test.input, result, test.expected)
			continue
		}
		for i := range result.Args {
			if result.Args[i] != test.expected.Args[i] {
				t.Errorf("ParseInput(%q) = %v, expected %v", test.input, result, test.expected)
				break
			}
		}
	}
}
