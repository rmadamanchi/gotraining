package main

import (
	"fmt"
	"math/rand"
)

/* Developer */
type Developer struct {
	name string
}

func (d *Developer) develop(system string) {
	fmt.Printf("%s is developing %s\n", d.name, system)
}

/* Deployer */
type Deployer struct {
	name string
}

func (d *Deployer) deploy(system string, environment string) {
	fmt.Printf("%s is deploying %s to %s\n", d.name, system, environment)
}

/* Specialist */
type Specialist struct {
	Developer
	Deployer
}

/* main */

var developers = []string{"John", "Mary", "Vince"}
var deployers = []string{"Tina", "Adam", "Emma"}
var specialists = []string{"Steve", "Alicia"}

func main() {
	var systems = []string{"saturn", "neptune", "pluto"}
	var secretsytems = []string{"mars", "venus"}
	var environments = []string{"dev", "staging", "prod"}

	for _, system := range systems {
		fmt.Printf("---Starting System %s\n", system)
		developer := FindDeveloper()
		developer.develop(system)

		for _, environment := range environments {
			deployer := FindDeployer()
			deployer.deploy(system, environment)
		}
		fmt.Printf("---Completed System %s\n", system)
	}

	for _, secretsystem := range secretsytems {
		fmt.Printf("---Starting Secret System %s\n", secretsystem)
		specialist := FindSpecialist()
		specialist.develop(secretsystem)

		for _, environment := range environments {
			specialist.deploy(secretsystem, environment)
		}
		fmt.Printf("---Completed Secret System %s\n", secretsystem)
	}
}

func FindDeveloper() Developer {
	name := developers[rand.Intn(len(developers))]
	return Developer{name}
}

func FindDeployer() Deployer {
	name := deployers[rand.Intn(len(deployers))]
	return Deployer{name}
}

func FindSpecialist() Specialist {
	name := specialists[rand.Intn(len(specialists))]
	return Specialist{name}
}
