package k8sconfig

import (
	"log/slog"
	"os"
	"testing"

	"github.com/blueorb/config/k8s"
	"github.com/blueorb/config/sample"
	"github.com/stretchr/testify/assert"
)

// test the architecture artefact generation
func TestArchitectureK8SGeneration(t *testing.T) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	architecture := sample.DefineArchitecture(log)

	path := "." + string(os.PathSeparator) + "config"

	k8s.GenerateK8SFiles(log, architecture, path)
}

// test the architecture verification
func TestArchitectureVerification(t *testing.T) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	architecture := sample.DefineArchitecture(log)

	architecture.Check(log)

	assert.Empty(t, architecture.Errors())
}
