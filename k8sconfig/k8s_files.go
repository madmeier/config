package k8sconfig

import (
	"log/slog"
	"os"
)

func (c *svcConfig) generateK8SFiles(log *slog.Logger, path string) {
	generateFile(log, c, path, "namespace", c.generateK8SNamespace())
	generateFile(log, c, path, "service", c.generateK8SService())
	generateFile(log, c, path, "deployment", c.generateK8SDeployment())
	generateFile(log, c, path, "svc-config-map", c.generateK8SConfigMap())

	return
}

func generateFile(log *slog.Logger, c *svcConfig, path string, fileType string, content string) {
	fileName := path + "/" + c.serviceName + "_" + fileType + ".yaml"

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
