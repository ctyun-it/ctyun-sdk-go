package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeExtendApi 扩容云硬盘
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=11992&data=87&isNormal=1
type EcsVolumeExtendApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeExtendApi(client *ctyunsdk.CtyunClient) *EcsVolumeExtendApi {
	return &EcsVolumeExtendApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/volume/extend",
		},
	}
}

func (this *EcsVolumeExtendApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeExtendRequest) (*EcsVolumeExtendResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVolumeExtendRealRequest{
		DiskSize:    req.DiskSize,
		DiskID:      req.DiskID,
		RegionID:    req.RegionID,
		ClientToken: req.ClientToken,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVolumeExtendRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVolumeExtendResponse{
		MasterOrderID:        realResponse.MasterOrderID,
		MasterOrderNO:        realResponse.MasterOrderNO,
		MasterResourceID:     realResponse.MasterResourceID,
		MasterResourceStatus: realResponse.MasterResourceStatus,
	}, nil
}

type EcsVolumeExtendRealRequest struct {
	DiskSize    *int    `json:"diskSize,omitempty"`
	DiskID      *string `json:"diskID,omitempty"`
	RegionID    *string `json:"regionID,omitempty"`
	ClientToken *string `json:"clientToken,omitempty"`
}

type EcsVolumeExtendRequest struct {
	DiskSize    *int
	DiskID      *string
	RegionID    *string
	ClientToken *string
}

type EcsVolumeExtendRealResponse struct {
	MasterOrderID        string `json:"masterOrderID,omitempty"`
	MasterOrderNO        string `json:"masterOrderNO,omitempty"`
	MasterResourceID     string `json:"masterResourceID,omitempty"`
	MasterResourceStatus string `json:"masterResourceStatus,omitempty"`
}

type EcsVolumeExtendResponse struct {
	MasterOrderID        string
	MasterOrderNO        string
	MasterResourceID     string
	MasterResourceStatus string
}
