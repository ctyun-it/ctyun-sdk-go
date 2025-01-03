package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupUpdateApi 更新云主机组信息
// https://www.ctyun.cn/document/10026730/10039555
type EcsAffinityGroupUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupUpdateApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupUpdateApi {
	return &EcsAffinityGroupUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group-update",
		},
	}
}

func (this *EcsAffinityGroupUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupUpdateRequest) (*EcsAffinityGroupUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupUpdateRealRequest{
		RegionID:          req.RegionID,
		AffinityGroupID:   req.AffinityGroupID,
		AffinityGroupName: req.AffinityGroupName,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAffinityGroupUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAffinityGroupUpdateResponse{
		AffinityGroupID: realResponse.AffinityGroupID,
	}, nil
}

type EcsAffinityGroupUpdateRealRequest struct {
	RegionID          *string `json:"regionID,omitempty"`
	AffinityGroupID   *string `json:"affinityGroupID,omitempty"`
	AffinityGroupName *string `json:"affinityGroupName,omitempty"`
}

type EcsAffinityGroupUpdateRequest struct {
	RegionID          *string
	AffinityGroupID   *string
	AffinityGroupName *string
}

type EcsAffinityGroupUpdateRealResponse struct {
	AffinityGroupID string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupUpdateResponse struct {
	AffinityGroupID string
}
