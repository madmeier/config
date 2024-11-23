package k8sconfig

import (
	"log/slog"
	"os"
	"testing"

	"github.com/blueorb/config/config"
	"github.com/blueorb/config/k8s"
)

func TestConfig(t *testing.T) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	config := config.NewConfig(
		log,
		"mirco",
		8000,
		"madmeier/mirco:1.0.8",
		config.WithNamespace("blue-orb"),
		config.WithReadinessProbe(
			config.ProbeOptions{
				InitialDelaySeconds: 5,
				TimeoutSeconds:      1,
				PeriodSeconds:       10,
				SuccessThreshold:    1,
				FailureThreshold:    3,
			},
		),
		config.WithLivenessProbe(
			config.ProbeOptions{
				InitialDelaySeconds: 5,
				TimeoutSeconds:      1,
				PeriodSeconds:       10,
				SuccessThreshold:    1,
				FailureThreshold:    3,
			},
		),
		config.WithStartupProbe(
			config.ProbeOptions{
				InitialDelaySeconds: 5,
				TimeoutSeconds:      1,
				PeriodSeconds:       10,
				SuccessThreshold:    1,
				FailureThreshold:    5,
			},
		),
		config.WithEnvironmentVariable(
			"FeatureOne", "enabled",
		),
		config.WithEnvironmentVariable(
			"FeatureTwo", "enabled",
		),
	)

	path := "." + string(os.PathSeparator) + "config"

	k8s.GenerateK8SFiles(log, config, path)
}
