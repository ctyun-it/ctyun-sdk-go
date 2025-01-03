package main

import (
	"context"
	"encoding/json"
	"fmt"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctecs/common"
)

func listRegions(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := common.NewApis(client)
	response, err := apis.RegionListRegionsApi.Do(context.Background(), credential, &common.RegionListRequest{
		RegionName: "**资源池", // 资源池名称
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func listZones(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := common.NewApis(client)
	response, err := apis.RegionGetZonesApi.Do(context.Background(), credential, &common.RegionGetZonesRequest{
		RegionID: "********************************", // 资源池ID
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func main() {
	credential, _ := ctyunsdk.NewCredential("ak****", "sk****")
	listRegions(*credential)
	listZones(*credential)
}
