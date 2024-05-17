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
		PayVoucherPrice: 1819.99,
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
		InstanceIdList:  "77493826-d038-2a9c-f684-e2f6adabeba3",
		SecurityGroupId: "sg-tdzefke02r",
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
		InstanceId: "77493826-d038-2a9c-f684-e2f6adabeba3",
		Force:      false,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func UnsubscribeInstance(credential ctyunsdk.Credential) {
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
func UpdateFlavorSpec(credential ctyunsdk.Credential) {
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

func main() {
	credential, _ := ctyunsdk.NewCredential("ak", "sk")
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
	UnsubscribeInstance(*credential)
	UpdateFlavorSpec(*credential)
	listInstanceVolumes(*credential)
	keypairCreate(*credential)
	keypairImport(*credential)
	keypairDetail(*credential)
	keypairAttach(*credential)
	keypairDetach(*credential)
	keypairDelete(*credential)
}
