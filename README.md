[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/omelette/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/omelette/releases/tag/v1.0.0)


# Omelette

An agent for running the unit tests on the CLI environment for the Java platform.

## Usage

```sh
omelette [OPTIONS] <PROJECT_DIRs...>
OPTIONS
    -c, --classpath <PATH>      specifies classpath separated with colon.
    -d, --delete-tempfiles      delete temporary files after running.
    -e, --excludes <REGEXP>     specifies target exclusion rules for unit tests. Default is "" (no filtering).
    -i, --includes <REGEXP>     specifies target inclusion rules for unit tests. Default is "" (no filtering).
    -n, --no-coverage           no calculating coverage of the unit test codes.
    -v, --verbose               verbose mode.

    -h, --help                  print this message.
ARGUMENTS
    PRODUCT_CODES_DIR           specifies the directory contains the product codes.
    TEST_CODES_DIR              specifies the directory contains the unit test codes.
```


## About

Icons made by [Nhor Phai](https://www.flaticon.com/authors/nhor-phai) from [www.flaticon.com](https://www.flaticon.com/)
