package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := CleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("CleanInput(%q) == %v, expected %v",
				cs.input,
				len(actual),
				len(cs.expected),
			)
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("CleanInput(%q) == %v, expected %v",
					cs.input,
					len(actualWord),
					len(expectedWord),
				)
			}
		}
	}
}
