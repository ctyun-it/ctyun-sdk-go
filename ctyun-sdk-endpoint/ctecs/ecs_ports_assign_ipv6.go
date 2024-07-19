package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsAssignIpv6Api
type EcsPortsAssignIpv6Api struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsAssignIpv6Api(client *ctyunsdk.CtyunClient) *EcsPortsAssignIpv6Api {
	return &EcsPortsAssignIpv6Api{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/assign-ipv6",
		},
	}
}

func (this *EcsPortsAssignIpv6Api) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsAssignIpv6Request) (*EcsPortsAssignIpv6Response, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsAssignIpv6RealRequest{
		ClientToken:        req.ClientToken,
		RegionID:           req.RegionID,
		NetworkInterfaceID: req.NetworkInterfaceID,
		Ipv6AddressesCount: req.Ipv6AddressesCount,
		Ipv6Addresses:      req.Ipv6Addresses,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsAssignIpv6RealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsAssignIpv6Response{}, nil
}

type EcsPortsAssignIpv6RealRequest struct {
	ClientToken        string   `json:"clientToken,omitempty"`
	RegionID           string   `json:"regionID,omitempty"`
	NetworkInterfaceID string   `json:"networkInterfaceID,omitempty"`
	Ipv6AddressesCount *int     `json:"ipv6AddressesCount,omitempty"`
	Ipv6Addresses      []string `json:"ipv6Addresses,omitempty"`
}

type EcsPortsAssignIpv6Request struct {
	ClientToken        string
	RegionID           string
	NetworkInterfaceID string
	Ipv6AddressesCount *int
	Ipv6Addresses      []string
}

type EcsPortsAssignIpv6RealResponse struct {
}

type EcsPortsAssignIpv6Response struct {
}
