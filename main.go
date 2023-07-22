package main

import (
	"fmt"
	"os"
	"sonarqube-goclient/rest"
)

func main() {
	user := os.Getenv("SONARQUBE_USER")
	pass := os.Getenv("SONARQUBE_PASS")
	url := os.Getenv("SONARQUBE_URL")
	c := rest.NewClient(url, user, pass)
	param := rest.ProjectStatusOptions{}
	list, _ := c.QualityGates.ProjectStatus(&param)
	fmt.Printf("%#v", list)
}
