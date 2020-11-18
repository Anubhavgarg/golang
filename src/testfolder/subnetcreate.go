package main

import (
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
	"log"
	"time"
)

func createnet(x chan simple, name string) {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}
	project := project // TODO: Update placeholder value.
	rb := &compute.Network{
		Name: name,
		AutoCreateSubnetworks: false,
		Mtu: 1460,
		RoutingConfig: &compute.NetworkRoutingConfig{
			RoutingMode: "REGIONAL",
		},
		Description: "creating the network",
		ForceSendFields: []string{"AutoCreateSubnetworks"},
	}

	resp, err := computeService.Networks.Insert(project, rb).Context(ctx).Do()
	if err != nil {
		res1B, _ := json.Marshal(err)
		res1 := simple{
		}
		json.Unmarshal([]byte(string(res1B)), &res1)
		x <- res1
	} else {
		res1B, _ := json.Marshal(resp)
		res := simple{}
		json.Unmarshal([]byte(string(res1B)), &res)
		x <- res
	}
}
func createSubnet(x chan simple, name string, IP string, instance string) {
	ctx := context.Background()
	c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	computeService, err := compute.New(c)
	if err != nil {
		log.Fatal(err)
	}
	project := project
	region := region
	rb := &compute.Subnetwork{
		IpCidrRange: IP,
		Name: name,
		Description: "checkibng",
		PrivateIpGoogleAccess: false,
		EnableFlowLogs: false,
		Network: "projects/western-notch-185412/global/networks/"+ namea,
	}
	resp, err := computeService.Subnetworks.Insert(project, region, rb).Context(ctx).Do()

	if err != nil {
		res1B, _ := json.Marshal(err)
		res1 := simple{
		}
		json.Unmarshal([]byte(string(res1B)), &res1)
		x <- res1
	} else {
		time.Sleep(30 * time.Second)
		res1B, _ := json.Marshal(resp)
		res := simple{}
		json.Unmarshal([]byte(string(res1B)), &res)
		finished3 := make(chan bool)
		go createinstance(finished3,"testing" + name,instance,res.TargetLink)
		<-finished3
		x <- res
	}
}

