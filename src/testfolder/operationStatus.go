package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
	"time"
)


func operationInfo(ctx context.Context,computeService *compute.Service,
	globalRegion string,project string, operation string, regionName string, zoneName string) (*compute.Operation,error) {
	if globalRegion == "global" {
		return computeService.GlobalOperations.Get(project, operation).Context(ctx).Do()
	} else if globalRegion == "region"  {
		return computeService.RegionOperations.Get(project, regionName,
			operation).Context(ctx).Do()
	} else {
		return computeService.ZoneOperations.Get(project, zoneName,
			operation).Context(ctx).Do()
	}
}
const timeInterval = 	5 * 60* time.Second
func OperationStatus(ctx context.Context,computeService *compute.Service,
	projectId string, resourceId string, message string, globalRegion string,
	regionName string,zoneName string  )  *response {
	timeout := time.After(timeInterval)
	tick := time.Tick(5 * time.Second)
	for {
		resp,err :=operationInfo(ctx,computeService,globalRegion,projectId,resourceId,
			regionName,zoneName)
		aa,_ := json.Marshal(resp)
		fmt.Println(string(aa))
		select {
		case <-timeout:
			errorMessage := "timeout happened while " + message
			return ResponseCreation(&errorMessage, true)
		case <-tick:
			if err != nil {
				return MarshallUnmarshallCreationResponse(err, true)
			} else if resp.Status == "DONE" && resp.Progress == 100 {
				return MarshallUnmarshallCreationResponse(resp,false)
			}

		}

	}
}
