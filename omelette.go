package omelette

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type TestRunner interface {
	Name() string
	AppendClasspath(path string)
	Perform(config *Config) error
	DestDir() string
	Postprocess(config *Config)
}

type Project interface {
	Name() string
	BaseDir() string
	ProductCodesDir() string
	TestCodesDir() string
	String() string
}

type FakeProject struct {
	productCodeDir string
	testCodeDir    string
	name           string
}

func NewFakeProject(name, productDir, testDir string) *FakeProject {
	return &FakeProject{name: name, productCodeDir: productDir, testCodeDir: testDir}
}

func (fp *FakeProject) String() string {
	return fmt.Sprintf("FakeProject{ name: %s, product: %s, test: %s}", fp.name, fp.productCodeDir, fp.testCodeDir)
}

func (fp *FakeProject) Name() string {
	return fp.name
}

func (fp *FakeProject) BaseDir() string {
	wd, _ := os.Getwd()
	return wd
}

func (fp *FakeProject) ProductCodesDir() string {
	return fp.productCodeDir
}

func (fp *FakeProject) TestCodesDir() string {
	return fp.testCodeDir
}

type BasicProject struct {
	path    string
	name    string
	builder BuilderType
}

func NewProject(path string) (Project, error) {
	base := filepath.Base(path)
	if base == "." || base == ".." {
		if abs, err := filepath.Abs(path); err == nil {
			path = abs
		}
	}
	builder, err := ParseType(path)
	if err != nil {
		return nil, err
	}
	return &BasicProject{name: filepath.Base(path), path: path, builder: builder}, err
}

func (project *BasicProject) String() string {
	return fmt.Sprintf("BasicProject{ path: %s, builder: %s }", project.path, project.builder.Name())
}

func (project *BasicProject) Name() string {
	return project.name
}

func (project *BasicProject) BaseDir() string {
	return project.path
}

func (project *BasicProject) ProductCodesDir() string {
	return project.builder.ProductCodeDir(project)
}

func (project *BasicProject) TestCodesDir() string {
	return project.builder.TestCodeDir(project)
}

func CollectTestTarget(project Project) ([]string, error) {
	base := project.TestCodesDir()
	targets, err := CollectSuffix(base, "Test.class")
	if err != nil {
		return []string{}, err
	}
	return convertToClassNames(base, targets), nil
}

func convertToClassNames(base string, targets []string) []string {
	results := []string{}
	for _, c := range targets {
		results = append(results, toClassName(base, c))
	}
	return results
}

func toClassName(base, path string) string {
	relPath, err := filepath.Rel(base, path)
	if err != nil {
		relPath = path
	}
	if strings.HasSuffix(relPath, ".class") {
		relPath = strings.TrimSuffix(relPath, ".class")
	}
	return strings.ReplaceAll(relPath, "/", ".")
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
