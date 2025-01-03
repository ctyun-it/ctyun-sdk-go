package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotDetailsApi
type EcsSnapshotDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotDetailsApi(client *ctyunsdk.CtyunClient) *EcsSnapshotDetailsApi {
	return &EcsSnapshotDetailsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/snapshot/details",
		},
	}
}

func (this *EcsSnapshotDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotDetailsRequest) (*EcsSnapshotDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	builder.AddParam("regionID", *req.RegionID).
		AddParam("snapshotID", *req.SnapshotID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsSnapshotDetailsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EcsSnapshotDetailsResultsResponse
	for _, res := range realResponse.Results {
		var members []EcsSnapshotDetailsMembersResponse
		for _, member := range res.Members {
			members = append(members, EcsSnapshotDetailsMembersResponse{
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
		results = append(results, EcsSnapshotDetailsResultsResponse{
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

	return &EcsSnapshotDetailsResponse{
		Results: results,
	}, nil
}

type EcsSnapshotDetailsRealRequest struct {
	RegionID   *string `json:"regionID,omitempty"`
	SnapshotID *string `json:"snapshotID,omitempty"`
}

type EcsSnapshotDetailsRequest struct {
	RegionID   *string
	SnapshotID *string
}

type EcsSnapshotDetailsMembersRealResponse struct {
	DiskType           string `json:"diskType,omitempty"`
	DiskID             string `json:"diskID,omitempty"`
	DiskName           string `json:"diskName,omitempty"`
	IsBootable         bool   `json:"isBootable,omitempty"`
	IsEncrypt          bool   `json:"isEncrypt,omitempty"`
	DiskSize           int    `json:"diskSize,omitempty"`
	DiskSnapshotID     string `json:"diskSnapshotID,omitempty"`
	DiskSnapshotStatus string `json:"diskSnapshotStatus,omitempty"`
}

type EcsSnapshotDetailsResultsRealResponse struct {
	SnapshotID          string                                  `json:"snapshotID,omitempty"`
	InstanceID          string                                  `json:"instanceID,omitempty"`
	InstanceName        string                                  `json:"instanceName,omitempty"`
	AzName              string                                  `json:"azName,omitempty"`
	SnapshotName        string                                  `json:"snapshotName,omitempty"`
	InstanceStatus      string                                  `json:"instanceStatus,omitempty"`
	SnapshotStatus      string                                  `json:"snapshotStatus,omitempty"`
	SnapshotDescription string                                  `json:"snapshotDescription,omitempty"`
	ProjectID           string                                  `json:"projectID,omitempty"`
	CreatedTime         string                                  `json:"createdTime,omitempty"`
	UpdatedTime         string                                  `json:"updatedTime,omitempty"`
	ImageID             string                                  `json:"imageID,omitempty"`
	Memory              int                                     `json:"memory,omitempty"`
	Cpu                 int                                     `json:"cpu,omitempty"`
	FlavorID            string                                  `json:"flavorID,omitempty"`
	Members             []EcsSnapshotDetailsMembersRealResponse `json:"members,omitempty"`
}

type EcsSnapshotDetailsRealResponse struct {
	Results []EcsSnapshotDetailsResultsRealResponse `json:"results,omitempty"`
}

type EcsSnapshotDetailsMembersResponse struct {
	DiskType           string
	DiskID             string
	DiskName           string
	IsBootable         bool
	IsEncrypt          bool
	DiskSize           int
	DiskSnapshotID     string
	DiskSnapshotStatus string
}

type EcsSnapshotDetailsResultsResponse struct {
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
	Members             []EcsSnapshotDetailsMembersResponse
}

type EcsSnapshotDetailsResponse struct {
	Results []EcsSnapshotDetailsResultsResponse
}
