package main

import (
	"fmt"
	"math/rand"
)

/* Interfaces */
type developer interface {
	develop(system string)
}

type deployer interface {
	deploy(system string, environment string)
}

type specialist interface {
	developer
	deployer
}

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

/** Specialist **/
type Specialist struct {
	name string
}

func (s *Specialist) develop(system string) {
	fmt.Printf("%s is developing %s\n", s.name, system)
}

func (s *Specialist) deploy(system string, environment string) {
	fmt.Printf("%s is deploying %s to %s\n", s.name, system, environment)
}

/* main */

var developers = []string{"John", "Mary", "Vince"}
var deployers = []string{"Tina", "Adam", "Emma"}
var specialists = []string{"Steve", "Alicia"}

func main() {
	var systems = []string{"saturn", "neptune", "pluto"}
	var specialsystems = []string{"mars", "venus"}
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

	for _, system := range specialsystems {
		fmt.Printf("---Starting Special System %s\n", system)
		specialist := FindSpecialist()
		specialist.develop(system)

		for _, environment := range environments {
			specialist.deploy(system, environment)
		}
		fmt.Printf("---Completed Special System %s\n", system)
	}
}

func FindDeveloper() developer {
	name := developers[rand.Intn(len(developers))]
	return &Developer{name}
}

func FindDeployer() deployer {
	name := deployers[rand.Intn(len(deployers))]
	return &Deployer{name}
}

func FindSpecialist() specialist {
	name := specialists[rand.Intn(len(specialists))]
	return &Specialist{name}
}
