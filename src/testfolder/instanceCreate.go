package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
)
func createInstanceParams(ctx context.Context, computeService *compute.Service, name string,
	networkName string, subNetwork string, projectName string, zone string) *response{
	fmt.Println("projects/" + projectName + "/zones/" + zone  + "/machineTypes/n1-standard-1")
	rb := &compute.Instance{
		MachineType: "projects/" + projectName + "/zones/" + zone  + "/machineTypes/n1-standard-1",
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
	resp, err := computeService.Instances.Insert(projectName, zone, rb).Context(ctx).Do()
	if err != nil {
		return MarshallUnmarshallCreationResponse(err, true)
	} else {
		responseNetCreation := MarshallUnmarshallCreationResponse(resp,false)
		statusCheck := OperationStatus(ctx,computeService, projectName,
			responseNetCreation.Id, "creation of instances", "zonal","",zone)
		return statusCheck
	}

}

// createServiceAccount creates a service account.
func 	Createinstance( ctx context.Context, computeService *compute.Service, name string,
	networkName string, subNetwork string, projectName string, zone string) *response {
	instanceCreation := createInstanceParams(ctx, computeService, name, networkName, subNetwork, projectName, zone)
	return instanceCreation
}