package main

import (
	"context"
	"encoding/json"
	"fmt"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctebm"
)

func ebmListInstanceExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	pageNo := 1
	pageSize := 1
	asc := false
	response, err := apis.EbmListInstanceApi.Do(context.Background(), credential, &ctebm.EbmListInstanceRequest{
		RegionID: "41f64827f25f468595ffa3a5deb5d15d",
		AzName:   "cn-anshun-t1a",

		ResourceID:       "38b6da76824a4106b9c5c990c090477f",
		Ip:               "192.168.0.149",
		InstanceName:     "openapi-regactor1",
		VpcID:            "cd20b32e-d35b-48b4-ab51-eb0ad9593562",
		SubnetID:         "1e4c6c1d-d658-4cbd-a526-c531257b4bb0",
		DeviceType:       "physical.t1.large.1",
		DeviceUUIDList:   "d-wrkawlwo9u0mlmtakyrt8indygww",
		QueryContent:     "192.168.0.210",
		InstanceUUIDList: "ss-udogivw2k85vmmtsoiydvewfvzox",
		InstanceUUID:     "ss-wcwkjt374iq5ce8orekenziqdfs5",
		Status:           "RUNNING",
		Sort:             "expire_time",
		Asc:              &asc,
		VipID:            "havip-b3lbmyhj27",
		VolumeUUID:       "458a9dc3-682f-477e-be4f-0f4d62ca8f8f",
		PageNo:           &pageNo,
		PageSize:         &pageSize,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ebmUpdateExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmUpdateApi.Do(context.Background(), credential, &ctebm.EbmUpdateRequest{
		RegionID:    "41f64827f25f468595ffa3a5deb5d15d",
		AzName:      "cn-anshun-t1a",
		DisplayName: "test-name1",
		//Description:  "test-desc",
		InstanceUUID: "ss-x8mdld9wlwyw8sjdtpe7aypssye2",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmPowerOnExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmPowerOnApi.Do(context.Background(), credential, &ctebm.EbmPowerOnRequest{
		RegionID:     "41f64827f25f468595ffa3a5deb5d15d",
		AzName:       "cn-anshun-t1a",
		InstanceUUID: "ss-x8mdld9wlwyw8sjdtpe7aypssye2",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmPowerOffExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmPowerOffApi.Do(context.Background(), credential, &ctebm.EbmPowerOffRequest{
		RegionID:     "41f64827f25f468595ffa3a5deb5d15d",
		AzName:       "cn-anshun-t1a",
		InstanceUUID: "ss-8jdjvpzfcd5hc9sqbi5breolb9po",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmRebootExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmRebootApi.Do(context.Background(), credential, &ctebm.EbmRebootRequest{
		RegionID:     "41f64827f25f468595ffa3a5deb5d15d",
		AzName:       "cn-anshun-t1a",
		InstanceUUID: "ss-x8mdld9wlwyw8sjdtpe7aypssye2",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmRebuildExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	RedoRaid := false
	response, err := apis.EbmRebuildApi.Do(context.Background(), credential, &ctebm.EbmRebuildRequest{
		RegionID:             "41f64827f25f468595ffa3a5deb5d15d",
		AzName:               "cn-anshun-t1a",
		InstanceUUID:         "ss-8jdjvpzfcd5hc9sqbi5breolb9po",
		Hostname:             "host-pm-3301",
		Password:             "CHNchn@0716+++",
		ImageUUID:            "im-as6g7uju3cesx8n7qru8vqn2iqkf",
		SystemVolumeRaidUUID: "r-wtzluqacgzzxgunnabdkpnpjew3d",
		RedoRaid:             &RedoRaid,
		//UserData:             "ZWNobyBoZWxsbyBnb3N0YWNrIQ==",
		KeyName: "KeyPair-openapi-sdk",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmChangePasswordExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmChangePasswordApi.Do(context.Background(), credential, &ctebm.EbmChangePasswordRequest{
		RegionID:     "41f64827f25f468595ffa3a5deb5d15d",
		AzName:       "cn-anshun-t1a",
		InstanceUUID: "ss-x8mdld9wlwyw8sjdtpe7aypssye2",
		NewPassword:  "CHNchn@0716++",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmRenewExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	cycleCount := 1
	response, err := apis.EbmRenewApi.Do(context.Background(), credential, &ctebm.EbmRenewRequest{
		RegionID:        "41f64827f25f468595ffa3a5deb5d15d",
		AzName:          "cn-anshun-t1a",
		InstanceUUID:    "ss-x8mdld9wlwyw8sjdtpe7aypssye2",
		PayVoucherPrice: 0,
		CycleType:       "MONTH",
		CycleCount:      &cycleCount,
		ClientToken:     "uGTWoCHLiO1223",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmDeleteExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmDeleteApi.Do(context.Background(), credential, &ctebm.EbmDeleteRequest{
		RegionID:     "41f64827f25f468595ffa3a5deb5d15d",
		AzName:       "cn-anshun-t1a",
		InstanceUUID: "ss-8jdjvpzfcd5hc9sqbi5breolb9po",
		ClientToken:  "asfawfawawd213",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmDeviceStockListExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	count := 4

	response, err := apis.EbmDeviceStockListApi.Do(context.Background(), credential, &ctebm.EbmDeviceStockListRequest{
		RegionID:   "41f64827f25f468595ffa3a5deb5d15d",
		AzName:     "cn-anshun-t1a",
		DeviceType: "physical.t2.large.1",
		Count:      &count,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmRaidTypeListExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmRaidTypeListApi.Do(context.Background(), credential, &ctebm.EbmRaidTypeListRequest{
		RegionID:   "41f64827f25f468595ffa3a5deb5d15d",
		AzName:     "cn-anshun-t1a",
		DeviceType: "physical.t2.large.1",
		VolumeType: "data",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmImageListExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)

	pageNo := 2
	pageSize := 5
	response, err := apis.EbmImageListApi.Do(context.Background(), credential, &ctebm.EbmImageListRequest{
		RegionID:   "41f64827f25f468595ffa3a5deb5d15d",
		AzName:     "cn-anshun-t1a",
		DeviceType: "physical.t2.large.1",
		ImageType:  "private",
		ImageUUID:  "im-1beuuaavadgpstt5gdh3roz5n1sy",
		OsName:     "CentOS",
		OsVersion:  "7.6",
		OsType:     "linux",
		PageNo:     &pageNo,
		PageSize:   &pageSize,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmDeviceTypeListExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmDeviceTypeListApi.Do(context.Background(), credential, &ctebm.EbmDeviceTypeListRequest{
		RegionID:   "41f64827f25f468595ffa3a5deb5d15d",
		AzName:     "cn-anshun-t1a",
		DeviceType: "physical.t29.large",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmCreateInstanceExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)
	apis := ctebm.NewApis(client)
	master := true
	system_disk_size := 20
	data_disk_size := 20
	bandwidth := 100
	autoRenewStatus := 1
	payVoucherPrice := 20.0
	cycleCount := 1
	orderCount := 1

	networkCardList := make([]ctebm.EbmCreateInstanceNetworkCardListRequest, 0)
	network_card := ctebm.EbmCreateInstanceNetworkCardListRequest{
		Title:    "ens1",
		Master:   &master,
		SubnetID: "f1c272f8-96c8-46c9-b38e-75efa96829b0",
		//FixedIP:  "192.168.0.111",
	}
	networkCardList = append(networkCardList, network_card)
	diskList := make([]ctebm.EbmCreateInstanceDiskListRequest, 0)
	system_disk := ctebm.EbmCreateInstanceDiskListRequest{
		Title:    "system-disk-test",
		DiskType: "system",
		Type:     "SSD",
		Size:     &system_disk_size,
	}
	diskList = append(diskList, system_disk)
	data_disk := ctebm.EbmCreateInstanceDiskListRequest{
		Title:    "data-disk-test",
		DiskType: "data",
		Type:     "SSD",
		Size:     &data_disk_size,
	}
	diskList = append(diskList, data_disk)

	response, err := apis.EbmCreateInstanceApi.Do(context.Background(), credential, &ctebm.EbmCreateInstanceRequest{
		RegionID:     "41f64827f25f468595ffa3a5deb5d15d",
		AzName:       "cn-anshun-t1a",
		DeviceType:   "physical.t2.large.1",
		InstanceName: "中文%",
		Hostname:     "host-openapi-sdk-test",
		ImageUUID:    "im-xevpi6apqilz1bixmogofyref9qm",
		Password:     "CHNchn@0715++",
		VpcID:        "0a58b57c-e4c8-47fa-ad42-8d1dcef1ddeb",
		ExtIP:        "0",
		ProjectID:    "0",
		IpType:       "ipv4",
		BandWidth:    &bandwidth,
		//PublicIP:     "29dc38cb-b687-4006-b47b-680c63d35db0",
		//SecurityGroupID: "sg-zkmzl3s8zv,sg-sb6uhk7iju",
		//DiskList:             diskList,
		NetworkCardList:      networkCardList,
		SystemVolumeRaidUUID: "r-wtzluqacgzzxgunnabdkpnpjew3d",
		//DataVolumeRaidUUID:   "",
		//UserData:           "",
		//KeyName:            "qyo84!*ymd",
		PayVoucherPrice:    &payVoucherPrice,
		AutoRenewStatus:    &autoRenewStatus,
		InstanceChargeType: "ORDER_ON_CYCLE",
		CycleCount:         &cycleCount,
		CycleType:          "MONTH",
		OrderCount:         &orderCount,
		ClientToken:        "li-test-0717123zfsad1123f",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmCreateInstanceT3Exec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)
	apis := ctebm.NewApis(client)
	master := false
	system_disk_size := 100
	data_disk_size := 20
	bandwidth := 100
	payVoucherPrice := 20.0
	autoRenewStatus := 1
	cycleCount := 2
	orderCount := 1

	networkCardList := make([]ctebm.EbmCreateInstanceNetworkCardListRequest, 0)
	network_card := ctebm.EbmCreateInstanceNetworkCardListRequest{
		Title:    "ens1",
		Master:   &master,
		SubnetID: "9b32fcdd-59e2-4754-a504-0fc6b26fbf39",
		FixedIP:  "192.168.0.111",
	}
	networkCardList = append(networkCardList, network_card)
	diskList := make([]ctebm.EbmCreateInstanceDiskListRequest, 0)
	system_disk := ctebm.EbmCreateInstanceDiskListRequest{
		Title:    "system-disk-test",
		DiskType: "system",
		Type:     "SSD",
		Size:     &system_disk_size,
	}
	diskList = append(diskList, system_disk)
	data_disk := ctebm.EbmCreateInstanceDiskListRequest{
		Title:    "data-disk-test",
		DiskType: "data",
		Type:     "SSD",
		Size:     &data_disk_size,
	}
	diskList = append(diskList, data_disk)

	response, err := apis.EbmCreateInstanceApi.Do(context.Background(), credential, &ctebm.EbmCreateInstanceRequest{
		RegionID:        "abae16caac3311ed9fbc0242ac110003",
		AzName:          "cn-anshun-t3a",
		DeviceType:      "physical.s5.2xlarge1",
		InstanceName:    "中文%",
		Hostname:        "host-openapi-sdk-test",
		ImageUUID:       "im-xevpi6apqilz1bixmogofyref9qm",
		Password:        "CHNchn@0715++",
		VpcID:           "62b0f9b0-5895-427e-a48c-d485a720dac8",
		ExtIP:           "2",
		ProjectID:       "0",
		IpType:          "ipv4",
		BandWidth:       &bandwidth,
		PublicIP:        "8843e0c0-fca0-49e9-a1e9-970779a03916",
		SecurityGroupID: "5c12806f-3c31-4974-857c-29b0a7ba473a",
		DiskList:        diskList,
		NetworkCardList: networkCardList,
		//SystemVolumeRaidUUID: "r-wtzluqacgzzxgunnabdkpnpjew3d",
		//DataVolumeRaidUUID:   "",
		//UserData:           "",
		//KeyName:            "qyo84!*ymd",
		PayVoucherPrice: &payVoucherPrice,
		AutoRenewStatus: &autoRenewStatus,
		//InstanceChargeType: "ORDER_ON_CYCLE",
		CycleCount:  &cycleCount,
		CycleType:   "MONTH",
		OrderCount:  &orderCount,
		ClientToken: "li-test-07175123521352131",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmDescribeInstanceExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmDescribeInstanceApi.Do(context.Background(), credential, &ctebm.EbmDescribeInstanceRequest{
		RegionID:     "41f64827f25f468595ffa3a5deb5d15d",
		AzName:       "default",
		InstanceUUID: "ss-x8mdld9wlwyw8sjdtpe7aypssye2",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmDescribeInstanceT3Exec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmDescribeInstanceApi.Do(context.Background(), credential, &ctebm.EbmDescribeInstanceRequest{
		RegionID:     "abae16caac3311ed9fbc0242ac110003",
		AzName:       "cn-anshun-t3a",
		InstanceUUID: "ss-9qxckm0qncd31fvxnfxeiapt9jkm",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ebmDestroyInstanceExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentTest)

	apis := ctebm.NewApis(client)
	response, err := apis.EbmDestroyInstanceApi.Do(context.Background(), credential, &ctebm.EbmDestroyInstanceRequest{
		RegionID:     "41f64827f25f468595ffa3a5deb5d15d",
		AzName:       "cn-anshun-t1a",
		InstanceUUID: "ss-8jdjvpzfcd5hc9sqbi5breolb9po",
		ClientToken:  "li-test-0717123214312",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func main() {
	credential, _ := ctyunsdk.NewCredential("ak", "sk")
	ebmDeviceTypeListExec(*credential)
	ebmListInstanceExec(*credential)
	ebmDescribeInstanceExec(*credential)
	ebmDeviceStockListExec(*credential)
	ebmRaidTypeListExec(*credential)
	ebmImageListExec(*credential)
	ebmRebootExec(*credential)
	ebmPowerOffExec(*credential)
	ebmPowerOnExec(*credential)
	ebmChangePasswordExec(*credential)
	ebmUpdateExec(*credential)
	ebmRenewExec(*credential)
	ebmDeleteExec(*credential)
	ebmRebuildExec(*credential)
	ebmDestroyInstanceExec(*credential)

	ebmCreateInstanceExec(*credential)
	ebmCreateInstanceT3Exec(*credential)
	ebmDescribeInstanceT3Exec(*credential)
}
