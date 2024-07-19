package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmChangePasswordApi D:\Project\go-sdk-auto-write\docs\物理机密码变更
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4584&data=97&isNormal=1
type EbmChangePasswordApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmChangePasswordApi(client *ctyunsdk.CtyunClient) *EbmChangePasswordApi {
	return &EbmChangePasswordApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/change-password",
		},
	}
}

func (this *EbmChangePasswordApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmChangePasswordRequest) (*EbmChangePasswordResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EbmChangePasswordRealRequest{
		RegionID:     req.RegionID,
		AzName:       req.AzName,
		InstanceUUID: req.InstanceUUID,
		NewPassword:  req.NewPassword,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmChangePasswordRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EbmChangePasswordResponse{}, nil
}

type EbmChangePasswordRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	AzName       string `json:"azName,omitempty"`
	InstanceUUID string `json:"instanceUUID,omitempty"`
	NewPassword  string `json:"newPassword,omitempty"`
}

type EbmChangePasswordRequest struct {
	RegionID     string
	AzName       string
	InstanceUUID string
	NewPassword  string
}

type EbmChangePasswordRealResponse struct {
}

type EbmChangePasswordResponse struct {
}
