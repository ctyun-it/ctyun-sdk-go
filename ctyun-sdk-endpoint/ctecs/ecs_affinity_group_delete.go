package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupDeleteApi 删除云主机组
// https://www.ctyun.cn/document/10026730/10039550
type EcsAffinityGroupDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupDeleteApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupDeleteApi {
	return &EcsAffinityGroupDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group-delete",
		},
	}
}

func (this *EcsAffinityGroupDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupDeleteRequest) (*EcsAffinityGroupDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupDeleteRealRequest{
		RegionID:        req.RegionID,
		AffinityGroupID: req.AffinityGroupID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAffinityGroupDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAffinityGroupDeleteResponse{
		AffinityGroupID: realResponse.AffinityGroupID,
	}, nil
}

type EcsAffinityGroupDeleteRealRequest struct {
	RegionID        *string `json:"regionID,omitempty"`
	AffinityGroupID *string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupDeleteRequest struct {
	RegionID        *string
	AffinityGroupID *string
}

type EcsAffinityGroupDeleteRealResponse struct {
	AffinityGroupID string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupDeleteResponse struct {
	AffinityGroupID string
}
