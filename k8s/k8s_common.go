package k8s

import "github.com/blueorb/config/config"

func k8sAppName(c *config.SvcConfig) string {
	return c.ServiceName() + "-app"
}

func k8sServiceName(c *config.SvcConfig) string {
	return c.ServiceName() + "-svc"
}

func k8sServiceConfigName(c *config.SvcConfig) string {
	return c.ServiceName() + "-cfg"
}

func k8sSvcContainerName(c *config.SvcConfig) string {
	return c.ServiceName() + "-container"
}

func k8sSvcDeploymentName(c *config.SvcConfig) string {
	return c.ServiceName() + "-deployment"
}
