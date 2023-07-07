package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrorNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		if err != want {
			t.Errorf("got %q want %q", err, want)
		}
	})

	t.Run("add word that doesn't exists", func(t *testing.T) {
		dictionary.Add("test2", "this is just a test2")
		got, _ := dictionary.Search("test2")
		want := "this is just a test2"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test2")
		}
	})

	t.Run("add word that already exists", func(t *testing.T) {
		err := dictionary.Add("test", "this is just a test")
		want := ErrorWordExists

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		if err != want {
			t.Errorf("got %q want %q", err, want)
		}
	})

	t.Run("update word that exists", func(t *testing.T) {
		dictionary.Update("test", "this is just a test updated")
		got, _ := dictionary.Search("test")
		want := "this is just a test updated"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	})

	t.Run("update word that doesn't exists", func(t *testing.T) {
		err := dictionary.Update("test3", "this is just a test3")
		want := ErrorWordDoesNotExist

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		if err != want {
			t.Errorf("got %q want %q", err, want)
		}
	})

	t.Run("delete word that exists", func(t *testing.T) {
		dictionary.Delete("test")
		_, err := dictionary.Search("test")
		want := ErrorNotFound

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		if err != want {
			t.Errorf("got %q want %q", err, want)
		}
	})
}
