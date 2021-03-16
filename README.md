# `chc-code-test`

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)
[![Check the NOTICE](https://img.shields.io/badge/Check%20the-NOTICE-420C3B.svg)](./NOTICE)

## Overview

This contains my short stab at solving the prompt below. See Usage for running the code; check out [my notes](./notes.md) for some analysis.

## Usage

This should work in theory. I'm too lazy to test the full initialization process from scratch.

```shell
# clone the repo
git clone git@github.com:thecjharries/chc-code-test.git
cd chc-code-test
# ensures you have the deps
make bootstrap
# runs the code (read it to find out what that does)
make run
# shows test coverage (only slightly gamed)
make coverage
```

You'll want it in your `GOPATH` or you should know how to set up things so you don't have to be in your `GOPATH`.

## Prompt

### Background

There are many problems in programming which involve collections of objects which are very similar but differ in their containing ability. For example, XML has elements which are standalone and others which may contain child elements. Another common case is that of an employee hierarchy. All people in a company are employees, yet some are managers and so, in a sense, contain collections of other employees.

### Problem Statement

For the following office hierarchy, provide concrete implementations of Manager and Employee. Write a test program which does two things:
Print out an ASCII employee tree (any format you want)
Print out the total salary requirements for the entire company.

**Extra Credit**: Sort employees alphabetically (hint: there is an easy way to do this)

**Mega Extra Credit**: Build unit tests for the code

**SuperMega Extra Credit**: Create the hierarchy by reading in from a properties file or json/yaml file

Please use copious comments and document fully all your assumptions and reasoning. It’s better to have a clear, documented approach that doesn’t compile than to have something that works but is hacked up and undocumented.

### Sample output

```
Jeff
    Employees of: Jeff
Dave
    Employees of: Dave
        Andy
        Dan
        Jason
        Rick
        Suzanne

Total salary: 275000
```
