package ctebm

import (
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
)

// Apis api的接口
type Apis struct {
	EbmPowerOnApi          *EbmPowerOnApi
	EbmPowerOffApi         *EbmPowerOffApi
	EbmListInstanceApi     *EbmListInstanceApi
	EbmUpdateApi           *EbmUpdateApi
	EbmRebuildApi          *EbmRebuildApi
	EbmRebootApi           *EbmRebootApi
	EbmChangePasswordApi   *EbmChangePasswordApi
	EbmRenewApi            *EbmRenewApi
	EbmDeleteApi           *EbmDeleteApi
	EbmDeviceStockListApi  *EbmDeviceStockListApi
	EbmRaidTypeListApi     *EbmRaidTypeListApi
	EbmImageListApi        *EbmImageListApi
	EbmDeviceTypeListApi   *EbmDeviceTypeListApi
	EbmCreateInstanceApi   *EbmCreateInstanceApi
	EbmDescribeInstanceApi *EbmDescribeInstanceApi
	EbmDestroyInstanceApi  *EbmDestroyInstanceApi
}

// NewApis 构建
func NewApis(client *ctyunsdk.CtyunClient) *Apis {
	builder := ctyunsdk.NewApiHookBuilder()
	for _, hook := range client.Config.ApiHooks {
		builder.AddHooks(hook)
	}
	client.RegisterEndpoint(ctyunsdk.EnvironmentDev, EndpointEbmTest)
	client.RegisterEndpoint(ctyunsdk.EnvironmentTest, EndpointEbmTest)
	client.RegisterEndpoint(ctyunsdk.EnvironmentProd, EndpointEbmProd)
	return &Apis{
		EbmPowerOnApi:          NewEbmPowerOnApi(client),
		EbmPowerOffApi:         NewEbmPowerOffApi(client),
		EbmListInstanceApi:     NewEbmListInstanceApi(client),
		EbmUpdateApi:           NewEbmUpdateApi(client),
		EbmRebuildApi:          NewEbmRebuildApi(client),
		EbmRebootApi:           NewEbmRebootApi(client),
		EbmChangePasswordApi:   NewEbmChangePasswordApi(client),
		EbmRenewApi:            NewEbmRenewApi(client),
		EbmDeleteApi:           NewEbmDeleteApi(client),
		EbmDeviceStockListApi:  NewEbmDeviceStockListApi(client),
		EbmRaidTypeListApi:     NewEbmRaidTypeListApi(client),
		EbmImageListApi:        NewEbmImageListApi(client),
		EbmDeviceTypeListApi:   NewEbmDeviceTypeListApi(client),
		EbmCreateInstanceApi:   NewEbmCreateInstanceApi(client),
		EbmDescribeInstanceApi: NewEbmDescribeInstanceApi(client),
		EbmDestroyInstanceApi:  NewEbmDestroyInstanceApi(client),
	}
}
