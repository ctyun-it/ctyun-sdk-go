package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsUpdateApi
type EcsPortsUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsUpdateApi(client *ctyunsdk.CtyunClient) *EcsPortsUpdateApi {
	return &EcsPortsUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/update",
		},
	}
}

func (this *EcsPortsUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsUpdateRequest) (*EcsPortsUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsUpdateRealRequest{
		ClientToken:        req.ClientToken,
		RegionID:           req.RegionID,
		NetworkInterfaceID: req.NetworkInterfaceID,
		Name:               req.Name,
		Description:        req.Description,
		SecurityGroupIDs:   req.SecurityGroupIDs,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsUpdateResponse{}, nil
}

type EcsPortsUpdateRealRequest struct {
	ClientToken        string   `json:"clientToken,omitempty"`
	RegionID           string   `json:"regionID,omitempty"`
	NetworkInterfaceID string   `json:"networkInterfaceID,omitempty"`
	Name               string   `json:"name,omitempty"`
	Description        string   `json:"description,omitempty"`
	SecurityGroupIDs   []string `json:"securityGroupIDs,omitempty"`
}

type EcsPortsUpdateRequest struct {
	ClientToken        string
	RegionID           string
	NetworkInterfaceID string
	Name               string
	Description        string
	SecurityGroupIDs   []string
}

type EcsPortsUpdateRealResponse struct {
}

type EcsPortsUpdateResponse struct {
}
