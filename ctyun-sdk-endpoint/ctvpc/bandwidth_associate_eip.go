package ctvpc

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// BandwidthAssociateEipApi 添加EIP至共享带宽
// https://www.ctyun.cn/document/10026761/10040809
type BandwidthAssociateEipApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewBandwidthAssociateEipApi(client *ctyunsdk.CtyunClient) *BandwidthAssociateEipApi {
	return &BandwidthAssociateEipApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/bandwidth/associate-eip",
		},
	}
}

func (this *BandwidthAssociateEipApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *BandwidthAssociateEipRequest) (*BandwidthAssociateEipResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	requestContent := bandwidthAssociateEipRealRequest{
		RegionID:    req.RegionID,
		ClientToken: req.ClientToken,
		BandwidthID: req.BandwidthID,
		EipIDs:      req.EipIDs,
	}
	_, err := builder.WriteJson(requestContent)
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtvpc, builder)
	if err != nil {
		return nil, err
	}

	err = response.ParseByStandardModelWithCheck(nil)
	if err != nil {
		return nil, err
	}

	return &BandwidthAssociateEipResponse{}, nil
}

type bandwidthAssociateEipRealRequest struct {
	RegionID    string   `json:"regionID"`
	ClientToken string   `json:"clientToken"`
	BandwidthID string   `json:"bandwidthID"`
	EipIDs      []string `json:"eipIDs"`
}

type BandwidthAssociateEipRequest struct {
	RegionID    string
	ClientToken string
	BandwidthID string
	EipIDs      []string
}

type BandwidthAssociateEipResponse struct {
}
