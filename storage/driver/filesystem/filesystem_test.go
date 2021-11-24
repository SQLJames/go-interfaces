package filesystem

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	assertCorrectMessage := func (t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Making sure the extension in the new function is properly formatted", func(t *testing.T) {
		filesystem := New(os.TempDir(),"json")
		expected := ".json"
		assertCorrectMessage(t,filesystem.extension,expected)
	})
	t.Run("Checking the filepath is properly formatted.", func(t *testing.T) {
		filesystem := New("test","json")
		got := filesystem.tablePath("test")
		expected := "test/test.json"
		assertCorrectMessage(t,got,expected)
	})
}