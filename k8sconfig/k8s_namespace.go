package k8sconfig

import "fmt"

func (c *svcConfig) generateK8SNamespace() (namespace string) {
	namespace = fmt.Sprintf(k8sNameSpaceTemplate, c.namespace)

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
