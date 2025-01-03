package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupUsageApi
type EcsBackupUsageApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupUsageApi(client *ctyunsdk.CtyunClient) *EcsBackupUsageApi {
	return &EcsBackupUsageApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-usage",
		},
	}
}

func (this *EcsBackupUsageApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupUsageRequest) (*EcsBackupUsageResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupUsageRealRequest{
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

	var realResponse EcsBackupUsageRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupUsageResponse{
		Usage: realResponse.Usage,
	}, nil
}

type EcsBackupUsageRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	InstanceBackupID *string `json:"instanceBackupID,omitempty"`
}

type EcsBackupUsageRequest struct {
	RegionID         *string
	InstanceBackupID *string
}

type EcsBackupUsageRealResponse struct {
	Usage int `json:"usage,omitempty"`
}

type EcsBackupUsageResponse struct {
	Usage int
}
