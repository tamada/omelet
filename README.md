[![codebeat badge](https://codebeat.co/badges/23134092-de46-44aa-942c-5d4a070eaf3c)](https://codebeat.co/projects/github-com-tamada-omelette-master)
[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/omelette/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.1.0-yellowgreen.svg)](https://github.com/tamada/omelette/releases/tag/v1.1.0)

# Omelette

An agent for running the unit tests on the CLI environment for the Java platform.

## :speaking_head: Overview

In the Java platform, we usually run the unit tests through some build tool, such as [Maven](https://maven.apache.org), [Gradle](https://gradle.org), and so on.
However, it is hard to run the unit tests on the CLI environment.

Calculating test coverages is also complicated for novice programmers.
Since, it requires understanding how to use several libraries (unit test library, coverage measuring library, and the build tool).

Running the unit tests and calculating test coverages usually requires software projects.
It is generally tiresome for calculating test coverages of toy programs with their unit tests.

Therefore, we developed `omelette` for running unit tests and calculating test coverages in the CLI environment for the Java platform.


## :runner: Usage

```sh
omelette version 1.0.0
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
    PROJECT_NAME                specifies the project name for destination file. Default is "unknown".
```

### :briefcase: Requirements

* Runtime
    * bash completion 2.x or later.
    * Java 10 or later.
        * [JUnit 4](https://junit.org/junit4/) 4.13, [hamcrest-all](https://mvnrepository.com/artifact/org.hamcrest/hamcrest-all) 1.3
        * [JaCoCo](https://www.eclemma.org/jacoco/) 0.8.5
* Development
    * Go lang 10.x or later.
    * Dependent Libraries
        * [pflag](https://github.com/spf13/pflag) v1.0.5

## :anchor: Install

### :beer: Homebrew (macOS)

```sh
brew install tamada/brew/omelette
```

### Go lang

```
go get github.com/tamada/omelette
```

After downloading `omelette`, run the following script.

```sh
cd ~/go/src/github.com/tamada/omelette; ./bin/download_dependencies.sh
```

### :hammer_and_wrench: Build from source

```sh
git clone https://github.com/tamada/omelette.git
cd omelette
make
```

## :smile: About

### :scroll: License

[Do What The F*ck You Want To Public License](https://github.com/tamada/omelette/blob/master/LICENSE)

[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/omelette/blob/master/LICENSE)

* This license permits
    * :+1: Commercial use,
    * :+1: Modification,
    * :+1: Distribution, and
    * :+1: Private use.

### :man_office_worker: Developers :woman_office_worker:

* [Haruaki Tamada](https://github.com/tamada) [:globe_with_meridians:](https://tamada.github.io)

## :question: Why does the product names omelette?

Because the lunch was omelette, when I developed this product.

### :handshake: Attributions

Icons made by [Nhor Phai](https://www.flaticon.com/authors/nhor-phai) from [www.flaticon.com](https://www.flaticon.com/).
