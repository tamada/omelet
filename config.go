package omelet

import (
	"log"
	"os"
)

const OMELET_HOME_ENV_NAME = "OMELET_HOME"

type Config struct {
	filter         Filter
	currentDir     string
	removeTempFile bool
	home           string
	verbose        bool
}

func NewConfig(filter Filter, removeTempFileFlag bool, verboseFlag bool) *Config {
	config := &Config{filter: filter, removeTempFile: removeTempFileFlag, verbose: verboseFlag}
	wd, _ := os.Getwd()
	config.currentDir = wd
	config.home = getOmeletHome()
	return config
}

func (config *Config) Filter(targets []string) []string {
	return config.filter.Filter(targets)
}

func (config *Config) CurrentDir() string {
	return config.currentDir
}

func (config *Config) OmeletHome() string {
	return config.home
}

func (config *Config) PrintIfVerbose(formatter string, args ...interface{}) {
	if config.verbose {
		log.Printf(formatter, args...)
	}
}

func getOmeletHome() string {
	path := os.Getenv(OMELET_HOME_ENV_NAME)
	if path != "" && ExistsFile(path) {
		return path
	}
	wd, _ := os.Getwd()
	return wd
}

func ExistsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
}
