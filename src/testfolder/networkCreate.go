package main

import (
	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
)

func   computeNetworkParams(networkName string, description string,
	computeService *compute.Service, ctx context.Context, mtu int64) *response{
	paramsComputeNetwork := &compute.Network{
		Name: networkName,
		AutoCreateSubnetworks: false,
		Mtu: mtu,
		RoutingConfig: &compute.NetworkRoutingConfig{
			RoutingMode: "REGIONAL",
		},
		Description: description,
		ForceSendFields: []string{"AutoCreateSubnetworks"},
	}
	resp, err := computeService.Networks.Insert(projectName, paramsComputeNetwork).Context(ctx).Do()
	if err != nil {
		return MarshallUnmarshallCreationResponse(err, true)
	} else {
		responseNetCreation := MarshallUnmarshallCreationResponse(resp,false)
		statusCheck := OperationStatus(ctx,computeService, projectName,
			responseNetCreation.Id, "creation of network", "global","","")
			return statusCheck
	}
}

func NetworkCreate(ctx context.Context,computeService *compute.Service,
	description string) *response  {
	paramsComputeNetwork := computeNetworkParams(networkName,description,
		computeService, ctx, 1460)
	return paramsComputeNetwork
}
func CreatenetSubNetMachine(ctx context.Context, x chan *response,
	networkName string, networkCreate bool, subnetworkCreate bool,
	instancesCreate bool, projectName string) {
	computeService,_ := createComputeService(ctx)
	if !networkCreate {
		errorMessage := "Network creation is not enabled"
		a := ResponseCreation(&errorMessage, true)
		x <- a
	}
	netWorkCreateResponse := NetworkCreate(ctx,computeService,
		"creating the network")
	if(netWorkCreateResponse.isError) {
		x<- netWorkCreateResponse
	}
	if(!subnetworkCreate) {
		errorMessage := "Network creation is done but subnetwork is not enabled"
		subnetworkCreateResponse:= ResponseCreation(&errorMessage, true)
		x<- subnetworkCreateResponse
	}
	finished1 := make(chan *response)
	finished2 := make(chan *response)
	go SubnetworkCreation(ctx, computeService,
		networkName,"subnet creation",
		"10.0.1.0/24",subnetworkName,projectName, regionName,
		instancesCreate,netWorkCreateResponse,finished1)
	go SubnetworkCreation(ctx, computeService,
		networkName,"subnet creation",
		"10.0.2.0/25",subnetworkNametwo,projectName, regionName,
		instancesCreate,netWorkCreateResponse,finished2)
	response1 := <-finished1
	response2 := <-finished2
	if response1.isError {
		x <- response1
	}else if response2.isError {
		x <- response2
	} else {
		x <- &response{
			isError: false,
			Message: "Everything is done",
		}
	}
}
func subnetworkParams(ctx context.Context, computeService *compute.Service,
	networkName string, description string, IP string,
	name string, projectName string, regionName string) *response{
	subnetparams := &compute.Subnetwork{
				IpCidrRange: IP,
				Name: name,
				Description: description,
				PrivateIpGoogleAccess: false,
				EnableFlowLogs: false,
				Network: "projects/"  + projectName + "/global/networks/"+ networkName,
			}
	resp, err := computeService.Subnetworks.Insert(projectName, regionName,subnetparams).
		Context(ctx).Do()
	if err != nil {
		return MarshallUnmarshallCreationResponse(err, true)
	} else {
		responseNetCreation := MarshallUnmarshallCreationResponse(resp,false)
		statusCheck := OperationStatus(ctx,computeService, projectName,
			responseNetCreation.Id, "creation of network", "region",
			regionName,"")
		return statusCheck
	}
}

func SubnetworkCreation(ctx context.Context, computeService *compute.Service,
	networkName string, description string, IP string, subnetworkName string,
	projectname string, regionName string,instancesCreate bool, netWorkCreateResponse *response,x chan *response) {
	paramsComputeNetwork := subnetworkParams(ctx, computeService,networkName,description,
		IP,subnetworkName,projectname, regionName)
	if(paramsComputeNetwork.isError) {
		x <-paramsComputeNetwork
		//return paramsComputeNetwork
	}
	if !instancesCreate {
		errorMessage := "Network creation and subnetwork creation is done but instance is not created as it is not enabled"
		res := ResponseCreation(&errorMessage, true)
		x <-res
	}
	instanceCreationResponse := Createinstance(ctx, computeService, "instancecreation" + subnetworkName, netWorkCreateResponse.TargetLink,
		paramsComputeNetwork.TargetLink,projectName,
		zoneValue)
	if(instanceCreationResponse.isError) {
		x <-instanceCreationResponse
	}
	errorMessage := ""
	r := ResponseCreation(&errorMessage, false)
	x <-r
}