package omelet

import (
	"io"
	"math"
	"os"
	"strings"
	"testing"
)

func TestJacocoToOmelet(t *testing.T) {
	reader, _ := os.Open("testdata/nmaker_jacoco.csv")
	writer, _ := os.OpenFile("testdata/result_tmp.csv", os.O_WRONLY|os.O_CREATE, 0600)
	defer reader.Close()
	defer writer.Close()
	defer os.Remove("testdata/result_tmp.csv")
	strWriter := &strings.Builder{}
	JacocoToOmelet(reader, io.MultiWriter(writer, strWriter))
	entries := strings.Split(strings.TrimSpace(strWriter.String()), "\n")
	if len(entries) != 16 {
		t.Errorf("the length of result did not match, wont 16, got %d", len(entries))
	}
	if !strings.HasPrefix(entries[0], "Package name,Class name") {
		t.Errorf("the file header is invalid, got \"%s\"", entries[0])
	}
	if entries[1] != "total,-,0.8249,464,2186,0.5962,42,62,0.6906,112,250,0.6812,44,94,0.7907,18,68" {
		t.Errorf("the entries[1] did not match")
	}
}

func TestParseJacocoItem(t *testing.T) {
	testdata := []struct {
		giveString      string
		wontPackage     string
		wontClassName   string
		wontInstruction *Coverage
		wontBranch      *Coverage
		wontLine        *Coverage
		wontComplexity  *Coverage
		wontMethod      *Coverage
	}{
		{"JaCoCo Coverage Report,nmaker,Fraction,44,94,3,7,7,15,5,9,2,7\r\nJaCoCo Coverage Report,nmaker,Number,14,116,2,8,4,22,4,17,2,14\n", "nmaker", "Number", NewCoverage("14", "116"), NewCoverage("2", "8"), NewCoverage("4", "22"), NewCoverage("4", "17"), NewCoverage("2", "14")},
	}
	for _, td := range testdata {
		items, err := readJacocoCsv(strings.NewReader(td.giveString))
		if err != nil {
			t.Errorf("Unknown error: %v", err)
		}
		if len(items) != 1 { // the first line is ignored.
			t.Errorf("the length of read JacocoItem did not match, wont 1, got %d", len(items))
		}
		item := items[0]
		if item.packageName != td.wontPackage || item.className != td.wontClassName {
			t.Errorf("read package name or class name did not match, wont %s.%s, got %s.%s", td.wontPackage, td.wontClassName, item.packageName, item.className)
		}
	}
}

func TestCoverage(t *testing.T) {
	testdata := []struct {
		missed       string
		covered      string
		naFlag       bool
		wontCoverage float64
		wontString   string
	}{
		{"0", "0", true, -1.0, "n/a"},
		{"1", "4", false, 0.8, "0.8000"},
		{"a", "4", false, 1.0, "1.0000"},
		{"b", "b", true, -1.0, "n/a"},
	}

	for _, td := range testdata {
		metric := NewCoverage(td.missed, td.covered)
		if metric.IsNA() != td.naFlag {
			t.Errorf("coverage should be NA flag, wont %v, got %v", td.naFlag, metric.IsNA())
		}
		coverage := metric.Coverage()
		if math.Abs(coverage-td.wontCoverage) > 0.0001 {
			t.Errorf("coverage value did not match, wont %f, got %f", td.wontCoverage, coverage)
		}
		if td.wontString != metric.CoverageString() {
			t.Errorf("coverage string did not match, wont %s, got %s", td.wontString, metric.CoverageString())
		}
	}
}
