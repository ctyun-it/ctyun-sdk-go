package ctecs

import (
	"context"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsAvailabilityZonesDetailsApi 查询指定资源池的可用区信息
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=25&api=8314&data=87&isNormal=1

type EcsAvailabilityZonesDetailsApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsAvailabilityZonesDetails(client *ctyunsdk.CtyunClient) *EcsAvailabilityZonesDetailsApi {
	return &EcsAvailabilityZonesDetailsApi{
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ecs/availability-zones/details",
		},
		client: client,
	}
}

func (this *EcsAvailabilityZonesDetailsApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsAvailabilityZonesDetailsRequest) (*EcsAvailabilityZonesDetailsResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	builder.AddParam("regionID", *req.RegionID)
	resp, requestError := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if requestError != nil {
		return nil, requestError
	}

	var realResponse EcsAvailabilityZonesDetailsRealResponse
	err := resp.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var azList []EcsAvailabilityZonesDetailsAzListResponse
	for _, az_info := range realResponse.AzList {
		azList = append(azList, EcsAvailabilityZonesDetailsAzListResponse{
			AzID:   az_info.AzID,
			AzName: az_info.AzName,
		})
	}

	return &EcsAvailabilityZonesDetailsResponse{
		AzList: azList,
	}, nil
}

type EcsAvailabilityZonesDetailsRequest struct {
	RegionID *string
}

type EcsAvailabilityZonesDetailsAzListRealResponse struct {
	AzID   string `json:"azID,omitempty"`
	AzName string `json:"azName,omitempty"`
}

type EcsAvailabilityZonesDetailsRealResponse struct {
	AzList []EcsAvailabilityZonesDetailsAzListRealResponse `json:"azList,omitempty"`
}

type EcsAvailabilityZonesDetailsAzListResponse struct {
	AzID   string
	AzName string
}

type EcsAvailabilityZonesDetailsResponse struct {
	AzList []EcsAvailabilityZonesDetailsAzListResponse
}
