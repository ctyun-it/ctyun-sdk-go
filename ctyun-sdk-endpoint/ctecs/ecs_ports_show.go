package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsShowApi
type EcsPortsShowApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsShowApi(client *ctyunsdk.CtyunClient) *EcsPortsShowApi {
	return &EcsPortsShowApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/ports/show",
		},
	}
}

func (this *EcsPortsShowApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsShowRequest) (*EcsPortsShowResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	builder.
		AddParam("regionID", *req.RegionID).
		AddParam("networkInterfaceID", *req.NetworkInterfaceID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsShowRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsShowResponse{
		NetworkInterfaceName: realResponse.NetworkInterfaceName,
		NetworkInterfaceID:   realResponse.NetworkInterfaceID,
		VpcID:                realResponse.VpcID,
		SubnetID:             realResponse.SubnetID,
		Role:                 realResponse.Role,
		MacAddress:           realResponse.MacAddress,
		PrimaryPrivateIp:     realResponse.PrimaryPrivateIp,
		Ipv6Addresses:        realResponse.Ipv6Addresses,
		InstanceID:           realResponse.InstanceID,
		InstanceType:         realResponse.InstanceType,
		Description:          realResponse.Description,
		SecurityGroupIds:     realResponse.SecurityGroupIds,
		SecondaryPrivateIps:  realResponse.SecondaryPrivateIps,
		AdminStatus:          realResponse.AdminStatus,
		AssociatedEip: EcsPortsShowAssociatedEipResponse{
			Id:   realResponse.AssociatedEip.Id,
			Name: realResponse.AssociatedEip.Name,
		},
	}, nil
}

type EcsPortsShowRealRequest struct {
	RegionID           *string `json:"regionID,omitempty"`
	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty"`
}

type EcsPortsShowRequest struct {
	RegionID           *string
	NetworkInterfaceID *string
}

type EcsPortsShowAssociatedEipRealResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type EcsPortsShowRealResponse struct {
	NetworkInterfaceName string                                `json:"networkInterfaceName,omitempty"`
	NetworkInterfaceID   string                                `json:"networkInterfaceID,omitempty"`
	VpcID                string                                `json:"vpcID,omitempty"`
	SubnetID             string                                `json:"subnetID,omitempty"`
	Role                 int                                   `json:"role,omitempty"`
	MacAddress           string                                `json:"macAddress,omitempty"`
	PrimaryPrivateIp     string                                `json:"primaryPrivateIp,omitempty"`
	Ipv6Addresses        []string                              `json:"ipv6Addresses,omitempty"`
	InstanceID           string                                `json:"instanceID,omitempty"`
	InstanceType         string                                `json:"instanceType,omitempty"`
	Description          string                                `json:"description,omitempty"`
	SecurityGroupIds     []string                              `json:"securityGroupIds,omitempty"`
	SecondaryPrivateIps  []string                              `json:"secondaryPrivateIps,omitempty"`
	AdminStatus          string                                `json:"adminStatus,omitempty"`
	AssociatedEip        EcsPortsShowAssociatedEipRealResponse `json:"associatedEip,omitempty"`
}

type EcsPortsShowAssociatedEipResponse struct {
	Id   string
	Name string
}

type EcsPortsShowResponse struct {
	NetworkInterfaceName string
	NetworkInterfaceID   string
	VpcID                string
	SubnetID             string
	Role                 int
	MacAddress           string
	PrimaryPrivateIp     string
	Ipv6Addresses        []string
	InstanceID           string
	InstanceType         string
	Description          string
	SecurityGroupIds     []string
	SecondaryPrivateIps  []string
	AdminStatus          string
	AssociatedEip        EcsPortsShowAssociatedEipResponse
}
