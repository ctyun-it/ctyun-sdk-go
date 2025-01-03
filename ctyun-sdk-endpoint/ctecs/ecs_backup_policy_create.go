package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupPolicyCreateApi 创建云主机备份策略
type EcsBackupPolicyCreateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

// NewEcsBackupPolicyCreateApi 创建 EcsBackupPolicyCreateApi 实例
func NewEcsBackupPolicyCreateApi(client *ctyunsdk.CtyunClient) *EcsBackupPolicyCreateApi {
	return &EcsBackupPolicyCreateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-policy/create",
		},
	}
}

// Do 执行备份策略创建操作
func (api *EcsBackupPolicyCreateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupPolicyCreateRequest) (*EcsBackupPolicyCreateResponse, ctyunsdk.CtyunRequestError) {
	builder := api.WithCredential(&credential)

	// 构建请求体
	realRequest := &EcsBackupPolicyCreateRealRequest{
		RegionID:           req.RegionID,
		PolicyName:         req.PolicyName,
		CycleType:          req.CycleType,
		CycleDay:           req.CycleDay,
		CycleWeek:          req.CycleWeek,
		Time:               req.Time,
		Status:             req.Status,
		RetentionType:      req.RetentionType,
		RetentionDay:       req.RetentionDay,
		RetentionNum:       req.RetentionNum,
		ProjectID:          req.ProjectID,
		TotalBackup:        req.TotalBackup,
		AdvRetentionStatus: req.AdvRetentionStatus,
		AdvRetention: EcsBackupPolicyCreateAdvRetentionRealRequest{
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
	var realResponse EcsBackupPolicyCreateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	// 构建返回结果
	return &EcsBackupPolicyCreateResponse{
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
		AdvRetention: EcsBackupPolicyCreateAdvRetentionResponse{
			AdvDay:   realResponse.AdvRetention.AdvDay,
			AdvWeek:  realResponse.AdvRetention.AdvWeek,
			AdvMonth: realResponse.AdvRetention.AdvMonth,
			AdvYear:  realResponse.AdvRetention.AdvYear,
		},
	}, nil
}

// 其他结构体定义
type EcsBackupPolicyCreateAdvRetentionRealRequest struct {
	AdvDay   *int `json:"advDay,omitempty"`
	AdvWeek  *int `json:"advWeek,omitempty"`
	AdvMonth *int `json:"advMonth,omitempty"`
	AdvYear  *int `json:"advYear,omitempty"`
}

type EcsBackupPolicyCreateRealRequest struct {
	RegionID           *string                                      `json:"regionID,omitempty"`
	PolicyName         *string                                      `json:"policyName,omitempty"`
	CycleType          *string                                      `json:"cycleType,omitempty"`
	CycleDay           *int                                         `json:"cycleDay,omitempty"`
	CycleWeek          *string                                      `json:"cycleWeek,omitempty"`
	Time               *string                                      `json:"time,omitempty"`
	Status             *int                                         `json:"status,omitempty"`
	RetentionType      *string                                      `json:"retentionType,omitempty"`
	RetentionDay       *int                                         `json:"retentionDay,omitempty"`
	RetentionNum       *int                                         `json:"retentionNum,omitempty"`
	ProjectID          *string                                      `json:"projectID,omitempty"`
	TotalBackup        *bool                                        `json:"totalBackup,omitempty"`
	AdvRetentionStatus *bool                                        `json:"advRetentionStatus,omitempty"`
	AdvRetention       EcsBackupPolicyCreateAdvRetentionRealRequest `json:"advRetention,omitempty"`
}

type EcsBackupPolicyCreateAdvRetentionRequest struct {
	AdvDay   *int
	AdvWeek  *int
	AdvMonth *int
	AdvYear  *int
}

type EcsBackupPolicyCreateRequest struct {
	RegionID           *string
	PolicyName         *string
	CycleType          *string
	CycleDay           *int
	CycleWeek          *string
	Time               *string
	Status             *int
	RetentionType      *string
	RetentionDay       *int
	RetentionNum       *int
	ProjectID          *string
	TotalBackup        *bool
	AdvRetentionStatus *bool
	AdvRetention       EcsBackupPolicyCreateAdvRetentionRequest
}

type EcsBackupPolicyCreateAdvRetentionRealResponse struct {
	AdvDay   int `json:"advDay,omitempty"`
	AdvWeek  int `json:"advWeek,omitempty"`
	AdvMonth int `json:"advMonth,omitempty"`
	AdvYear  int `json:"advYear,omitempty"`
}

type EcsBackupPolicyCreateRealResponse struct {
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
	AdvRetention       EcsBackupPolicyCreateAdvRetentionRealResponse `json:"advRetention,omitempty"`
}

type EcsBackupPolicyCreateAdvRetentionResponse struct {
	AdvDay   int
	AdvWeek  int
	AdvMonth int
	AdvYear  int
}

type EcsBackupPolicyCreateResponse struct {
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
	AdvRetention       EcsBackupPolicyCreateAdvRetentionResponse
}
