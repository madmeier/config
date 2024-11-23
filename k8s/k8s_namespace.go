package k8s

import (
	"fmt"
	"log/slog"

	"github.com/blueorb/config/config"
)

func generateK8SNamespace(log *slog.Logger, c *config.SvcConfig) (namespace string) {
	namespace = fmt.Sprintf(k8sNameSpaceTemplate, c.Namespace())

	return
}

// 1 k82 service name
// 2 name_space
// 3 port
// 4 app
const k8sNameSpaceTemplate = `
apiVersion: v1
kind: Namespace
metadata:
  name: %[1]s
`
