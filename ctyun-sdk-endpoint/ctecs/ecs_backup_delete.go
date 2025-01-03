package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupDeleteApi
type EcsBackupDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupDeleteApi(client *ctyunsdk.CtyunClient) *EcsBackupDeleteApi {
	return &EcsBackupDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup-delete",
		},
	}
}

func (this *EcsBackupDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupDeleteRequest) (*EcsBackupDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupDeleteRealRequest{
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

	var realResponse EcsBackupDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupDeleteResponse{
		InstanceBackupID: realResponse.InstanceBackupID,
	}, nil
}

type EcsBackupDeleteRealRequest struct {
	RegionID         *string `json:"regionID,omitempty"`
	InstanceBackupID *string `json:"instanceBackupID,omitempty"`
}

type EcsBackupDeleteRequest struct {
	RegionID         *string
	InstanceBackupID *string
}

type EcsBackupDeleteRealResponse struct {
	InstanceBackupID string `json:"instanceBackupID,omitempty"`
}

type EcsBackupDeleteResponse struct {
	InstanceBackupID string
}
