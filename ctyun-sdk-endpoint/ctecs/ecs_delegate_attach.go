package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsDelegateAttachApi
type EcsDelegateAttachApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsDelegateAttachApi(client *ctyunsdk.CtyunClient) *EcsDelegateAttachApi {
	return &EcsDelegateAttachApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/delegate/attach",
		},
	}
}

func (this *EcsDelegateAttachApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsDelegateAttachRequest) (*EcsDelegateAttachResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsDelegateAttachRealRequest{
		RegionID:     req.RegionID,
		InstanceID:   req.InstanceID,
		DelegateName: req.DelegateName,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsDelegateAttachRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsDelegateAttachResponse{
		InstanceID:   realResponse.InstanceID,
		DelegateName: realResponse.DelegateName,
	}, nil
}

type EcsDelegateAttachRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	InstanceID   string `json:"instanceID,omitempty"`
	DelegateName string `json:"delegateName,omitempty"`
}

type EcsDelegateAttachRequest struct {
	RegionID     string
	InstanceID   string
	DelegateName string
}

type EcsDelegateAttachRealResponse struct {
	InstanceID   string `json:"instanceID,omitempty"`
	DelegateName string `json:"delegateName,omitempty"`
}

type EcsDelegateAttachResponse struct {
	InstanceID   string
	DelegateName string
}
