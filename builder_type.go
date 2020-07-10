package omelette

import (
	"fmt"
	"path/filepath"
)

type BuilderType interface {
	ProductCodeDir(project *BasicProject) string
	TestCodeDir(project *BasicProject) string
	//	Classpath() []string
	Name() string
}

func ParseType(projectPath string) (BuilderType, error) {
	parsers := []struct {
		fileName  string
		generator func() BuilderType
	}{
		{"pom.xml", func() BuilderType { return new(MavenType) }},
		{"build.gradle", func() BuilderType { return new(GradleType) }},
	}
	for _, parser := range parsers {
		if ExistsFile(filepath.Join(projectPath, parser.fileName)) {
			return parser.generator(), nil
		}
	}
	return nil, fmt.Errorf("%s: unknown builder", projectPath)
}

type GradleType struct {
}

func (gt *GradleType) Name() string {
	return "gradle"
}

func (gt *GradleType) ProductCodeDir(project *BasicProject) string {
	return filepath.Join(project.BaseDir(), "build/classes/java/main")
}

func (gt *GradleType) TestCodeDir(project *BasicProject) string {
	return filepath.Join(project.BaseDir(), "build/classes/java/test")
}

type MavenType struct {
}

func (mt *MavenType) Name() string {
	return "maven"
}

func (mt *MavenType) ProductCodeDir(project *BasicProject) string {
	return filepath.Join(project.BaseDir(), "target/classes")
}

func (mt *MavenType) TestCodeDir(project *BasicProject) string {
	return filepath.Join(project.BaseDir(), "target/test-classes")
}
