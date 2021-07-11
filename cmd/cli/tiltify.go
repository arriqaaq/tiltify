package cli

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/arriqaaq/tiltify/pkg/models"
	"github.com/arriqaaq/tiltify/pkg/questions"

	// "go.starlark.net/starlark"

	"github.com/knadh/stuffbin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	configFile = "Tiltfile"
)

// Tiltify represents the structure for all app wide functions and structs.
type Tiltify struct {
	Logger  *logrus.Logger
	Fs      stuffbin.FileSystem
	Version string
}

// NewTiltify initializes an instance of Tiltify which holds app wide configuration.
func NewTiltify(logger *logrus.Logger, fs stuffbin.FileSystem, buildVersion string) *Tiltify {
	tiltify := &Tiltify{
		Logger:  logger,
		Fs:      fs,
		Version: buildVersion,
	}
	return tiltify
}

func (t *Tiltify) ensureInstalled() error {
	_, err := exec.LookPath("kind")
	if err != nil {
		return fmt.Errorf("kind not installed. Please install kind with these instructions: https://kind.sigs.k8s.io/")
	}
	_, err = exec.LookPath("tilt")
	if err != nil {
		return fmt.Errorf("tilt is not installed. Please install tilt with these instructions: https://docs.tilt.dev/install.html")
	}
	_, err = exec.LookPath("ctlptl")
	if err != nil {
		return fmt.Errorf("ctlptl is not installed. Please install ctlptl with these instructions: https://github.com/tilt-dev/ctlptl")
	}
	return nil
}

// initLoad loads app modules with Tiltify.
func (t *Tiltify) initLoad(fn cli.ActionFunc) cli.ActionFunc {
	return func(c *cli.Context) error {
		// check for modules
		err := t.ensureInstalled()
		if err != nil {
			log.Fatal(err)
		}
		return fn(c)
	}
}

// InitProject initializes git repo and copies a sample config
func (t *Tiltify) InitProject() cli.Command {
	return cli.Command{
		Name:    "init",
		Aliases: []string{"i"},
		Usage:   "Initializes a kind cluster.",
		Action:  t.initLoad(t.init),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "output, o",
				Usage: "Config file name.",
			},
		},
	}
}

func (t *Tiltify) init(cliCtx *cli.Context) error {
	var (
		output = cliCtx.String("output")
	)

	if output != "" {
		configFile = output
	}
	projectDir := filepath.Clean(configFile)

	var workload = models.NewWorkload()

	// Create default helm deployments
	defaultClusterDeployments := questions.GatherDefaultClusterDeploymentInfo()
	if len(defaultClusterDeployments) == 0 {
		return fmt.Errorf(fmt.Sprintf("No helm deployments specified in configuration."))
	}
	workload.HelmRemoteDeployments = append(workload.HelmRemoteDeployments, defaultClusterDeployments...)

	// Check local yaml deployments
	localDeployments := questions.GatherDeploymentsInfo()
	if len(localDeployments) != 0 {
		workload.Deployments = append(workload.Deployments, localDeployments...)
	}

	// Check local kustomize deployments
	kustomizeDeployments := questions.GatherKustomizeDeploymentsInfo()
	if len(kustomizeDeployments) != 0 {
		workload.KustomizeDeployments = append(workload.KustomizeDeployments, kustomizeDeployments...)
	}

	// Check local helm chart deployments
	helmLocalDeployments := questions.GatherHelmLocalDeploymentsInfo()
	if len(helmLocalDeployments) != 0 {
		workload.HelmLocalDeployments = append(workload.HelmLocalDeployments, helmLocalDeployments...)
	}

	// Check remote helm chart deployments
	helmRemoteDeployments := questions.GatherHelmRemoteDeploymentsInfo()
	if len(helmRemoteDeployments) != 0 {
		workload.HelmRemoteDeployments = append(workload.HelmRemoteDeployments, helmRemoteDeployments...)
	}

	err := createResource(models.Resource(workload), projectDir, models.Helm, t.Fs, configFile)
	if err != nil {
		return err
	}

	log.Printf("Your default cluster configuration is created at %s", configFile)
	return nil
}
