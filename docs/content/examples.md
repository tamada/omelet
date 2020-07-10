---
title: ":horse_racing: Examples"
date: 2020-07-10
---

## nmaker

An example for running `omelette` with no dependencies.

```sh
git clone https://github.com/tamada/nmaker.git
cd nmaker
mvn package # -> build the project.
omelette -v .  # run omelette in the verbose mode, and generate nmaker_coverage.csv in the current directory.
```


## 9rules

An example for running `omelette` with dependency libraries.

```sh
git clone https://github.com/tamada/9rules.git
cd 9rules
mvn package # -> download dependencies, copy them into target/lib directory, then build the project.
omelette -c target/lib . # -> generate 9rules_coverage.csv in the current directory.
```
