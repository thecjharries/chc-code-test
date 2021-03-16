# Random Notes

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)
[![Check the NOTICE](https://img.shields.io/badge/Check%20the-NOTICE-420C3B.svg)](./NOTICE)

## Before Starting

### Motivation

#### Python vs Go

I could be done quickly if I chose Python. It's a nested dict. I want to learn something new from this experience.

#### Primary Goals

1. I want to learn how to load YAML with Go. I do JSON all the time; YAML should be, in theory, identical.
2. I'm curious how initializing will work during unmarshaling. I like to build tree nodes with a depth member for convenience. As I start, I don't know if that's possible or not. I'm very curious to learn.

### Repo Contents

I populated the repo with copypasta from my other work. It's common for me to start like this; I prefer working from skeletons when possible.

## During

### Assumptions

* Ignoring git flow because I'm being lazy and there's no need at the start.
* `Employee.Salary` is going to be an `int` because I'm lazy and the example was a whole number. Changing the type doesn't really matter for what I'm doing.
* Reports are trees not cyclic graphs. eg if Alice reports to Bob and Bob reports to Charlie, Charlie cannot report to Alice.

### General Notes

* Godoc is unlike the typical Javadoc paradigm other languages (including Python) employ. I'm not super comfortable with it because I haven't compiled and used its output yet.
* I'm writing this in GoLand which has a lot of syntactic sugar and QoL functionality. I could write this in and for `vim`. Or I could use an IDE I pay for to make the process very simple. I'm not getting paid for this so I'm not making it hard on myself.

## Post-Mortem

### Primary Goals

1. Yeah, working with YAML is just as easy as working with JSON
2. I sidestepped this by seeding the graph print with depth. There might be a way to pass some builder function to some `Unmarshal` function but I was too lazy to mess with that.

### General Notes

* I didn't test any of the output. Normally I separate strings from printing, eg I'd have a `String` method and a `Print` method so I can verify `String` is proper and just ensure `Print` is doing the right system calls. I was too lazy and tbh I forgot about `String` until right now, writing my post-mortem.
* This was a pleasant and easy exercise to do.
* I don't do a lot of documentation inside of functions. There's pros and cons to it. I try to keep my functions small enough that you don't need to document internally. If you're doing some complicated logic, it makes sense to break it down. I'm doing Data Structures 101 here so it makes no sense.
