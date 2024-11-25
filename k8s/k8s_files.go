package k8s

import (
	"log/slog"
	"os"

	"github.com/blueorb/config/config"
)

func GenerateK8SFiles(log *slog.Logger, architecture *config.ArchConfig, path string) {
	err := architecture.Check(log)
	if err != nil {
		log.Error("architecture check failed - generation aborted", "error", err)
		return
	}

	for _, env := range architecture.Environments() {

		envPath := path + "/" + env.Name()

		log.Info("generating files for environment", "environment", env.Name())

		for _, namespace := range architecture.NameSpaces() {
			generateFile(log, namespace, envPath, "namespace", generateK8SNamespace(log, env, namespace))
		}

		for _, cfgMap := range architecture.ConfigMaps() {
			generateFile(log, cfgMap, envPath, "configmap", generateK8SConfigMap(log, env, cfgMap))
		}

	}
	return
}

func generateFile(log *slog.Logger, element config.ArchElement, path string, fileType string, content string) {
	fileName := path + "/" + element.Name() + "_" + fileType + ".yaml"

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
