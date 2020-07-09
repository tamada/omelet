package omelette

/*
Coverager generates arguments for computing coverages.
*/
type Coverager interface {
	Args(tr TestRunner, project Project, config *Config) string
	ToOmeletteFormat(tr TestRunner, project Project, config *Config) error
}

type NoCoverager struct {
}

func (nc *NoCoverager) Args(tr TestRunner, project Project, config *Config) string {
	return ""
}

func (nc *NoCoverager) ToOmeletteFormat(tr TestRunner, project Project, config *Config) error {
	return nil
}

func NewCoverager(kind string) Coverager {
	if kind == "jacoco" {
		return new(JacocoCoverager)
	}
	return new(NoCoverager)
}
