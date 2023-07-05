package main

import (
	"fmt"
	"os"
	"sonarqube-goclient/client"
	"sonarqube-goclient/rest/qualitygates"
)

func main() {
	user := os.Getenv("SONARQUBE_USER")
	pass := os.Getenv("SONARQUBE_PASS")
	url := os.Getenv("SONARQUBE_URL")
	c := client.NewClient(url, user, pass)
	rest := qualitygates.ApiQualityGates{c}
	opt := qualitygates.QualityGateOptions{Project: "awf-robot"}
	gate, _ := rest.GetByProject(&opt)

	fmt.Printf("%s", gate.QualityGate.Id)
}
