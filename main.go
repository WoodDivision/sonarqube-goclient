package main

import (
	"fmt"
	"os"
	"sonarqube-goclient/rest"
	"sonarqube-goclient/rest/qualitygates"
)

func main() {
	user := os.Getenv("SONARQUBE_USER")
	pass := os.Getenv("SONARQUBE_PASS")
	url := os.Getenv("SONARQUBE_URL")
	c := rest.NewClient(url, user, pass)
	opt := qualitygates.GetByProjectOptions{Project: "awf-robot"}
	gate, _ := c.QualityGates.GetByProject(&opt)

	fmt.Printf("%s", gate.QualityGate.Name)
}
