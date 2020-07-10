package omelette

import "regexp"

type Filter interface {
	Filter(targets []string) []string
}

type noFilter struct {
}

func (filter *noFilter) Filter(targets []string) []string {
	return targets
}

type multipleFilter struct {
	filters []Filter
}

func NewFilter(includesPattern, excludesPattern string) (Filter, error) {
	filter := new(multipleFilter)
	iFilter, err1 := NewIncludesFilter(includesPattern)
	eFilter, err2 := NewExcludesFilter(excludesPattern)
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	filter.filters = []Filter{iFilter, eFilter}
	return filter, nil
}

func (mFilter *multipleFilter) Filter(targets []string) []string {
	results := targets
	for _, filter := range mFilter.filters {
		results = filter.Filter(results)
	}
	return results
}

type IncludesFilter struct {
	origin  string
	pattern *regexp.Regexp
}

func NewIncludesFilter(pattern string) (Filter, error) {
	if pattern == "" {
		return new(noFilter), nil
	}
	regexp, err := regexp.Compile(pattern)
	return &IncludesFilter{origin: pattern, pattern: regexp}, err
}

func (filter *IncludesFilter) Filter(targets []string) []string {
	results := []string{}
	for _, item := range targets {
		if filter.pattern.Match([]byte(item)) {
			results = append(results, item)
		}
	}
	return results
}

type ExcludesFilter struct {
	origin  string
	pattern *regexp.Regexp
}

func NewExcludesFilter(pattern string) (Filter, error) {
	if pattern == "" {
		return new(noFilter), nil
	}
	regexp, err := regexp.Compile(pattern)
	return &ExcludesFilter{origin: pattern, pattern: regexp}, err
}

func (filter *ExcludesFilter) Filter(targets []string) []string {
	results := []string{}
	for _, item := range targets {
		if !filter.pattern.Match([]byte(item)) {
			results = append(results, item)
		}
	}
	return results
}
