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
	opt := rest.SearchOptions{GateName: "Critical Issue"}
	search, _ := c.QualityGates.Search(&opt)
	fmt.Printf("%#v", search)
}
