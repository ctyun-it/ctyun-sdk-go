package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupDetailsApi 查询云主机所在云主机组
// https://www.ctyun.cn/document/10026730/10106058
type EcsAffinityGroupDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupDetailsApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupDetailsApi {
	return &EcsAffinityGroupDetailsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/affinity-group/details",
		},
	}
}

func (this *EcsAffinityGroupDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupDetailsRequest) (*EcsAffinityGroupDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", *req.RegionID).
		AddParam("instanceID", *req.InstanceID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAffinityGroupDetailsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAffinityGroupDetailsResponse{
		PolicyTypeName:    realResponse.PolicyTypeName,
		AffinityGroupName: realResponse.AffinityGroupName,
		AffinityGroupID:   realResponse.AffinityGroupID,
	}, nil
}

type EcsAffinityGroupDetailsRealRequest struct {
	RegionID   *string `json:"regionID,omitempty"`
	InstanceID *string `json:"instanceID,omitempty"`
}

type EcsAffinityGroupDetailsRequest struct {
	RegionID   *string
	InstanceID *string
}

type EcsAffinityGroupDetailsRealResponse struct {
	PolicyTypeName    string `json:"policyTypeName,omitempty"`
	AffinityGroupName string `json:"affinityGroupName,omitempty"`
	AffinityGroupID   string `json:"affinityGroupID,omitempty"`
}

type EcsAffinityGroupDetailsResponse struct {
	PolicyTypeName    string
	AffinityGroupName string
	AffinityGroupID   string
}
