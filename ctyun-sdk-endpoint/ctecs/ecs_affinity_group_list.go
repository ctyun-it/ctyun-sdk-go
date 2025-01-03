package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAffinityGroupListApi 查询云主机组列表或详情
// https://www.ctyun.cn/document/10026730/10106059
type EcsAffinityGroupListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAffinityGroupListApi(client *ctyunsdk.CtyunClient) *EcsAffinityGroupListApi {
	return &EcsAffinityGroupListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/affinity-group/list",
		},
	}
}

func (this *EcsAffinityGroupListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAffinityGroupListRequest) (*EcsAffinityGroupListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsAffinityGroupListRealRequest{
		RegionID:        req.RegionID,
		AffinityGroupID: req.AffinityGroupID,
		QueryContent:    req.QueryContent,
		PageNo:          req.PageNo,
		PageSize:        req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsAffinityGroupListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EcsAffinityGroupListResultsResponse
	for _, res := range realResponse.Results {
		results = append(results, EcsAffinityGroupListResultsResponse{
			AffinityGroupID:   res.AffinityGroupID,
			AffinityGroupName: res.AffinityGroupName,
			AffinityGroupPolicy: EcsAffinityGroupListAffinityGroupPolicyResponse{
				PolicyType:     res.AffinityGroupPolicy.PolicyType,
				PolicyTypeName: res.AffinityGroupPolicy.PolicyTypeName,
			},
		})
	}

	return &EcsAffinityGroupListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type EcsAffinityGroupListRealRequest struct {
	RegionID        *string `json:"regionID,omitempty"`
	AffinityGroupID *string `json:"affinityGroupID,omitempty"`
	QueryContent    *string `json:"queryContent,omitempty"`
	PageNo          *int    `json:"pageNo,omitempty"`
	PageSize        *int    `json:"pageSize,omitempty"`
}

type EcsAffinityGroupListRequest struct {
	RegionID        *string
	AffinityGroupID *string
	QueryContent    *string
	PageNo          *int
	PageSize        *int
}

type EcsAffinityGroupListAffinityGroupPolicyRealResponse struct {
	PolicyType     int    `json:"policyType,omitempty"`
	PolicyTypeName string `json:"policyTypeName,omitempty"`
}

type EcsAffinityGroupListResultsRealResponse struct {
	AffinityGroupID     string                                              `json:"affinityGroupID,omitempty"`
	AffinityGroupName   string                                              `json:"affinityGroupName,omitempty"`
	AffinityGroupPolicy EcsAffinityGroupListAffinityGroupPolicyRealResponse `json:"affinityGroupPolicy,omitempty"`
	CreatedTime         string                                              `json:"createdTime,omitempty"`
	UpdatedTime         string                                              `json:"updatedTime,omitempty"`
	Deleted             bool                                                `json:"deleted,omitempty"`
}

type EcsAffinityGroupListRealResponse struct {
	CurrentCount int                                       `json:"currentCount,omitempty"`
	TotalCount   int                                       `json:"totalCount,omitempty"`
	TotalPage    int                                       `json:"totalPage,omitempty"`
	Results      []EcsAffinityGroupListResultsRealResponse `json:"results,omitempty"`
}

type EcsAffinityGroupListAffinityGroupPolicyResponse struct {
	PolicyType     int
	PolicyTypeName string
}

type EcsAffinityGroupListResultsResponse struct {
	AffinityGroupID     string
	AffinityGroupName   string
	AffinityGroupPolicy EcsAffinityGroupListAffinityGroupPolicyResponse
	CreatedTime         string
	UpdatedTime         string
	Deleted             bool
}

type EcsAffinityGroupListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EcsAffinityGroupListResultsResponse
}
