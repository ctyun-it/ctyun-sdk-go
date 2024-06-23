package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupStatusApi
type EcsBackupStatusApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupStatusApi(client *ctyunsdk.CtyunClient) *EcsBackupStatusApi {
	return &EcsBackupStatusApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup/status",
		},
	}
}

func (this *EcsBackupStatusApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupStatusRequest) (*EcsBackupStatusResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupStatusRealRequest{
		RegionID:         req.RegionID,
		InstanceBackupID: req.InstanceBackupID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupStatusRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupStatusResponse{
		InstanceBackupStatus: realResponse.InstanceBackupStatus,
	}, nil
}

type EcsBackupStatusRealRequest struct {
	RegionID         string `json:"regionID,omitempty"`
	InstanceBackupID string `json:"instanceBackupID,omitempty"`
}

type EcsBackupStatusRequest struct {
	RegionID         string
	InstanceBackupID string
}

type EcsBackupStatusRealResponse struct {
	InstanceBackupStatus string `json:"instanceBackupStatus,omitempty"`
}

type EcsBackupStatusResponse struct {
	InstanceBackupStatus string
}
