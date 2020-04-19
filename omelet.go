package omelet

import (
	"os"
	"path/filepath"
	"strings"
)

type Project struct {
	Path    string
	Builder BuilderType
}

func NewProject(path string) (*Project, error) {
	builder, err := ParseType(path)
	return &Project{Path: path, Builder: builder}, err
}

func (project *Project) Name() string {
	return filepath.Base(project.Path)
}

func (project *Project) ProductCodesDir() string {
	return project.Builder.ProductCodeDir(project)
}

func (project *Project) TestCodesDir() string {
	return project.Builder.TestCodeDir(project)
}

func CollectSuffix(basePath, suffix string) ([]string, error) {
	return CollectFiles(basePath, func(path string) bool {
		return strings.HasSuffix(path, suffix)
	})
}

func CollectClassFiles(basePath string) ([]string, error) {
	return CollectSuffix(basePath, ".class")
}

func CollectFiles(basePath string, predicate func(path string) bool) ([]string, error) {
	list := []string{}
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && predicate(path) {
			list = append(list, path)
		}
		return nil
	})
	return list, err
}
