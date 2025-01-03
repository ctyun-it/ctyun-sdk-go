package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsGpuDriverListApi 查询GPU云主机驱动版本
// https://www.ctyun.cn/document/10026730/10548907

type EcsGpuDriverListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsGpuDriverListApi(client *ctyunsdk.CtyunClient) *EcsGpuDriverListApi {
	return &EcsGpuDriverListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/gpu-driver/list",
		},
	}
}

func (this *EcsGpuDriverListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsGpuDriverListRequest) (*EcsGpuDriverListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsGpuDriverListRealRequest{
		RegionID: req.RegionID,
		FlavorID: req.FlavorID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsGpuDriverListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsGpuDriverListResponse{
		GpuDriverList: realResponse.GpuDriverList,
	}, nil
}

type EcsGpuDriverListRealRequest struct {
	RegionID *string `json:"regionID,omitempty"`
	FlavorID *string `json:"flavorID,omitempty"`
}

type EcsGpuDriverListRequest struct {
	RegionID *string
	FlavorID *string
}

type EcsGpuDriverListRealResponse struct {
	GpuDriverList []string `json:"gpuDriverList,omitempty"`
}

type EcsGpuDriverListResponse struct {
	GpuDriverList []string
}
