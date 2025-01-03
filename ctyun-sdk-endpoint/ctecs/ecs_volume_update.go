package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVolumeUpdateApi 修改云硬盘属性
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=11979&data=87&isNormal=1
type EcsVolumeUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVolumeUpdateApi(client *ctyunsdk.CtyunClient) *EcsVolumeUpdateApi {
	return &EcsVolumeUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/volume/update",
		},
	}
}

func (this *EcsVolumeUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVolumeUpdateRequest) (*EcsVolumeUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsVolumeUpdateRealRequest{
		RegionID: req.RegionID,
		DiskName: req.DiskName,
		DiskID:   req.DiskID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVolumeUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVolumeUpdateResponse{}, nil
}

type EcsVolumeUpdateRealRequest struct {
	RegionID *string `json:"regionID,omitempty"`
	DiskName *string `json:"diskName,omitempty"`
	DiskID   *string `json:"diskID,omitempty"`
}

type EcsVolumeUpdateRequest struct {
	RegionID *string
	DiskName *string
	DiskID   *string
}

type EcsVolumeUpdateRealResponse struct {
}

type EcsVolumeUpdateResponse struct {
}
