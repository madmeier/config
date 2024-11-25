package k8s

import (
	"fmt"
	"log/slog"

	"github.com/blueorb/config/config"
)

func generateK8SNamespace(log *slog.Logger, env *config.EnvConfig, ns *config.NameSpaceConfig) (namespace string) {
	namespace = fmt.Sprintf(k8sNameSpaceTemplate, ns.Name())

	return
}

// 1 name_space
const k8sNameSpaceTemplate = `
apiVersion: v1
kind: Namespace
metadata:
  name: %[1]s
`
