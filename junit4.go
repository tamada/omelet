package omelette

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const JUNIT4_LIB_DIR = "lib/junit4"

type JUnit4TestRunner struct {
	Project    Project
	Coverager  Coverager
	Classpaths []string
	dest       string
}

func NewJUnit4TestRunner(project Project, coverager Coverager) *JUnit4TestRunner {
	j4tr := &JUnit4TestRunner{Project: project, Classpaths: []string{}, Coverager: coverager, dest: os.TempDir()}
	j4tr.AppendClasspath(project.ProductCodesDir())
	j4tr.AppendClasspath(project.TestCodesDir())
	return j4tr
}

func (j4tr *JUnit4TestRunner) Name() string {
	return "junit4"
}

func (j4tr *JUnit4TestRunner) DestDir() string {
	return j4tr.dest
}

func (j4tr *JUnit4TestRunner) AppendClasspath(path string) {
	wd, _ := os.Getwd()
	for _, cp := range strings.Split(path, ":") {
		info, err := os.Stat(cp)
		if err != nil {
			continue
		}
		if info.IsDir() {
			j4tr.Classpaths = appendJarsOrDir(j4tr.Classpaths, cp, wd)
		} else {
			j4tr.Classpaths = append(j4tr.Classpaths, toAbsolute(cp, wd))
		}
	}
}

func buildCli(j4tr *JUnit4TestRunner, config *Config) (string, error) {
	j4tr.AppendClasspath(filepath.Join(config.OmeletteHome(), JUNIT4_LIB_DIR))
	classpath := strings.Join(j4tr.Classpaths, ":")
	coverage := j4tr.Coverager.Args(j4tr, j4tr.Project, config)
	arguments, err := CollectTestTarget(j4tr.Project)
	if err != nil {
		return "", err
	}
	arguments = config.Filter(arguments)
	unitTests := strings.Join(arguments, " ")
	cli := fmt.Sprintf("-cp %s %s org.junit.runner.JUnitCore %s", classpath, coverage, unitTests)
	config.PrintIfVerbose("execute: java %s", cli)
	return cli, nil
}

func (j4tr *JUnit4TestRunner) Perform(config *Config) error {
	if err := j4tr.executeUnitTests(config); err != nil {
		return err
	}
	return j4tr.Coverager.ToOmeletteFormat(j4tr, j4tr.Project, config)
}

func (j4tr *JUnit4TestRunner) executeUnitTests(config *Config) error {
	config.PrintIfVerbose("running unit test codes with JUnit4...")
	cli, err := buildCli(j4tr, config)
	if err != nil {
		return err
	}
	args := strings.Split(cli, " ")
	cmd := exec.Command("java", args...)
	cmd.Dir = j4tr.Project.BaseDir()
	stderr := bytes.NewBuffer([]byte{})
	cmd.Stderr = stderr
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("%s", stderr.String())
		return fmt.Errorf("JUnit4 failed: %s", err.Error())
	}
	config.PrintIfVerbose("output: %s", string(out))
	return nil
}

func (j4tr JUnit4TestRunner) Postprocess(config *Config) {
	if !config.removeTempFile {
		return
	}
	config.PrintIfVerbose("remove %s_jacoco.{exec, csv} in %s", j4tr.Project.Name(), j4tr.DestDir())
	os.Remove(filepath.Join(j4tr.DestDir(), j4tr.Project.Name()+"_jacoco.exec"))
	os.Remove(filepath.Join(j4tr.DestDir(), j4tr.Project.Name()+"_jacoco.csv"))
}

func toAbsolute(path, wd string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(wd, path)
}

func appendJarsOrDir(paths []string, path, wd string) []string {
	jars, err := CollectFiles(path, func(path string) bool {
		return strings.HasSuffix(path, ".jar")
	})
	if err != nil {
		return paths
	}
	if len(jars) == 0 {
		return append(paths, toAbsolute(path, wd))
	}
	for _, jar := range jars {
		paths = append(paths, toAbsolute(jar, wd))
	}
	return paths
}
