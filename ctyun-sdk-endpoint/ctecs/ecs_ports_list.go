package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
	"strconv"
)

// EcsPortsListApi
type EcsPortsListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsListApi(client *ctyunsdk.CtyunClient) *EcsPortsListApi {
	return &EcsPortsListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/ports/list",
		},
	}
}

func (this *EcsPortsListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsListRequest) (*EcsPortsListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", req.RegionID).
		AddParam("vpcID", req.VpcID).
		AddParam("deviceID", req.DeviceID).
		AddParam("subnetID", req.SubnetID).
		AddParam("pageNumber", strconv.Itoa(req.PageNumber)).
		AddParam("pageNo", strconv.Itoa(req.PageNo)).
		AddParam("pageSize", strconv.Itoa(req.PageSize))

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var associatedEip []EcsPortsListAssociatedEipResponse
	for _, res := range realResponse.AssociatedEip {
		associatedEip = append(associatedEip, EcsPortsListAssociatedEipResponse{
			Id:   res.Id,
			Name: res.Name,
			Ip:   res.Ip,
		})
	}

	return &EcsPortsListResponse{
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
		AssociatedEip:        associatedEip,
	}, nil
}

type EcsPortsListRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	VpcID      string `json:"vpcID,omitempty"`
	DeviceID   string `json:"deviceID,omitempty"`
	SubnetID   string `json:"subnetID,omitempty"`
	PageNumber int    `json:"pageNumber,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
	PageNo     int    `json:"pageNo,omitempty"`
}

type EcsPortsListRequest struct {
	RegionID   string
	VpcID      string
	DeviceID   string
	SubnetID   string
	PageNumber int
	PageSize   int
	PageNo     int
}

type EcsPortsListAssociatedEipRealResponse struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Ip   string `json:"ip,omitempty"`
}

type EcsPortsListRealResponse struct {
	NetworkInterfaceName string                                  `json:"networkInterfaceName,omitempty"`
	NetworkInterfaceID   string                                  `json:"networkInterfaceID,omitempty"`
	VpcID                string                                  `json:"vpcID,omitempty"`
	SubnetID             string                                  `json:"subnetID,omitempty"`
	Role                 int                                     `json:"role,omitempty"`
	MacAddress           string                                  `json:"macAddress,omitempty"`
	PrimaryPrivateIp     string                                  `json:"primaryPrivateIp,omitempty"`
	Ipv6Addresses        []string                                `json:"ipv6Addresses,omitempty"`
	InstanceID           string                                  `json:"instanceID,omitempty"`
	InstanceType         string                                  `json:"instanceType,omitempty"`
	Description          string                                  `json:"description,omitempty"`
	SecurityGroupIds     []string                                `json:"securityGroupIds,omitempty"`
	SecondaryPrivateIps  []string                                `json:"secondaryPrivateIps,omitempty"`
	AdminStatus          string                                  `json:"adminStatus,omitempty"`
	AssociatedEip        []EcsPortsListAssociatedEipRealResponse `json:"associatedEip,omitempty"`
}

type EcsPortsListAssociatedEipResponse struct {
	Id   string
	Name string
	Ip   string
}

type EcsPortsListResponse struct {
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
	AssociatedEip        []EcsPortsListAssociatedEipResponse
}
