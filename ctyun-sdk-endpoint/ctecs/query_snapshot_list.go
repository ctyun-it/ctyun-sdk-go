package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// QuerySnapshotListApi
type QuerySnapshotListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewQuerySnapshotListApi(client *ctyunsdk.CtyunClient) *QuerySnapshotListApi {
	return &QuerySnapshotListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/list",
		},
	}
}

func (this *QuerySnapshotListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *QuerySnapshotListRequest) (*QuerySnapshotListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&QuerySnapshotListRealRequest{
		RegionID:       req.RegionID,
		ProjectID:      req.ProjectID,
		PageNo:         req.PageNo,
		PageSize:       req.PageSize,
		InstanceID:     req.InstanceID,
		SnapshotStatus: req.SnapshotStatus,
		SnapshotID:     req.SnapshotID,
		QueryContent:   req.QueryContent,
		SnapshotName:   req.SnapshotName,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse QuerySnapshotListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []QuerySnapshotListResultsResponse
	for _, res := range realResponse.Results {
		var members []QuerySnapshotListMembersResponse
		for _, member := range res.Members {
			members = append(members, QuerySnapshotListMembersResponse{
				DiskType:           member.DiskType,
				DiskID:             member.DiskID,
				DiskName:           member.DiskName,
				IsBootable:         member.IsBootable,
				IsEncrypt:          member.IsEncrypt,
				DiskSize:           member.DiskSize,
				DiskSnapshotID:     member.DiskSnapshotID,
				DiskSnapshotStatus: member.DiskSnapshotStatus,
			})
		}
		results = append(results, QuerySnapshotListResultsResponse{
			SnapshotID:          res.SnapshotID,
			InstanceID:          res.InstanceID,
			InstanceName:        res.InstanceName,
			AzName:              res.AzName,
			SnapshotName:        res.SnapshotName,
			InstanceStatus:      res.InstanceStatus,
			SnapshotStatus:      res.SnapshotStatus,
			SnapshotDescription: res.SnapshotDescription,
			ProjectID:           res.ProjectID,
			CreatedTime:         res.CreatedTime,
			UpdatedTime:         res.UpdatedTime,
			ImageID:             res.ImageID,
			Memory:              res.Memory,
			Cpu:                 res.Cpu,
			FlavorID:            res.FlavorID,
			Members:             members,
		})
	}

	return &QuerySnapshotListResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type QuerySnapshotListRealRequest struct {
	RegionID       string `json:"regionID,omitempty"`
	ProjectID      string `json:"projectID,omitempty"`
	PageNo         int    `json:"pageNo,omitempty"`
	PageSize       int    `json:"pageSize,omitempty"`
	InstanceID     string `json:"instanceID,omitempty"`
	SnapshotStatus string `json:"snapshotStatus,omitempty"`
	SnapshotID     string `json:"snapshotID,omitempty"`
	QueryContent   string `json:"queryContent,omitempty"`
	SnapshotName   string `json:"snapshotName,omitempty"`
}

type QuerySnapshotListRequest struct {
	RegionID       string
	ProjectID      string
	PageNo         int
	PageSize       int
	InstanceID     string
	SnapshotStatus string
	SnapshotID     string
	QueryContent   string
	SnapshotName   string
}

type QuerySnapshotListMembersRealResponse struct {
	DiskType           string `json:"diskType,omitempty"`
	DiskID             string `json:"diskID,omitempty"`
	DiskName           string `json:"diskName,omitempty"`
	IsBootable         bool   `json:"isBootable,omitempty"`
	IsEncrypt          bool   `json:"isEncrypt,omitempty"`
	DiskSize           int    `json:"diskSize,omitempty"`
	DiskSnapshotID     string `json:"diskSnapshotID,omitempty"`
	DiskSnapshotStatus string `json:"diskSnapshotStatus,omitempty"`
}

type QuerySnapshotListResultsRealResponse struct {
	SnapshotID          string                                 `json:"snapshotID,omitempty"`
	InstanceID          string                                 `json:"instanceID,omitempty"`
	InstanceName        string                                 `json:"instanceName,omitempty"`
	AzName              string                                 `json:"azName,omitempty"`
	SnapshotName        string                                 `json:"snapshotName,omitempty"`
	InstanceStatus      string                                 `json:"instanceStatus,omitempty"`
	SnapshotStatus      string                                 `json:"snapshotStatus,omitempty"`
	SnapshotDescription string                                 `json:"snapshotDescription,omitempty"`
	ProjectID           string                                 `json:"projectID,omitempty"`
	CreatedTime         string                                 `json:"createdTime,omitempty"`
	UpdatedTime         string                                 `json:"updatedTime,omitempty"`
	ImageID             string                                 `json:"imageID,omitempty"`
	Memory              int                                    `json:"memory,omitempty"`
	Cpu                 int                                    `json:"cpu,omitempty"`
	FlavorID            string                                 `json:"flavorID,omitempty"`
	Members             []QuerySnapshotListMembersRealResponse `json:"members,omitempty"`
}

type QuerySnapshotListRealResponse struct {
	CurrentCount int                                    `json:"currentCount,omitempty"`
	TotalCount   int                                    `json:"totalCount,omitempty"`
	TotalPage    int                                    `json:"totalPage,omitempty"`
	Results      []QuerySnapshotListResultsRealResponse `json:"results,omitempty"`
}

type QuerySnapshotListMembersResponse struct {
	DiskType           string
	DiskID             string
	DiskName           string
	IsBootable         bool
	IsEncrypt          bool
	DiskSize           int
	DiskSnapshotID     string
	DiskSnapshotStatus string
}

type QuerySnapshotListResultsResponse struct {
	SnapshotID          string
	InstanceID          string
	InstanceName        string
	AzName              string
	SnapshotName        string
	InstanceStatus      string
	SnapshotStatus      string
	SnapshotDescription string
	ProjectID           string
	CreatedTime         string
	UpdatedTime         string
	ImageID             string
	Memory              int
	Cpu                 int
	FlavorID            string
	Members             []QuerySnapshotListMembersResponse
}

type QuerySnapshotListResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []QuerySnapshotListResultsResponse
}
