package k8s

import (
	"log/slog"
	"os"

	"github.com/blueorb/config/config"
)

func GenerateK8SFiles(log *slog.Logger, c *config.SvcConfig, path string) {
	generateFile(log, c, path, "namespace", generateK8SNamespace(log, c))
	generateFile(log, c, path, "service", generateK8SService(log, c))
	generateFile(log, c, path, "deployment", generateK8SDeployment(log, c))
	generateFile(log, c, path, "svc-config-map", generateK8SConfigMap(log, c))

	return
}

func generateFile(log *slog.Logger, c *config.SvcConfig, path string, fileType string, content string) {
	fileName := path + "/" + c.Name() + "_" + fileType + ".yaml"

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Error("error creating directory", "error", err, "directory", path)
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Error("error creating file", "error", err, "fileName", fileName)
		return
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Error("error writing to file", "error", err, "fileName", fileName)
		return
	}

	return
}
