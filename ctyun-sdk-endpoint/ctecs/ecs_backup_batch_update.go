package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBackupBatchUpdateApi
type EcsBackupBatchUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBackupBatchUpdateApi(client *ctyunsdk.CtyunClient) *EcsBackupBatchUpdateApi {
	return &EcsBackupBatchUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/backup/batch-update",
		},
	}
}

func (this *EcsBackupBatchUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBackupBatchUpdateRequest) (*EcsBackupBatchUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var updateInfo []EcsBackupBatchUpdateUpdateInfoRealRequest
	for _, request := range req.UpdateInfo {
		updateInfo = append(updateInfo, EcsBackupBatchUpdateUpdateInfoRealRequest{
			InstanceBackupID:          request.InstanceBackupID,
			InstanceBackupName:        request.InstanceBackupName,
			InstanceBackupDescription: request.InstanceBackupDescription,
		})
	}

	_, err := builder.WriteJson(&EcsBackupBatchUpdateRealRequest{
		RegionID:   req.RegionID,
		UpdateInfo: updateInfo,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsBackupBatchUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBackupBatchUpdateResponse{
		InstanceBackupIDList: realResponse.InstanceBackupIDList,
	}, nil
}

type EcsBackupBatchUpdateUpdateInfoRealRequest struct {
	InstanceBackupID          string `json:"instanceBackupID,omitempty"`
	InstanceBackupName        string `json:"instanceBackupName,omitempty"`
	InstanceBackupDescription string `json:"instanceBackupDescription,omitempty"`
}

type EcsBackupBatchUpdateRealRequest struct {
	RegionID   string                                      `json:"regionID,omitempty"`
	UpdateInfo []EcsBackupBatchUpdateUpdateInfoRealRequest `json:"updateInfo,omitempty"`
}

type EcsBackupBatchUpdateUpdateInfoRequest struct {
	InstanceBackupID          string
	InstanceBackupName        string
	InstanceBackupDescription string
}

type EcsBackupBatchUpdateRequest struct {
	RegionID   string
	UpdateInfo []EcsBackupBatchUpdateUpdateInfoRequest
}

type EcsBackupBatchUpdateRealResponse struct {
	InstanceBackupIDList []string `json:"instanceBackupIDList,omitempty"`
}

type EcsBackupBatchUpdateResponse struct {
	InstanceBackupIDList []string
}
