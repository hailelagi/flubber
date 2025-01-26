package storage

import (
	"os"
	"testing"

	"github.com/anishathalye/porcupine"
)

func visualizeTempFile(t *testing.T, model porcupine.Model, info porcupine.LinearizationInfo) {
	file, err := os.CreateTemp("", "*.html")
	if err != nil {
		t.Fatalf("failed to create temp file")
	}
	err = porcupine.Visualize(model, info, file)
	if err != nil {
		t.Fatalf("visualization failed")
	}
	t.Logf("wrote visualization to %s", file.Name())
}
