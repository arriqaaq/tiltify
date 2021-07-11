package models

// Workload represents the structure to represent all configs and resources to deploy an application.
type Workload struct {
	Deployments           []Deployment           `json:"deployment" yaml:"deployment"`
	KustomizeDeployments  []KustomizeDeployment  `json:"kustomize_deployments" yaml:"kustomize_deployments"`
	HelmLocalDeployments  []HelmLocalDeployment  `json:"helm_local_deployments" yaml:"helm_local_deployments"`
	HelmRemoteDeployments []HelmRemoteDeployment `json:"helm_remote_deployments" yaml:"helm_remote_deployments"`
}

// HelmRemoteDeployment represents configuration options for the helm deployment spec.
type HelmRemoteDeployment struct {
	Name            string   `json:"name" yaml:"name"`
	ReleaseName     string   `json:"release_name" yaml:"release_name"`
	RepoName        string   `json:"repo_name" yaml:"repo_name"`
	RepoURL         string   `json:"repo_url" yaml:"repo_url"`
	Version         string   `json:"version" yaml:"version"`
	Namespace       string   `json:"namespace" yaml:"namespace"`
	CreateNamespace string   `json:"create_namespace" yaml:"create_namespace"`
	Set             []string `json:"set" yaml:"set"`
}

// Deployment represents configuration options for the helm deployment spec.
type Deployment struct {
	Name  string   `json:"name" yaml:"name"`
	Files []string `json:"files" yaml:"files"`
}

// KustomizeDeployment represents configuration options for the helm Kustomizedeployment spec.
type KustomizeDeployment struct {
	Name string `json:"name" yaml:"name"`
	Dir  string `json:"dir" yaml:"dir"`
}

// HelmLocalDeployment represents configuration options for the helm deployment spec.
type HelmLocalDeployment struct {
	Name string `json:"name" yaml:"name"`
	Dir  string `json:"dir" yaml:"dir"`
}

// Resource is a set of common actions performed on Resource Types.
type Resource interface {
	GetMetaData() ResourceMeta
}

// ResourceMeta contains metadata for preparing resource manifests.
type ResourceMeta struct {
	Config       interface{}
	TemplatePath string
}
