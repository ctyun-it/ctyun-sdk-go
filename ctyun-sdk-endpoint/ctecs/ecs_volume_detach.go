package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeDetachApi 卸载云硬盘
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=11991&data=87&isNormal=1
type EcsVolumeDetachApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeDetachApi(client *ctyunsdk.CtyunClient) *EcsVolumeDetachApi {
	return &EcsVolumeDetachApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/volume/detach",
		},
	}
}

func (this *EcsVolumeDetachApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeDetachRequest) (*EcsVolumeDetachResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVolumeDetachRealRequest{
		DiskID:     req.DiskID,
		RegionID:   req.RegionID,
		InstanceID: req.InstanceID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVolumeDetachRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVolumeDetachResponse{
		DiskJobID:     realResponse.DiskJobID,
		DiskRequestID: realResponse.DiskRequestID,
	}, nil
}

type EcsVolumeDetachRealRequest struct {
	DiskID     *string `json:"diskID,omitempty"`
	RegionID   *string `json:"regionID,omitempty"`
	InstanceID *string `json:"instanceID,omitempty"`
}

type EcsVolumeDetachRequest struct {
	DiskID     *string
	RegionID   *string
	InstanceID *string
}

type EcsVolumeDetachRealResponse struct {
	DiskJobID     string `json:"diskJobID,omitempty"`
	DiskRequestID string `json:"diskRequestID,omitempty"`
}

type EcsVolumeDetachResponse struct {
	DiskJobID     string
	DiskRequestID string
}
