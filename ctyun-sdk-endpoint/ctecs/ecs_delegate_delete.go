package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsDelegateDeleteApi
type EcsDelegateDeleteApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsDelegateDeleteApi(client *ctyunsdk.CtyunClient) *EcsDelegateDeleteApi {
	return &EcsDelegateDeleteApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/delegate/delete",
		},
	}
}

func (this *EcsDelegateDeleteApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsDelegateDeleteRequest) (*EcsDelegateDeleteResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsDelegateDeleteRealRequest{
		RegionID:   req.RegionID,
		InstanceID: req.InstanceID,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsDelegateDeleteRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsDelegateDeleteResponse{
		InstanceID: realResponse.InstanceID,
	}, nil
}

type EcsDelegateDeleteRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	InstanceID string `json:"instanceID,omitempty"`
}

type EcsDelegateDeleteRequest struct {
	RegionID   string
	InstanceID string
}

type EcsDelegateDeleteRealResponse struct {
	InstanceID string `json:"instanceID,omitempty"`
}

type EcsDelegateDeleteResponse struct {
	InstanceID string
}
