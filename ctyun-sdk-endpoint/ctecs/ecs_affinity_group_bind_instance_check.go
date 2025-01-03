package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupBindInstanceCheckApi 云主机加入主机组校验
// https://www.ctyun.cn/document/10026730/10106270
type EcsAffinityGroupBindInstanceCheckApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupBindInstanceCheckApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupBindInstanceCheckApi {
	return &EcsAffinityGroupBindInstanceCheckApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group/bind-instance-check",
		},
	}
}

func (this *EcsAffinityGroupBindInstanceCheckApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupBindInstanceCheckRequest) (*EcsAffinityGroupBindInstanceCheckResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupBindInstanceCheckRealRequest{
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

	var realResponse EcsAffinityGroupBindInstanceCheckRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAffinityGroupBindInstanceCheckResponse{
		NeedMigrate: realResponse.NeedMigrate,
	}, nil
}

type EcsAffinityGroupBindInstanceCheckRealRequest struct {
	RegionID        *string `json:"regionID,omitempty"`
	InstanceID      *string `json:"instanceID,omitempty"`
	AffinityGroupID *string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupBindInstanceCheckRequest struct {
	RegionID        *string
	InstanceID      *string
	AffinityGroupID *string
}

type EcsAffinityGroupBindInstanceCheckRealResponse struct {
	NeedMigrate int `json:"needMigrate,omitempty"`
}

type EcsAffinityGroupBindInstanceCheckResponse struct {
	NeedMigrate int
}
