package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsDetachApi
type EcsPortsDetachApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsDetachApi(client *ctyunsdk.CtyunClient) *EcsPortsDetachApi {
	return &EcsPortsDetachApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/detach",
		},
	}
}

func (this *EcsPortsDetachApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsDetachRequest) (*EcsPortsDetachResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsDetachRealRequest{
		ClientToken:        req.ClientToken,
		RegionID:           req.RegionID,
		NetworkInterfaceID: req.NetworkInterfaceID,
		InstanceID:         req.InstanceID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsDetachRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsDetachResponse{}, nil
}

type EcsPortsDetachRealRequest struct {
	ClientToken        *string `json:"clientToken,omitempty"`
	RegionID           *string `json:"regionID,omitempty"`
	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty"`
	InstanceID         *string `json:"instanceID,omitempty"`
}

type EcsPortsDetachRequest struct {
	ClientToken        *string
	RegionID           *string
	NetworkInterfaceID *string
	InstanceID         *string
}

type EcsPortsDetachRealResponse struct {
}

type EcsPortsDetachResponse struct {
}
