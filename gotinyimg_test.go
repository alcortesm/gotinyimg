package gotinyimg_test

import (
	"testing"

	"github.com/alcortesm/gotinyimg"

	"golang.org/x/text/language"
)

func TestHelloKnowLang(t *testing.T) {
	got, err := gotinyimg.Hello(language.English)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "Hello, world!"
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestHelloUnknowLang(t *testing.T) {
	got, err := gotinyimg.Hello(language.Spanish)
	if err == nil {
		t.Fatalf("expected an error, got none; message was %s", got)
	}
}
