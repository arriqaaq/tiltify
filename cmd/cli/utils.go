package cli

import (
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/arriqaaq/tiltify/pkg/models"
	"github.com/knadh/stuffbin"
)

// createWorkload fetches metadata about the workload and creates the Tiltfile.
func createWorkload(resource models.Resource, rootDir string, workload string, fs stuffbin.FileSystem, dest string) error {
	var (
		template           = resource.GetMetaData().TemplatePath
		config             = resource.GetMetaData().Config
		op       io.Writer = os.Stdout
	)

	if rootDir != "" {
		f, err := os.Create(dest)
		if err != nil {
			return err
		}
		op = f
	}

	return saveWorkload(template, op, config, fs)
}

func saveWorkload(template string, dest io.Writer, config interface{}, fs stuffbin.FileSystem) error {
	// parse template file
	tmpl, err := parse(template, fs)
	if err != nil {
		return err
	}

	err = writeTemplate(tmpl, config, dest)
	if err != nil {
		return err
	}

	return nil
}

func parse(src string, fs stuffbin.FileSystem) (*template.Template, error) {
	// read template file
	tmpl := template.New(src)
	// load main template
	f, err := fs.Read(src)
	if err != nil {
		return nil, fmt.Errorf("error reading template file %s: %v", src, err)
	}
	return tmpl.Parse(string(f))
}

func writeTemplate(tmpl *template.Template, config interface{}, dest io.Writer) error {
	err := tmpl.Execute(dest, config)
	if err != nil {
		return err
	}
	return nil
}
