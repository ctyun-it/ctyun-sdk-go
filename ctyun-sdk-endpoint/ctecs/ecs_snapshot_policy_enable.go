package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyEnableApi 启用云主机快照策略
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9591&data=87&isNormal=1
type EcsSnapshotPolicyEnableApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyEnableApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyEnableApi {
	return &EcsSnapshotPolicyEnableApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/enable",
		},
	}
}

func (this *EcsSnapshotPolicyEnableApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyEnableRequest) (*EcsSnapshotPolicyEnableResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyEnableRealRequest{
		RegionID:         req.RegionID,
		SnapshotPolicyID: req.SnapshotPolicyID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotPolicyEnableRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotPolicyEnableResponse{
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

type EcsSnapshotPolicyEnableRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
}

type EcsSnapshotPolicyEnableRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
}

type EcsSnapshotPolicyEnableRealResponse struct {
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

type EcsSnapshotPolicyEnableResponse struct {
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
