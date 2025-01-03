package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsEipDeleteApi
type EcsEipDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsEipDeleteApi(client *ctyunsdk.CtyunClient) *EcsEipDeleteApi {
	return &EcsEipDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/eip/delete",
		},
	}
}

func (this *EcsEipDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsEipDeleteRequest) (*EcsEipDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsEipDeleteRealRequest{
		ClientToken: req.ClientToken,
		RegionID:    req.RegionID,
		ProjectID:   req.ProjectID,
		EipID:       req.EipID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsEipDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsEipDeleteResponse{
		MasterOrderID:        realResponse.MasterOrderID,
		MasterOrderNO:        realResponse.MasterOrderNO,
		RegionID:             realResponse.RegionID,
		MasterResourceStatus: realResponse.MasterResourceStatus,
	}, nil
}

type EcsEipDeleteRealRequest struct {
	ClientToken *string `json:"clientToken,omitempty"`
	RegionID    *string `json:"regionID,omitempty"`
	ProjectID   *string `json:"projectID,omitempty"`
	EipID       *string `json:"eipID,omitempty"`
}

type EcsEipDeleteRequest struct {
	ClientToken *string
	RegionID    *string
	ProjectID   *string
	EipID       *string
}

type EcsEipDeleteRealResponse struct {
	MasterOrderID        string `json:"masterOrderID,omitempty"`
	MasterOrderNO        string `json:"masterOrderNO,omitempty"`
	RegionID             string `json:"regionID,omitempty"`
	MasterResourceStatus string `json:"masterResourceStatus,omitempty"`
}

type EcsEipDeleteResponse struct {
	MasterOrderID        string
	MasterOrderNO        string
	RegionID             string
	MasterResourceStatus string
}
