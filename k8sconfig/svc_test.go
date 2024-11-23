package k8sconfig

import (
	"log/slog"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(log)

	config := NewConfig(
		log,
		"mirco",
		8000,
		"madmeier/mirco:1.0.8",
		WithNamespace("blue-orb"),
		WithReadinessProbe(
			ProbeOptions{
				InitialDelaySeconds: 5,
				TimeoutSeconds:      1,
				PeriodSeconds:       10,
				SuccessThreshold:    1,
				FailureThreshold:    3,
			},
		),
		WithLivenessProbe(
			ProbeOptions{
				InitialDelaySeconds: 5,
				TimeoutSeconds:      1,
				PeriodSeconds:       10,
				SuccessThreshold:    1,
				FailureThreshold:    3,
			},
		),
		WithStartupProbe(
			ProbeOptions{
				InitialDelaySeconds: 5,
				TimeoutSeconds:      1,
				PeriodSeconds:       10,
				SuccessThreshold:    1,
				FailureThreshold:    5,
			},
		),
		WithEnvironmentVariable(
			"FeatureOne", "enabled",
		),
		WithEnvironmentVariable(
			"FeatureTwo", "enabled",
		),
	)

	path := "." + string(os.PathSeparator) + "config"

	config.generateK8SFiles(log, path)
}
