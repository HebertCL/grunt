package gcp

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

type ComputeService struct {
	Service   *compute.Service
	ProjectID string
	Fields    string
	Zone      string
}

func NewComputeClient(credentials []byte) (*compute.Service, error) {
	ctx := context.Background()

	client := option.WithCredentialsJSON(credentials)

	svc, err := compute.NewService(ctx, client)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return svc, nil
}

func (cs ComputeService) ListInstances() {
	ctx := context.Background()
	req := cs.Service.Instances.List(cs.ProjectID, cs.Zone).Fields()

	if err := req.Pages(ctx, func(page *compute.InstanceList) error {
		for _, instance := range page.Items {
			fmt.Printf("%#v\n", instance)
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
