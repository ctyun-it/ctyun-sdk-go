package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsDeleteApi
type EcsPortsDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsDeleteApi(client *ctyunsdk.CtyunClient) *EcsPortsDeleteApi {
	return &EcsPortsDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/delete",
		},
	}
}

func (this *EcsPortsDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsDeleteRequest) (*EcsPortsDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsDeleteRealRequest{
		ClientToken:        req.ClientToken,
		RegionID:           req.RegionID,
		NetworkInterfaceID: req.NetworkInterfaceID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsDeleteResponse{}, nil
}

type EcsPortsDeleteRealRequest struct {
	ClientToken        string `json:"clientToken,omitempty"`
	RegionID           string `json:"regionID,omitempty"`
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty"`
}

type EcsPortsDeleteRequest struct {
	ClientToken        string
	RegionID           string
	NetworkInterfaceID string
}

type EcsPortsDeleteRealResponse struct {
}

type EcsPortsDeleteResponse struct {
}
