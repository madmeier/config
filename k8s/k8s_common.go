package k8s

import "github.com/blueorb/config/config"

func k8sAppName(c *config.SvcConfig) string {
	return c.Name() + "-app"
}

func k8sServiceName(c *config.SvcConfig) string {
	return c.Name() + "-svc"
}

func k8sServiceConfigName(c *config.SvcConfig) string {
	return c.Name() + "-cfg"
}

func k8sSvcContainerName(c *config.SvcConfig) string {
	return c.Name() + "-container"
}

func k8sSvcDeploymentName(c *config.SvcConfig) string {
	return c.Name() + "-deployment"
}
