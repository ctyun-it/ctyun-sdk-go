package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsBatchResetPasswordApi
type EcsBatchResetPasswordApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsBatchResetPasswordApi(client *ctyunsdk.CtyunClient) *EcsBatchResetPasswordApi {
	return &EcsBatchResetPasswordApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/batch-reset-password",
		},
	}
}

func (this *EcsBatchResetPasswordApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsBatchResetPasswordRequest) (*EcsBatchResetPasswordResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	var updatePwdInfo []EcsBatchResetPasswordUpdatePwdInfoRealRequest
	for _, request := range req.UpdatePwdInfo {
		updatePwdInfo = append(updatePwdInfo, EcsBatchResetPasswordUpdatePwdInfoRealRequest{
			InstanceID:  request.InstanceID,
			NewPassword: request.NewPassword,
		})
	}

	_, err := builder.WriteJson(&EcsBatchResetPasswordRealRequest{
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

	var realResponse EcsBatchResetPasswordRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsBatchResetPasswordResponse{
		InstanceIDList: realResponse.InstanceIDList,
	}, nil
}

type EcsBatchResetPasswordUpdatePwdInfoRealRequest struct {
	InstanceID  string `json:"instanceID,omitempty"`
	NewPassword string `json:"newPassword,omitempty"`
}

type EcsBatchResetPasswordRealRequest struct {
	RegionID      string                                          `json:"regionID,omitempty"`
	UpdatePwdInfo []EcsBatchResetPasswordUpdatePwdInfoRealRequest `json:"updatePwdInfo,omitempty"`
}

type EcsBatchResetPasswordUpdatePwdInfoRequest struct {
	InstanceID  string
	NewPassword string
}

type EcsBatchResetPasswordRequest struct {
	RegionID      string
	UpdatePwdInfo []EcsBatchResetPasswordUpdatePwdInfoRequest
}

type EcsBatchResetPasswordRealResponse struct {
	InstanceIDList string `json:"instanceIDList,omitempty"`
}

type EcsBatchResetPasswordResponse struct {
	InstanceIDList string
}
