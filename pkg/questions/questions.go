package questions

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/arriqaaq/tiltify/pkg/models"
	"github.com/arriqaaq/tiltify/pkg/models/helm"
)

const (
	deploymentKafka       = "kafka"
	deploymentPrometheus  = "prometheus"
	deploymentCertManager = "cert-manager"
	deploymentCapsule     = "capsule"
)

var (
	defaultDeployments = []string{deploymentCertManager, deploymentPrometheus, deploymentKafka, deploymentCapsule}
)

// isInt checks if value is integer
func isInt(val interface{}) error {
	// the reflect value of the result
	value := (val).(string)
	// if the value passed in not int
	_, err := strconv.Atoi(value)
	if err != nil {
		return errors.New("Please enter a number")
	}
	return nil
}

func GatherClusterInfo() bool {
	var wantCluster bool

	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Do you want to automatically create a Cluster?",
		Help:    "Cluster creates a simple kind cluster to handle your workloads",
		Default: true,
	}, &wantCluster, survey.WithValidator(survey.Required)))

	return wantCluster
}

func GatherDefaultClusterDeploymentInfo() []models.HelmRemoteDeployment {
	deployments := []string{}
	exitOnInterrupt(survey.AskOne(&survey.MultiSelect{
		Message: "What services would you like to be installed on the cluster?",
		Options: defaultDeployments,
	}, &deployments))

	deploymentList := []models.HelmRemoteDeployment{}

	for _, dep := range deployments {
		switch dep {
		case deploymentCertManager:
			deploymentList = append(deploymentList, helm.NewCertDeployment())
		case deploymentKafka:
			deploymentList = append(deploymentList, helm.NewKafkaDeployment())
		case deploymentPrometheus:
			deploymentList = append(deploymentList, helm.NewPrometheusDeployment())
		case deploymentCapsule:
			deploymentList = append(deploymentList, helm.NewCapsuleDeployment())
		}
	}
	return deploymentList
}

func exitOnInterrupt(err error) error {
	if err == terminal.InterruptErr {
		fmt.Println("quitting tiltify...")
		os.Exit(0)
	}
	return err
}
