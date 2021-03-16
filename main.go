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
	"strings"
)

// Expose print function for testing to game 100% coverage
var zPrint = fmt.Printf

const sampleYamlPath string = "./sample.yaml"

// This is a graph of employees
// Each employee is a tree of reports
type EmployeeGraph struct {
	Employees []Employee `yaml:"employees"`
}

// Walk and print the employee graph while collecting total salary
func (g *EmployeeGraph) Print() {
	salary := 0
	for _, employee := range g.Employees {
		salary += employee.PrintAndCollectSalary(0)
	}
	_, _ = zPrint(fmt.Sprintf("Total Salary: %d\n", salary))
}

// Each employee has a salary and reports (possibly)
type Employee struct {
	Name      string     `yaml:"name"`
	Salary    int        `yaml:"salary"`
	Employees []Employee `yaml:"employees"`
}

// Print the employee name, their employees, and aggregate salaries
// Note that I'm not forcing depth to be >= 0; I'm being lazy
func (e *Employee) PrintAndCollectSalary(depth int) int {
	_, _ = zPrint(fmt.Sprintf("%s%s\n", strings.Repeat("  ", depth), e.Name))
	salary := e.Salary
	if 0 < len(e.Employees) {
		_, _ = zPrint(fmt.Sprintf("%sEmployees of %s\n", strings.Repeat("  ", depth + 1), e.Name))
	}
	for _, employee := range e.Employees {
		salary += employee.PrintAndCollectSalary(depth + 2)
	}
	return salary
}

// Primary runner
func main() {
	sampleEmployeeGraph := loadSampleYaml()
	sampleEmployeeGraph.Print()
}

// Load the contents of the sample file into an Employee Graph
func loadSampleYaml() (graph EmployeeGraph) {
	// Skipping the error because I'm not getting paid for this
	data, _ := ioutil.ReadFile(sampleYamlPath)
	// Skipping the error because I'm not getting paid for this
	_ = yaml.Unmarshal(data, &graph)
	return
}
