---
title: ":runner: Usage"
date: 2020-04-14
---

## :paperclip: CLI

```sh
$ omelette --help
omelette version 1.0.0
omelette [OPTIONS] <PROJECT_DIR>
    or
omelette [OPTIONS] -p <PRODUCT_CODE_DIR> -t <TEST_CODE_DIR> [PROJECT_NAME]
OPTIONS
    -c, --classpath <PATH>      specifies classpath separated with colon.
    -d, --delete-tempfiles      delete temporary files after running.
    -e, --excludes <REGEXP>     specifies target exclusion rules for unit tests. Default is "" (no filtering).
    -i, --includes <REGEXP>     specifies target inclusion rules for unit tests. Default is "" (no filtering).
    -n, --no-coverage           no calculating coverage of test codes.
    -p, --product-code <DIR>    specifies the directory contains the product codes.
    -t, --test-code <DIR>       specifies the directory contains the test codes.
    -v, --verbose               verbose mode.

    -h, --help                  print this message.
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
