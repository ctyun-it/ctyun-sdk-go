package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsUnassignIpv6Api
type EcsPortsUnassignIpv6Api struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsUnassignIpv6Api(client *ctyunsdk.CtyunClient) *EcsPortsUnassignIpv6Api {
	return &EcsPortsUnassignIpv6Api{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/unassign-ipv6",
		},
	}
}

func (this *EcsPortsUnassignIpv6Api) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsUnassignIpv6Request) (*EcsPortsUnassignIpv6Response, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsUnassignIpv6RealRequest{
		ClientToken:        req.ClientToken,
		RegionID:           req.RegionID,
		NetworkInterfaceID: req.NetworkInterfaceID,
		Ipv6Addresses:      req.Ipv6Addresses,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsUnassignIpv6RealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsUnassignIpv6Response{}, nil
}

type EcsPortsUnassignIpv6RealRequest struct {
	ClientToken        string   `json:"clientToken,omitempty"`
	RegionID           string   `json:"regionID,omitempty"`
	NetworkInterfaceID string   `json:"networkInterfaceID,omitempty"`
	Ipv6Addresses      []string `json:"ipv6Addresses,omitempty"`
}

type EcsPortsUnassignIpv6Request struct {
	ClientToken        string
	RegionID           string
	NetworkInterfaceID string
	Ipv6Addresses      []string
}

type EcsPortsUnassignIpv6RealResponse struct {
}

type EcsPortsUnassignIpv6Response struct {
}
