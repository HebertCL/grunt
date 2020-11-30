package main

import (
	"io/ioutil"
	"log"

	"github.com/HebertCL/grunt/pkg/gcp"
)

func main() {
	creds, err := ioutil.ReadFile("/Users/hebertcl/.gsutil/terraform-crabi-dev.json")
	if err != nil {
		log.Fatal(err)
	}

	client, err := gcp.NewComputeClient(creds)
	if err != nil {
		log.Fatal(err)
	}

	svc := &gcp.ComputeService{
		Service:   client,
		ProjectID: "crabi-dev",
		Fields:    "",
		Zone:      "us-west2-a",
	}

	svc.ListInstances()
}
