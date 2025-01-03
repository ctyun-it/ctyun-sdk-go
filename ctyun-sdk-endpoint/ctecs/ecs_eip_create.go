package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsEipCreateApi
type EcsEipCreateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsEipCreateApi(client *ctyunsdk.CtyunClient) *EcsEipCreateApi {
	return &EcsEipCreateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/eip/create",
		},
	}
}

func (this *EcsEipCreateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsEipCreateRequest) (*EcsEipCreateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsEipCreateRealRequest{
		ClientToken:       req.ClientToken,
		RegionID:          req.RegionID,
		ProjectID:         req.ProjectID,
		CycleType:         req.CycleType,
		CycleCount:        req.CycleCount,
		Name:              req.Name,
		Bandwidth:         req.Bandwidth,
		BandwidthID:       req.BandwidthID,
		DemandBillingType: req.DemandBillingType,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsEipCreateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsEipCreateResponse{
		MasterOrderID:        realResponse.MasterOrderID,
		MasterOrderNO:        realResponse.MasterOrderNO,
		MasterResourceStatus: realResponse.MasterResourceStatus,
		MasterResourceID:     realResponse.MasterResourceID,
		RegionID:             realResponse.RegionID,
		EipID:                realResponse.EipID,
	}, nil
}

type EcsEipCreateRealRequest struct {
	ClientToken       *string `json:"clientToken,omitempty"`
	RegionID          *string `json:"regionID,omitempty"`
	ProjectID         *string `json:"projectID,omitempty"`
	CycleType         *string `json:"cycleType,omitempty"`
	CycleCount        *int    `json:"cycleCount,omitempty"`
	Name              *string `json:"name,omitempty"`
	Bandwidth         *int    `json:"bandwidth,omitempty"`
	BandwidthID       *string `json:"bandwidthID,omitempty"`
	DemandBillingType *string `json:"demandBillingType,omitempty"`
}

type EcsEipCreateRequest struct {
	ClientToken       *string
	RegionID          *string
	ProjectID         *string
	CycleType         *string
	CycleCount        *int
	Name              *string
	Bandwidth         *int
	BandwidthID       *string
	DemandBillingType *string
}

type EcsEipCreateRealResponse struct {
	MasterOrderID        string `json:"masterOrderID,omitempty"`
	MasterOrderNO        string `json:"masterOrderNO,omitempty"`
	MasterResourceStatus string `json:"masterResourceStatus,omitempty"`
	MasterResourceID     string `json:"masterResourceID,omitempty"`
	RegionID             string `json:"regionID,omitempty"`
	EipID                string `json:"eipID,omitempty"`
}

type EcsEipCreateResponse struct {
	MasterOrderID        string
	MasterOrderNO        string
	MasterResourceStatus string
	MasterResourceID     string
	RegionID             string
	EipID                string
}
