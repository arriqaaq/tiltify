package helm

import (
	"github.com/arriqaaq/tiltify/pkg/models"
)

func NewKafkaDeployment() models.HelmRemoteDeployment {
	return models.HelmRemoteDeployment{
		Name:            "kafka",
		ReleaseName:     "kafka",
		RepoName:        "strimzi",
		RepoURL:         "https://strimzi.io/charts",
		Namespace:       "vdp-kafka",
		Version:         "0.8.2",
		Set:             []string{"installCRDs=true"},
		CreateNamespace: "True",
	}
}
