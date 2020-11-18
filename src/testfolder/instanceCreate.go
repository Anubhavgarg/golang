package main

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"log"
)
// createServiceAccount creates a service account.
func 	createinstance( yahoo chan bool, name string,
	networkName string, subNetwork string) {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}
	rb := &compute.Instance{
		MachineType: "projects/western-notch-185412/zones/us-central1-a/machineTypes/n1-standard-1",
		Name: name,
		Disks: []*compute.AttachedDisk{
			{
				Type:       "PERSISTENT",
				Boot:       true,
				Mode:       "READ_WRITE",
				AutoDelete: true,
				DeviceName: "instance-1",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					SourceImage: "projects/ubuntu-os-cloud/global/images/ubuntu-1604-xenial-v20201014",
					DiskType:    "projects/western-notch-185412/zones/us-central1-a/diskTypes/pd-standard",
					DiskSizeGb:  10,
				},
				DiskEncryptionKey: &compute.CustomerEncryptionKey{},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			&compute.NetworkInterface{
				Network: networkName,
				Subnetwork: subNetwork,
			},
		},
	}
	_, err = computeService.Instances.Insert(project, zone, rb).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}
	yahoo <- true

}