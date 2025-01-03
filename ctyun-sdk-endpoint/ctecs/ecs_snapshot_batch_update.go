package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsSnapshotBatchUpdateApi
type EcsSnapshotBatchUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsSnapshotBatchUpdateApi(client *ctyunsdk.CtyunClient) *EcsSnapshotBatchUpdateApi {
	return &EcsSnapshotBatchUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/batch-update",
		},
	}
}

func (this *EcsSnapshotBatchUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsSnapshotBatchUpdateRequest) (*EcsSnapshotBatchUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var updateInfo []EcsSnapshotBatchUpdateUpdateInfoRealRequest
	for _, request := range req.UpdateInfo {
		updateInfo = append(updateInfo, EcsSnapshotBatchUpdateUpdateInfoRealRequest{
			SnapshotID:          request.SnapshotID,
			SnapshotName:        request.SnapshotName,
			SnapshotDescription: request.SnapshotDescription,
		})
	}

	_, err := builder.WriteJson(&EcsSnapshotBatchUpdateRealRequest{
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

	var realResponse EcsSnapshotBatchUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsSnapshotBatchUpdateResponse{
		SnapshotIDList: realResponse.SnapshotIDList,
	}, nil
}

type EcsSnapshotBatchUpdateUpdateInfoRealRequest struct {
	SnapshotID          *string `json:"snapshotID,omitempty"`
	SnapshotName        *string `json:"snapshotName,omitempty"`
	SnapshotDescription *string `json:"snapshotDescription,omitempty"`
}

type EcsSnapshotBatchUpdateRealRequest struct {
	RegionID   *string                                       `json:"regionID,omitempty"`
	UpdateInfo []EcsSnapshotBatchUpdateUpdateInfoRealRequest `json:"updateInfo,omitempty"`
}

type EcsSnapshotBatchUpdateUpdateInfoRequest struct {
	SnapshotID          *string
	SnapshotName        *string
	SnapshotDescription *string
}

type EcsSnapshotBatchUpdateRequest struct {
	RegionID   *string
	UpdateInfo []EcsSnapshotBatchUpdateUpdateInfoRequest
}

type EcsSnapshotBatchUpdateRealResponse struct {
	SnapshotIDList []string `json:"snapshotIDList,omitempty"`
}

type EcsSnapshotBatchUpdateResponse struct {
	SnapshotIDList []string
}
