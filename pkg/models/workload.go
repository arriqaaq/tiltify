package models

func (w Workload) GetMetaData() ResourceMeta {
	return ResourceMeta{
		TemplatePath: TiltTemplatePath,
		Config:       w,
	}
}

func NewWorkload() Workload {
	return Workload{
		Deployments:           make([]Deployment, 0, 1),
		KustomizeDeployments:  make([]KustomizeDeployment, 0, 1),
		HelmLocalDeployments:  make([]HelmLocalDeployment, 0, 1),
		HelmRemoteDeployments: make([]HelmRemoteDeployment, 0, 1),
	}
}
