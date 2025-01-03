package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupPolicyUpdateApi 创建云主机备份策略
type EcsBackupPolicyUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

// NewEcsBackupPolicyUpdateApi 创建 EcsBackupPolicyUpdateApi 实例
func NewEcsBackupPolicyUpdateApi(client *ctyunsdk.CtyunClient) *EcsBackupPolicyUpdateApi {
	return &EcsBackupPolicyUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/update",
		},
	}
}

// Do 执行备份策略创建操作
func (api *EcsBackupPolicyUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupPolicyUpdateRequest) (*EcsBackupPolicyUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := api.WithCredential(&credential)

	// 构建请求体
	realRequest := &EcsBackupPolicyUpdateRealRequest{
		RegionID:           req.RegionID,
		PolicyID:           req.PolicyID,
		PolicyName:         req.PolicyName,
		CycleType:          req.CycleType,
		CycleDay:           req.CycleDay,
		CycleWeek:          req.CycleWeek,
		Time:               req.Time,
		Status:             req.Status,
		RetentionType:      req.RetentionType,
		RetentionDay:       req.RetentionDay,
		RetentionNum:       req.RetentionNum,
		TotalBackup:        req.TotalBackup,
		AdvRetentionStatus: req.AdvRetentionStatus,
		AdvRetention: EcsBackupPolicyUpdateAdvRetentionRealRequest{
			AdvDay:   req.AdvRetention.AdvDay,
			AdvWeek:  req.AdvRetention.AdvWeek,
			AdvMonth: req.AdvRetention.AdvMonth,
			AdvYear:  req.AdvRetention.AdvYear,
		},
	}

	// 写入请求体
	_, err := builder.WriteJson(realRequest)
	if err != nil {
		return nil, err
	}

	// 发送请求并获取响应
	response, err := api.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var realResponse EcsBackupPolicyUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	// 构建返回结果
	return &EcsBackupPolicyUpdateResponse{
		Status:             realResponse.Status,
		PolicyName:         realResponse.PolicyName,
		RetentionType:      realResponse.RetentionType,
		RetentionDay:       realResponse.RetentionDay,
		RetentionNum:       realResponse.RetentionNum,
		RegionID:           realResponse.RegionID,
		CycleType:          realResponse.CycleType,
		CycleDay:           realResponse.CycleDay,
		CycleWeek:          realResponse.CycleWeek,
		PolicyID:           realResponse.PolicyID,
		Time:               realResponse.Time,
		ProjectID:          realResponse.ProjectID,
		TotalBackup:        realResponse.TotalBackup,
		AdvRetentionStatus: realResponse.AdvRetentionStatus,
		AdvRetention: EcsBackupPolicyUpdateAdvRetentionResponse{
			AdvDay:   realResponse.AdvRetention.AdvDay,
			AdvWeek:  realResponse.AdvRetention.AdvWeek,
			AdvMonth: realResponse.AdvRetention.AdvMonth,
			AdvYear:  realResponse.AdvRetention.AdvYear,
		},
	}, nil
}

// 其他结构体定义
type EcsBackupPolicyUpdateAdvRetentionRealRequest struct {
	AdvDay   *int `json:"advDay,omitempty"`
	AdvWeek  *int `json:"advWeek,omitempty"`
	AdvMonth *int `json:"advMonth,omitempty"`
	AdvYear  *int `json:"advYear,omitempty"`
}

type EcsBackupPolicyUpdateRealRequest struct {
	RegionID           *string                                      `json:"regionID,omitempty"`
	PolicyID           *string                                      `json:"policyID,omitempty"`
	PolicyName         *string                                      `json:"policyName,omitempty"`
	CycleType          *string                                      `json:"cycleType,omitempty"`
	CycleDay           *int                                         `json:"cycleDay,omitempty"`
	CycleWeek          *string                                      `json:"cycleWeek,omitempty"`
	Time               *string                                      `json:"time,omitempty"`
	Status             *int                                         `json:"status,omitempty"`
	RetentionType      *string                                      `json:"retentionType,omitempty"`
	RetentionDay       *int                                         `json:"retentionDay,omitempty"`
	RetentionNum       *int                                         `json:"retentionNum,omitempty"`
	TotalBackup        *bool                                        `json:"totalBackup,omitempty"`
	AdvRetentionStatus *bool                                        `json:"advRetentionStatus,omitempty"`
	AdvRetention       EcsBackupPolicyUpdateAdvRetentionRealRequest `json:"advRetention,omitempty"`
}

type EcsBackupPolicyUpdateAdvRetentionRequest struct {
	AdvDay   *int
	AdvWeek  *int
	AdvMonth *int
	AdvYear  *int
}

type EcsBackupPolicyUpdateRequest struct {
	RegionID           *string
	PolicyID           *string
	PolicyName         *string
	CycleType          *string
	CycleDay           *int
	CycleWeek          *string
	Time               *string
	Status             *int
	RetentionType      *string
	RetentionDay       *int
	RetentionNum       *int
	TotalBackup        *bool
	AdvRetentionStatus *bool
	AdvRetention       EcsBackupPolicyUpdateAdvRetentionRequest
}

type EcsBackupPolicyUpdateAdvRetentionRealResponse struct {
	AdvDay   int `json:"advDay,omitempty"`
	AdvWeek  int `json:"advWeek,omitempty"`
	AdvMonth int `json:"advMonth,omitempty"`
	AdvYear  int `json:"advYear,omitempty"`
}

type EcsBackupPolicyUpdateRealResponse struct {
	Status             int                                           `json:"status,omitempty"`
	PolicyName         string                                        `json:"policyName,omitempty"`
	RetentionType      string                                        `json:"retentionType,omitempty"`
	RetentionDay       int                                           `json:"retentionDay,omitempty"`
	RetentionNum       int                                           `json:"retentionNum,omitempty"`
	RegionID           string                                        `json:"regionID,omitempty"`
	CycleType          string                                        `json:"cycleType,omitempty"`
	CycleDay           int                                           `json:"cycleDay,omitempty"`
	CycleWeek          string                                        `json:"cycleWeek,omitempty"`
	PolicyID           string                                        `json:"policyID,omitempty"`
	Time               string                                        `json:"time,omitempty"`
	ProjectID          string                                        `json:"projectID,omitempty"`
	TotalBackup        bool                                          `json:"totalBackup,omitempty"`
	AdvRetentionStatus bool                                          `json:"advRetentionStatus,omitempty"`
	AdvRetention       EcsBackupPolicyUpdateAdvRetentionRealResponse `json:"advRetention,omitempty"`
}

type EcsBackupPolicyUpdateAdvRetentionResponse struct {
	AdvDay   int
	AdvWeek  int
	AdvMonth int
	AdvYear  int
}

type EcsBackupPolicyUpdateResponse struct {
	Status             int
	PolicyName         string
	RetentionType      string
	RetentionDay       int
	RetentionNum       int
	RegionID           string
	CycleType          string
	CycleDay           int
	CycleWeek          string
	PolicyID           string
	Time               string
	ProjectID          string
	TotalBackup        bool
	AdvRetentionStatus bool
	AdvRetention       EcsBackupPolicyUpdateAdvRetentionResponse
}
