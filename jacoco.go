package omelette

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const JACOCO_LIB_DIR = "lib"

/*
JacocoCoverager is an implementation of Coverager.
*/
type JacocoCoverager struct {
}

func jacocoCsvToOmelette(from, to string) error {
	reader, err1 := os.Open(from)
	if err1 != nil {
		return err1
	}
	defer reader.Close()
	writer, err2 := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0600)
	if err2 != nil {
		return err2
	}
	return jacocoToOmelette(reader, writer)
}

func (jc *JacocoCoverager) ToOmeletteFormat(tr TestRunner, project Project, config *Config) error {
	if err := executeJacocoReport(tr, project, config); err != nil {
		return err
	}
	config.PrintIfVerbose("convert jacoco csv to omelette csv format")
	from := filepath.Join(tr.DestDir(), project.Name()+"_jacoco.csv")
	wd, _ := os.Getwd()
	to := filepath.Join(wd, project.Name()+"_coverage.csv")
	return jacocoCsvToOmelette(from, to)
}

func executeJacocoReport(tr TestRunner, project Project, config *Config) error {
	config.PrintIfVerbose("report JaCoCo coverages in csv format")
	wd := config.OmeletteHome()
	execFile := fmt.Sprintf("%s/%s_jacoco.exec", tr.DestDir(), project.Name())
	csvFile := fmt.Sprintf("%s/%s_jacoco.csv", tr.DestDir(), project.Name())
	cli := fmt.Sprintf("-jar %s/%s/jacococli.jar report %s --classfiles %s --csv %s", wd, JACOCO_LIB_DIR, execFile, project.ProductCodesDir(), csvFile)
	config.PrintIfVerbose("execute: java %s", cli)
	cmd := exec.Command("java", strings.Split(cli, " ")...)
	out, err := cmd.Output()
	cmd.Run()
	config.PrintIfVerbose("output: %s", string(out))
	return err
}

func (jc *JacocoCoverager) Args(tr TestRunner, project Project, config *Config) string {
	omelette := config.OmeletteHome()
	jacocoPath := filepath.Join(omelette, JACOCO_LIB_DIR, "jacocoagent.jar")
	dest := filepath.Join(tr.DestDir(), project.Name()+"_jacoco.exec")
	return fmt.Sprintf(`-javaagent:%s=destfile=%s`, jacocoPath, dest)
}

type coverage struct {
	Missed  int
	Covered int
}

type jacocoItem struct {
	packageName string
	className   string
	instruction *coverage
	branch      *coverage
	line        *coverage
	complexity  *coverage
	method      *coverage
}

func (coverage *coverage) CoverageString() string {
	if coverage.IsNA() {
		return "n/a"
	}
	return strconv.FormatFloat(coverage.Coverage(), 'f', 4, 64)
}

func (c *coverage) IsNA() bool {
	return c.Covered == 0 && c.Missed == 0
}

func (c *coverage) Add(other *coverage) *coverage {
	return &coverage{Missed: c.Missed + other.Missed, Covered: c.Covered + other.Covered}
}

func (c *coverage) Coverage() float64 {
	if c.IsNA() {
		return -1
	}
	return float64(c.Covered) / float64(c.Covered+c.Missed)
}

func newCoverage(missed, covered string) *coverage {
	missedValue, err := strconv.Atoi(missed)
	if err != nil {
		missedValue = 0
	}
	coveredValue, err := strconv.Atoi(covered)
	if err != nil {
		coveredValue = 0
	}
	return &coverage{Missed: missedValue, Covered: coveredValue}
}

func parse(line []string) *jacocoItem {
	item := &jacocoItem{packageName: line[1], className: line[2]}
	item.instruction = newCoverage(line[3], line[4])
	item.branch = newCoverage(line[5], line[6])
	item.line = newCoverage(line[7], line[8])
	item.complexity = newCoverage(line[9], line[10])
	item.method = newCoverage(line[11], line[12])
	return item
}

func readJacocoCsv(in io.Reader) ([]*jacocoItem, error) {
	results := []*jacocoItem{}
	reader := csv.NewReader(in)
	reader.Read() // ignore the first line.
	for {
		line, err := reader.Read()
		if err != nil {
			break
		}
		results = append(results, parse(line))
	}
	return results, nil
}

func (item *jacocoItem) Add(other *jacocoItem) {
	item.instruction = item.instruction.Add(other.instruction)
	item.branch = item.branch.Add(other.branch)
	item.line = item.line.Add(other.line)
	item.complexity = item.complexity.Add(other.complexity)
	item.method = item.method.Add(other.method)
}

func computeCoverageByPackage(items []*jacocoItem) []*jacocoItem {
	coverages := map[string]*jacocoItem{}
	for _, item := range items {
		value, ok := coverages[item.packageName]
		if !ok {
			value = &jacocoItem{packageName: item.packageName, className: "-", instruction: new(coverage), branch: new(coverage), complexity: new(coverage), line: new(coverage), method: new(coverage)}
		}
		value.Add(item)
		coverages[item.packageName] = value
	}
	total := &jacocoItem{packageName: "total", className: "-", instruction: new(coverage), branch: new(coverage), complexity: new(coverage), line: new(coverage), method: new(coverage)}
	results := []*jacocoItem{total}
	for _, values := range coverages {
		results = append(results, values)
		total.Add(values)
	}
	return results
}

func appendCoverage(slice []string, coverage *coverage) []string {
	slice = append(slice, coverage.CoverageString())
	slice = append(slice, strconv.Itoa(coverage.Missed))
	slice = append(slice, strconv.Itoa(coverage.Covered))
	return slice
}

func toStringSlice(item *jacocoItem) []string {
	results := []string{item.packageName, item.className}
	results = appendCoverage(results, item.instruction)
	results = appendCoverage(results, item.branch)
	results = appendCoverage(results, item.line)
	results = appendCoverage(results, item.complexity)
	results = appendCoverage(results, item.method)
	return results
}

func csvHeader() []string {
	return []string{"Package name", "Class name", "Instruction coverage", "Instruction missed", "Instruction covered", "Branch coverage", "Branch missed", "Branch covered", "Line coverage", "Line missed", "Line covered", "Complexity coverage", "Complexity missed", "Complexity covered", "Method coverage", "Method missed", "Method covered"}
}

func writeFile(out io.Writer, items []*jacocoItem) error {
	packageCoverage := computeCoverageByPackage(items)
	writer := csv.NewWriter(out)
	writer.Write(csvHeader())
	for _, item := range packageCoverage {
		writer.Write(toStringSlice(item))
	}
	for _, item := range items {
		writer.Write(toStringSlice(item))
	}
	writer.Flush()
	return nil
}

func jacocoToOmelette(reader io.Reader, writer io.Writer) error {
	items, err := readJacocoCsv(reader)
	if err != nil {
		return err
	}
	return writeFile(writer, items)
}
