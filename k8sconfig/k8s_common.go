package k8sconfig

func (c *svcConfig) k8sAppName() string {
	return c.serviceName + "-app"
}

func (c *svcConfig) k8sServiceName() string {
	return c.serviceName + "-svc"
}

func (c *svcConfig) k8sServiceConfigName() string {
	return c.serviceName + "-cfg"
}

func (c *svcConfig) k8sSvcContainerName() string {
	return c.serviceName + "-container"
}

func (c *svcConfig) k8sSvcDeploymentName() string {
	return c.serviceName + "-deployment"
}
