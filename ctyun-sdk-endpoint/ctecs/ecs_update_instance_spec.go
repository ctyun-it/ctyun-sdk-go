package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsUpdateInstanceSpecApi 云主机修改带宽或规格
// https://www.ctyun.cn/document/10026730/10106617

type EcsUpdateInstanceSpecApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsUpdateInstanceSpecApi(client *ctyunsdk.CtyunClient) *EcsUpdateInstanceSpecApi {
	return &EcsUpdateInstanceSpecApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/update-instance-spec",
		},
	}
}

func (this *EcsUpdateInstanceSpecApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsUpdateInstanceSpecRequest) (*EcsUpdateInstanceSpecResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsUpdateInstanceSpecRealRequest{
		RegionID:    req.RegionID,
		InstanceID:  req.InstanceID,
		Bandwidth:   req.Bandwidth,
		FlavorID:    req.FlavorID,
		ClientToken: req.ClientToken,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsUpdateInstanceSpecRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsUpdateInstanceSpecResponse{
		MasterOrderID: realResponse.MasterOrderID,
		MasterOrderNO: realResponse.MasterOrderNO,
		RegionID:      realResponse.RegionID,
	}, nil
}

type EcsUpdateInstanceSpecRealRequest struct {
	RegionID    string `json:"regionID,omitempty"`
	InstanceID  string `json:"instanceID,omitempty"`
	Bandwidth   int    `json:"bandwidth,omitempty"`
	FlavorID    string `json:"flavorID,omitempty"`
	ClientToken string `json:"clientToken,omitempty"`
}

type EcsUpdateInstanceSpecRequest struct {
	RegionID    string
	InstanceID  string
	Bandwidth   int
	FlavorID    string
	ClientToken string
}

type EcsUpdateInstanceSpecRealResponse struct {
	MasterOrderID string `json:"masterOrderID,omitempty"`
	MasterOrderNO string `json:"masterOrderNO,omitempty"`
	RegionID      string `json:"regionID,omitempty"`
}

type EcsUpdateInstanceSpecResponse struct {
	MasterOrderID string
	MasterOrderNO string
	RegionID      string
}
