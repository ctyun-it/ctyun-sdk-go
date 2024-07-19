package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmPowerOnApi 物理机开机
// https://www.ctyun.cn/document/10027724/10040133
type EbmPowerOnApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmPowerOnApi(client *ctyunsdk.CtyunClient) *EbmPowerOnApi {
	return &EbmPowerOnApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/power-on",
		},
	}
}

func (this *EbmPowerOnApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmPowerOnRequest) (*EbmPowerOnResponse, ctyunsdk.CtyunRequestError) {
	// 鉴权
	builder := this.WithCredential(&credential)

	// 请求OpenAPI服务
	_, err := builder.WriteJson(&ebmPowerOnRealRequest{
		RegionID:     req.RegionID,
		AzName:       req.AzName,
		InstanceUUID: req.InstanceUUID,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse ebmPowerOnRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EbmPowerOnResponse{}, nil
}

type ebmPowerOnRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	AzName       string `json:"azName,omitempty"`
	InstanceUUID string `json:"instanceUUID,omitempty"`
}

type EbmPowerOnRequest struct {
	RegionID     string
	AzName       string
	InstanceUUID string
}

type ebmPowerOnRealResponse struct{}

type EbmPowerOnResponse struct{}
