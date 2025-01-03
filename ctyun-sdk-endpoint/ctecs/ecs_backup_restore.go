package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupRestoreApi
type EcsBackupRestoreApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupRestoreApi(client *ctyunsdk.CtyunClient) *EcsBackupRestoreApi {
	return &EcsBackupRestoreApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-restore",
		},
	}
}

func (this *EcsBackupRestoreApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupRestoreRequest) (*EcsBackupRestoreResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupRestoreRealRequest{
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

	var realResponse EcsBackupRestoreRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupRestoreResponse{
		InstanceBackupID: realResponse.InstanceBackupID,
	}, nil
}

type EcsBackupRestoreRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	InstanceBackupID *string `json:"instanceBackupID,omitempty"`
}

type EcsBackupRestoreRequest struct {
	RegionID         *string
	InstanceBackupID *string
}

type EcsBackupRestoreRealResponse struct {
	InstanceBackupID string `json:"instanceBackupID,omitempty"`
}

type EcsBackupRestoreResponse struct {
	InstanceBackupID string
}
