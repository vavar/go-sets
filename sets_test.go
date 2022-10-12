package sets_test

import (
	"encoding/json"
	"testing"

	"github.com/vavar/go-sets"
)

func TestSet_ShouldBeUniqueWhenDuplicate(t *testing.T) {
	s := sets.Set[string]{}

	s.Add("11")
	s.Add("22")
	s.Add("11")

	if _, ok := s["11"]; !ok {
		t.Error("insert key not found")
	}

	if _, ok := s["22"]; !ok {
		t.Error("insert key not found")
	}

	sl := s.Slice()

	if len(sl) != 2 {
		t.Error("length check failed")
	}
}

func TestSet_ShouldRemoveKeySuccess(t *testing.T) {
	s := sets.Set[string]{}

	s.Add("11")
	s.Add("22")

	s.Remove("11")
	s.Remove("22")

	sl := s.Slice()

	if len(sl) != 0 {
		t.Error("length check failed")
	}
}

func TestSet_MarshalJSON(t *testing.T) {
	s := sets.Set[string]{}

	s.Add("11")

	b, err := json.Marshal(s)
	if err != nil {
		t.Errorf("JSON Marshal should not be error: %#v", err)
	}
	expected := `["11"]`
	if string(b) != expected {
		t.Errorf("Output failed exp: %s act: %s", expected, string(b))
	}
}

func TestSet_UnmarshalJSON(t *testing.T) {

	s := sets.Set[string]{}
	err := json.Unmarshal([]byte(`["11"]`), &s)

	if err != nil {
		t.Errorf("JSON Unmarshal should not be error: %#v", err)
	}

	sl := s.Slice()

	if len(sl) == 0 {
		t.Error("length check failed")
	}
}
