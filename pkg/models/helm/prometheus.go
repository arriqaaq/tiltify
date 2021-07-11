package helm

import (
	"github.com/arriqaaq/tiltify/pkg/models"
)

func NewPrometheusDeployment() models.HelmRemoteDeployment {
	return models.HelmRemoteDeployment{
		Name:            "kube-prometheus-stack",
		ReleaseName:     "kube-prometheus-stack",
		RepoName:        "prometheus-community",
		RepoURL:         "https://prometheus-community.github.io/helm-charts",
		Namespace:       "monitoring",
		Set:             []string{"installCRDs=true"},
		CreateNamespace: "True",
	}
}
