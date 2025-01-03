package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupBindInstanceApi 云主机组中添加单台云主机
// https://www.ctyun.cn/document/10026730/10106277
type EcsAffinityGroupBindInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupBindInstanceApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupBindInstanceApi {
	return &EcsAffinityGroupBindInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group/bind-instance",
		},
	}
}

func (this *EcsAffinityGroupBindInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupBindInstanceRequest) (*EcsAffinityGroupBindInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupBindInstanceRealRequest{
		RegionID:        req.RegionID,
		InstanceID:      req.InstanceID,
		AffinityGroupID: req.AffinityGroupID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAffinityGroupBindInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAffinityGroupBindInstanceResponse{}, nil
}

type EcsAffinityGroupBindInstanceRealRequest struct {
	RegionID        *string `json:"regionID,omitempty"`
	InstanceID      *string `json:"instanceID,omitempty"`
	AffinityGroupID *string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupBindInstanceRequest struct {
	RegionID        *string
	InstanceID      *string
	AffinityGroupID *string
}

type EcsAffinityGroupBindInstanceRealResponse struct {
}

type EcsAffinityGroupBindInstanceResponse struct {
}
