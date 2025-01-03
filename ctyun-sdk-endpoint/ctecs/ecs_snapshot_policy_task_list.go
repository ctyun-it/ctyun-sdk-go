package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotPolicyTaskListApi 查询云主机快照任务列表
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=9598&data=87&isNormal=1
type EcsSnapshotPolicyTaskListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotPolicyTaskListApi(client *ctyunsdk.CtyunClient) *EcsSnapshotPolicyTaskListApi {
	return &EcsSnapshotPolicyTaskListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot-policy/task-list",
		},
	}
}

func (this *EcsSnapshotPolicyTaskListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotPolicyTaskListRequest) (*EcsSnapshotPolicyTaskListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsSnapshotPolicyTaskListRealRequest{
		RegionID:         req.RegionID,
		SnapshotPolicyID: req.SnapshotPolicyID,
		PageNo:           req.PageNo,
		PageSize:         req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotPolicyTaskListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var taskList []EcsSnapshotPolicyTaskListTaskListResponse
	for _, res := range realResponse.TaskList {
		taskList = append(taskList, EcsSnapshotPolicyTaskListTaskListResponse{
			TaskID:         res.TaskID,
			TaskStatus:     res.TaskStatus,
			SnapshotStatus: res.SnapshotStatus,
			InstanceID:     res.InstanceID,
			SnapshotID:     res.SnapshotID,
			SnapshotName:   res.SnapshotName,
			CreateTime:     res.CreateTime,
			CompleteTime:   res.CompleteTime,
		})
	}

	return &EcsSnapshotPolicyTaskListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		TaskList:     taskList,
	}, nil
}

type EcsSnapshotPolicyTaskListRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	SnapshotPolicyID *string `json:"snapshotPolicyID,omitempty"`
	PageNo           *int    `json:"pageNo,omitempty"`
	PageSize         *int    `json:"pageSize,omitempty"`
}

type EcsSnapshotPolicyTaskListRequest struct {
	RegionID         *string
	SnapshotPolicyID *string
	PageNo           *int
	PageSize         *int
}

type EcsSnapshotPolicyTaskListTaskListRealResponse struct {
	TaskID         string `json:"taskID,omitempty"`
	TaskStatus     string `json:"taskStatus,omitempty"`
	SnapshotStatus string `json:"snapshotStatus,omitempty"`
	InstanceID     string `json:"instanceID,omitempty"`
	SnapshotID     string `json:"snapshotID,omitempty"`
	SnapshotName   string `json:"snapshotName,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
	CompleteTime   string `json:"completeTime,omitempty"`
}

type EcsSnapshotPolicyTaskListRealResponse struct {
	CurrentCount int                                             `json:"currentCount,omitempty"`
	TotalCount   int                                             `json:"totalCount,omitempty"`
	TotalPage    int                                             `json:"totalPage,omitempty"`
	TaskList     []EcsSnapshotPolicyTaskListTaskListRealResponse `json:"taskList,omitempty"`
}

type EcsSnapshotPolicyTaskListTaskListResponse struct {
	TaskID         string
	TaskStatus     string
	SnapshotStatus string
	InstanceID     string
	SnapshotID     string
	SnapshotName   string
	CreateTime     string
	CompleteTime   string
}

type EcsSnapshotPolicyTaskListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	TaskList     []EcsSnapshotPolicyTaskListTaskListResponse
}
