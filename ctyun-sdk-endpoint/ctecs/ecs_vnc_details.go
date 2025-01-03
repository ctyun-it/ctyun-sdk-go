package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsVncDetailsApi
type EcsVncDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsVncDetailsApi(client *ctyunsdk.CtyunClient) *EcsVncDetailsApi {
	return &EcsVncDetailsApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/vnc/details",
		},
	}
}

func (this *EcsVncDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsVncDetailsRequest) (*EcsVncDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	builder.AddParam("regionID", *req.RegionID)
	builder.AddParam("instanceID", *req.InstanceID)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsVncDetailsRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	return &EcsVncDetailsResponse{
		Token: realResponse.Token,
	}, nil
}

type EcsVncDetailsRealRequest struct {
	RegionID   *string `json:"regionID,omitempty"`
	InstanceID *string `json:"instanceID,omitempty"`
}

type EcsVncDetailsRequest struct {
	RegionID   *string
	InstanceID *string
}

type EcsVncDetailsRealResponse struct {
	Token string `json:"token,omitempty"`
}

type EcsVncDetailsResponse struct {
	Token string
}
