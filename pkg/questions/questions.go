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

func GatherDeploymentsInfo() []models.Deployment {
	var deployments = []models.Deployment{}

	// check for yaml deployments
	var hasDeployment bool
	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Would you like to configure yaml deployments?",
	}, &hasDeployment, survey.WithValidator(survey.Required)))
	if hasDeployment {
		var deploymentsLen = 0
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "How many deployments do you want to configure?",
			Help:    "Deployments represent different k8s yaml configurations.",
			Default: "1",
		}, &deploymentsLen, survey.WithValidator(isInt)))
		for i := 0; i < deploymentsLen; i++ {
			var dep = models.Deployment{}
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the name of deployment?",
				Help:    "Name of the deployment to be configured.",
				Default: "mydeployment",
			}, &dep.Name, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the location of the k8s yaml file?",
				Help:    "Location of the k8s yaml file to be deployed.",
			}, &dep.File, survey.WithValidator(survey.Required)))
			deployments = append(deployments, dep)
		}
	}

	return deployments
}

func GatherKustomizeDeploymentsInfo() []models.KustomizeDeployment {
	var deployments = []models.KustomizeDeployment{}
	var hasKustomizeDeployment bool
	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Would you like to configure kustomize deployments?",
	}, &hasKustomizeDeployment, survey.WithValidator(survey.Required)))
	if hasKustomizeDeployment {
		var deploymentsLen = 0
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "How many kustomize deployments do you want to configure?",
			Help:    "Deployments represent different kustomize configurations.",
			Default: "1",
		}, &deploymentsLen, survey.WithValidator(isInt)))
		for i := 0; i < deploymentsLen; i++ {
			var dep = models.KustomizeDeployment{}
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the name of the kustomize deployment?",
				Help:    "Name of the kustomize deployment to be configured.",
				Default: "mydeployment",
			}, &dep.Name, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the location of the kustomize directory?",
				Help:    "Location of the kustomize repo to be deployed.",
			}, &dep.Dir, survey.WithValidator(survey.Required)))
			deployments = append(deployments, dep)
		}
	}

	return deployments
}

func GatherHelmLocalDeploymentsInfo() []models.HelmLocalDeployment {
	var deployments = []models.HelmLocalDeployment{}

	// check for helm charts locally
	var hasHelmLocalCharts bool
	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Would you like to configure helm charts installed locally?",
	}, &hasHelmLocalCharts, survey.WithValidator(survey.Required)))
	if hasHelmLocalCharts {
		var deploymentsLen = 0
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "How many helm charts do you want to configure?",
			Help:    "Helm charts represent different helm charts directories.",
			Default: "1",
		}, &deploymentsLen, survey.WithValidator(isInt)))
		for i := 0; i < deploymentsLen; i++ {
			var dep = models.HelmLocalDeployment{}
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the name of helm chart?",
				Help:    "Name of the chart to be configured.",
				Default: "mydeployment",
			}, &dep.Name, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the location of the helm chart?",
				Help:    "Location of the helm chart to be deployed.",
			}, &dep.Dir, survey.WithValidator(survey.Required)))
			deployments = append(deployments, dep)
		}
	}

	return deployments
}

func GatherHelmRemoteDeploymentsInfo() []models.HelmRemoteDeployment {
	var deployments = []models.HelmRemoteDeployment{}
	var hasHelmRemoteCharts bool
	exitOnInterrupt(survey.AskOne(&survey.Confirm{
		Message: "Would you like to configure helm charts installed locally?",
	}, &hasHelmRemoteCharts, survey.WithValidator(survey.Required)))

	if hasHelmRemoteCharts {
		var deploymentsLen = 0
		exitOnInterrupt(survey.AskOne(&survey.Input{
			Message: "How many remote helm charts do you want to configure?",
			Help:    "Helm charts represent different remote helm charts.",
			Default: "1",
		}, &deploymentsLen, survey.WithValidator(isInt)))
		for i := 0; i < deploymentsLen; i++ {
			var dep = models.HelmRemoteDeployment{}
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the name of helm chart?",
				Help:    "Name of the chart to be configured.",
			}, &dep.Name, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the release name of helm chart?",
			}, &dep.ReleaseName, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the repo name of helm chart?",
			}, &dep.RepoName, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the repo url of helm chart?",
			}, &dep.RepoName, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the version of the helm chart being deployed?",
			}, &dep.RepoName, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Input{
				Message: "What's the namespace this helm chart should be deployed into?",
			}, &dep.Namespace, survey.WithValidator(survey.Required)))
			exitOnInterrupt(survey.AskOne(&survey.Select{
				Message: "Would you like this namespace to be created?",
				Options: []string{"True", "False"},
			}, &dep.CreateNamespace, survey.WithValidator(survey.Required)))

			deployments = append(deployments, dep)
		}
	}
	return deployments
}

func exitOnInterrupt(err error) error {
	if err == terminal.InterruptErr {
		fmt.Println("quitting tiltify...")
		os.Exit(0)
	}
	return err
}
