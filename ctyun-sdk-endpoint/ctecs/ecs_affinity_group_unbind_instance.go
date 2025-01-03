package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupUnbindInstanceApi 云主机组中删除单台云主机
// https://www.ctyun.cn/document/10026730/10106275
type EcsAffinityGroupUnbindInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupUnbindInstanceApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupUnbindInstanceApi {
	return &EcsAffinityGroupUnbindInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group/unbind-instance",
		},
	}
}

func (this *EcsAffinityGroupUnbindInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupUnbindInstanceRequest) (*EcsAffinityGroupUnbindInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupUnbindInstanceRealRequest{
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

	var realResponse EcsAffinityGroupUnbindInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAffinityGroupUnbindInstanceResponse{}, nil
}

type EcsAffinityGroupUnbindInstanceRealRequest struct {
	RegionID        *string `json:"regionID,omitempty"`
	InstanceID      *string `json:"instanceID,omitempty"`
	AffinityGroupID *string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupUnbindInstanceRequest struct {
	RegionID        *string
	InstanceID      *string
	AffinityGroupID *string
}

type EcsAffinityGroupUnbindInstanceRealResponse struct {
}

type EcsAffinityGroupUnbindInstanceResponse struct {
}
