package omelet

import (
	"testing"
)

func matchSlice(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func TestCollectFiles(t *testing.T) {
	testdata := []struct {
		projectPath          string
		wontName             string
		wontProductFileCount int
		wontTestFileCount    int
		wontProductFiles     []string
		wontTestFiles        []string
	}{
		{"testdata/mavenproject", "mavenproject", 1, 2, []string{"A.class"}, []string{"B.class", "C.class"}},
		{"testdata/gradleproject", "gradleproject", 3, 4, []string{"D.class", "E.clsas", "F.class"}, []string{"G.class", "H.class", "I.class", "J.class"}},
	}

	for _, td := range testdata {
		project, err := NewProject(td.projectPath)
		if err != nil {
			t.Errorf(err.Error())
			continue
		}
		if project.Name() != td.wontName {
			t.Errorf("%s: project name did not match, wont %s, got %s", td.projectPath, td.wontName, project.Name())
		}
		gotProducts, _ := CollectClassFiles(project.ProductCodesDir())
		gotTests, _ := CollectClassFiles(project.TestCodesDir())
		if len(gotProducts) != td.wontProductFileCount {
			t.Errorf("%s: file count (product codes) did not match, wont %d, got %d", td.projectPath, td.wontProductFileCount, len(gotProducts))
		}
		if len(gotTests) != td.wontTestFileCount {
			t.Errorf("%s: file count (test codes) did not match, wont %d, got %d", td.projectPath, td.wontTestFileCount, len(gotTests))
		}
		if matchSlice(gotProducts, td.wontProductFiles) {
			t.Errorf("%s: product files did not match, wont %v, got %v", td.projectPath, td.wontProductFiles, gotProducts)
		}
		if matchSlice(gotTests, td.wontTestFiles) {
			t.Errorf("%s: test files did not match, wont %v, got %v", td.projectPath, td.wontTestFiles, gotTests)
		}
	}
}
