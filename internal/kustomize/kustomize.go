package kustomize

import (
	"sigs.k8s.io/kustomize/api/types"
)

// GetKustomization returns a Kustomization pinning each Spinnaker microservice
// to the specified version.
func GetKustomization(version string) types.Kustomization {
	kust := types.Kustomization{
		TypeMeta: types.TypeMeta{
			Kind:       types.KustomizationKind,
			APIVersion: types.KustomizationVersion,
		},
		Namespace: "spinnaker",
		CommonLabels: map[string]string{
			"app.kubernetes.io/component":  "server",
			"app.kubernetes.io/managed-by": "kleat",
			"app.kubernetes.io/part-of":    "spinnaker",
			"app.kubernetes.io/version":    version,
		},
		Images: imageOverrides(version),
		Resources: []string{
			"github.com/spinnaker/kustomization-base/core",
			"github.com/spinnaker/kustomization-base/fiat",
		},
		ConfigMapGenerator: nil,
		SecretGenerator:    nil,
	}
	return kust
}

func imageOverrides(tag string) []types.Image {
	var svcs = []string{
		"clouddriver",
		"deck",
		"echo",
		"fiat",
		"front50",
		"gate",
		"igor",
		"kayenta",
		"monitoring-daemon",
		"orca",
		"rosco",
	}
	imgs := make([]types.Image, len(svcs))
	for i, svc := range svcs {
		imgs[i] = types.Image{
			Name:   "gcr.io/spinnaker-marketplace/" + svc,
			NewTag: tag,
		}
	}
	return imgs
}
