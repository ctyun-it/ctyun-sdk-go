package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupUpdateApi
type EcsBackupUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupUpdateApi(client *ctyunsdk.CtyunClient) *EcsBackupUpdateApi {
	return &EcsBackupUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup/update",
		},
	}
}

func (this *EcsBackupUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupUpdateRequest) (*EcsBackupUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsBackupUpdateRealRequest{
		RegionID:                  req.RegionID,
		InstanceBackupID:          req.InstanceBackupID,
		InstanceBackupName:        req.InstanceBackupName,
		InstanceBackupDescription: req.InstanceBackupDescription,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupUpdateResponse{
		InstanceBackupID: realResponse.InstanceBackupID,
	}, nil
}

type EcsBackupUpdateRealRequest struct {
	RegionID                  string `json:"regionID,omitempty"`
	InstanceBackupID          string `json:"instanceBackupID,omitempty"`
	InstanceBackupName        string `json:"instanceBackupName,omitempty"`
	InstanceBackupDescription string `json:"instanceBackupDescription,omitempty"`
}

type EcsBackupUpdateRequest struct {
	RegionID                  string
	InstanceBackupID          string
	InstanceBackupName        string
	InstanceBackupDescription string
}

type EcsBackupUpdateRealResponse struct {
	InstanceBackupID string `json:"instanceBackupID,omitempty"`
}

type EcsBackupUpdateResponse struct {
	InstanceBackupID string
}
