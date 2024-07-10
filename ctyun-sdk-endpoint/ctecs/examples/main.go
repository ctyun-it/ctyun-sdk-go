package main

import (
	"context"
	"encoding/json"
	"fmt"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctecs"
)

func createInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	networkCardList := make([]ctecs.EcsCreateInstanceNetworkCardListRequest, 0)
	network_card := ctecs.EcsCreateInstanceNetworkCardListRequest{
		NicName:  "nic-test",
		IsMaster: true,
		SubnetId: "subnet-4c4333pc67",
	}
	networkCardList = append(networkCardList, network_card)
	dataDiskList := make([]ctecs.EcsCreateInstanceDataDiskListRequest, 0)
	data_disk := ctecs.EcsCreateInstanceDataDiskListRequest{
		DiskMode: "VBD",
		DiskName: "data-disk-test",
		DiskType: "SATA",
		DiskSize: 20,
	}
	dataDiskList = append(dataDiskList, data_disk)
	labelList := make([]ctecs.EcsCreateInstanceLabelListRequest, 0)
	label := ctecs.EcsCreateInstanceLabelListRequest{
		LabelKey:   "label-key-test",
		LabelValue: "label-value-test",
	}
	labelList = append(labelList, label)
	response, err := apis.EcsCreateInstanceApi.Do(context.Background(), credential, &ctecs.EcsCreateInstanceRequest{
		ClientToken:     "ecs-create-instance-test-02",
		RegionId:        "bb9fdb42056f11eda1610242ac110002",
		AzName:          "cn-huadong1-jsnj1A-public-ctcloud",
		InstanceName:    "ecm-go-test",
		DisplayName:     "ecm-go-test",
		FlavorId:        "b6779240-5649-803b-4a4c-8fc59d310ecf",
		ImageType:       1,
		ImageId:         "939c131f-a986-420f-a3b2-57feb9995e47",
		BootDiskType:    "SATA",
		BootDiskSize:    40,
		VpcId:           "vpc-chz0ilszsp",
		OnDemand:        false,
		NetworkCardList: networkCardList,
		ExtIp:           "1",
		ProjectID:       "0",
		SecGroupList:    []string{"sg-bqv0t629h6", "sg-bqv0t629h6"},
		DataDiskList:    dataDiskList,
		IpVersion:       "ipv4",
		Bandwidth:       50,
		UserPassword:    "qyo84!*ymd",
		CycleCount:      1,
		CycleType:       "MONTH",
		AutoRenewStatus: 0,
		UserData:        "YmF0Y2hDcmVhdGVUZXN0MDgwMw==",
		PayVoucherPrice: 114.00,
		LabelList:       labelList,
		MonitorService:  true,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))

}

func ecsBatchCreateInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	networkCardList := make([]ctecs.EcsBatchCreateInstancesNetworkCardListRequest, 0)
	network_card := ctecs.EcsBatchCreateInstancesNetworkCardListRequest{
		NicName:  "nic-test",
		IsMaster: true,
		SubnetID: "subnet-4c4333pc67",
		FixedIP:  "192.168.0.26",
	}
	networkCardList = append(networkCardList, network_card)
	dataDiskList := make([]ctecs.EcsBatchCreateInstancesDataDiskListRequest, 0)
	data_disk := ctecs.EcsBatchCreateInstancesDataDiskListRequest{
		DiskMode: "VBD",
		DiskName: "data-disk-test",
		DiskType: "SATA",
		DiskSize: 20,
	}
	dataDiskList = append(dataDiskList, data_disk)
	labelList := make([]ctecs.EcsBatchCreateInstancesLabelListRequest, 0)
	label := ctecs.EcsBatchCreateInstancesLabelListRequest{
		LabelKey:   "label-key-test",
		LabelValue: "label-value-test",
	}
	labelList = append(labelList, label)
	response, err := apis.EcsBatchCreateInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchCreateInstancesRequest{
		ClientToken:     "ecs-batch-create-instances-test-02",
		RegionID:        "bb9fdb42056f11eda1610242ac110002",
		AzName:          "cn-huadong1-jsnj1A-public-ctcloud",
		InstanceName:    "go-sdk-test",
		DisplayName:     "go-sdk-test",
		FlavorID:        "b6779240-5649-803b-4a4c-8fc59d310ecf",
		ImageType:       1,
		ImageID:         "939c131f-a986-420f-a3b2-57feb9995e47",
		BootDiskType:    "SATA",
		BootDiskSize:    40,
		VpcID:           "vpc-chz0ilszsp",
		OnDemand:        false,
		NetworkCardList: networkCardList,
		ExtIP:           "1",
		ProjectID:       "0",
		SecGroupList:    []string{"sg-bqv0t629h6", "sg-bqv0t629h6"},
		DataDiskList:    dataDiskList,
		IpVersion:       "ipv4",
		Bandwidth:       1,
		UserPassword:    "qyo84!*ymd",
		CycleCount:      1,
		CycleType:       "MONTH",
		AutoRenewStatus: 0,
		UserData:        "YmF0Y2hDcmVhdGVUZXN0MDgwMw==",
		PayVoucherPrice: 114.00,
		LabelList:       labelList,
		MonitorService:  true,
		OrderCount:      1,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))

}

func describeInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsDescribeInstancesApi.Do(context.Background(), credential, &ctecs.EcsDescribeInstancesRequest{
		RegionId:        "bb9fdb42056f11eda1610242ac110002",
		AzName:          "cn-huadong1-jsnj1A-public-ctcloud",
		ProjectId:       "0",
		PageNo:          1,
		PageSize:        10,
		State:           "active",
		Keyword:         "ecm-57fd",
		InstanceName:    "ecm-57fd",
		InstanceIdList:  "0fec78e4-1889-803f-b2a7-515c1c40b788",
		SecurityGroupId: "sg-tdzefke02r",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsListInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	labelList := make([]ctecs.EcsListInstancesLabelListRequest, 0)
	label_list := ctecs.EcsListInstancesLabelListRequest{
		LabelKey:   "label-key-test",
		LabelValue: "label-value-test",
	}
	labelList = append(labelList, label_list)

	response, err := apis.EcsListInstancesApi.Do(context.Background(), credential, &ctecs.EcsListInstancesRequest{
		RegionID:        "bb9fdb42056f11eda1610242ac110002",
		AzName:          "cn-huadong1-jsnj1A-public-ctcloud",
		ProjectID:       "0",
		PageNo:          1,
		PageSize:        10,
		State:           "active",
		Keyword:         "ecm-57fd",
		InstanceName:    "ecm-57fd",
		InstanceIDList:  "0fec78e4-1889-803f-b2a7-515c1c40b788",
		SecurityGroupID: "sg-tdzefke02r",
		VpcID:           "vpc-chz0ilszsp",
		//ResourceID:      "d57ec586da6c497ea1e1b04e08ad9a8b",
		//LabelList:       labelList,
		Sort: "expiredTime",
		Asc:  true,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func listFlavors(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsFlavorListApi.Do(context.Background(), credential, &ctecs.EcsFlavorListRequest{
		RegionId:     "bb9fdb42056f11eda1610242ac110002",
		AzName:       "cn-huadong1-jsnj1A-public-ctcloud",
		FlavorType:   "CPU_KS1",
		FlavorName:   "ks1.medium.2",
		FlavorCpu:    1,
		FlavorArch:   "arm",
		FlavorSeries: "ks",
		FlavorId:     "b6779240-5649-803b-4a4c-8fc59d310ecf",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func instanceDetail(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	response, err := apis.EcsInstanceDetailsApi.Do(context.Background(), credential, &ctecs.EcsInstanceDetailsRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func instanceJoinSecurityGroup(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	response, err := apis.EcsJoinSecurityGroupApi.Do(context.Background(), credential, &ctecs.EcsJoinSecurityGroupRequest{
		RegionId:           "bb9fdb42056f11eda1610242ac110002",
		SecurityGroupId:    "sg-tdzefke02r",
		InstanceId:         "77493826-d038-2a9c-f684-e2f6adabeba3",
		NetworkInterfaceId: "port-pja7l0zfvk",
		Action:             "joinSecurityGroup",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func instanceLeaveSecurityGroup(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	response, err := apis.EcsLeaveSecurityGroupApi.Do(context.Background(), credential, &ctecs.EcsLeaveSecurityGroupRequest{
		RegionId:        "bb9fdb42056f11eda1610242ac110002",
		SecurityGroupId: "sg-tdzefke02r",
		InstanceId:      "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func instanceQueryAsyncResult(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsQueryAsyncResultApi.Do(context.Background(), credential, &ctecs.EcsQueryAsyncResultRequest{
		RegionId: "bb9fdb42056f11eda1610242ac110002",
		JobId:    "",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func resetInstancePassword(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsResetPasswordApi.Do(context.Background(), credential, &ctecs.EcsResetPasswordRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
		NewPassword: "test-test-test-960",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func startInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsStartInstanceApi.Do(context.Background(), credential, &ctecs.EcsStartInstanceRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func stopInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsStopInstanceApi.Do(context.Background(), credential, &ctecs.EcsStopInstanceRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		InstanceId: "63afb617-b8f5-d482-9ecd-6d8bb9124d4e",
		Force:      false,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func unsubscribeInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsUnsubscribeInstanceApi.Do(context.Background(), credential, &ctecs.EcsUnsubscribeInstanceRequest{
		ClientToken: "unsubscribe-instance-test",
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func updateFlavorSpec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsUpdateFlavorSpecApi.Do(context.Background(), credential, &ctecs.EcsUpdateFlavorSpecRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
		FlavorId:    "b6779240-5649-803b-4a4c-8fc59d310ecf",
		ClientToken: "update-flavor-spec-test",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func listInstanceVolumes(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsVolumeListApi.Do(context.Background(), credential, &ctecs.EcsVolumeListRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
		PageNo:     1,
		PageSize:   10,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func keypairCreate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	res, err := apis.KeypairCreateApi.Do(context.Background(), credential, &ctecs.KeypairCreateRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		KeyPairName: "keypair-test",
		ProjectId:   "0",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func keypairDelete(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	res, err := apis.KeypairDeleteApi.Do(context.Background(), credential, &ctecs.KeypairDeleteRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		KeyPairName: "keypair-test",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func keypairDetail(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	res, err := apis.KeypairDetailApi.Do(context.Background(), credential, &ctecs.KeypairDetailRequest{
		RegionId:     "bb9fdb42056f11eda1610242ac110002",
		KeyPairName:  "keypair-test",
		ProjectId:    "0",
		QueryContent: "keypair-test",
		PageNo:       1,
		PageSize:     10,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func keypairImport(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	res, err := apis.KeypairImportApi.Do(context.Background(), credential, &ctecs.KeypairImportRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		KeyPairName: "keypair-test",
		PublicKey:   "",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func keypairAttach(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	res, err := apis.KeypairAttachApi.Do(context.Background(), credential, &ctecs.KeypairAttachRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		KeyPairName: "keypair-test",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))

}

func keypairDetach(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	res, err := apis.KeypairDetachApi.Do(context.Background(), credential, &ctecs.KeypairDetachRequest{
		RegionId:    "bb9fdb42056f11eda1610242ac110002",
		KeyPairName: "keypair-test",
		InstanceId:  "77493826-d038-2a9c-f684-e2f6adabeba3",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func rebootInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsRebootInstanceApi.Do(context.Background(), credential, &ctecs.EcsRebootInstanceRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		InstanceID: "0fec78e4-1889-803f-b2a7-515c1c40b788",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func batchRebootInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsBatchRebootInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchRebootInstanceRequest{
		RegionID:       "bb9fdb42056f11eda1610242ac110002",
		InstanceIDList: "0fec78e4-1889-803f-b2a7-515c1c40b788",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func rebuildInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsRebuildInstanceApi.Do(context.Background(), credential, &ctecs.EcsRebuildInstanceRequest{
		RegionID:       "bb9fdb42056f11eda1610242ac110002",
		InstanceID:     "0fec78e4-1889-803f-b2a7-515c1c40b788",
		Password:       "rebuildTest195%",
		ImageID:        "b1d896e1-c977-4fd4-b6c2-5432549977be",
		UserData:       "UmVidWlsZFRlc3QyMDIyMTEyNDEzMTE=",
		InstanceName:   "ecm-3300",
		MonitorService: true,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func batchRebuildInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	rebuildInfoList := make([]ctecs.EcsBatchRebuildInstancesRebuildInfoRequest, 0)
	rebuildInfo := ctecs.EcsBatchRebuildInstancesRebuildInfoRequest{
		InstanceID:     "63afb617-b8f5-d482-9ecd-6d8bb9124d4e",
		Password:       "rebuildTest195%",
		ImageID:        "b1d896e1-c977-4fd4-b6c2-5432549977be",
		UserData:       "UmVidWlsZFRlc3QyMDIyMTEyNDEzMTE=",
		InstanceName:   "ecm-3300",
		MonitorService: true,
	}
	rebuildInfoList = append(rebuildInfoList, rebuildInfo)
	response, err := apis.EcsBatchRebuildInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchRebuildInstancesRequest{
		RegionID:    "bb9fdb42056f11eda1610242ac110002",
		RebuildInfo: rebuildInfoList,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func batchUnsubscribeInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	response, err := apis.EcsBatchUnsubscribeInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchUnsubscribeInstanceRequest{
		ClientToken:    "batch-unsubscribe-instance",
		RegionID:       "bb9fdb42056f11eda1610242ac110002",
		InstanceIDList: "96b254b1-b472-e72d-eb7f-05b61d973ad4",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsBatchStopInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBatchStopInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchStopInstancesRequest{
		RegionID:       "bb9fdb42056f11eda1610242ac110002",
		InstanceIDList: "adc614e0-e838-d73f-0618-a6d51d09070a,5ae83d07-0389-22fd-def8-995090ee3d5a",
		Force:          true,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsUpdateNetworkSpec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsUpdateNetworkSpecApi.Do(context.Background(), credential, &ctecs.EcsUpdateNetworkSpecRequest{
		RegionID:    "bb9fdb42056f11eda1610242ac110002",
		InstanceID:  "93366056-b08f-4b9b-8e47-c50d92f2d4fd",
		Bandwidth:   100,
		ClientToken: "ea1b9004-f450-11ec-8d4f-00155de3fd73",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsUpdateInstanceSpec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsUpdateInstanceSpecApi.Do(context.Background(), credential, &ctecs.EcsUpdateInstanceSpecRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		InstanceID: "93366056-b08f-4b9b-8e47-c50d92f2d4fd",
		Bandwidth:  2,
		//FlavorID:    "00ebe3aa-aac0-1d99-0b9e-4d391c5e06d5",
		ClientToken: "bdfse888-8ed8-88b8-88cb-888f8b8cf8fa",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsGpuDriverList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsGpuDriverListApi.Do(context.Background(), credential, &ctecs.EcsGpuDriverListRequest{
		RegionID: "bb9fdb42056f11eda1610242ac110002",
		FlavorID: "5cf44a7e-e23c-4199-9ebf-226650646e5a",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsFlavorListByFamilies(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsFlavorListByFamiliesApi.Do(context.Background(), credential, &ctecs.EcsFlavorListByFamiliesRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		AzName:       "cn-huadong1-jsnj1A-public-ctcloud",
		FlavorFamily: "s7",
		PageNo:       1,
		PageSize:     10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAvailabilityZonesDetails(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	response, err := apis.EcsAvailabilityZonesDetailsApi.Do(context.Background(), credential, &ctecs.EcsAvailabilityZonesDetailsRequest{
		RegionID: "bb9fdb42056f11eda1610242ac110002",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsFlavorFamiliesList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	response, err := apis.EcsFlavorFamiliesListApi.Do(context.Background(), credential, &ctecs.EcsFlavorFamiliesListRequest{
		RegionID: "bb9fdb42056f11eda1610242ac110002",
		AzName:   "cn-huadong1-jsnj1A-public-ctcloud",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func updateInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.UpdateInstanceApi.Do(context.Background(), credential, &ctecs.UpdateInstanceRequest{
		RegionID:    "bb9fdb42056f11eda1610242ac110002",
		InstanceID:  "cfe4d576-4e2c-efd0-e823-250664d95d8f",
		DisplayName: "java-sdk-ecm-1447",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func queryVncDetails(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.QueryVncDetailsApi.Do(context.Background(), credential, &ctecs.QueryVncDetailsRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		InstanceID: "cfe4d576-4e2c-efd0-e823-250664d95d8f",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func getVolumeStatistics(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.GetVolumeStatisticsApi.Do(context.Background(), credential, &ctecs.GetVolumeStatisticsRequest{
		RegionID:  "bb9fdb42056f11eda1610242ac110002",
		ProjectID: "0",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func batchUpdateInstancesPassword(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	UpdatePwdInfo := make([]ctecs.BatchUpdateInstancesPasswordUpdatePwdInfoRequest, 0)
	updatePwdInfo := ctecs.BatchUpdateInstancesPasswordUpdatePwdInfoRequest{
		InstanceID:  "cfe4d576-4e2c-efd0-e823-250664d95d8f",
		NewPassword: "Ctyun_GoSDK_1",
	}

	UpdatePwdInfo = append(UpdatePwdInfo, updatePwdInfo)

	response, err := apis.BatchUpdateInstancesPasswordApi.Do(context.Background(), credential, &ctecs.BatchUpdateInstancesPasswordRequest{
		RegionID:      "bb9fdb42056f11eda1610242ac110002",
		UpdatePwdInfo: UpdatePwdInfo,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func listInstanceBackupPolicy(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.ListInstanceBackupPolicyApi.Do(context.Background(), credential, &ctecs.ListInstanceBackupPolicyRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		PolicyID:   "3e251bce0d1411efb0a10242ac110002",
		PolicyName: "policy_test_0508",
		ProjectID:  "0",
		PageNo:     1,
		PageSize:   10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func listInstanceBackupPolicyBindInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.ListInstanceBackupPolicyBindInstancesApi.Do(context.Background(), credential, &ctecs.ListInstanceBackupPolicyBindInstancesRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		PolicyID:     "3e251bce0d1411efb0a10242ac110002",
		InstanceName: "",
		PageNo:       1,
		PageSize:     10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func instanceBackupPolicyBindInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.InstanceBackupPolicyBindInstancesApi.Do(context.Background(), credential, &ctecs.InstanceBackupPolicyBindInstancesRequest{
		RegionID:       "bb9fdb42056f11eda1610242ac110002",
		PolicyID:       "3e251bce0d1411efb0a10242ac110002",
		InstanceIDList: "4bde19ee-1e3a-bb84-9ee2-0e55de396a8e",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func instanceBackupPolicyUnbindInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.InstanceBackupPolicyUnbindInstancesApi.Do(context.Background(), credential, &ctecs.InstanceBackupPolicyUnbindInstancesRequest{
		RegionID:       "bb9fdb42056f11eda1610242ac110002",
		PolicyID:       "3e251bce0d1411efb0a10242ac110002",
		InstanceIDList: "4bde19ee-1e3a-bb84-9ee2-0e55de396a8e",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func instanceBackupPolicyBindRepo(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.InstanceBackupPolicyBindRepoApi.Do(context.Background(), credential, &ctecs.InstanceBackupPolicyBindRepoRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		RepositoryID: "98658ed6-e699-426c-af3a-f6b6343a9829",
		PolicyID:     "3e251bce0d1411efb0a10242ac110002",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func instanceBackupPolicyUnbindRepo(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.InstanceBackupPolicyUnbindRepoApi.Do(context.Background(), credential, &ctecs.InstanceBackupPolicyUnbindRepoRequest{
		RegionID: "bb9fdb42056f11eda1610242ac110002",
		PolicyID: "3e251bce0d1411efb0a10242ac110002",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func createSnapshot(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.CreateSnapshotApi.Do(context.Background(), credential, &ctecs.CreateSnapshotRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		InstanceID:   "4bde19ee-1e3a-bb84-9ee2-0e55de396a8e",
		SnapshotName: "test-go-sdk",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func restoreSnapshot(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.RestoreSnapshotApi.Do(context.Background(), credential, &ctecs.RestoreSnapshotRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		SnapshotID: "a2c7ef3c-4290-15ed-bb6e-a03d67a46394",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func updateSnapshot(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.UpdateSnapshotApi.Do(context.Background(), credential, &ctecs.UpdateSnapshotRequest{
		RegionID:            "bb9fdb42056f11eda1610242ac110002",
		SnapshotID:          "a2c7ef3c-4290-15ed-bb6e-a03d67a46394",
		SnapshotName:        "napshot_update_01",
		SnapshotDescription: "snapshot_des",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func batchUpdateSnapshot(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	UpdateInfo := make([]ctecs.BatchUpdateSnapshotUpdateInfoRequest, 0)
	updateInfo := ctecs.BatchUpdateSnapshotUpdateInfoRequest{
		SnapshotID:          "a2c7ef3c-4290-15ed-bb6e-a03d67a46394",
		SnapshotName:        "snapshot_update_batch01",
		SnapshotDescription: "snapshot_update_des",
	}

	UpdateInfo = append(UpdateInfo, updateInfo)

	response, err := apis.BatchUpdateSnapshotApi.Do(context.Background(), credential, &ctecs.BatchUpdateSnapshotRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		UpdateInfo: UpdateInfo,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func querySnapshotList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.QuerySnapshotListApi.Do(context.Background(), credential, &ctecs.QuerySnapshotListRequest{
		RegionID:       "bb9fdb42056f11eda1610242ac110002",
		ProjectID:      "0",
		PageNo:         1,
		PageSize:       10,
		InstanceID:     "4bde19ee-1e3a-bb84-9ee2-0e55de396a8e",
		SnapshotStatus: "",
		SnapshotID:     "",
		QueryContent:   "",
		SnapshotName:   "",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func querySnapshotDetails(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.QuerySnapshotDetailsApi.Do(context.Background(), credential, &ctecs.QuerySnapshotDetailsRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		SnapshotID: "a2c7ef3c-4290-15ed-bb6e-a03d67a46394",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func querySnapshotStatistics(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.QuerySnapshotStatisticsApi.Do(context.Background(), credential, &ctecs.QuerySnapshotStatisticsRequest{
		RegionID:       "bb9fdb42056f11eda1610242ac110002",
		InstanceIDList: "a2c7ef3c-4290-15ed-bb6e-a03d67a46394",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func querySnapshotStatus(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.QuerySnapshotStatusApi.Do(context.Background(), credential, &ctecs.QuerySnapshotStatusRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		SnapshotID: "a2c7ef3c-4290-15ed-bb6e-a03d67a46394",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func deleteSnapshot(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.DeleteSnapshotApi.Do(context.Background(), credential, &ctecs.DeleteSnapshotRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		SnapshotID: "a2c7ef3c-4290-15ed-bb6e-a03d67a46394",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVmDiskLatestMetricData(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsVmDiskLatestMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmDiskLatestMetricDataRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		DeviceIDList: []string{"0fec78e4-1889-803f-b2a7-515c1c40b788"},
		PageNo:       1,
		Page:         1,
		PageSize:     10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVmCpuLatestMetricData(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsVmCpuLatestMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmCpuLatestMetricDataRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		DeviceIDList: []string{"0fec78e4-1889-803f-b2a7-515c1c40b788"},
		PageNo:       1,
		Page:         1,
		PageSize:     10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVmNetworkHistoryMetricData(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsVmNetworkHistoryMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmNetworkHistoryMetricDataRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		DeviceIDList: []string{"e2a016ec-543b-08f1-38c4-13d9dac55b5a"},
		Period:       14400,
		StartTime:    "1717402682",
		EndTime:      "1717575482",
		PageNo:       1,
		Page:         1,
		PageSize:     10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVmMemHistoryMetricData(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsVmMemHistoryMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmMemHistoryMetricDataRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		DeviceIDList: []string{"e2a016ec-543b-08f1-38c4-13d9dac55b5a"},
		Period:       14400,
		StartTime:    "1717402682",
		EndTime:      "1717575482",
		PageNo:       1,
		Page:         1,
		PageSize:     10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsVmDiskHistoryMetricData(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsVmDiskHistoryMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmDiskHistoryMetricDataRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		DeviceIDList: []string{"e2a016ec-543b-08f1-38c4-13d9dac55b5a"},
		Period:       14400,
		StartTime:    "1717402682",
		EndTime:      "1717575482",
		PageNo:       1,
		Page:         1,
		PageSize:     10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsVmCpuHistoryMetricData(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsVmCpuHistoryMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmCpuHistoryMetricDataRequest{
		RegionID:     "bb9fdb42056f11eda1610242ac110002",
		DeviceIDList: []string{"e2a016ec-543b-08f1-38c4-13d9dac55b5a"},
		Period:       14400,
		StartTime:    "1717402682",
		EndTime:      "1717575482",
		PageNo:       1,
		Page:         1,
		PageSize:     10,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupCreate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupCreateApi.Do(context.Background(), credential, &ctecs.EcsBackupCreateRequest{
		RegionID:                  "bb9fdb42056f11eda1610242ac110002",
		InstanceID:                "bd39aca6-e10e-5fab-b8ce-cebd4fe79aae",
		InstanceBackupName:        "backup-061101",
		InstanceBackupDescription: "creat_test01",
		RepositoryID:              "00dbd561-163f-43fb-80d2-ee9744219f9c",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupUpdate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupUpdateApi.Do(context.Background(), credential, &ctecs.EcsBackupUpdateRequest{
		RegionID:                  "bb9fdb42056f11eda1610242ac110002",
		InstanceBackupID:          "d7e44422-12ff-acc5-e87c-b80a7dd659b1",
		InstanceBackupName:        "update-test-sdk-02",
		InstanceBackupDescription: "api_update_test01",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupBatchUpdate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	UpdateInfo := make([]ctecs.EcsBackupBatchUpdateUpdateInfoRequest, 0)
	updateInfo := ctecs.EcsBackupBatchUpdateUpdateInfoRequest{
		InstanceBackupID:          "d7e44422-12ff-acc5-e87c-b80a7dd659b1",
		InstanceBackupName:        "update-test-sdk-03",
		InstanceBackupDescription: "api_update_test03",
	}

	UpdateInfo = append(UpdateInfo, updateInfo)

	response, err := apis.EcsBackupBatchUpdateApi.Do(context.Background(), credential, &ctecs.EcsBackupBatchUpdateRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		UpdateInfo: UpdateInfo,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupUsage(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupUsageApi.Do(context.Background(), credential, &ctecs.EcsBackupUsageRequest{
		RegionID:         "bb9fdb42056f11eda1610242ac110002",
		InstanceBackupID: "d7e44422-12ff-acc5-e87c-b80a7dd659b1",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupStatistics(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupStatisticsApi.Do(context.Background(), credential, &ctecs.EcsBackupStatisticsRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		InstanceID: "4ae65791-f489-66ef-dd13-73ecbca89672",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupInstanceResource(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupInstanceResourceApi.Do(context.Background(), credential, &ctecs.EcsBackupInstanceResourceRequest{
		RegionID: "bb9fdb42056f11eda1610242ac110002",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupListApi.Do(context.Background(), credential, &ctecs.EcsBackupListRequest{
		RegionID:             "bb9fdb42056f11eda1610242ac110002",
		PageNo:               1,
		PageSize:             10,
		InstanceID:           "de70ef00-1ea0-459a-b74d-b06272561a32",
		RepositoryID:         "de70ef00-1ea0-459a-b74d-b06272561a32",
		InstanceBackupID:     "ed48dc25-d6bb-48e6-b202-3e36ee6321a3",
		QueryContent:         "backup-test01",
		InstanceBackupStatus: "ACTIVE",
		ProjectID:            "0",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupDetails(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupDetailsApi.Do(context.Background(), credential, &ctecs.EcsBackupDetailsRequest{
		RegionID:         "bb9fdb42056f11eda1610242ac110002",
		InstanceBackupID: "d7e44422-12ff-acc5-e87c-b80a7dd659b1",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupInstanceDetails(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupInstanceDetailsApi.Do(context.Background(), credential, &ctecs.EcsBackupInstanceDetailsRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		InstanceID: "4ae65791-f489-66ef-dd13-73ecbca89672",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupStatus(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupStatusApi.Do(context.Background(), credential, &ctecs.EcsBackupStatusRequest{
		RegionID:         "bb9fdb42056f11eda1610242ac110002",
		InstanceBackupID: "d7e44422-12ff-acc5-e87c-b80a7dd659b1",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupRestore(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupRestoreApi.Do(context.Background(), credential, &ctecs.EcsBackupRestoreRequest{
		RegionID:         "bb9fdb42056f11eda1610242ac110002",
		InstanceBackupID: "d7e44422-12ff-acc5-e87c-b80a7dd659b1",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupDelete(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupDeleteApi.Do(context.Background(), credential, &ctecs.EcsBackupDeleteRequest{
		RegionID:         "bb9fdb42056f11eda1610242ac110002",
		InstanceBackupID: "b1d7cd00-4ab9-b1c0-4aac-b72184c21ae4",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsAgentBatchAction(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	ActionInfo := make([]ctecs.EcsAgentBatchActionActionInfoRequest, 0)
	actionInfo := ctecs.EcsAgentBatchActionActionInfoRequest{
		InstanceID:    "bd39aca6-e10e-5fab-b8ce-cebd4fe79aae",
		SystemType:    "linux",
		SystemArch:    "amd64",
		SystemVersion: "1.27.6",
	}

	ActionInfo = append(ActionInfo, actionInfo)

	response, err := apis.EcsAgentBatchActionApi.Do(context.Background(), credential, &ctecs.EcsAgentBatchActionRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		Action:     "update",
		ActionInfo: ActionInfo,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsCreate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsCreateApi.Do(context.Background(), credential, &ctecs.EcsPortsCreateRequest{
		ClientToken:             "ports_create061801",
		RegionID:                "bb9fdb42056f11eda1610242ac110002",
		SubnetID:                "subnet-3o8uvvp6h4",
		PrimaryPrivateIp:        "172.16.0.141",
		Ipv6Addresses:           []string{"240e:978:497c:ec00:cd74:fd9d:c45d:4131"},
		SecurityGroupIds:        []string{"sg-n7nu88xfbq"},
		SecondaryPrivateIpCount: 1,
		SecondaryPrivateIps:     []string{"172.16.0.210"},
		Name:                    "nic-test01",
		Description:             "dec-test",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsDelete(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsDeleteApi.Do(context.Background(), credential, &ctecs.EcsPortsDeleteRequest{
		ClientToken:        "delete-ports-test-01",
		RegionID:           "bb9fdb42056f11eda1610242ac110002",
		NetworkInterfaceID: "port-qp8i3s4c2h",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsUpdate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsUpdateApi.Do(context.Background(), credential, &ctecs.EcsPortsUpdateRequest{
		ClientToken:        "update-port-test-01",
		RegionID:           "bb9fdb42056f11eda1610242ac110002",
		NetworkInterfaceID: "port-vfa9afga2b",
		Name:               "nic-update-name",
		Description:        "nic_update_description",
		SecurityGroupIDs:   []string{"sg-tdzefke02r"},
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsListApi.Do(context.Background(), credential, &ctecs.EcsPortsListRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		VpcID:      "vpc-r5i4zghgvq",
		DeviceID:   "a628a7d9-ef97-3b16-8a0a-4a794fcdxxxx",
		SubnetID:   "subnet-r5i4zghgvq",
		PageNumber: 1,
		PageSize:   10,
		PageNo:     1,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsShow(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsShowApi.Do(context.Background(), credential, &ctecs.EcsPortsShowRequest{
		RegionID:           "bb9fdb42056f11eda1610242ac110002",
		NetworkInterfaceID: "port-vfa9afga2b",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsAttach(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsAttachApi.Do(context.Background(), credential, &ctecs.EcsPortsAttachRequest{
		ClientToken:        "attach_test01",
		RegionID:           "bb9fdb42056f11eda1610242ac110002",
		AzName:             "cn-huadong1-jsnj1A-public-ctcloud",
		ProjectID:          "0",
		NetworkInterfaceID: "port-vfa9afga2b",
		InstanceID:         "4bde19ee-1e3a-bb84-9ee2-0e55de396a8e",
		InstanceType:       3,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsDetach(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsDetachApi.Do(context.Background(), credential, &ctecs.EcsPortsDetachRequest{
		ClientToken:        "ports-detach-01",
		RegionID:           "bb9fdb42056f11eda1610242ac110002",
		NetworkInterfaceID: "port-vfa9afga2b",
		InstanceID:         "4bde19ee-1e3a-bb84-9ee2-0e55de396a8e",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsAssignIpv6(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsAssignIpv6Api.Do(context.Background(), credential, &ctecs.EcsPortsAssignIpv6Request{
		ClientToken:        "assign-ipv6-01",
		RegionID:           "bb9fdb42056f11eda1610242ac110002",
		NetworkInterfaceID: "port-vfa9afga2b",
		Ipv6AddressesCount: 1,
		//Ipv6Addresses:      []string{""},
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsUnassignIpv6(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsUnassignIpv6Api.Do(context.Background(), credential, &ctecs.EcsPortsUnassignIpv6Request{
		ClientToken:        "unassign-ipv6-01",
		RegionID:           "bb9fdb42056f11eda1610242ac110002",
		NetworkInterfaceID: "port-vfa9afga2b",
		Ipv6Addresses:      []string{"240e:978:49f5:3100:bb42:3ddf:5960:98c7"},
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsAssignSecondaryPrivateIps(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsAssignSecondaryPrivateIpsApi.Do(context.Background(), credential, &ctecs.EcsPortsAssignSecondaryPrivateIpsRequest{
		ClientToken:             "assign-secondary-private-ips-01",
		RegionID:                "bb9fdb42056f11eda1610242ac110002",
		NetworkInterfaceID:      "port-vfa9afga2b",
		SecondaryPrivateIpCount: 1,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsPortsUnassignSecondaryPrivateIps(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsPortsUnassignSecondaryPrivateIpsApi.Do(context.Background(), credential, &ctecs.EcsPortsUnassignSecondaryPrivateIpsRequest{
		ClientToken:         "unassign-secondary-private-ips-01",
		RegionID:            "bb9fdb42056f11eda1610242ac110002",
		NetworkInterfaceID:  "port-vfa9afga2b",
		SecondaryPrivateIps: []string{"192.168.0.5"},
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsEipCreate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsEipCreateApi.Do(context.Background(), credential, &ctecs.EcsEipCreateRequest{
		ClientToken:       "create-eip-test",
		RegionID:          "bb9fdb42056f11eda1610242ac110002",
		ProjectID:         "0",
		CycleType:         "month",
		CycleCount:        1,
		Name:              "eip-name",
		Bandwidth:         5,
		BandwidthID:       "bandwidth-7hzv449r2j",
		DemandBillingType: "bandwidth",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsEipDelete(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsEipDeleteApi.Do(context.Background(), credential, &ctecs.EcsEipDeleteRequest{
		ClientToken: "delete-eip-test",
		RegionID:    "bb9fdb42056f11eda1610242ac110002",
		ProjectID:   "0",
		EipID:       "eip-lunma2v53e",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsShareInterfaceAttach(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	response, err := apis.EcsShareInterfaceAttachApi.Do(context.Background(), credential, &ctecs.EcsShareInterfaceAttachRequest{
		RegionID:   "bb9fdb42056f11eda1610242ac110002",
		InstanceID: "4bde19ee-1e3a-bb84-9ee2-0e55de396a8e",
		SubnetID:   "subnet-3o8uvvp6h4",
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsBackupCreateInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	networkCardList := make([]ctecs.EcsBackupCreateInstanceNetworkCardListRequest, 0)
	network_card := ctecs.EcsBackupCreateInstanceNetworkCardListRequest{
		NicName:  "nic-test-061901",
		IsMaster: true,
		SubnetID: "subnet-4c4333pc67",
	}
	networkCardList = append(networkCardList, network_card)
	labelList := make([]ctecs.EcsBackupCreateInstanceLabelListRequest, 0)
	label := ctecs.EcsBackupCreateInstanceLabelListRequest{
		LabelKey:   "label-key-test",
		LabelValue: "label-value-test",
	}
	labelList = append(labelList, label)
	response, err := apis.EcsBackupCreateInstanceApi.Do(context.Background(), credential, &ctecs.EcsBackupCreateInstanceRequest{
		ClientToken:      "ecs-backup-create-instance-test-061901",
		RegionID:         "bb9fdb42056f11eda1610242ac110002",
		AzName:           "cn-huadong1-jsnj1A-public-ctcloud",
		InstanceName:     "ecm-go-test-061901",
		DisplayName:      "ecm-go-test-061901",
		InstanceBackupID: "e718f4b2-0ff2-e486-b322-ea206fbce240",
		FlavorID:         "34e1b6f6-e974-1575-20b2-172ba0e0bf83",
		VpcID:            "vpc-chz0ilszsp",
		OnDemand:         false,
		NetworkCardList:  networkCardList,
		ExtIP:            "1",
		IpVersion:        "ipv4",
		Bandwidth:        50,
		ProjectID:        "0",
		Ipv6AddressID:    "",
		SecGroupList:     []string{"sg-ku5edgbitc"},
		EipID:            "",
		AffinityGroupID:  "",
		KeyPairID:        "",
		UserPassword:     "qyo84!*ymd",
		CycleCount:       1,
		CycleType:        "MONTH",
		AutoRenewStatus:  0,
		UserData:         "YmF0Y2hDcmVhdGVUZXN0MDgwMw==",
		PayVoucherPrice:  1819.50,
		LabelList:        labelList,
		MonitorService:   true,
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
	deleteSnapshot(*credential)
	querySnapshotStatus(*credential)
	querySnapshotStatistics(*credential)
	querySnapshotDetails(*credential)
	querySnapshotList(*credential)
	batchUpdateSnapshot(*credential)
	updateSnapshot(*credential)
	restoreSnapshot(*credential)
	createSnapshot(*credential)
	instanceBackupPolicyUnbindRepo(*credential)
	instanceBackupPolicyBindRepo(*credential)
	instanceBackupPolicyUnbindInstances(*credential)
	instanceBackupPolicyBindInstances(*credential)
	listInstanceBackupPolicyBindInstances(*credential)
	listInstanceBackupPolicy(*credential)
	batchUpdateInstancesPassword(*credential)
	getVolumeStatistics(*credential)
	queryVncDetails(*credential)
	updateInstance(*credential)
	ecsFlavorFamiliesList(*credential)
	ecsAvailabilityZonesDetails(*credential)
	ecsListInstances(*credential)
	ecsBatchCreateInstance(*credential)
	ecsFlavorListByFamilies(*credential)
	ecsGpuDriverList(*credential)
	ecsUpdateInstanceSpec(*credential)
	ecsUpdateNetworkSpec(*credential)
	ecsBatchStopInstances(*credential)
	ecsBackupCreateInstance(*credential)
	ecsShareInterfaceAttach(*credential)
	ecsEipDelete(*credential)
	ecsEipCreate(*credential)
	ecsPortsUnassignSecondaryPrivateIps(*credential)
	ecsPortsAssignSecondaryPrivateIps(*credential)
	ecsPortsUnassignIpv6(*credential)
	ecsPortsAssignIpv6(*credential)
	ecsPortsDetach(*credential)
	ecsPortsAttach(*credential)
	ecsPortsShow(*credential)
	ecsPortsList(*credential)
	ecsPortsUpdate(*credential)
	ecsPortsDelete(*credential)
	ecsPortsCreate(*credential)
	ecsAgentBatchAction(*credential)
	ecsBackupDelete(*credential)
	ecsBackupRestore(*credential)
	ecsBackupStatus(*credential)
	ecsBackupInstanceDetails(*credential)
	ecsBackupDetails(*credential)
	ecsBackupList(*credential)
	ecsBackupInstanceResource(*credential)
	ecsBackupStatistics(*credential)
	ecsBackupUsage(*credential)
	ecsBackupBatchUpdate(*credential)
	ecsBackupUpdate(*credential)
	ecsBackupCreate(*credential)
	ecsVmCpuHistoryMetricData(*credential)
	ecsVmDiskHistoryMetricData(*credential)
	ecsVmMemHistoryMetricData(*credential)
	ecsVmNetworkHistoryMetricData(*credential)
	ecsVmCpuLatestMetricData(*credential)
	ecsVmDiskLatestMetricData(*credential)
	listFlavors(*credential)
	createInstance(*credential)
	describeInstances(*credential)
	instanceDetail(*credential)
	instanceJoinSecurityGroup(*credential)
	instanceLeaveSecurityGroup(*credential)
	instanceQueryAsyncResult(*credential)
	resetInstancePassword(*credential)
	startInstance(*credential)
	stopInstance(*credential)
	unsubscribeInstance(*credential)
	updateFlavorSpec(*credential)
	listInstanceVolumes(*credential)
	keypairCreate(*credential)
	keypairImport(*credential)
	keypairDetail(*credential)
	keypairAttach(*credential)
	keypairDetach(*credential)
	keypairDelete(*credential)
	rebootInstance(*credential)
	batchRebootInstances(*credential)
	rebuildInstance(*credential)
	batchRebuildInstances(*credential)
	batchUnsubscribeInstances(*credential)
}
