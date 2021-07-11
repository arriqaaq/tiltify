package helm

import (
	"github.com/arriqaaq/tiltify/pkg/models"
)

func NewCertDeployment() models.HelmRemoteDeployment {
	return models.HelmRemoteDeployment{
		Name:            "cert-manager",
		ReleaseName:     "cert-manager",
		RepoName:        "jetstack",
		RepoURL:         "https://charts.jetstack.io",
		Namespace:       "cert-manager",
		Version:         "1.4.0",
		Set:             []string{"installCRDs=true"},
		CreateNamespace: "True",
	}
}
