package common

import ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"

// Apis api的接口
type Apis struct {
	RegionListRegionsApi *RegionListRegionsApi
	RegionGetZonesApi    *RegionGetZonesApi
}

// NewApis 构建
func NewApis(client *ctyunsdk.CtyunClient) *Apis {
	client.RegisterEndpoint(ctyunsdk.EnvironmentDev, EndpointCtecsTest)
	client.RegisterEndpoint(ctyunsdk.EnvironmentTest, EndpointCtecsTest)
	client.RegisterEndpoint(ctyunsdk.EnvironmentProd, EndpointCtecsProd)
	return &Apis{
		RegionListRegionsApi: NewRegionListRegionsApi(client),
		RegionGetZonesApi:    NewRegionGetZonesApi(client),
	}
}
