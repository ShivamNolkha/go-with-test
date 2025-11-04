package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorretMessage(t, got, want)

	})
	t.Run("say 'Hello, World' when an empty string in supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorretMessage(t, got, want)
	})
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Shivam", "Spanish")
		want := "Hola, Shivam"
		assertCorretMessage(t, got, want)
	})
	t.Run("in hindi", func(t *testing.T){
		got := Hello("Shivam", "Hindi")
		want := "Namaste, Shivam"
		assertCorretMessage(t, got, want)
	})
}

func assertCorretMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
