package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsUnassignSecondaryPrivateIpsApi
type EcsPortsUnassignSecondaryPrivateIpsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsUnassignSecondaryPrivateIpsApi(client *ctyunsdk.CtyunClient) *EcsPortsUnassignSecondaryPrivateIpsApi {
	return &EcsPortsUnassignSecondaryPrivateIpsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/unassign-secondary-private-ips",
		},
	}
}

func (this *EcsPortsUnassignSecondaryPrivateIpsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsUnassignSecondaryPrivateIpsRequest) (*EcsPortsUnassignSecondaryPrivateIpsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsUnassignSecondaryPrivateIpsRealRequest{
		ClientToken:         req.ClientToken,
		RegionID:            req.RegionID,
		NetworkInterfaceID:  req.NetworkInterfaceID,
		SecondaryPrivateIps: req.SecondaryPrivateIps,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsUnassignSecondaryPrivateIpsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsUnassignSecondaryPrivateIpsResponse{}, nil
}

type EcsPortsUnassignSecondaryPrivateIpsRealRequest struct {
	ClientToken         *string   `json:"clientToken,omitempty"`
	RegionID            *string   `json:"regionID,omitempty"`
	NetworkInterfaceID  *string   `json:"networkInterfaceID,omitempty"`
	SecondaryPrivateIps *[]string `json:"secondaryPrivateIps,omitempty"`
}

type EcsPortsUnassignSecondaryPrivateIpsRequest struct {
	ClientToken         *string
	RegionID            *string
	NetworkInterfaceID  *string
	SecondaryPrivateIps *[]string
}

type EcsPortsUnassignSecondaryPrivateIpsRealResponse struct {
}

type EcsPortsUnassignSecondaryPrivateIpsResponse struct {
}
