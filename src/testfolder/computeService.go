package main

import (
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"golang.org/x/net/context"
	"log"
)

func createComputeService(ctx context.Context) (*compute.Service,error) {
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	computeService, err := compute.New(c)
	if err != nil {
		return nil,err
	}
	return computeService,nil
}
