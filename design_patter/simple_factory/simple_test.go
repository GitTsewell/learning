package simple_factory

import "testing"

func TestNewApi(t *testing.T) {
	api := NewApi()
	s := api.Say("tsewell")
	if s != "hi,my name is tsewell" {
		t.Fatal("newApi test fail")
	}
}
