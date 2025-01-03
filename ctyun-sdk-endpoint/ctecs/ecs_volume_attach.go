package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeAttachApi 挂载云硬盘
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=11993&data=87&isNormal=1
type EcsVolumeAttachApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeAttachApi(client *ctyunsdk.CtyunClient) *EcsVolumeAttachApi {
	return &EcsVolumeAttachApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/volume/attach",
		},
	}
}

func (this *EcsVolumeAttachApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeAttachRequest) (*EcsVolumeAttachResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVolumeAttachRealRequest{
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

	var realResponse EcsVolumeAttachRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVolumeAttachResponse{
		DiskJobID:     realResponse.DiskJobID,
		DiskRequestID: realResponse.DiskRequestID,
	}, nil
}

type EcsVolumeAttachRealRequest struct {
	DiskID     *string `json:"diskID,omitempty"`
	RegionID   *string `json:"regionID,omitempty"`
	InstanceID *string `json:"instanceID,omitempty"`
}

type EcsVolumeAttachRequest struct {
	DiskID     *string
	RegionID   *string
	InstanceID *string
}

type EcsVolumeAttachRealResponse struct {
	DiskJobID     string `json:"diskJobID,omitempty"`
	DiskRequestID string `json:"diskRequestID,omitempty"`
}

type EcsVolumeAttachResponse struct {
	DiskJobID     string
	DiskRequestID string
}
