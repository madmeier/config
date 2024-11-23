package k8sconfig

import (
	"log/slog"
	"os"
	"testing"

	"github.com/blueorb/config/k8s"
	"github.com/blueorb/config/sample"
)

func TestConfig(t *testing.T) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	config := sample.MicroService(log)

	path := "." + string(os.PathSeparator) + "config"

	k8s.GenerateK8SFiles(log, config, path)
}
