package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyDetailsApi 查询云主机快照策略详情
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9603&data=87&isNormal=1
type EcsSnapshotPolicyDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyDetailsApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyDetailsApi {
	return &EcsSnapshotPolicyDetailsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/snapshot-policy/details",
		},
	}
}

func (this *EcsSnapshotPolicyDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyDetailsRequest) (*EcsSnapshotPolicyDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", *req.RegionID).
		AddParam("snapshotPolicyID", *req.SnapshotPolicyID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotPolicyDetailsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotPolicyDetailsResponse{
		SnapshotPolicyID:     realResponse.SnapshotPolicyID,
		SnapshotPolicyStatus: realResponse.SnapshotPolicyStatus,
		SnapshotPolicyName:   realResponse.SnapshotPolicyName,
		SnapshotTime:         realResponse.SnapshotTime,
		RetentionType:        realResponse.RetentionType,
		RetentionDay:         realResponse.RetentionDay,
		RetentionNum:         realResponse.RetentionNum,
		CycleType:            realResponse.CycleType,
		CycleDay:             realResponse.CycleDay,
		CycleWeek:            realResponse.CycleWeek,
		ResourceCount:        realResponse.ResourceCount,
	}, nil
}

type EcsSnapshotPolicyDetailsRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
}

type EcsSnapshotPolicyDetailsRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
}

type EcsSnapshotPolicyDetailsRealResponse struct {
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

type EcsSnapshotPolicyDetailsResponse struct {
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
