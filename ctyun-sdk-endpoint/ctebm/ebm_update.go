package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmUpdateApi D:\Project\go-sdk-auto-write\docs\物理机更新
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4912&data=97&isNormal=1
type EbmUpdateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmUpdateApi(client *ctyunsdk.CtyunClient) *EbmUpdateApi {
	return &EbmUpdateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/update",
		},
	}
}

func (this *EbmUpdateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmUpdateRequest) (*EbmUpdateResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EbmUpdateRealRequest{
		RegionID:     req.RegionID,
		AzName:       req.AzName,
		DisplayName:  req.DisplayName,
		Description:  req.Description,
		InstanceUUID: req.InstanceUUID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmUpdateRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EbmUpdateResponse{}, nil
}

type EbmUpdateRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	AzName       string `json:"azName,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Description  string `json:"description,omitempty"`
	InstanceUUID string `json:"instanceUUID,omitempty"`
}

type EbmUpdateRequest struct {
	RegionID     string
	AzName       string
	DisplayName  string
	Description  string
	InstanceUUID string
}

type EbmUpdateRealResponse struct {
}

type EbmUpdateResponse struct {
}
