package partial

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJsonShouldApply(t *testing.T) {
	type A struct {
		ID uint
		Name string
	}

	patch := `{"iD": 2, "name": "test"}`
	
	var pa PartialJSON[A]
	err := json.Unmarshal([]byte(patch), &pa)
	if err != nil {
		t.Fatal(err)
	}

	var a A

	err = pa.Apply(&a)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, a.ID, uint(2))

	// Find on database
	a = A{
		ID: 2,
		Name: "Not a Test",
	}

	pa.Apply(&a)
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, a.Name, "test")
	require.Equal(t, a.ID, uint(2))
}
