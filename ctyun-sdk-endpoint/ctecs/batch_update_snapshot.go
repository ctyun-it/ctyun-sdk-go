package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// BatchUpdateSnapshotApi
type BatchUpdateSnapshotApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewBatchUpdateSnapshotApi(client *ctyunsdk.CtyunClient) *BatchUpdateSnapshotApi {
	return &BatchUpdateSnapshotApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/snapshot/batch-update",
		},
	}
}

func (this *BatchUpdateSnapshotApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *BatchUpdateSnapshotRequest) (*BatchUpdateSnapshotResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var updateInfo []BatchUpdateSnapshotUpdateInfoRealRequest
	for _, request := range req.UpdateInfo {
		updateInfo = append(updateInfo, BatchUpdateSnapshotUpdateInfoRealRequest{
			SnapshotID:          request.SnapshotID,
			SnapshotName:        request.SnapshotName,
			SnapshotDescription: request.SnapshotDescription,
		})
	}

	_, err := builder.WriteJson(&BatchUpdateSnapshotRealRequest{
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

	var realResponse BatchUpdateSnapshotRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &BatchUpdateSnapshotResponse{
		SnapshotIDList: realResponse.SnapshotIDList,
	}, nil
}

type BatchUpdateSnapshotUpdateInfoRealRequest struct {
	SnapshotID          string `json:"snapshotID,omitempty"`
	SnapshotName        string `json:"snapshotName,omitempty"`
	SnapshotDescription string `json:"snapshotDescription,omitempty"`
}

type BatchUpdateSnapshotRealRequest struct {
	RegionID   string                                     `json:"regionID,omitempty"`
	UpdateInfo []BatchUpdateSnapshotUpdateInfoRealRequest `json:"updateInfo,omitempty"`
}

type BatchUpdateSnapshotUpdateInfoRequest struct {
	SnapshotID          string
	SnapshotName        string
	SnapshotDescription string
}

type BatchUpdateSnapshotRequest struct {
	RegionID   string
	UpdateInfo []BatchUpdateSnapshotUpdateInfoRequest
}

type BatchUpdateSnapshotRealResponse struct {
	SnapshotIDList []string `json:"snapshotIDList,omitempty"`
}

type BatchUpdateSnapshotResponse struct {
	SnapshotIDList []string
}
