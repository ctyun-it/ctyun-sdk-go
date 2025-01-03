package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyDisableApi 停用云主机快照策略
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9579&data=87&isNormal=1
type EcsSnapshotPolicyDisableApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyDisableApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyDisableApi {
	return &EcsSnapshotPolicyDisableApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/disable",
		},
	}
}

func (this *EcsSnapshotPolicyDisableApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyDisableRequest) (*EcsSnapshotPolicyDisableResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyDisableRealRequest{
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

	var realResponse EcsSnapshotPolicyDisableRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotPolicyDisableResponse{
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

type EcsSnapshotPolicyDisableRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
}

type EcsSnapshotPolicyDisableRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
}

type EcsSnapshotPolicyDisableRealResponse struct {
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

type EcsSnapshotPolicyDisableResponse struct {
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
