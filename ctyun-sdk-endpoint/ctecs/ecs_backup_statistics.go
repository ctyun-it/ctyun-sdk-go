package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupStatisticsApi
type EcsBackupStatisticsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupStatisticsApi(client *ctyunsdk.CtyunClient) *EcsBackupStatisticsApi {
	return &EcsBackupStatisticsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/backup/statistics",
		},
	}
}

func (this *EcsBackupStatisticsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupStatisticsRequest) (*EcsBackupStatisticsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupStatisticsRealRequest{
		RegionID:   req.RegionID,
		InstanceID: req.InstanceID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupStatisticsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupStatisticsResponse{
		TotalDiskSize: realResponse.TotalDiskSize,
	}, nil
}

type EcsBackupStatisticsRealRequest struct {
	RegionID   *string `json:"regionID,omitempty"`
	InstanceID *string `json:"instanceID,omitempty"`
}

type EcsBackupStatisticsRequest struct {
	RegionID   *string
	InstanceID *string
}

type EcsBackupStatisticsRealResponse struct {
	TotalDiskSize int `json:"totalDiskSize,omitempty"`
}

type EcsBackupStatisticsResponse struct {
	TotalDiskSize int
}
