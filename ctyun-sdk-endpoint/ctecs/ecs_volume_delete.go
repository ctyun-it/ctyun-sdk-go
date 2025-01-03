package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeDeleteApi 释放云硬盘
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=11994&data=87&isNormal=1
type EcsVolumeDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeDeleteApi(client *ctyunsdk.CtyunClient) *EcsVolumeDeleteApi {
	return &EcsVolumeDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/volume/delete",
		},
	}
}

func (this *EcsVolumeDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeDeleteRequest) (*EcsVolumeDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVolumeDeleteRealRequest{
		ClientToken: req.ClientToken,
		DiskID:      req.DiskID,
		RegionID:    req.RegionID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVolumeDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVolumeDeleteResponse{
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
	}, nil
}

type EcsVolumeDeleteRealRequest struct {
	ClientToken *string `json:"clientToken,omitempty"`
	DiskID      *string `json:"diskID,omitempty"`
	RegionID    *string `json:"regionID,omitempty"`
}

type EcsVolumeDeleteRequest struct {
	ClientToken *string
	DiskID      *string
	RegionID    *string
}

type EcsVolumeDeleteRealResponse struct {
	MasterOrderID string `json:"masterOrderID,omitempty"`
	MasterOrderNO string `json:"masterOrderNO,omitempty"`
}

type EcsVolumeDeleteResponse struct {
	MasterOrderID string
	MasterOrderNO string
}
