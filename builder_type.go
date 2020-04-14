package omlet

import (
	"fmt"
	"os"
	"path/filepath"
)

type BuilderType interface {
	ProductCodeDir(project *Project) string
	TestCodeDir(project *Project) string
	Name() string
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
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
		if FileExists(filepath.Join(projectPath, parser.fileName)) {
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

func (gt *GradleType) ProductCodeDir(project *Project) string {
	return filepath.Join(project.Path, "build/classes/java/main")
}

func (gt *GradleType) TestCodeDir(project *Project) string {
	return filepath.Join(project.Path, "build/classes/java/test")
}

type MavenType struct {
}

func (mt *MavenType) Name() string {
	return "maven"
}

func (mt *MavenType) ProductCodeDir(project *Project) string {
	return filepath.Join(project.Path, "target/classes")
}

func (mt *MavenType) TestCodeDir(project *Project) string {
	return filepath.Join(project.Path, "target/test-classes")
}
