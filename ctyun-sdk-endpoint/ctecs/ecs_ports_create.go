package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsCreateApi
type EcsPortsCreateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsCreateApi(client *ctyunsdk.CtyunClient) *EcsPortsCreateApi {
	return &EcsPortsCreateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/create",
		},
	}
}

func (this *EcsPortsCreateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsCreateRequest) (*EcsPortsCreateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsCreateRealRequest{
		ClientToken:             req.ClientToken,
		RegionID:                req.RegionID,
		SubnetID:                req.SubnetID,
		PrimaryPrivateIp:        req.PrimaryPrivateIp,
		Ipv6Addresses:           req.Ipv6Addresses,
		SecurityGroupIds:        req.SecurityGroupIds,
		SecondaryPrivateIpCount: req.SecondaryPrivateIpCount,
		SecondaryPrivateIps:     req.SecondaryPrivateIps,
		Name:                    req.Name,
		Description:             req.Description,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsCreateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsCreateResponse{
		VpcID:                realResponse.VpcID,
		SubnetID:             realResponse.SubnetID,
		NetworkInterfaceID:   realResponse.NetworkInterfaceID,
		NetworkInterfaceName: realResponse.NetworkInterfaceName,
		MacAddress:           realResponse.MacAddress,
		Description:          realResponse.Description,
		Ipv6Address:          realResponse.Ipv6Address,
		SecurityGroupIds:     realResponse.SecurityGroupIds,
		SecondaryPrivateIps:  realResponse.SecondaryPrivateIps,
		PrivateIpAddress:     realResponse.PrivateIpAddress,
		InstanceOwnerID:      realResponse.InstanceOwnerID,
		InstanceType:         realResponse.InstanceType,
		InstanceID:           realResponse.InstanceID,
	}, nil
}

type EcsPortsCreateRealRequest struct {
	ClientToken             *string   `json:"clientToken,omitempty"`
	RegionID                *string   `json:"regionID,omitempty"`
	SubnetID                *string   `json:"subnetID,omitempty"`
	PrimaryPrivateIp        *string   `json:"primaryPrivateIp,omitempty"`
	Ipv6Addresses           *[]string `json:"ipv6Addresses,omitempty"`
	SecurityGroupIds        *[]string `json:"securityGroupIds,omitempty"`
	SecondaryPrivateIpCount *int      `json:"secondaryPrivateIpCount,omitempty"`
	SecondaryPrivateIps     *[]string `json:"secondaryPrivateIps,omitempty"`
	Name                    *string   `json:"name,omitempty"`
	Description             *string   `json:"description,omitempty"`
}

type EcsPortsCreateRequest struct {
	ClientToken             *string
	RegionID                *string
	SubnetID                *string
	PrimaryPrivateIp        *string
	Ipv6Addresses           *[]string
	SecurityGroupIds        *[]string
	SecondaryPrivateIpCount *int
	SecondaryPrivateIps     *[]string
	Name                    *string
	Description             *string
}

type EcsPortsCreateRealResponse struct {
	VpcID                string   `json:"vpcID,omitempty"`
	SubnetID             string   `json:"subnetID,omitempty"`
	NetworkInterfaceID   string   `json:"networkInterfaceID,omitempty"`
	NetworkInterfaceName string   `json:"networkInterfaceName,omitempty"`
	MacAddress           string   `json:"macAddress,omitempty"`
	Description          string   `json:"description,omitempty"`
	Ipv6Address          []string `json:"ipv6Address,omitempty"`
	SecurityGroupIds     []string `json:"securityGroupIds,omitempty"`
	SecondaryPrivateIps  []string `json:"secondaryPrivateIps,omitempty"`
	PrivateIpAddress     string   `json:"privateIpAddress,omitempty"`
	InstanceOwnerID      string   `json:"instanceOwnerID,omitempty"`
	InstanceType         string   `json:"instanceType,omitempty"`
	InstanceID           string   `json:"instanceID,omitempty"`
}

type EcsPortsCreateResponse struct {
	VpcID                string
	SubnetID             string
	NetworkInterfaceID   string
	NetworkInterfaceName string
	MacAddress           string
	Description          string
	Ipv6Address          []string
	SecurityGroupIds     []string
	SecondaryPrivateIps  []string
	PrivateIpAddress     string
	InstanceOwnerID      string
	InstanceType         string
	InstanceID           string
}
