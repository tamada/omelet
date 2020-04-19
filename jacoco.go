package omelet

import (
	"encoding/csv"
	"io"
	"strconv"
)

type Coverage struct {
	Missed  int
	Covered int
}

type JacocoItem struct {
	packageName string
	className   string
	instruction *Coverage
	branch      *Coverage
	line        *Coverage
	complexity  *Coverage
	method      *Coverage
}

func (coverage *Coverage) CoverageString() string {
	if coverage.IsNA() {
		return "n/a"
	}
	return strconv.FormatFloat(coverage.Coverage(), 'f', 4, 64)
}

func (coverage *Coverage) IsNA() bool {
	return coverage.Covered == 0 && coverage.Missed == 0
}

func (coverage *Coverage) Add(other *Coverage) *Coverage {
	return &Coverage{Missed: coverage.Missed + other.Missed, Covered: coverage.Covered + other.Covered}
}

func (coverage *Coverage) Coverage() float64 {
	if coverage.IsNA() {
		return -1
	}
	return float64(coverage.Covered) / float64(coverage.Covered+coverage.Missed)
}

func NewCoverage(missed, covered string) *Coverage {
	missedValue, err := strconv.Atoi(missed)
	if err != nil {
		missedValue = 0
	}
	coveredValue, err := strconv.Atoi(covered)
	if err != nil {
		coveredValue = 0
	}
	return &Coverage{Missed: missedValue, Covered: coveredValue}
}

func parse(line []string) *JacocoItem {
	item := &JacocoItem{packageName: line[1], className: line[2]}
	item.instruction = NewCoverage(line[3], line[4])
	item.branch = NewCoverage(line[5], line[6])
	item.line = NewCoverage(line[7], line[8])
	item.complexity = NewCoverage(line[9], line[10])
	item.method = NewCoverage(line[11], line[12])
	return item
}

func readJacocoCsv(in io.Reader) ([]*JacocoItem, error) {
	results := []*JacocoItem{}
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

func (item *JacocoItem) Add(other *JacocoItem) {
	item.instruction = item.instruction.Add(other.instruction)
	item.branch = item.branch.Add(other.branch)
	item.line = item.line.Add(other.line)
	item.complexity = item.complexity.Add(other.complexity)
	item.method = item.method.Add(other.method)
}

func computeCoverageByPackage(items []*JacocoItem) []*JacocoItem {
	coverages := map[string]*JacocoItem{}
	for _, item := range items {
		value, ok := coverages[item.packageName]
		if !ok {
			value = &JacocoItem{packageName: item.packageName, className: "-", instruction: new(Coverage), branch: new(Coverage), complexity: new(Coverage), line: new(Coverage), method: new(Coverage)}
		}
		value.Add(item)
		coverages[item.packageName] = value
	}
	total := &JacocoItem{packageName: "total", className: "-", instruction: new(Coverage), branch: new(Coverage), complexity: new(Coverage), line: new(Coverage), method: new(Coverage)}
	results := []*JacocoItem{total}
	for _, values := range coverages {
		results = append(results, values)
		total.Add(values)
	}
	return results
}

func appendCoverage(slice []string, coverage *Coverage) []string {
	slice = append(slice, coverage.CoverageString())
	slice = append(slice, strconv.Itoa(coverage.Missed))
	slice = append(slice, strconv.Itoa(coverage.Covered))
	return slice
}

func toStringSlice(item *JacocoItem) []string {
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

func writeFile(out io.Writer, items []*JacocoItem) error {
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

func JacocoToOmelet(reader io.Reader, writer io.Writer) error {
	items, err := readJacocoCsv(reader)
	if err != nil {
		return err
	}
	return writeFile(writer, items)
}
