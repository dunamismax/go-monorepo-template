package demopubliclibrary

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello, World!"
	got := HelloWorld()
	if got != want {
		t.Errorf("HelloWorld() = %q, want %q", got, want)
	}
}
