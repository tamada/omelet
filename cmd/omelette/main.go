package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/tamada/omelette"
)

const VERSION = "1.1.0"

type filterOpts struct {
	includes string
	excludes string
}

type configOpts struct {
	deleteTempFileFlag bool
	verboseFlag        bool
	filter             *filterOpts
}

type projectOpts struct {
	productDir string
	testDir    string
}

type options struct {
	noCoverageFlag bool
	helpFlag       bool
	classpaths     []string
	project        *projectOpts
	config         *configOpts
}

func existsDirectory(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func (opts *options) validate(args []string) error {
	if len(args) == 0 && !opts.project.availableDirs() {
		return fmt.Errorf("product-code and test-code options are mandatory if no arguments")
	}
	if len(args) == 1 {
		if !opts.project.availableDirs() && !existsDirectory(args[0]) {
			return fmt.Errorf("%s: directory not found", args[0])
		}
	}
	if len(args) > 1 {
		return fmt.Errorf("%v: too many arguments", args)
	}
	return nil
}

func (configOpts *configOpts) buildConfig() (*omelette.Config, error) {
	filter, err := omelette.NewFilter(configOpts.filter.includes, configOpts.filter.excludes)
	if err != nil {
		return nil, err
	}
	return omelette.NewConfig(filter, configOpts.deleteTempFileFlag, configOpts.verboseFlag), nil
}

func buildCoverager(opts *options, config *omelette.Config) omelette.Coverager {
	if opts.noCoverageFlag {
		return &omelette.NoCoverager{}
	}
	return omelette.NewCoverager("jacoco")
}

func (projectOpts *projectOpts) availableDirs() bool {
	return projectOpts.productDir != "" && projectOpts.testDir != "" && existsDirectory(projectOpts.productDir) && existsDirectory(projectOpts.testDir)
}

func (projectOpts *projectOpts) buildFakeProject(args []string, config *omelette.Config) (omelette.Project, error) {
	name := "unknown"
	if len(args) == 1 {
		name = args[0]
	}
	project := omelette.NewFakeProject(name, projectOpts.productDir, projectOpts.testDir)
	config.PrintIfVerbose("%s: %s", args[0], project.String())
	return project, nil
}

func (projectOpts *projectOpts) buildProject(args []string, config *omelette.Config) (omelette.Project, error) {
	if len(args) == 1 && existsDirectory(args[0]) {
		project, err := omelette.NewProject(args[0])
		if err == nil {
			config.PrintIfVerbose("%s: %s", args[0], project.String())
			return project, nil
		}
	}
	if projectOpts.availableDirs() {
		return projectOpts.buildFakeProject(args, config)
	}
	return nil, fmt.Errorf("%s: cannot build project", args[0])
}

func (opts *options) build(args []string) (omelette.TestRunner, *omelette.Config, error) {
	config, err := opts.config.buildConfig()
	if err != nil {
		return nil, config, err
	}
	project, err := opts.project.buildProject(args, config)
	if err != nil {
		return nil, config, err
	}
	runner := omelette.NewJUnit4TestRunner(project, buildCoverager(opts, config))
	return runner, config, err
}

func perform(args []string, opts *options) int {
	runner, config, err := opts.build(args)
	if err != nil {
		return printError(err, 3)
	}
	for _, cp := range opts.classpaths {
		runner.AppendClasspath(cp)
	}
	if err := runner.Perform(config); err != nil {
		return printError(err, 4)
	}
	runner.Postprocess(config)
	return 0
}

func helpMessage() string {
	return fmt.Sprintf(`omelette version %s
omelette [OPTIONS] <PROJECT_DIR>
    or
omelette [OPTIONS] -p <PRODUCT_CODE_DIR> -t <TEST_CODE_DIR> [PROJECT_NAME]

OPTIONS
    -c, --classpath <PATH>      specifies classpath list separated with a colon, or defines several options.
    -d, --delete-tempfiles      deletes temporary files after running.
    -e, --excludes <REGEXP>     specifies target exclusion rules for unit tests. Default is "" (no filtering).
    -i, --includes <REGEXP>     specifies target inclusion rules for unit tests. Default is "" (no filtering).
    -n, --no-coverage           calculates no coverage of test codes.
    -p, --product-code <DIR>    specifies the directory contains the product codes.
    -t, --test-code <DIR>       specifies the directory contains the test codes.
    -v, --verbose               verbose mode.

    -h, --help                  prints this message.
ARGUMENTS
    PROJECT_DIR                 specifies the directory contains the product codes and the unit test codes.
    PROJECT_NAME                specifies the project name for destination file. Default is "unknown".`, VERSION)
}

func printError(err error, statusCode int) int {
	if err == nil {
		return 0
	}
	fmt.Println(err.Error())
	return statusCode
}

func buildFlagSet() (*flag.FlagSet, *options) {
	opts := &options{project: &projectOpts{}, config: &configOpts{filter: &filterOpts{}}}
	flags := flag.NewFlagSet("omelette", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage()) }
	flags.StringVarP(&opts.config.filter.includes, "includes", "i", "", "specifies target includesion rules for unit test codes")
	flags.StringVarP(&opts.config.filter.excludes, "excludes", "e", "", "specifies target excludesion rules for unit test codes")
	flags.StringVarP(&opts.project.productDir, "product-code", "p", "", "specifies the directory contains the product codes.")
	flags.StringVarP(&opts.project.testDir, "test-code", "t", "", "specifies the directory contains the test codes.")
	flags.StringSliceVarP(&opts.classpaths, "classpath", "c", []string{}, "specifies classpath separated with colon")
	flags.BoolVarP(&opts.config.deleteTempFileFlag, "delete-tempfiles", "d", false, "deletes temporary files after running")
	flags.BoolVarP(&opts.config.verboseFlag, "verbose", "v", false, "verbose mode")
	flags.BoolVarP(&opts.noCoverageFlag, "no-coverage", "n", false, "calculats no coverage of test codes")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "prints this message")
	return flags, opts
}

func parseOptions(args []string) ([]string, *options, error) {
	flags, opts := buildFlagSet()
	err := flags.Parse(args)
	if err != nil {
		return nil, nil, err
	}
	return flags.Args()[1:], opts, nil
}

func goMain(args []string) int {
	args, opts, err := parseOptions(args)
	if err != nil {
		return printError(fmt.Errorf("args parse error: %s", err.Error()), 1)
	}
	if opts.helpFlag {
		return printError(fmt.Errorf(helpMessage()), 0)
	}
	if err := opts.validate(args); err != nil {
		status := printError(err, 2)
		fmt.Println(helpMessage())
		return status
	}
	return perform(args, opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
