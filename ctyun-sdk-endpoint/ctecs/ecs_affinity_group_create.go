package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupCreateApi 创建云主机组
// https://www.ctyun.cn/document/10026730/10106207
type EcsAffinityGroupCreateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupCreateApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupCreateApi {
	return &EcsAffinityGroupCreateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group/create",
		},
	}
}

func (this *EcsAffinityGroupCreateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupCreateRequest) (*EcsAffinityGroupCreateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupCreateRealRequest{
		RegionID:          req.RegionID,
		AffinityGroupName: req.AffinityGroupName,
		PolicyType:        req.PolicyType,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAffinityGroupCreateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsAffinityGroupCreateResponse{
		AffinityGroupID:   realResponse.AffinityGroupID,
		AffinityGroupName: realResponse.AffinityGroupName,
		AffinityGroupPolicy: EcsAffinityGroupCreateAffinityGroupPolicyResponse{
			PolicyType:     realResponse.AffinityGroupPolicy.PolicyType,
			PolicyTypeName: realResponse.AffinityGroupPolicy.PolicyTypeName,
		},
	}, nil
}

type EcsAffinityGroupCreateRealRequest struct {
	RegionID          *string `json:"regionID,omitempty"`
	AffinityGroupName *string `json:"affinityGroupName,omitempty"`
	PolicyType        *int    `json:"policyType,omitempty"`
}

type EcsAffinityGroupCreateRequest struct {
	RegionID          *string
	AffinityGroupName *string
	PolicyType        *int
}

type EcsAffinityGroupCreateAffinityGroupPolicyRealResponse struct {
	PolicyType     int    `json:"policyType,omitempty"`
	PolicyTypeName string `json:"policyTypeName,omitempty"`
}

type EcsAffinityGroupCreateRealResponse struct {
	AffinityGroupID     string                                                `json:"affinityGroupID,omitempty"`
	AffinityGroupName   string                                                `json:"affinityGroupName,omitempty"`
	AffinityGroupPolicy EcsAffinityGroupCreateAffinityGroupPolicyRealResponse `json:"affinityGroupPolicy,omitempty"`
}

type EcsAffinityGroupCreateAffinityGroupPolicyResponse struct {
	PolicyType     int
	PolicyTypeName string
}

type EcsAffinityGroupCreateResponse struct {
	AffinityGroupID     string
	AffinityGroupName   string
	AffinityGroupPolicy EcsAffinityGroupCreateAffinityGroupPolicyResponse
}
