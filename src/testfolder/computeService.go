package main

import (
	"encoding/base64"
	"encoding/json"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)
type ServiceKey struct {
	TypeGoogleCredentialsFile string `json:"type"`
	ProjectID                 string `json:"project_id"`
	PrivateKeyID              string `json:"private_key_id"`
	PrivateKey                string `json:"private_key"`
	ClientEmail               string `json:"client_email"`
	ClientID                  string `json:"client_id"`
	AuthURI                   string `json:"auth_uri"`
	TokenURI                  string `json:"token_uri"`
	AuthProviderX509CertURL   string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL         string `json:"client_x509_cert_url"`
}

func createComputeService(ctx context.Context, tokenStr string,scopes []string) (*compute.Service,error) {
	raw, _ := base64.StdEncoding.DecodeString(tokenStr)
	var serviceKey *ServiceKey
	json.Unmarshal([]byte(raw), &serviceKey)
	data1, err := json.Marshal(serviceKey)
	conf, err := google.JWTConfigFromJSON(data1, scopes...)
	if err != nil {
		return nil, err
	}
	c := conf.Client(ctx)
	computeService, err := compute.New(c)
	if err != nil {
		return nil,err
	}
	return computeService,nil
}
