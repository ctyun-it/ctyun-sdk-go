package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsFlavorListByFamiliesApi
type EcsFlavorListByFamiliesApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsFlavorListByFamiliesApi(client *ctyunsdk.CtyunClient) *EcsFlavorListByFamiliesApi {
	return &EcsFlavorListByFamiliesApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/flavor/list-by-families",
		},
	}
}

func (this *EcsFlavorListByFamiliesApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsFlavorListByFamiliesRequest) (*EcsFlavorListByFamiliesResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	_, err := builder.WriteJson(&EcsFlavorListByFamiliesRealRequest{
		RegionID:     req.RegionID,
		AzName:       req.AzName,
		FlavorFamily: req.FlavorFamily,
		PageNo:       req.PageNo,
		PageSize:     req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsFlavorListByFamiliesRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EcsFlavorListByFamiliesResultsResponse
	for _, res := range realResponse.Results {
		results = append(results, EcsFlavorListByFamiliesResultsResponse{
			InstanceID:   res.InstanceID,
			InstanceName: res.InstanceName,
			Flavor: EcsFlavorListByFamiliesFlavorResponse{
				FlavorID:     res.Flavor.FlavorID,
				FlavorName:   res.Flavor.FlavorName,
				FlavorCPU:    res.Flavor.FlavorCPU,
				FlavorRAM:    res.Flavor.FlavorRAM,
				GpuType:      res.Flavor.GpuType,
				GpuCount:     res.Flavor.GpuCount,
				GpuVendor:    res.Flavor.GpuVendor,
				VideoMemSize: res.Flavor.VideoMemSize,
			},
		})
	}

	return &EcsFlavorListByFamiliesResponse{
		CurrentCount: realResponse.CurrentCount,
		TotalCount:   realResponse.TotalCount,
		TotalPage:    realResponse.TotalPage,
		Results:      results,
	}, nil
}

type EcsFlavorListByFamiliesRealRequest struct {
	RegionID     string `json:"regionID,omitempty"`
	AzName       string `json:"azName,omitempty"`
	FlavorFamily string `json:"flavorFamily,omitempty"`
	PageNo       int    `json:"pageNo,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

type EcsFlavorListByFamiliesRequest struct {
	RegionID     string
	AzName       string
	FlavorFamily string
	PageNo       int
	PageSize     int
}

type EcsFlavorListByFamiliesFlavorRealResponse struct {
	FlavorID     string `json:"flavorID,omitempty"`
	FlavorName   string `json:"flavorName,omitempty"`
	FlavorCPU    int    `json:"flavorCPU,omitempty"`
	FlavorRAM    int    `json:"flavorRAM,omitempty"`
	GpuType      string `json:"gpuType,omitempty"`
	GpuCount     int    `json:"gpuCount,omitempty"`
	GpuVendor    string `json:"gpuVendor,omitempty"`
	VideoMemSize int    `json:"videoMemSize,omitempty"`
}

type EcsFlavorListByFamiliesResultsRealResponse struct {
	InstanceID   string                                    `json:"instanceID,omitempty"`
	InstanceName string                                    `json:"instanceName,omitempty"`
	Flavor       EcsFlavorListByFamiliesFlavorRealResponse `json:"flavor,omitempty"`
}

type EcsFlavorListByFamiliesRealResponse struct {
	CurrentCount int                                          `json:"currentCount,omitempty"`
	TotalCount   int                                          `json:"totalCount,omitempty"`
	TotalPage    int                                          `json:"totalPage,omitempty"`
	Results      []EcsFlavorListByFamiliesResultsRealResponse `json:"results,omitempty"`
}

type EcsFlavorListByFamiliesFlavorResponse struct {
	FlavorID     string
	FlavorName   string
	FlavorCPU    int
	FlavorRAM    int
	GpuType      string
	GpuCount     int
	GpuVendor    string
	VideoMemSize int
}

type EcsFlavorListByFamiliesResultsResponse struct {
	InstanceID   string
	InstanceName string
	Flavor       EcsFlavorListByFamiliesFlavorResponse
}

type EcsFlavorListByFamiliesResponse struct {
	CurrentCount int
	TotalCount   int
	TotalPage    int
	Results      []EcsFlavorListByFamiliesResultsResponse
}
