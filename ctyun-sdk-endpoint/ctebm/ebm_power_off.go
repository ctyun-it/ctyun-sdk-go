package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmPowerOffApi 物理机关机
// https://www.ctyun.cn/document/10027724/10040134
type EbmPowerOffApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmPowerOffApi(client *ctyunsdk.CtyunClient) *EbmPowerOffApi {
	return &EbmPowerOffApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/power-off",
		},
	}
}

func (this *EbmPowerOffApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmPowerOffRequest) (*EbmPowerOffResponse, ctyunsdk.CtyunRequestError) {
	// 构造鉴权参数
	builder := this.WithCredential(&credential)

	// 请求OpenAPI服务
	_, err := builder.WriteJson(&ebmPowerOffRealRequest{
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

	var realResponse ebmPowerOffRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EbmPowerOffResponse{}, nil
}

// 实际调用OpenAPI接口请求参数，参数名首字母大写
type ebmPowerOffRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	AzName       string `json:"azName,omitempty"`
	InstanceUUID string `json:"instanceUUID,omitempty"`
}

// 对外提供接口请求参数
type EbmPowerOffRequest struct {
	RegionID     string
	AzName       string
	InstanceUUID string
}

// 实际调用OpenAPI接口响应参数
type ebmPowerOffRealResponse struct{}

// 对外提供接口响应参数
type EbmPowerOffResponse struct{}
