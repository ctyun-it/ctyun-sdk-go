package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmRebootApi 物理机重启
// https://www.ctyun.cn/document/10027724/10040135
type EbmRebootApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmRebootApi(client *ctyunsdk.CtyunClient) *EbmRebootApi {
	return &EbmRebootApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/reboot",
		},
	}
}

func (this *EbmRebootApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmRebootRequest) (*EbmRebootResponse, ctyunsdk.CtyunRequestError) {
	// 鉴权
	builder := this.WithCredential(&credential)

	// 请求OpenAPI服务
	_, err := builder.WriteJson(&ebmRebootRealRequest{
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

	var realResponse ebmRebootRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EbmRebootResponse{}, nil
}

type ebmRebootRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	AzName       string `json:"azName,omitempty"`
	InstanceUUID string `json:"instanceUUID,omitempty"`
}

type EbmRebootRequest struct {
	RegionID     string
	AzName       string
	InstanceUUID string
}

type ebmRebootRealResponse struct{}

type EbmRebootResponse struct{}
