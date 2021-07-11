package helm

import (
	"github.com/arriqaaq/tiltify/pkg/models"
)

func NewCapsuleDeployment() models.HelmRemoteDeployment {
	return models.HelmRemoteDeployment{
		Name:            "capsule",
		ReleaseName:     "capsule",
		RepoName:        "clastix",
		RepoURL:         "https://clastix.github.io/charts",
		Namespace:       "capsule-system",
		Set:             []string{"installCRDs=true"},
		CreateNamespace: "True",
	}
}
