package models

import "fmt"

const (
	TemplatesDir = "pkg/templates"
	Helm         = "helm"
)

var (
	TiltTemplatePath = fmt.Sprintf("%s/tilt.tmpl", TemplatesDir)
)
