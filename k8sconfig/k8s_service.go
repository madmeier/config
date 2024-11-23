package k8sconfig

import "fmt"

func (c *svcConfig) generateK8SService() (service string) {
	service = fmt.Sprintf(k8sServiceBaseTemplate, c.k8sServiceName(), c.namespace, c.port, c.k8sAppName())

	return
}

// 1 k82 service name
// 2 name_space
// 3 port
// 4 app
const k8sServiceBaseTemplate = `
apiVersion: v1
kind: Service
metadata:
  name: %[1]s
  namespace: %[2]s
spec:
  selector:
    app: %[4]s  # This matches the labels in the Deployment
  ports:
    - protocol: TCP
      port: %[3]d        # Port on which the service will be available internally
      targetPort: %[3]d  # Port that the container listens on
  type: ClusterIP  # Use ClusterIP for internal communication or change to LoadBalancer for external access
 
`
