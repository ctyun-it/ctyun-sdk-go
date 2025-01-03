package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyUpdateApi 修改云主机快照策略
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9549&data=87&isNormal=1
type EcsSnapshotPolicyUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyUpdateApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyUpdateApi {
	return &EcsSnapshotPolicyUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/update",
		},
	}
}

func (this *EcsSnapshotPolicyUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyUpdateRequest) (*EcsSnapshotPolicyUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyUpdateRealRequest{
		RegionID:           req.RegionID,
		SnapshotPolicyID:   req.SnapshotPolicyID,
		SnapshotPolicyName: req.SnapshotPolicyName,
		SnapshotTime:       req.SnapshotTime,
		CycleType:          req.CycleType,
		CycleDay:           req.CycleDay,
		CycleWeek:          req.CycleWeek,
		RetentionType:      req.RetentionType,
		RetentionDay:       req.RetentionDay,
		RetentionNum:       req.RetentionNum,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotPolicyUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotPolicyUpdateResponse{
		SnapshotPolicyStatus: realResponse.SnapshotPolicyStatus,
		SnapshotPolicyName:   realResponse.SnapshotPolicyName,
		RetentionType:        realResponse.RetentionType,
		RetentionDay:         realResponse.RetentionDay,
		RetentionNum:         realResponse.RetentionNum,
		CycleType:            realResponse.CycleType,
		CycleDay:             realResponse.CycleDay,
		CycleWeek:            realResponse.CycleWeek,
		SnapshotPolicyID:     realResponse.SnapshotPolicyID,
		SnapshotTime:         realResponse.SnapshotTime,
	}, nil
}

type EcsSnapshotPolicyUpdateRealRequest struct {
	RegionID           *string `json:"regionID,omitempty"`
	SnapshotPolicyID   *string `json:"snapshotPolicyID,omitempty"`
	SnapshotPolicyName *string `json:"snapshotPolicyName,omitempty"`
	SnapshotTime       *string `json:"snapshotTime,omitempty"`
	CycleType          *string `json:"cycleType,omitempty"`
	CycleDay           *int    `json:"cycleDay,omitempty"`
	CycleWeek          *string `json:"cycleWeek,omitempty"`
	RetentionType      *string `json:"retentionType,omitempty"`
	RetentionDay       *int    `json:"retentionDay,omitempty"`
	RetentionNum       *int    `json:"retentionNum,omitempty"`
}

type EcsSnapshotPolicyUpdateRequest struct {
	RegionID           *string
	SnapshotPolicyID   *string
	SnapshotPolicyName *string
	SnapshotTime       *string
	CycleType          *string
	CycleDay           *int
	CycleWeek          *string
	RetentionType      *string
	RetentionDay       *int
	RetentionNum       *int
}

type EcsSnapshotPolicyUpdateRealResponse struct {
	SnapshotPolicyStatus int    `json:"snapshotPolicyStatus,omitempty"`
	SnapshotPolicyName   string `json:"snapshotPolicyName,omitempty"`
	RetentionType        string `json:"retentionType,omitempty"`
	RetentionDay         string `json:"retentionDay,omitempty"`
	RetentionNum         string `json:"retentionNum,omitempty"`
	CycleType            string `json:"cycleType,omitempty"`
	CycleDay             int    `json:"cycleDay,omitempty"`
	CycleWeek            string `json:"cycleWeek,omitempty"`
	SnapshotPolicyID     string `json:"snapshotPolicyID,omitempty"`
	SnapshotTime         string `json:"snapshotTime,omitempty"`
}

type EcsSnapshotPolicyUpdateResponse struct {
	SnapshotPolicyStatus int
	SnapshotPolicyName   string
	RetentionType        string
	RetentionDay         string
	RetentionNum         string
	CycleType            string
	CycleDay             int
	CycleWeek            string
	SnapshotPolicyID     string
	SnapshotTime         string
}
