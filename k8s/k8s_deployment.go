package k8s

import (
	"fmt"
	"log/slog"

	"github.com/blueorb/config/config"
)

func generateOptions(log *slog.Logger, p *config.ProbeOptions) (opts string) {
	if p == nil {
		return
	}

	if p.FailureThreshold != 0 {
		opts = fmt.Sprintf(k8sProbeOptionsTemplate, "failureThreshold", p.FailureThreshold)
	}
	if p.InitialDelaySeconds != 0 {
		opts += fmt.Sprintf(k8sProbeOptionsTemplate, "initialDelaySeconds", p.InitialDelaySeconds)
	}
	if p.PeriodSeconds != 0 {
		opts += fmt.Sprintf(k8sProbeOptionsTemplate, "periodSeconds", p.PeriodSeconds)
	}
	if p.SuccessThreshold != 0 {
		opts += fmt.Sprintf(k8sProbeOptionsTemplate, "successThreshold", p.SuccessThreshold)
	}
	if p.TerminationGracePeriodSeconds != 0 {
		opts += fmt.Sprintf(k8sProbeOptionsTemplate, "terminationGracePeriodSeconds", p.TerminationGracePeriodSeconds)
	}

	return
}

func generateK8SDeployment(log *slog.Logger, c *config.SvcConfig) (cfgMap string) {
	cfgMap = fmt.Sprintf(
		k8sDeploymentBaseTemplate,
		c.Name(),
		c.NameSpace(),
		c.Port(),
		c.Image(),
		c.ServiceAccount(),
		k8sAppName(c),
		k8sSvcDeploymentName(c),
		k8sSvcContainerName(c),
	)

	if c.LivenessProbe() {
		cfgMap += fmt.Sprintf(k8sProbeTemplate, "liveness", c.Port(), "/is-alive")
		cfgMap += generateOptions(log, c.LivenessProbeOptions())
	}

	if c.ReadinessProbe() {
		cfgMap += fmt.Sprintf(k8sProbeTemplate, "readiness", c.Port(), "/is-ready")
		cfgMap += generateOptions(log, c.ReadinessProbeOptions())
	}

	if c.StartupProbe() {
		cfgMap += fmt.Sprintf(k8sProbeTemplate, "startup", c.Port(), "/is-ready")
		cfgMap += generateOptions(log, c.StartupProbeOptions())
	}

	return
}

// 1 service name
// 2 name_space
// 3 port
// 4 image
// 5 service account
// 6 app name
// 7 deployment name
// 8 container name
const k8sDeploymentBaseTemplate = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: %[7]s
  namespace: %[2]s
  labels:
    app: %[6]s
spec:
  replicas: 3  
  selector:
    matchLabels:
      app: %[6]s
  template:
    metadata:
      labels:
        app: %[6]s
    spec:
      serviceAccountName: %[5]s 
      containers:
      - name: %[8]s
        image: %[4]s
        ports:
        - containerPort: %[3]d`

// 1 type
// 2 port
// 3 path
const k8sProbeTemplate = `
        %[1]sProbe:
          httpGet:
            path: %[3]s
            port: %[2]d`

// 1 option name
// 2 option value
const k8sProbeOptionsTemplate = `
          %[1]s: %[2]d`

// 1 service name
const k8sEnvVarTemplate = `
        envFrom:
        - configMapRef:
            name: %[1]s-config`
