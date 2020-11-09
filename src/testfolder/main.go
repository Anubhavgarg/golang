
package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/iam/v1"
	"log"
)

	const projectName = "sdfdassasa"
	const displayName = "budifa"
	const project = "western-notch-185412"
const zone = "us-central1-a"

type ada struct {
	Description string `json:"description,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
// createServiceAccount creates a service account.
func createinstance( yahoo chan bool, name string) {
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
				&compute.AttachedDisk{
					Type: "PERSISTENT",
					Boot: true,
					Mode: "READ_WRITE",
					AutoDelete: true,
					DeviceName: "instance-1",
					InitializeParams: &compute.AttachedDiskInitializeParams{
						SourceImage: "projects/ubuntu-os-cloud/global/images/ubuntu-1604-xenial-v20201014",
						DiskType: "projects/western-notch-185412/zones/us-central1-a/diskTypes/pd-standard",
						DiskSizeGb: 10,
					},
					DiskEncryptionKey: &compute.CustomerEncryptionKey{},
				},
			},
			NetworkInterfaces: []*compute.NetworkInterface{
				&compute.NetworkInterface{
					Subnetwork: "projects/western-notch-185412/regions/us-central1/subnetworks/default",
					AccessConfigs: []*compute.AccessConfig{
						&compute.AccessConfig{
							Name: "External NAT",
							Type: "ONE_TO_ONE_NAT",
							NetworkTier: "PREMIUM",
						},
					},
					AliasIpRanges: nil,
				},
			},
		}
	data,err := computeService.Instances.Insert(project, zone, rb).Context(ctx).Do()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(data)
	yahoo <- true

}

func createServiceAccount(finished chan bool, projectID, name, displayName string) (*iam.ServiceAccount, error) {
	ctx := context.Background()
	service, err := iam.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("iam.NewService: %v", err)
	}

	request := &iam.CreateServiceAccountRequest{
		AccountId: name,
		ServiceAccount: &iam.ServiceAccount{
			DisplayName: displayName,
		},
	}
	account, err := service.Projects.ServiceAccounts.Create("projects/"+projectID, request).Do()
	res1B, _ := json.Marshal(account)
	fmt.Println("TEST" + string(res1B))
	if err != nil {
		return nil, fmt.Errorf("Projects.ServiceAccounts.Create: %v", err)
	}
	createinstance(finished, name)
	return account, nil
}
func main() {
	finished := make(chan bool)
	finished1 := make(chan bool)
	go createinstance(finished,"testing")
	 go createServiceAccount(finished1,project,projectName,displayName)
	//a, err :=  createServiceAccount(finished1,project,projectName,displayName)
	//res1B, _ := json.Marshal(a)
	//fmt.Println("TEST" + string(res1B))
	//res := ada{}
	//json.Unmarshal([]byte(string(res1B)), &res)
	<-finished
	<-finished1
}