package tests

import (
	"testing"

	"github.com/karkulevskiy/shareen/src/internal/lib"
)

// TestGenerateURL tests GenerateURL
func TestGenerateURL(t *testing.T) {
	cases := []struct {
		Name string
	}{
		{
			Name: "Valid URL",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			url := lib.GenerateURL()
			if url == "" {
				t.Error("Empty URL")
			}
		})
	}
}
