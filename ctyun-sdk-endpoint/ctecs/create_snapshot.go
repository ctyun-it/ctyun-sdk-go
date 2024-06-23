package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// CreateSnapshotApi
type CreateSnapshotApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewCreateSnapshotApi(client *ctyunsdk.CtyunClient) *CreateSnapshotApi {
	return &CreateSnapshotApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/create",
		},
	}
}

func (this *CreateSnapshotApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *CreateSnapshotRequest) (*CreateSnapshotResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&CreateSnapshotRealRequest{
		RegionID:     req.RegionID,
		InstanceID:   req.InstanceID,
		SnapshotName: req.SnapshotName,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse CreateSnapshotRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &CreateSnapshotResponse{
		JobID:          realResponse.JobID,
		SnapshotStatus: realResponse.SnapshotStatus,
		InstanceID:     realResponse.InstanceID,
		InstanceName:   realResponse.InstanceName,
		SnapshotID:     realResponse.SnapshotID,
		ProjectID:      realResponse.ProjectID,
		SnapshotName:   realResponse.SnapshotName,
	}, nil
}

type CreateSnapshotRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	InstanceID   string `json:"instanceID,omitempty"`
	SnapshotName string `json:"snapshotName,omitempty"`
}

type CreateSnapshotRequest struct {
	RegionID     string
	InstanceID   string
	SnapshotName string
}

type CreateSnapshotRealResponse struct {
	JobID          string `json:"jobID,omitempty"`
	SnapshotStatus string `json:"snapshotStatus,omitempty"`
	InstanceID     string `json:"instanceID,omitempty"`
	InstanceName   string `json:"instanceName,omitempty"`
	SnapshotID     string `json:"snapshotID,omitempty"`
	ProjectID      string `json:"projectID,omitempty"`
	SnapshotName   string `json:"snapshotName,omitempty"`
}

type CreateSnapshotResponse struct {
	JobID          string
	SnapshotStatus string
	InstanceID     string
	InstanceName   string
	SnapshotID     string
	ProjectID      string
	SnapshotName   string
}
