package iterations

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("expected '%q' but got '%q'", want, got)
	}
}

func BenchmarkRepeat(b *testing.B) {

	b.Run("Repeat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 5)
		}
	})

	b.Run("RepeatStrings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			RepeatStrings("a", 5)
		}
	})
}
