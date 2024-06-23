package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// QuerySnapshotDetailsApi
type QuerySnapshotDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewQuerySnapshotDetailsApi(client *ctyunsdk.CtyunClient) *QuerySnapshotDetailsApi {
	return &QuerySnapshotDetailsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/snapshot/details",
		},
	}
}

func (this *QuerySnapshotDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *QuerySnapshotDetailsRequest) (*QuerySnapshotDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&QuerySnapshotDetailsRealRequest{
		RegionID:   req.RegionID,
		SnapshotID: req.SnapshotID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse QuerySnapshotDetailsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []QuerySnapshotDetailsResultsResponse
	for _, res := range realResponse.Results {
		var members []QuerySnapshotDetailsMembersResponse
		for _, member := range res.Members {
			members = append(members, QuerySnapshotDetailsMembersResponse{
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
		results = append(results, QuerySnapshotDetailsResultsResponse{
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

	return &QuerySnapshotDetailsResponse{
		Results: results,
	}, nil
}

type QuerySnapshotDetailsRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	SnapshotID string `json:"snapshotID,omitempty"`
}

type QuerySnapshotDetailsRequest struct {
	RegionID   string
	SnapshotID string
}

type QuerySnapshotDetailsMembersRealResponse struct {
	DiskType           string `json:"diskType,omitempty"`
	DiskID             string `json:"diskID,omitempty"`
	DiskName           string `json:"diskName,omitempty"`
	IsBootable         bool   `json:"isBootable,omitempty"`
	IsEncrypt          bool   `json:"isEncrypt,omitempty"`
	DiskSize           int    `json:"diskSize,omitempty"`
	DiskSnapshotID     string `json:"diskSnapshotID,omitempty"`
	DiskSnapshotStatus string `json:"diskSnapshotStatus,omitempty"`
}

type QuerySnapshotDetailsResultsRealResponse struct {
	SnapshotID          string                                    `json:"snapshotID,omitempty"`
	InstanceID          string                                    `json:"instanceID,omitempty"`
	InstanceName        string                                    `json:"instanceName,omitempty"`
	AzName              string                                    `json:"azName,omitempty"`
	SnapshotName        string                                    `json:"snapshotName,omitempty"`
	InstanceStatus      string                                    `json:"instanceStatus,omitempty"`
	SnapshotStatus      string                                    `json:"snapshotStatus,omitempty"`
	SnapshotDescription string                                    `json:"snapshotDescription,omitempty"`
	ProjectID           string                                    `json:"projectID,omitempty"`
	CreatedTime         string                                    `json:"createdTime,omitempty"`
	UpdatedTime         string                                    `json:"updatedTime,omitempty"`
	ImageID             string                                    `json:"imageID,omitempty"`
	Memory              int                                       `json:"memory,omitempty"`
	Cpu                 int                                       `json:"cpu,omitempty"`
	FlavorID            string                                    `json:"flavorID,omitempty"`
	Members             []QuerySnapshotDetailsMembersRealResponse `json:"members,omitempty"`
}

type QuerySnapshotDetailsRealResponse struct {
	Results []QuerySnapshotDetailsResultsRealResponse `json:"results,omitempty"`
}

type QuerySnapshotDetailsMembersResponse struct {
	DiskType           string
	DiskID             string
	DiskName           string
	IsBootable         bool
	IsEncrypt          bool
	DiskSize           int
	DiskSnapshotID     string
	DiskSnapshotStatus string
}

type QuerySnapshotDetailsResultsResponse struct {
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
	Members             []QuerySnapshotDetailsMembersResponse
}

type QuerySnapshotDetailsResponse struct {
	Results []QuerySnapshotDetailsResultsResponse
}
