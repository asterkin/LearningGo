package greeting

import "testing"

func TestGreeting(t *testing.T) {
	var greetingTest = [] struct {
		args     []string
		expected string
	}{
		{ args: []string{},                           expected: "Hello World" },
		{ args: []string{"Golang"},                   expected: "Hello Golang"},
		{ args: []string{"Golang","and", "IntelliJ"}, expected: "Hello Golang and IntelliJ"},
	}

	for i, c := range greetingTest {
		var actual = Greeting(c.args)
		if actual != c.expected {
			t.Errorf("Greeting[%d]: expected '%s' actual '%s'", i, c.expected, actual)
		}
	}
}
