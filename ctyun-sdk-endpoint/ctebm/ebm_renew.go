package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmRenewApi D:\Project\go-sdk-auto-write\docs\物理机续订
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4910&data=97&isNormal=1
type EbmRenewApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmRenewApi(client *ctyunsdk.CtyunClient) *EbmRenewApi {
	return &EbmRenewApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ebm/renew",
		},
	}
}

func (this *EbmRenewApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmRenewRequest) (*EbmRenewResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EbmRenewRealRequest{
		RegionID:        req.RegionID,
		AzName:          req.AzName,
		InstanceUUID:    req.InstanceUUID,
		PayVoucherPrice: req.PayVoucherPrice,
		CycleType:       req.CycleType,
		CycleCount:      req.CycleCount,
		ClientToken:     req.ClientToken,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmRenewRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EbmRenewResponse{
		RegionID:      realResponse.RegionID,
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
	}, nil
}

type EbmRenewRealRequest struct {
	RegionID        string  `json:"regionID,omitempty"`
	AzName          string  `json:"azName,omitempty"`
	InstanceUUID    string  `json:"instanceUUID,omitempty"`
	PayVoucherPrice float64 `json:"payVoucherPrice,omitempty"`
	CycleType       string  `json:"cycleType,omitempty"`
	CycleCount      *int    `json:"cycleCount,omitempty"`
	ClientToken     string  `json:"clientToken,omitempty"`
}

type EbmRenewRequest struct {
	RegionID        string
	AzName          string
	InstanceUUID    string
	PayVoucherPrice float64
	CycleType       string
	CycleCount      *int
	ClientToken     string
}

type EbmRenewRealResponse struct {
	RegionID      string `json:"regionID,omitempty"`
	MasterOrderID string `json:"masterOrderID,omitempty"`
	MasterOrderNO string `json:"masterOrderNO,omitempty"`
}

type EbmRenewResponse struct {
	RegionID      string
	MasterOrderID string
	MasterOrderNO string
}
