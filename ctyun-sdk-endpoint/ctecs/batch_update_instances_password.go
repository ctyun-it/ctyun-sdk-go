package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// BatchUpdateInstancesPasswordApi
type BatchUpdateInstancesPasswordApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewBatchUpdateInstancesPasswordApi(client *ctyunsdk.CtyunClient) *BatchUpdateInstancesPasswordApi {
	return &BatchUpdateInstancesPasswordApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/batch-reset-password",
		},
	}
}

func (this *BatchUpdateInstancesPasswordApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *BatchUpdateInstancesPasswordRequest) (*BatchUpdateInstancesPasswordResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var updatePwdInfo []BatchUpdateInstancesPasswordUpdatePwdInfoRealRequest
	for _, request := range req.UpdatePwdInfo {
		updatePwdInfo = append(updatePwdInfo, BatchUpdateInstancesPasswordUpdatePwdInfoRealRequest{
			InstanceID:  request.InstanceID,
			NewPassword: request.NewPassword,
		})
	}

	_, err := builder.WriteJson(&BatchUpdateInstancesPasswordRealRequest{
		RegionID:      req.RegionID,
		UpdatePwdInfo: updatePwdInfo,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse BatchUpdateInstancesPasswordRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &BatchUpdateInstancesPasswordResponse{
		InstanceIDList: realResponse.InstanceIDList,
	}, nil
}

type BatchUpdateInstancesPasswordUpdatePwdInfoRealRequest struct {
	InstanceID  string `json:"instanceID,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type BatchUpdateInstancesPasswordRealRequest struct {
	RegionID      string                                                 `json:"regionID,omitempty"`
	UpdatePwdInfo []BatchUpdateInstancesPasswordUpdatePwdInfoRealRequest `json:"updatePwdInfo,omitempty"`
}

type BatchUpdateInstancesPasswordUpdatePwdInfoRequest struct {
	InstanceID  string
	NewPassword string
}

type BatchUpdateInstancesPasswordRequest struct {
	RegionID      string
	UpdatePwdInfo []BatchUpdateInstancesPasswordUpdatePwdInfoRequest
}

type BatchUpdateInstancesPasswordRealResponse struct {
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type BatchUpdateInstancesPasswordResponse struct {
	InstanceIDList string
}
