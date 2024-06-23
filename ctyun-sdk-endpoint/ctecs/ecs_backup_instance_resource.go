package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupInstanceResourceApi
type EcsBackupInstanceResourceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupInstanceResourceApi(client *ctyunsdk.CtyunClient) *EcsBackupInstanceResourceApi {
	return &EcsBackupInstanceResourceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-instance-resource",
		},
	}
}

func (this *EcsBackupInstanceResourceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupInstanceResourceRequest) (*EcsBackupInstanceResourceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupInstanceResourceRealRequest{
		RegionID: req.RegionID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupInstanceResourceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupInstanceResourceResponse{
		TotalVolumeSize:  realResponse.TotalVolumeSize,
		TotalBackupCount: realResponse.TotalBackupCount,
	}, nil
}

type EcsBackupInstanceResourceRealRequest struct {
	RegionID string `json:"regionID,omitempty"`
}

type EcsBackupInstanceResourceRequest struct {
	RegionID string
}

type EcsBackupInstanceResourceRealResponse struct {
	TotalVolumeSize  int `json:"totalVolumeSize,omitempty"`
	TotalBackupCount int `json:"totalBackupCount,omitempty"`
}

type EcsBackupInstanceResourceResponse struct {
	TotalVolumeSize  int
	TotalBackupCount int
}
