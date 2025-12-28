package deck

import (
	"reflect"
	"testing"
)

func TestErrors(t *testing.T) {
	if ErrDeckNotFound.Error() != "deck not found" {
		t.Fatalf("unexpected ErrDeckNotFound: %v", ErrDeckNotFound)
	}
	if ErrInvalidDeck.Error() != "deck name is required" {
		t.Fatalf("unexpected ErrInvalidDeck: %v", ErrInvalidDeck)
	}
}

func TestDeckJSONTags(t *testing.T) {
	typ := reflect.TypeOf(Deck{})
	want := map[string]string{
		"ID":          "id",
		"OwnerID":     "owner_id",
		"Name":        "name",
		"Description": "description",
		"IsPublic":    "is_public",
		"CreatedAt":   "created_at",
		"UpdatedAt":   "updated_at",
	}
	for field, tag := range want {
		f, ok := typ.FieldByName(field)
		if !ok {
			t.Fatalf("field %s not found", field)
		}
		got := f.Tag.Get("json")
		if got != tag {
			t.Errorf("field %s json tag = %q; want %q", field, got, tag)
		}
	}
}

func TestNewDeckRequestFields(t *testing.T) {
	typ := reflect.TypeOf(NewDeckRequest{})
	wantFields := []struct {
		name string
		kind reflect.Kind
	}{
		{"OwnerID", reflect.String},
		{"Name", reflect.String},
		{"Description", reflect.String},
		{"IsPublic", reflect.Bool},
	}
	for _, wf := range wantFields {
		f, ok := typ.FieldByName(wf.name)
		if !ok {
			t.Fatalf("field %s not found", wf.name)
		}
		if f.Type.Kind() != wf.kind {
			t.Errorf("field %s kind = %v; want %v", wf.name, f.Type.Kind(), wf.kind)
		}
	}
}
