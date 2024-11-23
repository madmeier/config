package k8sconfig

import "fmt"

func (c *svcConfig) generateK8SConfigMap() (cfgMap string) {
	cfgMap = fmt.Sprintf(k8sConfigMapBaseTemplate, c.k8sServiceConfigName(), c.namespace)

	for k, v := range c.environmentVariables {
		cfgMap += fmt.Sprintf(k8sConfigMapEnvVarTemplate, k, v)
	}

	return
}

// 1 k8s svc config name
// 2 config name
const k8sConfigMapBaseTemplate = `
apiVersion: v1
kind: ConfigMap
metadata:
  name: %[1]s
  namespace: %[2]s
data:`

// 1 var name
// 2 var value
const k8sConfigMapEnvVarTemplate = `
    %[1]s: "%[2]s"`
