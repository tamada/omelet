---
title: ":fork_and_knife: Usage"
date: 2020-04-14
---

## :paperclip: CLI

```sh
$ omelet --help
omelet version 1.0.0
omelet [OPTIONS] <PROJECT_DIRs...>
OPTIONS
    -c, --classpath <PATH>      specifies classpath separated with colon.
    -d, --delete-tempfiles      delete temporary files after running.
    -e, --excludes <REGEXP>     specifies target exclusion rules for unit tests. Default is "" (no filtering).
    -i, --includes <REGEXP>     specifies target inclusion rules for unit tests. Default is "" (no filtering).
    -n, --no-coverage           no calculating coverage of test codes.
    -v, --verbose               verbose mode.

    -h, --help                  print this message.
ARGUMENTS
    PRODUCT_CODES_DIR           specifies the directory contains the product codes.
    TEST_CODES_DIR              specifies the directory contains the unit test codes.
```
