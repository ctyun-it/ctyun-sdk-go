package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsCreateSnapshotApi
type EcsCreateSnapshotApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsCreateSnapshotApi(client *ctyunsdk.CtyunClient) *EcsCreateSnapshotApi {
	return &EcsCreateSnapshotApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/create",
		},
	}
}

func (this *EcsCreateSnapshotApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsCreateSnapshotRequest) (*EcsCreateSnapshotResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsCreateSnapshotRealRequest{
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

	var realResponse EcsCreateSnapshotRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsCreateSnapshotResponse{
		JobID:          realResponse.JobID,
		SnapshotStatus: realResponse.SnapshotStatus,
		InstanceID:     realResponse.InstanceID,
		InstanceName:   realResponse.InstanceName,
		SnapshotID:     realResponse.SnapshotID,
		ProjectID:      realResponse.ProjectID,
		SnapshotName:   realResponse.SnapshotName,
	}, nil
}

type EcsCreateSnapshotRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	InstanceID   string `json:"instanceID,omitempty"`
	SnapshotName string `json:"snapshotName,omitempty"`
}

type EcsCreateSnapshotRequest struct {
	RegionID     string
	InstanceID   string
	SnapshotName string
}

type EcsCreateSnapshotRealResponse struct {
	JobID          string `json:"jobID,omitempty"`
	SnapshotStatus string `json:"snapshotStatus,omitempty"`
	InstanceID     string `json:"instanceID,omitempty"`
	InstanceName   string `json:"instanceName,omitempty"`
	SnapshotID     string `json:"snapshotID,omitempty"`
	ProjectID      string `json:"projectID,omitempty"`
	SnapshotName   string `json:"snapshotName,omitempty"`
}

type EcsCreateSnapshotResponse struct {
	JobID          string
	SnapshotStatus string
	InstanceID     string
	InstanceName   string
	SnapshotID     string
	ProjectID      string
	SnapshotName   string
}
