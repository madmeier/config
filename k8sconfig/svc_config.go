package k8sconfig

import "log/slog"

type ConfigOption func(*svcConfig)

type ProbeOptions struct {
	// Number of seconds after the container has started before startup, liveness or readiness probes are initiated.
	// If a startup probe is defined, liveness and readiness probe delays do not begin until the startup probe has succeeded.
	// If the value of periodSeconds is greater than initialDelaySeconds then the initialDelaySeconds will be ignored.
	// Defaults to 0 seconds. Minimum value is 0.
	InitialDelaySeconds int32
	// How often (in seconds) to perform the probe. Default to 10 seconds. The minimum value is 1.
	// While a container is not Ready, the ReadinessProbe may be executed at times other than the configured periodSeconds interval.
	// This is to make the Pod ready faster.
	PeriodSeconds int32
	// timeoutSeconds: Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1.
	TimeoutSeconds int32
	// After a probe fails failureThreshold times in a row, Kubernetes considers that the overall check has failed:
	// the container is not ready/healthy/live. Defaults to 3. Minimum value is 1.
	// For the case of a startup or liveness probe, if at least failureThreshold probes have failed,
	// Kubernetes treats the container as unhealthy and triggers a restart for that specific container.
	// The kubelet honors the setting of terminationGracePeriodSeconds for that container.
	// For a failed readiness probe, the kubelet continues running the container that failed checks,
	// and also continues to run more probes; because the check failed, the kubelet sets the Ready condition on the Pod to false.
	FailureThreshold int32 // number of retries
	// successThreshold: Minimum consecutive successes for the probe to be considered successful after having failed.
	// Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.
	SuccessThreshold int32
	// Configure a grace period for the kubelet to wait between triggering a shut down of the failed container,
	// and then forcing the container runtime to stop that container. The default is to inherit the Pod-level
	// value for terminationGracePeriodSeconds (30 seconds if not specified), and the minimum value is 1.
	TerminationGracePeriodSeconds int32
}

type svcConfig struct {
	// name of the service
	serviceName string
	port        int32 // required
	// uses sandbox if not specified
	environment string
	// uses default if not specified
	namespace string
	// service account to use
	serviceAccount string
	// docker image to use
	image string
	// environment variables with default value
	environmentVariables  map[string]string
	readinessProbe        bool
	readinessProbeOptions *ProbeOptions
	livenessProbe         bool
	livenessProbeOptions  *ProbeOptions
	startupProbe          bool
	startupProbeOptions   *ProbeOptions
}

func WithEnvironmentVariable(name string, defaultValue string) ConfigOption {
	return func(s *svcConfig) {
		s.environmentVariables[name] = defaultValue
	}
}

func WithNamespace(namespace string) ConfigOption {
	return func(s *svcConfig) {
		s.namespace = namespace
	}
}

func WithServiceAccount(serviceAccount string) ConfigOption {
	return func(s *svcConfig) {
		s.serviceAccount = serviceAccount
	}
}

func WithReadinessProbe(probeOptions ...ProbeOptions) ConfigOption {
	return func(s *svcConfig) {
		s.readinessProbe = true
		if len(probeOptions) > 0 {
			s.readinessProbeOptions = &probeOptions[0]
		}
	}
}

func WithLivenessProbe(probeOptions ...ProbeOptions) ConfigOption {
	return func(s *svcConfig) {
		s.livenessProbe = true
		if len(probeOptions) > 0 {
			s.livenessProbeOptions = &probeOptions[0]
		}
	}
}

func WithStartupProbe(probeOptions ...ProbeOptions) ConfigOption {
	return func(s *svcConfig) {
		s.startupProbe = true
		if len(probeOptions) > 0 {
			s.startupProbeOptions = &probeOptions[0]
		}
	}
}

func WithSvcImage(image string) ConfigOption {
	return func(s *svcConfig) {
		s.image = image
	}
}

func getDefaultsvcConfig(serviceName string, servicePort int32, imageName string) *svcConfig {
	return &svcConfig{
		serviceName:          serviceName,
		port:                 servicePort,
		environment:          "sb",
		namespace:            "default",
		serviceAccount:       "default",
		environmentVariables: map[string]string{},
		image:                imageName,
	}
}

func extractServiceOptions(
	log *slog.Logger,
	serviceName string,
	servicePort int32,
	imageName string,
	options ...ConfigOption,
) (
	svcConfig *svcConfig,
) {
	log.Info("extracting config options", "number of options", len(options))

	svcConfig = getDefaultsvcConfig(serviceName, servicePort, imageName)
	for _, opt := range options {
		opt(svcConfig)
	}

	return
}

func NewConfig(log *slog.Logger, serviceName string, servicePort int32, imageName string, options ...ConfigOption) (config *svcConfig) {
	config = extractServiceOptions(log, serviceName, servicePort, imageName, options...)
	log.Info("config options", "svcConfig", config)
	return
}

// generate the kubernetes deployment yaml
func (c *svcConfig) generateKubernetes() string {
	return "kubernetes"
}

// generate the go config file
func (c *svcConfig) generateGo() string {
	return "go"
}
