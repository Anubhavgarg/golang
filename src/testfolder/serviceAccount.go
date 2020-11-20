package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/iam/v1"
)

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
json.Marshal(account)
if err != nil {
return nil, fmt.Errorf("Projects.ServiceAccounts.Create: %v", err)
}
//createinstance(finished, name,"S","ss")
return account, nil
}
