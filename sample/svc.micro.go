package sample

import (
	"log/slog"

	"github.com/blueorb/config/config"
)

func MicroServiceConfig(log *slog.Logger) (svc *config.SvcConfig) {
	return config.NewConfig(
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
}
