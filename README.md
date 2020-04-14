[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/omlet/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/omlet/releases/tag/v1.0.0)


# Omlet

Java Unit test Running Agent on CLI for the Java platform.

## Usage

```sh
omlet [OPTIONS] <PROJECT_DIR>
OPTIONS
    -d, --delete-temfiles     delete temporary files after running.
    -e, --excludes <REGEXP>   specifies target exclusion rules for unit tests. Default is "" (no filtering).
    -i, --includes <REGEXP>   specifies target inclusion rules for unit tests. Default is "" (no filtering).
    -n, --no-coverage         no calculating coverage of tset codes.
    -s, --source <DIR>        specifies the directory contains product codes.
                              Default is ${PROJECT_DIR}/src/main/java.
    -t, --unit-test <DIR>     specifies the directory contains unit test codes.
                              Default is ${PROJECT_DIR}/src/test/java.

    -h, --help                print this message.
ARGUMENTS
    PRODUCT_CODES_DIR         specifies the directory contains the product codes.
    TEST_CODES_DIR            specifies the directory contains the unit test codes.
```

## どうやって実現する？？？

* カバレッジは jacoco（[Java Agent](https://www.jacoco.org/jacoco/trunk/doc/agent.html)）で計測する．
    * includes, excludes などのオプションを生成する．
* Apache Ivy で依存jarファイルをダウンロードする．
