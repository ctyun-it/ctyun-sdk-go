package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsPortsAttachApi
type EcsPortsAttachApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsPortsAttachApi(client *ctyunsdk.CtyunClient) *EcsPortsAttachApi {
	return &EcsPortsAttachApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/ports/attach",
		},
	}
}

func (this *EcsPortsAttachApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsPortsAttachRequest) (*EcsPortsAttachResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsPortsAttachRealRequest{
		ClientToken:        req.ClientToken,
		RegionID:           req.RegionID,
		AzName:             req.AzName,
		ProjectID:          req.ProjectID,
		NetworkInterfaceID: req.NetworkInterfaceID,
		InstanceID:         req.InstanceID,
		InstanceType:       req.InstanceType,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsPortsAttachRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsPortsAttachResponse{}, nil
}

type EcsPortsAttachRealRequest struct {
	ClientToken        string `json:"clientToken,omitempty"`
	RegionID           string `json:"regionID,omitempty"`
	AzName             string `json:"azName,omitempty"`
	ProjectID          string `json:"projectID,omitempty"`
	NetworkInterfaceID string `json:"networkInterfaceID,omitempty"`
	InstanceID         string `json:"instanceID,omitempty"`
	InstanceType       *int   `json:"instanceType,omitempty"`
}

type EcsPortsAttachRequest struct {
	ClientToken        string
	RegionID           string
	AzName             string
	ProjectID          string
	NetworkInterfaceID string
	InstanceID         string
	InstanceType       *int
}

type EcsPortsAttachRealResponse struct {
}

type EcsPortsAttachResponse struct {
}
