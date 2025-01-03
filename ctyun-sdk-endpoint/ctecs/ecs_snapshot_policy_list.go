package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyListApi 查询云主机快照策略列表
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9600&data=87&isNormal=1
type EcsSnapshotPolicyListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyListApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyListApi {
	return &EcsSnapshotPolicyListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/list",
		},
	}
}

func (this *EcsSnapshotPolicyListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyListRequest) (*EcsSnapshotPolicyListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyListRealRequest{
		RegionID:             req.RegionID,
		PageNo:               req.PageNo,
		PageSize:             req.PageSize,
		SnapshotPolicyStatus: req.SnapshotPolicyStatus,
		QueryContent:         req.QueryContent,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotPolicyListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var snapshotPolicyList []EcsSnapshotPolicyListSnapshotPolicyListResponse
	for _, res := range realResponse.SnapshotPolicyList {
		snapshotPolicyList = append(snapshotPolicyList, EcsSnapshotPolicyListSnapshotPolicyListResponse{
			SnapshotPolicyID:     res.SnapshotPolicyID,
			SnapshotPolicyStatus: res.SnapshotPolicyStatus,
			SnapshotPolicyName:   res.SnapshotPolicyName,
			SnapshotTime:         res.SnapshotTime,
			RetentionType:        res.RetentionType,
			RetentionDay:         res.RetentionDay,
			RetentionNum:         res.RetentionNum,
			CycleType:            res.CycleType,
			CycleDay:             res.CycleDay,
			CycleWeek:            res.CycleWeek,
			ResourceCount:        res.ResourceCount,
		})
	}

	return &EcsSnapshotPolicyListResponse{
		CurrentCount:       realResponse.CurrentCount,
		TotalCount:         realResponse.TotalCount,
		TotalPage:          realResponse.TotalPage,
		SnapshotPolicyList: snapshotPolicyList,
	}, nil
}

type EcsSnapshotPolicyListRealRequest struct {
	RegionID             *string `json:"regionID,omitempty"`
	PageNo               *int    `json:"pageNo,omitempty"`
	PageSize             *int    `json:"pageSize,omitempty"`
	SnapshotPolicyStatus *int    `json:"snapshotPolicyStatus,omitempty"`
	QueryContent         *string `json:"queryContent,omitempty"`
}

type EcsSnapshotPolicyListRequest struct {
	RegionID             *string
	PageNo               *int
	PageSize             *int
	SnapshotPolicyStatus *int
	QueryContent         *string
}

type EcsSnapshotPolicyListSnapshotPolicyListRealResponse struct {
	SnapshotPolicyID     string `json:"snapshotPolicyID,omitempty"`
	SnapshotPolicyStatus int    `json:"snapshotPolicyStatus,omitempty"`
	SnapshotPolicyName   string `json:"snapshotPolicyName,omitempty"`
	SnapshotTime         string `json:"snapshotTime,omitempty"`
	RetentionType        string `json:"retentionType,omitempty"`
	RetentionDay         string `json:"retentionDay,omitempty"`
	RetentionNum         string `json:"retentionNum,omitempty"`
	CycleType            string `json:"cycleType,omitempty"`
	CycleDay             int    `json:"cycleDay,omitempty"`
	CycleWeek            string `json:"cycleWeek,omitempty"`
	ResourceCount        int    `json:"resourceCount,omitempty"`
}

type EcsSnapshotPolicyListRealResponse struct {
	CurrentCount       int                                                   `json:"currentCount,omitempty"`
	TotalCount         int                                                   `json:"totalCount,omitempty"`
	TotalPage          int                                                   `json:"totalPage,omitempty"`
	SnapshotPolicyList []EcsSnapshotPolicyListSnapshotPolicyListRealResponse `json:"snapshotPolicyList,omitempty"`
}

type EcsSnapshotPolicyListSnapshotPolicyListResponse struct {
	SnapshotPolicyID     string
	SnapshotPolicyStatus int
	SnapshotPolicyName   string
	SnapshotTime         string
	RetentionType        string
	RetentionDay         string
	RetentionNum         string
	CycleType            string
	CycleDay             int
	CycleWeek            string
	ResourceCount        int
}

type EcsSnapshotPolicyListResponse struct {
	CurrentCount       int
	TotalCount         int
	TotalPage          int
	SnapshotPolicyList []EcsSnapshotPolicyListSnapshotPolicyListResponse
}
