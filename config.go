package omelette

import (
	"log"
	"os"
)

const OMELETTE_HOME_ENV_NAME = "OMELETTE_HOME"

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
	config.home = getOmeletteHome()
	return config
}

func (config *Config) Filter(targets []string) []string {
	return config.filter.Filter(targets)
}

func (config *Config) CurrentDir() string {
	return config.currentDir
}

func (config *Config) OmeletteHome() string {
	return config.home
}

func (config *Config) PrintIfVerbose(formatter string, args ...interface{}) {
	if config.verbose {
		log.Printf(formatter, args...)
	}
}

func getOmeletteHome() string {
	path := os.Getenv(OMELETTE_HOME_ENV_NAME)
	if path != "" && ExistsDir(path) {
		return path
	}
	wd, _ := os.Getwd()
	return wd
}

func ExistsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func ExistsFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
}
