package shortUrl

import (
	"testing"
)

func TestEncode(t *testing.T) {
	short := Encode("hello")
	if short != 750201407 {
		t.Errorf("shortUrl Encode failed")
	}
}

func TestDecode(t *testing.T) {
	origin := Decode(750201407)
	if origin != "hello" {
		t.Errorf("shortUrl Decode failed")
	}
}
