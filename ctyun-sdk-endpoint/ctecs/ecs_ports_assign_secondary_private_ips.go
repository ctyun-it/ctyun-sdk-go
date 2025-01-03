package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsAssignSecondaryPrivateIpsApi
type EcsPortsAssignSecondaryPrivateIpsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsAssignSecondaryPrivateIpsApi(client *ctyunsdk.CtyunClient) *EcsPortsAssignSecondaryPrivateIpsApi {
	return &EcsPortsAssignSecondaryPrivateIpsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/assign-secondary-private-ips",
		},
	}
}

func (this *EcsPortsAssignSecondaryPrivateIpsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsAssignSecondaryPrivateIpsRequest) (*EcsPortsAssignSecondaryPrivateIpsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsAssignSecondaryPrivateIpsRealRequest{
		ClientToken:             req.ClientToken,
		RegionID:                req.RegionID,
		NetworkInterfaceID:      req.NetworkInterfaceID,
		SecondaryPrivateIps:     req.SecondaryPrivateIps,
		SecondaryPrivateIpCount: req.SecondaryPrivateIpCount,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsAssignSecondaryPrivateIpsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsAssignSecondaryPrivateIpsResponse{}, nil
}

type EcsPortsAssignSecondaryPrivateIpsRealRequest struct {
	ClientToken             *string   `json:"clientToken,omitempty"`
	RegionID                *string   `json:"regionID,omitempty"`
	NetworkInterfaceID      *string   `json:"networkInterfaceID,omitempty"`
	SecondaryPrivateIps     *[]string `json:"secondaryPrivateIps,omitempty"`
	SecondaryPrivateIpCount *int      `json:"secondaryPrivateIpCount,omitempty"`
}

type EcsPortsAssignSecondaryPrivateIpsRequest struct {
	ClientToken             *string
	RegionID                *string
	NetworkInterfaceID      *string
	SecondaryPrivateIps     *[]string
	SecondaryPrivateIpCount *int
}

type EcsPortsAssignSecondaryPrivateIpsRealResponse struct {
}

type EcsPortsAssignSecondaryPrivateIpsResponse struct {
}
