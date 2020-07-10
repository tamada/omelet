package omelette

import "testing"

func TestIncludesFilter(t *testing.T) {
	testdata := []struct {
		inputs    []string
		pattern   string
		errorFlag bool
		wontData  []string
	}{
		{[]string{"a", "ab", "abc", "abcd"}, ".*", false, []string{"a", "ab", "abc", "abcd"}},
		{[]string{"a", "ab", "abc", "abcd"}, "", false, []string{"a", "ab", "abc", "abcd"}},
		{[]string{"a", "ab", "abc", "abcd"}, "ABCD", false, []string{}},
		{[]string{"a", "ab", "abc", "abcd"}, "c", false, []string{"abc", "abcd"}},
	}

	for _, td := range testdata {
		filter, err := NewIncludesFilter(td.pattern)
		if (err != nil) != td.errorFlag {
			t.Errorf("NewIncludesFilter(%s) should (not) be error: %v", td.pattern, err)
			continue
		}
		gotData := filter.Filter(td.inputs)
		if len(gotData) != len(td.wontData) {
			t.Errorf("length of filter(%s) results did not match, wont %d, got %d", td.pattern, len(td.wontData), len(gotData))
			continue
		}
		for i := range gotData {
			if gotData[i] != td.wontData[i] {
				t.Errorf("item of gotData[%d] and wontData[%d] did not match, wont %s, got %s", i, i, td.wontData[i], gotData[i])
			}
		}
	}
}

func TestExcludesFilter(t *testing.T) {
	testdata := []struct {
		inputs    []string
		pattern   string
		errorFlag bool
		wontData  []string
	}{
		{[]string{"a", "ab", "abc", "abcd"}, ".*", false, []string{}},
		{[]string{"a", "ab", "abc", "abcd"}, "", false, []string{"a", "ab", "abc", "abcd"}},
		{[]string{"a", "ab", "abc", "abcd"}, "ABCD", false, []string{"a", "ab", "abc", "abcd"}},
		{[]string{"a", "ab", "abc", "abcd"}, "c", false, []string{"a", "ab"}},
	}

	for _, td := range testdata {
		filter, err := NewExcludesFilter(td.pattern)
		if (err != nil) != td.errorFlag {
			t.Errorf("NewExcludesFilter(%s) should (not) be error: %v", td.pattern, err)
			continue
		}
		gotData := filter.Filter(td.inputs)
		if len(gotData) != len(td.wontData) {
			t.Errorf("length of filter(%s) results did not match, wont %d, got %d", td.pattern, len(td.wontData), len(gotData))
			continue
		}
		for i := range gotData {
			if gotData[i] != td.wontData[i] {
				t.Errorf("item of gotData[%d] and wontData[%d] did not match, wont %s, got %s", i, i, td.wontData[i], gotData[i])
			}
		}
	}
}

func TestFilter(t *testing.T) {
	testdata := []struct {
		inputs    []string
		iPattern  string
		ePattern  string
		errorFlag bool
		wontData  []string
	}{
		{[]string{"a", "ab", "abc", "abcd"}, ".*", "", false, []string{"a", "ab", "abc", "abcd"}},
		{[]string{"a", "ab", "abc", "abcd"}, "b", "d", false, []string{"ab", "abc"}},
		{[]string{"a", "ab", "abc", "abcd"}, "ABCD", "a", false, []string{}},
		{[]string{"a", "ab", "abc", "abcd"}, "", "d", false, []string{"a", "ab", "abc"}},
	}

	for _, td := range testdata {
		filter, err := NewFilter(td.iPattern, td.ePattern)
		if (err != nil) != td.errorFlag {
			t.Errorf("NewFilter(%s, %s) should (not) be error: %v", td.iPattern, td.ePattern, err)
			continue
		}
		gotData := filter.Filter(td.inputs)
		if len(gotData) != len(td.wontData) {
			t.Errorf("length of filter(%s, %s) results did not match, wont %d, got %d", td.iPattern, td.ePattern, len(td.wontData), len(gotData))
			continue
		}
		for i := range gotData {
			if gotData[i] != td.wontData[i] {
				t.Errorf("item of gotData[%d] and wontData[%d] did not match, wont %s, got %s", i, i, td.wontData[i], gotData[i])
			}
		}
	}
}
