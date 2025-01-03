package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsShareInterfaceAttachApi
type EcsShareInterfaceAttachApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsShareInterfaceAttachApi(client *ctyunsdk.CtyunClient) *EcsShareInterfaceAttachApi {
	return &EcsShareInterfaceAttachApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/share-interface/attach",
		},
	}
}

func (this *EcsShareInterfaceAttachApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsShareInterfaceAttachRequest) (*EcsShareInterfaceAttachResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsShareInterfaceAttachRealRequest{
		RegionID:   req.RegionID,
		InstanceID: req.InstanceID,
		SubnetID:   req.SubnetID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsShareInterfaceAttachRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsShareInterfaceAttachResponse{
		NicID: realResponse.NicID,
	}, nil
}

type EcsShareInterfaceAttachRealRequest struct {
	RegionID   *string `json:"regionID,omitempty"`
	InstanceID *string `json:"instanceID,omitempty"`
	SubnetID   *string `json:"subnetID,omitempty"`
}

type EcsShareInterfaceAttachRequest struct {
	RegionID   *string
	InstanceID *string
	SubnetID   *string
}

type EcsShareInterfaceAttachRealResponse struct {
	NicID string `json:"nicID,omitempty"`
}

type EcsShareInterfaceAttachResponse struct {
	NicID string
}
