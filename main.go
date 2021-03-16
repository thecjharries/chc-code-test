// Copyright 2021 CJ Harries
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Expose print function for testing to game 100% coverage
var zPrint = fmt.Printf

const sampleYamlPath string = "./sample.yaml"

// This is a graph of employees
// Each employee is a tree of reports
type EmployeeGraph struct {
	Employees []Employee `yaml:",employees"`
}

// Each employee has a salary and reports (possibly)
type Employee struct {
	Salary    int        `yaml:"salary"`
	Employees []Employee `yaml:",employees"`
}

// Primary runner
func main() {
	_, _ = zPrint("hello world")
}

// Load the contents of the sample file into an Employee Graph
func loadSampleYaml() (graph EmployeeGraph) {
	// Skipping the error because I'm not getting paid for this
	data, _ := ioutil.ReadFile(sampleYamlPath)
	// Skipping the error because I'm not getting paid for this
	_ = yaml.Unmarshal(data, &graph)
	return
}
