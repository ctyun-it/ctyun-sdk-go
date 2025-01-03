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

	nic_name := "nic-test"
	nic_is_master := true
	subnet_id := "fe0c40b4-4176-4623-b076-1b71562da129"
	networkCardList := make([]ctecs.EcsCreateInstanceNetworkCardListRequest, 0)
	networkCard := ctecs.EcsCreateInstanceNetworkCardListRequest{
		NicName:  &nic_name,
		IsMaster: &nic_is_master,
		SubnetId: &subnet_id,
	}
	networkCardList = append(networkCardList, networkCard)
	data_disk_mode := "VBD"
	data_disk_name := "data-disk-test"
	data_disk_type := "SATA"
	data_disk_size := 20
	dataDiskList := make([]ctecs.EcsCreateInstanceDataDiskListRequest, 0)
	dataDisk := ctecs.EcsCreateInstanceDataDiskListRequest{
		DiskMode: &data_disk_mode,
		DiskName: &data_disk_name,
		DiskType: &data_disk_type,
		DiskSize: &data_disk_size,
	}
	dataDiskList = append(dataDiskList, dataDisk)
	label_key := "label-key-test"
	label_value := "label-value-test"
	labelList := make([]ctecs.EcsCreateInstanceLabelListRequest, 0)
	label := ctecs.EcsCreateInstanceLabelListRequest{
		LabelKey:   &label_key,
		LabelValue: &label_value,
	}
	labelList = append(labelList, label)
	client_token := ""
	region_id := ""
	az_name := ""
	instance_name := "ecm-go-test"
	display_name := "ecm-go-test"
	flavor_id := ""
	image_type := 1
	image_id := ""
	boot_disk_type := "SATA"
	boot_disk_size := 40
	vpc_id := ""
	on_demand := false
	ext_ip := "2"
	project_id := "0"
	sec_group_list := []string{""}
	ip_version := "ipv4"
	band_width := 50
	ipv6_address_id := ""
	eip_id := ""
	affinity_group_id := ""
	keypair_id := ""
	user_password := ""
	cycle_count := 1
	cycle_type := "MONTH"
	auto_renew_status := 0
	user_data := "YmF0Y2hDcmVhdGVUZXN0MDgwMw=="
	monitor_service := true
	pay_voucher_price := 1888.62

	response, err := apis.EcsCreateInstanceApi.Do(context.Background(), credential, &ctecs.EcsCreateInstanceRequest{
		ClientToken:     &client_token,
		RegionId:        &region_id,
		AzName:          &az_name,
		InstanceName:    &instance_name,
		DisplayName:     &display_name,
		FlavorId:        &flavor_id,
		ImageType:       &image_type,
		ImageId:         &image_id,
		BootDiskType:    &boot_disk_type,
		BootDiskSize:    &boot_disk_size,
		VpcId:           &vpc_id,
		OnDemand:        &on_demand,
		NetworkCardList: networkCardList,
		ExtIp:           &ext_ip,
		ProjectID:       &project_id,
		SecGroupList:    &sec_group_list,
		DataDiskList:    dataDiskList,
		IpVersion:       &ip_version,
		Bandwidth:       &band_width,
		Ipv6AddressID:   &ipv6_address_id,
		EipID:           &eip_id,
		AffinityGroupID: &affinity_group_id,
		KeyPairID:       &keypair_id,
		UserPassword:    &user_password,
		CycleCount:      &cycle_count,
		CycleType:       &cycle_type,
		AutoRenewStatus: &auto_renew_status,
		UserData:        &user_data,
		PayVoucherPrice: &pay_voucher_price,
		LabelList:       labelList,
		MonitorService:  &monitor_service,
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

	// 网卡列表
	networkCardList := make([]ctecs.EcsBatchCreateInstancesNetworkCardListRequest, 0)
	nic_name1 := "nic-test"
	nic_is_master1 := true
	subnet_id1 := "subnet-ps8hjw16vt"
	fixed_ip1 := "192.168.0.26"
	network_card := ctecs.EcsBatchCreateInstancesNetworkCardListRequest{
		NicName:  &nic_name1,
		IsMaster: &nic_is_master1,
		SubnetID: &subnet_id1,
		FixedIP:  &fixed_ip1,
	}
	networkCardList = append(networkCardList, network_card)

	// 数据盘列表
	dataDiskList := make([]ctecs.EcsBatchCreateInstancesDataDiskListRequest, 0)
	disk_mode1 := "VBD"
	disk_name1 := "data-disk-test"
	disk_type1 := "SATA"
	data_disk_size1 := 20
	data_disk := ctecs.EcsBatchCreateInstancesDataDiskListRequest{
		DiskMode: &disk_mode1,
		DiskName: &disk_name1,
		DiskType: &disk_type1,
		DiskSize: &data_disk_size1,
	}
	dataDiskList = append(dataDiskList, data_disk)

	// 标签列表
	labelList := make([]ctecs.EcsBatchCreateInstancesLabelListRequest, 0)
	label_key := "label-key-test"
	label_value := "label-value-test"
	label := ctecs.EcsBatchCreateInstancesLabelListRequest{
		LabelKey:   &label_key,
		LabelValue: &label_value,
	}
	labelList = append(labelList, label)

	// 具体请求
	client_token := "ecs-batch-create-instances-test-02"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	az_name := "cn-huadong1-jsnj1A-public-ctcloud"
	instance_name := "go-sdk-test"
	display_name := "go-sdk-test"
	flavor_id := "5622ce59-da34-cb43-ca0d-eef2a51475b3"
	image_type := 1
	image_id := "b78812b0-ff50-4816-b58f-5c4fbc230b08"
	boot_disk_type := "SATA"
	boot_disk_size := 40
	vpc_id := "vpc-riwxr5wpju"
	on_demand := false
	ext_ip := "1"
	project_id := "0"
	sec_group_list := []string{"sg-bqv0t629h6", "sg-bqv0t629h6"}
	ip_version := "ipv4"
	bandwidth := 1
	ipv6_address_id := ""
	eip_id := ""
	affinity_group_id := ""
	key_pair_id := ""
	password := "qyo84!*ymd"
	cycle_count := 1
	cycle_type := "MONTH"
	auto_renew_status := 0
	pay_voucher_price := 114.00
	monitor_service := true
	order_count := 2
	user_data := "YmF0Y2hDcmVhdGVUZXN0MDgwMw=="

	response, err := apis.EcsBatchCreateInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchCreateInstancesRequest{
		ClientToken:     &client_token,
		RegionID:        &region_id,
		AzName:          &az_name,
		InstanceName:    &instance_name,
		DisplayName:     &display_name,
		FlavorID:        &flavor_id,
		ImageType:       &image_type,
		ImageID:         &image_id,
		BootDiskType:    &boot_disk_type,
		BootDiskSize:    &boot_disk_size,
		VpcID:           &vpc_id,
		OnDemand:        &on_demand,
		NetworkCardList: networkCardList,
		ExtIP:           &ext_ip,
		ProjectID:       &project_id,
		SecGroupList:    &sec_group_list,
		DataDiskList:    dataDiskList,
		IpVersion:       &ip_version,
		Bandwidth:       &bandwidth,
		Ipv6AddressID:   &ipv6_address_id,
		EipID:           &eip_id,
		AffinityGroupID: &affinity_group_id,
		KeyPairID:       &key_pair_id,
		UserPassword:    &password,
		CycleCount:      &cycle_count,
		CycleType:       &cycle_type,
		AutoRenewStatus: &auto_renew_status,
		UserData:        &user_data,
		PayVoucherPrice: &pay_voucher_price,
		LabelList:       labelList,
		MonitorService:  &monitor_service,
		OrderCount:      &order_count,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	az_name := "cn-huadong1-jsnj1A-public-ctcloud"
	project_id := "0"
	page_no := 1
	page_size := 10
	state := "active"
	key_word := "ecm-57fd"
	instance_name := "ecm-57fd"
	instance_id_list := "0fec78e4-1889-803f-b2a7-515c1c40b788"
	security_group_id := "sg-tdzefke02r"
	response, err := apis.EcsDescribeInstancesApi.Do(context.Background(), credential, &ctecs.EcsDescribeInstancesRequest{
		RegionId:        &region_id,
		AzName:          &az_name,
		ProjectId:       &project_id,
		PageNo:          &page_no,
		PageSize:        &page_size,
		State:           &state,
		Keyword:         &key_word,
		InstanceName:    &instance_name,
		InstanceIdList:  &instance_id_list,
		SecurityGroupId: &security_group_id,
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
	label_key1 := "label-key-test"
	label_value1 := "label-value-test"
	label_list := ctecs.EcsListInstancesLabelListRequest{
		LabelKey:   &label_key1,
		LabelValue: &label_value1,
	}
	labelList = append(labelList, label_list)

	region_id := "bb9fdb42056f11eda1610242ac110002"
	az_name := "cn-huadong1-jsnj1A-public-ctcloud"
	project_id := "0"
	page_no := 1
	page_size := 10
	state := "active"
	keyword := "ecm-57fd"
	instance_name := "ecm-57fd"
	instance_id_list := "0fec78e4-1889-803f-b2a7-515c1c40b788"
	security_group_id := "sg-tdzefke02r"
	vpc_id := "vpc-chz0ilszsp"
	//resource_id := "d57ec586da6c497ea1e1b04e08ad9a8b"
	sort := "expiredTime"
	asc := true

	response, err := apis.EcsListInstancesApi.Do(context.Background(), credential, &ctecs.EcsListInstancesRequest{
		RegionID:        &region_id,
		AzName:          &az_name,
		ProjectID:       &project_id,
		PageNo:          &page_no,
		PageSize:        &page_size,
		State:           &state,
		Keyword:         &keyword,
		InstanceName:    &instance_name,
		InstanceIDList:  &instance_id_list,
		SecurityGroupID: &security_group_id,
		VpcID:           &vpc_id,
		//ResourceID:      &resource_id,
		//LabelList:       labelList,
		Sort: &sort,
		Asc:  &asc,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	az_name := "cn-huadong1-jsnj1A-public-ctcloud"
	flavor_type := "CPU_KS1"
	flavor_name := "ks1.medium.2"
	flavor_cpu := 1
	flavor_arch := "arm"
	flavor_series := "ks"
	flavor_id := "b6779240-5649-803b-4a4c-8fc59d310ecf"
	response, err := apis.EcsFlavorListApi.Do(context.Background(), credential, &ctecs.EcsFlavorListRequest{
		RegionId:     &region_id,
		AzName:       &az_name,
		FlavorType:   &flavor_type,
		FlavorName:   &flavor_name,
		FlavorCpu:    &flavor_cpu,
		FlavorArch:   &flavor_arch,
		FlavorSeries: &flavor_series,
		FlavorId:     &flavor_id,
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

	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	response, err := apis.EcsInstanceDetailsApi.Do(context.Background(), credential, &ctecs.EcsInstanceDetailsRequest{
		RegionId:   &region_id,
		InstanceId: &instance_id,
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

	region_id := "bb9fdb42056f11eda1610242ac110002"
	security_group_id := "sg-tdzefke02r"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	network_interface_id := "port-pja7l0zfvk"
	action := "joinSecurityGroup"
	response, err := apis.EcsJoinSecurityGroupApi.Do(context.Background(), credential, &ctecs.EcsJoinSecurityGroupRequest{
		RegionId:           &region_id,
		SecurityGroupId:    &security_group_id,
		InstanceId:         &instance_id,
		NetworkInterfaceId: &network_interface_id,
		Action:             &action,
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

	region_id := "bb9fdb42056f11eda1610242ac110002"
	security_group_id := "sg-tdzefke02r"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	response, err := apis.EcsLeaveSecurityGroupApi.Do(context.Background(), credential, &ctecs.EcsLeaveSecurityGroupRequest{
		RegionId:        &region_id,
		SecurityGroupId: &security_group_id,
		InstanceId:      &instance_id,
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

	region_id := "bb9fdb42056f11eda1610242ac110002"
	job_id := "XXX"
	response, err := apis.EcsQueryAsyncResultApi.Do(context.Background(), credential, &ctecs.EcsQueryAsyncResultRequest{
		RegionId: &region_id,
		JobId:    &job_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	new_password := "test-test-test-960"
	response, err := apis.EcsResetPasswordApi.Do(context.Background(), credential, &ctecs.EcsResetPasswordRequest{
		RegionId:    &region_id,
		InstanceId:  &instance_id,
		NewPassword: &new_password,
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

	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	response, err := apis.EcsStartInstanceApi.Do(context.Background(), credential, &ctecs.EcsStartInstanceRequest{
		RegionId:   &region_id,
		InstanceId: &instance_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "63afb617-b8f5-d482-9ecd-6d8bb9124d4e"
	force := false
	response, err := apis.EcsStopInstanceApi.Do(context.Background(), credential, &ctecs.EcsStopInstanceRequest{
		RegionId:   &region_id,
		InstanceId: &instance_id,
		Force:      &force,
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
	client_token := "unsubscribe-instance-test"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	delete_volume := false
	delete_eip := false
	response, err := apis.EcsUnsubscribeInstanceApi.Do(context.Background(), credential, &ctecs.EcsUnsubscribeInstanceRequest{
		ClientToken:  &client_token,
		RegionId:     &region_id,
		InstanceId:   &instance_id,
		DeleteVolume: &delete_volume,
		DeleteEIP:    &delete_eip,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	flavor_id := "b6779240-5649-803b-4a4c-8fc59d310ecf"
	client_token := "update-flavor-spec-test"
	pay_voucher_price := 100.00
	response, err := apis.EcsUpdateFlavorSpecApi.Do(context.Background(), credential, &ctecs.EcsUpdateFlavorSpecRequest{
		RegionId:        &region_id,
		InstanceId:      &instance_id,
		FlavorId:        &flavor_id,
		ClientToken:     &client_token,
		PayVoucherPrice: &pay_voucher_price,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	page_no := 1
	page_size := 10
	response, err := apis.EcsVolumeListApi.Do(context.Background(), credential, &ctecs.EcsVolumeListRequest{
		RegionId:   &region_id,
		InstanceId: &instance_id,
		PageNo:     &page_no,
		PageSize:   &page_size,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	keypair_name := "keypair-test"
	project_id := "0"
	res, err := apis.KeypairCreateApi.Do(context.Background(), credential, &ctecs.KeypairCreateRequest{
		RegionId:    &region_id,
		KeyPairName: &keypair_name,
		ProjectId:   &project_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	keypair_name := "keypair-test"
	res, err := apis.KeypairDeleteApi.Do(context.Background(), credential, &ctecs.KeypairDeleteRequest{
		RegionId:    &region_id,
		KeyPairName: &keypair_name,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	projetc_id := "0"
	query_content := "keypair-test"
	page_no := 1
	page_size := 10
	keypair_name := "keypair_test"
	res, err := apis.KeypairDetailApi.Do(context.Background(), credential, &ctecs.KeypairDetailRequest{
		RegionId:     &region_id,
		KeyPairName:  &keypair_name,
		ProjectId:    &projetc_id,
		QueryContent: &query_content,
		PageNo:       &page_no,
		PageSize:     &page_size,
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
	region_id := ""
	keypair_name := ""
	public_key := ""
	res, err := apis.KeypairImportApi.Do(context.Background(), credential, &ctecs.KeypairImportRequest{
		RegionId:    &region_id,
		KeyPairName: &keypair_name,
		PublicKey:   &public_key,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	keypair_name := "keypair-test"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	res, err := apis.KeypairAttachApi.Do(context.Background(), credential, &ctecs.KeypairAttachRequest{
		RegionId:    &region_id,
		KeyPairName: &keypair_name,
		InstanceId:  &instance_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	keypair_name := "keypair-test"
	instance_id := "77493826-d038-2a9c-f684-e2f6adabeba3"
	res, err := apis.KeypairDetachApi.Do(context.Background(), credential, &ctecs.KeypairDetachRequest{
		RegionId:    &region_id,
		KeyPairName: &keypair_name,
		InstanceId:  &instance_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "0fec78e4-1889-803f-b2a7-515c1c40b788"
	response, err := apis.EcsRebootInstanceApi.Do(context.Background(), credential, &ctecs.EcsRebootInstanceRequest{
		RegionID:   &region_id,
		InstanceID: &instance_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id_list := "0fec78e4-1889-803f-b2a7-515c1c40b788"
	response, err := apis.EcsBatchRebootInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchRebootInstanceRequest{
		RegionID:       &region_id,
		InstanceIDList: &instance_id_list,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "0fec78e4-1889-803f-b2a7-515c1c40b788"
	password := "rebuildTest195%"
	image_id := "b1d896e1-c977-4fd4-b6c2-5432549977be"
	user_data := "UmVidWlsZFRlc3QyMDIyMTEyNDEzMTE="
	instance_name := "ecm-3300"
	monitor_service := true
	pay_image := false
	response, err := apis.EcsRebuildInstanceApi.Do(context.Background(), credential, &ctecs.EcsRebuildInstanceRequest{
		RegionID:       &region_id,
		InstanceID:     &instance_id,
		Password:       &password,
		ImageID:        &image_id,
		UserData:       &user_data,
		InstanceName:   &instance_name,
		MonitorService: &monitor_service,
		PayImage:       &pay_image,
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
	instance_id := "63afb617-b8f5-d482-9ecd-6d8bb9124d4e"
	password := "rebuildTest195%"
	image_id := "b1d896e1-c977-4fd4-b6c2-5432549977be"
	user_data := "UmVidWlsZFRlc3QyMDIyMTEyNDEzMTE="
	instance_name := "ecm-3300"
	monitor_service := true
	rebuildInfoList := make([]ctecs.EcsBatchRebuildInstancesRebuildInfoRequest, 0)
	rebuildInfo := ctecs.EcsBatchRebuildInstancesRebuildInfoRequest{
		InstanceID:     &instance_id,
		Password:       &password,
		ImageID:        &image_id,
		UserData:       &user_data,
		InstanceName:   &instance_name,
		MonitorService: &monitor_service,
	}
	rebuildInfoList = append(rebuildInfoList, rebuildInfo)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	response, err := apis.EcsBatchRebuildInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchRebuildInstancesRequest{
		RegionID:    &region_id,
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
	client_token := "batch-unsubscribe-instance"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id_list := "96b254b1-b472-e72d-eb7f-05b61d973ad4"
	response, err := apis.EcsBatchUnsubscribeInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchUnsubscribeInstanceRequest{
		ClientToken:    &client_token,
		RegionID:       &region_id,
		InstanceIDList: &instance_id_list,
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

	region_id := "81f7728662dd11ec810800155d307d5b"
	instance_id_list := "dba4eaed-fc14-607b-495c-347922ac96fe,d6243178-bd6f-07ab-02a4-6b615f038625"
	force := false
	response, err := apis.EcsBatchStopInstancesApi.Do(context.Background(), credential, &ctecs.EcsBatchStopInstancesRequest{
		RegionID:       &region_id,
		InstanceIDList: &instance_id_list,
		Force:          &force,
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

	region_id := "81f7728662dd11ec810800155d307d5b"
	instance_id := "74d7447a-1275-ec32-cede-8fd285e43f76"
	bandwidth := 15
	client_token := "ea1b9004-f450-11ec-8d4f-00155de3fd73"
	response, err := apis.EcsUpdateNetworkSpecApi.Do(context.Background(), credential, &ctecs.EcsUpdateNetworkSpecRequest{
		RegionID:    &region_id,
		InstanceID:  &instance_id,
		Bandwidth:   &bandwidth,
		ClientToken: &client_token,
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

	region_id := "81f7728662dd11ec810800155d307d5b"
	instance_id := "74d7447a-1275-ec32-cede-8fd285e43f76"
	//bandwidth := 15
	client_token := "ea1b9004-f450-11ec-8d4f-00155de3fd73"
	flavor_id := "5f3ba144-98e4-ea85-ea9b-933394dded33"
	response, err := apis.EcsUpdateInstanceSpecApi.Do(context.Background(), credential, &ctecs.EcsUpdateInstanceSpecRequest{
		RegionID:   &region_id,
		InstanceID: &instance_id,
		//Bandwidth:  &bandwidth,
		FlavorID:    &flavor_id,
		ClientToken: &client_token,
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

	region_id := "bb9fdb42056f11eda1610242ac110002"
	flavor_id := "5cf44a7e-e23c-4199-9ebf-226650646e5a"
	response, err := apis.EcsGpuDriverListApi.Do(context.Background(), credential, &ctecs.EcsGpuDriverListRequest{
		RegionID: &region_id,
		FlavorID: &flavor_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	az_name := "cn-huadong1-jsnj1A-public-ctcloud"
	flavor_family := "s7"
	page_no := 1
	page_size := 10
	response, err := apis.EcsFlavorListByFamiliesApi.Do(context.Background(), credential, &ctecs.EcsFlavorListByFamiliesRequest{
		RegionID:     &region_id,
		AzName:       &az_name,
		FlavorFamily: &flavor_family,
		PageNo:       &page_no,
		PageSize:     &page_size,
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

	region_id := "bb9fdb42056f11eda1610242ac110002"
	response, err := apis.EcsAvailabilityZonesDetailsApi.Do(context.Background(), credential, &ctecs.EcsAvailabilityZonesDetailsRequest{
		RegionID: &region_id,
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

	region_id := "bb9fdb42056f11eda1610242ac110002"
	az_name := "cn-huadong1-jsnj1A-public-ctcloud"

	response, err := apis.EcsFlavorFamiliesListApi.Do(context.Background(), credential, &ctecs.EcsFlavorFamiliesListRequest{
		RegionID: &region_id,
		AzName:   &az_name,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsUpdateInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	disply_name := "java-sdk-ecm-1930"
	response, err := apis.EcsUpdateInstanceApi.Do(context.Background(), credential, &ctecs.EcsUpdateInstanceRequest{
		RegionID:    &region_id,
		InstanceID:  &instance_id,
		DisplayName: &disply_name,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsVncDetails(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	response, err := apis.EcsVncDetailsApi.Do(context.Background(), credential, &ctecs.EcsVncDetailsRequest{
		RegionID:   &region_id,
		InstanceID: &instance_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsVolumeStatistics(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	project_id := "0"
	response, err := apis.EcsVolumeStatisticsApi.Do(context.Background(), credential, &ctecs.EcsVolumeStatisticsRequest{
		RegionID:  &region_id,
		ProjectID: &project_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBatchResetPassword(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	new_password := "hello!world1"
	UpdatePwdInfo := make([]ctecs.EcsBatchResetPasswordUpdatePwdInfoRequest, 0)
	updatePwdInfo := ctecs.EcsBatchResetPasswordUpdatePwdInfoRequest{
		InstanceID:  &instance_id,
		NewPassword: &new_password,
	}

	UpdatePwdInfo = append(UpdatePwdInfo, updatePwdInfo)

	response, err := apis.EcsBatchResetPasswordApi.Do(context.Background(), credential, &ctecs.EcsBatchResetPasswordRequest{
		RegionID:      &region_id,
		UpdatePwdInfo: UpdatePwdInfo,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupPolicyListApi(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	policy_id := "3e251bce0d1411efb0a10242ac110002"
	policy_name := "policy_test_0508"
	project_id := "default"
	page_no := 1
	page_size := 10
	response, err := apis.EcsBackupPolicyListApi.Do(context.Background(), credential, &ctecs.EcsBackupPolicyListRequest{
		RegionID:   &region_id,
		PolicyID:   &policy_id,
		PolicyName: &policy_name,
		ProjectID:  &project_id,
		PageNo:     &page_no,
		PageSize:   &page_size,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupPolicyListInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	policy_id := "3e251bce0d1411efb0a10242ac110002"
	instance_name := ""
	page_no := 1
	page_size := 10
	apis := ctecs.NewApis(client)
	response, err := apis.EcsBackupPolicyListInstancesApi.Do(context.Background(), credential, &ctecs.EcsBackupPolicyListInstancesRequest{
		RegionID:     &region_id,
		PolicyID:     &policy_id,
		InstanceName: &instance_name,
		PageNo:       &page_no,
		PageSize:     &page_size,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupPolicyBindInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	policy_id := "3e251bce0d1411efb0a10242ac110002"
	instanceid_list := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	response, err := apis.EcsBackupPolicyBindInstancesApi.Do(context.Background(), credential, &ctecs.EcsBackupPolicyBindInstancesRequest{
		RegionID:       &region_id,
		PolicyID:       &policy_id,
		InstanceIDList: &instanceid_list,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupPolicyUnbindInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	policy_id := "3e251bce0d1411efb0a10242ac110002"
	instanceid_list := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	response, err := apis.EcsBackupPolicyUnbindInstancesApi.Do(context.Background(), credential, &ctecs.EcsBackupPolicyUnbindInstancesRequest{
		RegionID:       &region_id,
		PolicyID:       &policy_id,
		InstanceIDList: &instanceid_list,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupPolicyBindRepo(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	repository_id := "64f4675e-9a67-45bf-a283-3bcc6b518d2c"
	policy_id := "3e251bce0d1411efb0a10242ac110002"
	response, err := apis.EcsBackupPolicyBindRepoApi.Do(context.Background(), credential, &ctecs.EcsBackupPolicyBindRepoRequest{
		RegionID:     &region_id,
		RepositoryID: &repository_id,
		PolicyID:     &policy_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupPolicyUnbindRepo(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	policy_id := "3e251bce0d1411efb0a10242ac110002"
	response, err := apis.EcsBackupPolicyUnbindRepoApi.Do(context.Background(), credential, &ctecs.EcsBackupPolicyUnbindRepoRequest{
		RegionID: &region_id,
		PolicyID: &policy_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsCreateSnapshot(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	snapshot_name := "test-go-ss"
	response, err := apis.EcsCreateSnapshotApi.Do(context.Background(), credential, &ctecs.EcsCreateSnapshotRequest{
		RegionID:     &region_id,
		InstanceID:   &instance_id,
		SnapshotName: &snapshot_name,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsRestoreSnapshot(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshot_id := "d0e1ec51-a438-03a4-3249-0c083e7a14d6"
	response, err := apis.EcsRestoreSnapshotApi.Do(context.Background(), credential, &ctecs.EcsRestoreSnapshotRequest{
		RegionID:   &region_id,
		SnapshotID: &snapshot_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsUpdateSnapshot(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshot_id := "d0e1ec51-a438-03a4-3249-0c083e7a14d6"
	snapshot_name := "test-go-ss1"
	snaoshot_description := "snapshot_des1"
	response, err := apis.EcsUpdateSnapshotApi.Do(context.Background(), credential, &ctecs.EcsUpdateSnapshotRequest{
		RegionID:            &region_id,
		SnapshotID:          &snapshot_id,
		SnapshotName:        &snapshot_name,
		SnapshotDescription: &snaoshot_description,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotBatchUpdate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshot_id := "d0e1ec51-a438-03a4-3249-0c083e7a14d6"
	snapshot_name := "test-go-ss2"
	snaoshot_description := "snapshot_des2"
	UpdateInfo := make([]ctecs.EcsSnapshotBatchUpdateUpdateInfoRequest, 0)
	updateInfo := ctecs.EcsSnapshotBatchUpdateUpdateInfoRequest{
		SnapshotID:          &snapshot_id,
		SnapshotName:        &snapshot_name,
		SnapshotDescription: &snaoshot_description,
	}

	UpdateInfo = append(UpdateInfo, updateInfo)

	response, err := apis.EcsSnapshotBatchUpdateApi.Do(context.Background(), credential, &ctecs.EcsSnapshotBatchUpdateRequest{
		RegionID:   &region_id,
		UpdateInfo: UpdateInfo,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	project_id := ""
	instance_id := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	snapshot_status := "available"
	snapshot_id := "d0e1ec51-a438-03a4-3249-0c083e7a14d6"
	query_content := ""
	snapshot_name := "test-go-ss"
	page_no := 1
	page_size := 10
	response, err := apis.EcsSnapshotListApi.Do(context.Background(), credential, &ctecs.EcsSnapshotListRequest{
		RegionID:       &region_id,
		ProjectID:      &project_id,
		PageNo:         &page_no,
		PageSize:       &page_size,
		InstanceID:     &instance_id,
		SnapshotStatus: &snapshot_status,
		SnapshotID:     &snapshot_id,
		QueryContent:   &query_content,
		SnapshotName:   &snapshot_name,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotDetails(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshot_id := "d0e1ec51-a438-03a4-3249-0c083e7a14d6"
	response, err := apis.EcsSnapshotDetailsApi.Do(context.Background(), credential, &ctecs.EcsSnapshotDetailsRequest{
		RegionID:   &region_id,
		SnapshotID: &snapshot_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotStatistics(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instanceid_list := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	response, err := apis.EcsSnapshotStatisticsApi.Do(context.Background(), credential, &ctecs.EcsSnapshotStatisticsRequest{
		RegionID:       &region_id,
		InstanceIDList: &instanceid_list,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotStatus(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshot_id := "d0e1ec51-a438-03a4-3249-0c083e7a14d6"
	response, err := apis.EcsSnapshotStatusApi.Do(context.Background(), credential, &ctecs.EcsSnapshotStatusRequest{
		RegionID:   &region_id,
		SnapshotID: &snapshot_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotDelete(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshot_id := "d0e1ec51-a438-03a4-3249-0c083e7a14d6"
	response, err := apis.EcsSnapshotDeleteApi.Do(context.Background(), credential, &ctecs.EcsSnapshotDeleteRequest{
		RegionID:   &region_id,
		SnapshotID: &snapshot_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	device_id_list := []string{"0fec78e4-1889-803f-b2a7-515c1c40b788"}
	page_no := 1
	page := 1
	page_size := 10
	response, err := apis.EcsVmDiskLatestMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmDiskLatestMetricDataRequest{
		RegionID:     &region_id,
		DeviceIDList: &device_id_list,
		PageNo:       &page_no,
		Page:         &page,
		PageSize:     &page_size,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	device_id_list := []string{"0fec78e4-1889-803f-b2a7-515c1c40b788"}
	page_no := 1
	page := 1
	page_size := 10
	response, err := apis.EcsVmCpuLatestMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmCpuLatestMetricDataRequest{
		RegionID:     &region_id,
		DeviceIDList: &device_id_list,
		PageNo:       &page_no,
		Page:         &page,
		PageSize:     &page_size,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	device_id_list := []string{"0fec78e4-1889-803f-b2a7-515c1c40b788"}
	period := 14400
	start_time := "1717402682"
	end_time := ""
	page_no := 1
	page := 1
	page_size := 10
	response, err := apis.EcsVmNetworkHistoryMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmNetworkHistoryMetricDataRequest{
		RegionID:     &region_id,
		DeviceIDList: &device_id_list,
		Period:       &period,
		StartTime:    &start_time,
		EndTime:      &end_time,
		PageNo:       &page_no,
		Page:         &page,
		PageSize:     &page_size,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	device_id_list := []string{"0fec78e4-1889-803f-b2a7-515c1c40b788"}
	period := 14400
	start_time := "1717402682"
	end_time := "1717575482"
	page_no := 1
	page := 1
	page_size := 10
	response, err := apis.EcsVmMemHistoryMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmMemHistoryMetricDataRequest{
		RegionID:     &region_id,
		DeviceIDList: &device_id_list,
		Period:       &period,
		StartTime:    &start_time,
		EndTime:      &end_time,
		PageNo:       &page_no,
		Page:         &page,
		PageSize:     &page_size,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	device_id_list := []string{"e2a016ec-543b-08f1-38c4-13d9dac55b5a"}
	period := 14400
	start_time := "1717402682"
	end_time := "1717575482"
	page_no := 1
	page := 1
	page_size := 10
	response, err := apis.EcsVmDiskHistoryMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmDiskHistoryMetricDataRequest{
		RegionID:     &region_id,
		DeviceIDList: &device_id_list,
		Period:       &period,
		StartTime:    &start_time,
		EndTime:      &end_time,
		PageNo:       &page_no,
		Page:         &page,
		PageSize:     &page_size,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	device_id_list := []string{"e2a016ec-543b-08f1-38c4-13d9dac55b5a"}
	period := 14400
	start_time := "1717402682"
	end_time := "1717575482"
	page_no := 1
	page := 1
	page_size := 10
	response, err := apis.EcsVmCpuHistoryMetricDataApi.Do(context.Background(), credential, &ctecs.EcsVmCpuHistoryMetricDataRequest{
		RegionID:     &region_id,
		DeviceIDList: &device_id_list,
		Period:       &period,
		StartTime:    &start_time,
		EndTime:      &end_time,
		PageNo:       &page_no,
		Page:         &page,
		PageSize:     &page_size,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "63afb617-b8f5-d482-9ecd-6d8bb9124d4e"
	instance_backup_name := "backup-061103"
	instance_backup_dec := "creat_test01"
	repo_id := "64f4675e-9a67-45bf-a283-3bcc6b518d2c"
	response, err := apis.EcsBackupCreateApi.Do(context.Background(), credential, &ctecs.EcsBackupCreateRequest{
		RegionID:                  &region_id,
		InstanceID:                &instance_id,
		InstanceBackupName:        &instance_backup_name,
		InstanceBackupDescription: &instance_backup_dec,
		RepositoryID:              &repo_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_backup_id := "cebc0c31-6dcc-d5d6-b820-3afa708627ab"
	instance_backup_name := "update-test-sdk-02"
	instance_backup_dec := "api_update_test01"
	response, err := apis.EcsBackupUpdateApi.Do(context.Background(), credential, &ctecs.EcsBackupUpdateRequest{
		RegionID:                  &region_id,
		InstanceBackupID:          &instance_backup_id,
		InstanceBackupName:        &instance_backup_name,
		InstanceBackupDescription: &instance_backup_dec,
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
	instance_backup_id := "cebc0c31-6dcc-d5d6-b820-3afa708627ab"
	instance_backup_name := "update-test-sdk-03"
	instance_backup_desc := "api_update_test03"
	UpdateInfo := make([]ctecs.EcsBackupBatchUpdateUpdateInfoRequest, 0)
	updateInfo := ctecs.EcsBackupBatchUpdateUpdateInfoRequest{
		InstanceBackupID:          &instance_backup_id,
		InstanceBackupName:        &instance_backup_name,
		InstanceBackupDescription: &instance_backup_desc,
	}

	UpdateInfo = append(UpdateInfo, updateInfo)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	response, err := apis.EcsBackupBatchUpdateApi.Do(context.Background(), credential, &ctecs.EcsBackupBatchUpdateRequest{
		RegionID:   &region_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_backup_id := "f6b50900-d077-fca0-62fc-a15cd9255100"
	response, err := apis.EcsBackupUsageApi.Do(context.Background(), credential, &ctecs.EcsBackupUsageRequest{
		RegionID:         &region_id,
		InstanceBackupID: &instance_backup_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "63afb617-b8f5-d482-9ecd-6d8bb9124d4e"
	response, err := apis.EcsBackupStatisticsApi.Do(context.Background(), credential, &ctecs.EcsBackupStatisticsRequest{
		RegionID:   &region_id,
		InstanceID: &instance_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	response, err := apis.EcsBackupInstanceResourceApi.Do(context.Background(), credential, &ctecs.EcsBackupInstanceResourceRequest{
		RegionID: &region_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	page_no := 1
	page_size := 10
	instance_id := ""
	repo_id := ""
	instance_backup_id := ""
	query_content := ""
	instance_backup_status := ""
	project_id := ""
	response, err := apis.EcsBackupListApi.Do(context.Background(), credential, &ctecs.EcsBackupListRequest{
		RegionID:             &region_id,
		PageNo:               &page_no,
		PageSize:             &page_size,
		InstanceID:           &instance_id,
		RepositoryID:         &repo_id,
		InstanceBackupID:     &instance_backup_id,
		QueryContent:         &query_content,
		InstanceBackupStatus: &instance_backup_status,
		ProjectID:            &project_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_backup_id := "f6b50900-d077-fca0-62fc-a15cd9255100"
	response, err := apis.EcsBackupDetailsApi.Do(context.Background(), credential, &ctecs.EcsBackupDetailsRequest{
		RegionID:         &region_id,
		InstanceBackupID: &instance_backup_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "63afb617-b8f5-d482-9ecd-6d8bb9124d4e"
	response, err := apis.EcsBackupInstanceDetailsApi.Do(context.Background(), credential, &ctecs.EcsBackupInstanceDetailsRequest{
		RegionID:   &region_id,
		InstanceID: &instance_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_backup_id := "f6b50900-d077-fca0-62fc-a15cd9255100"
	response, err := apis.EcsBackupStatusApi.Do(context.Background(), credential, &ctecs.EcsBackupStatusRequest{
		RegionID:         &region_id,
		InstanceBackupID: &instance_backup_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_backup_id := "cebc0c31-6dcc-d5d6-b820-3afa708627ab"
	response, err := apis.EcsBackupRestoreApi.Do(context.Background(), credential, &ctecs.EcsBackupRestoreRequest{
		RegionID:         &region_id,
		InstanceBackupID: &instance_backup_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_backup_id := "cebc0c31-6dcc-d5d6-b820-3afa708627ab"
	response, err := apis.EcsBackupDeleteApi.Do(context.Background(), credential, &ctecs.EcsBackupDeleteRequest{
		RegionID:         &region_id,
		InstanceBackupID: &instance_backup_id,
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
	instance_id := ""
	system_type := ""
	system_arch := ""
	system_version := ""
	actionInfo := ctecs.EcsAgentBatchActionActionInfoRequest{
		InstanceID:    &instance_id,
		SystemType:    &system_type,
		SystemArch:    &system_arch,
		SystemVersion: &system_version,
	}

	ActionInfo = append(ActionInfo, actionInfo)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	action := "update"
	response, err := apis.EcsAgentBatchActionApi.Do(context.Background(), credential, &ctecs.EcsAgentBatchActionRequest{
		RegionID:   &region_id,
		Action:     &action,
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
	client_token := "ports_create061801"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	subnet_id := "subnet-4c4333pc67"
	primary_private_ip := "172.16.0.141"
	ipv6_addresses := []string{"240e:978:497c:ec00:cd74:fd9d:c45d:4131"}
	security_group_ids := []string{"sg-n7nu88xfbq"}
	secondary_private_ip_count := 1
	secondary_private_ips := []string{"172.16.0.210"}
	name := "nic-test01"
	description := "dec-test"
	response, err := apis.EcsPortsCreateApi.Do(context.Background(), credential, &ctecs.EcsPortsCreateRequest{
		ClientToken:             &client_token,
		RegionID:                &region_id,
		SubnetID:                &subnet_id,
		PrimaryPrivateIp:        &primary_private_ip,
		Ipv6Addresses:           &ipv6_addresses,
		SecurityGroupIds:        &security_group_ids,
		SecondaryPrivateIpCount: &secondary_private_ip_count,
		SecondaryPrivateIps:     &secondary_private_ips,
		Name:                    &name,
		Description:             &description,
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
	client_token := "delete-ports-test-01"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	network_interface_id := "port-1jtj14yppt"
	response, err := apis.EcsPortsDeleteApi.Do(context.Background(), credential, &ctecs.EcsPortsDeleteRequest{
		ClientToken:        &client_token,
		RegionID:           &region_id,
		NetworkInterfaceID: &network_interface_id,
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
	client_token := "update-port-test-01"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	network_interface_id := "port-webh3q806q"
	name := "nic-update-name"
	description := "nic_update_description"
	security_group_ids := []string{"sg-tdzefke02r"}
	response, err := apis.EcsPortsUpdateApi.Do(context.Background(), credential, &ctecs.EcsPortsUpdateRequest{
		ClientToken:        &client_token,
		RegionID:           &region_id,
		NetworkInterfaceID: &network_interface_id,
		Name:               &name,
		Description:        &description,
		SecurityGroupIDs:   &security_group_ids,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	vpc_id := ""
	device_id := ""
	subnet_id := ""
	page_number := 1
	page_size := 1
	page_no := 1
	response, err := apis.EcsPortsListApi.Do(context.Background(), credential, &ctecs.EcsPortsListRequest{
		RegionID:   &region_id,
		VpcID:      &vpc_id,
		DeviceID:   &device_id,
		SubnetID:   &subnet_id,
		PageNumber: &page_number,
		PageSize:   &page_size,
		PageNo:     &page_no,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	network_interface_id := "port-webh3q806q"
	response, err := apis.EcsPortsShowApi.Do(context.Background(), credential, &ctecs.EcsPortsShowRequest{
		RegionID:           &region_id,
		NetworkInterfaceID: &network_interface_id,
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
	client_token := "attach-port-test-01"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	az_name := "cn-huadong1-jsnj1A-public-ctcloud"
	project_id := "0"
	network_interface_id := "port-55ty7j56xi"
	instance_id := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	instance_type := 3
	response, err := apis.EcsPortsAttachApi.Do(context.Background(), credential, &ctecs.EcsPortsAttachRequest{
		ClientToken:        &client_token,
		RegionID:           &region_id,
		AzName:             &az_name,
		ProjectID:          &project_id,
		NetworkInterfaceID: &network_interface_id,
		InstanceID:         &instance_id,
		InstanceType:       &instance_type,
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
	client_token := "detach-port-test-01"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	network_interface_id := "port-55ty7j56xi"
	instance_id := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	response, err := apis.EcsPortsDetachApi.Do(context.Background(), credential, &ctecs.EcsPortsDetachRequest{
		ClientToken:        &client_token,
		RegionID:           &region_id,
		NetworkInterfaceID: &network_interface_id,
		InstanceID:         &instance_id,
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
	client_token := "assign-ipv6-01"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	network_interface_id := "port-webh3q806q"
	ipv6_addresses_count := 1
	response, err := apis.EcsPortsAssignIpv6Api.Do(context.Background(), credential, &ctecs.EcsPortsAssignIpv6Request{
		ClientToken:        &client_token,
		RegionID:           &region_id,
		NetworkInterfaceID: &network_interface_id,
		Ipv6AddressesCount: &ipv6_addresses_count,
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
	client_token := "unassign-ipv6-01"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	network_interface_id := "port-webh3q806q"
	ipv6_addresses := []string{"240e:978:49f5:3100:fb1e:6fed:927:e1e1"}
	response, err := apis.EcsPortsUnassignIpv6Api.Do(context.Background(), credential, &ctecs.EcsPortsUnassignIpv6Request{
		ClientToken:        &client_token,
		RegionID:           &region_id,
		NetworkInterfaceID: &network_interface_id,
		Ipv6Addresses:      &ipv6_addresses,
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
	client_token := "assign-secondary-private-ips-01"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	network_interface_id := "port-webh3q806q"
	secondary_private_ip_count := 1
	response, err := apis.EcsPortsAssignSecondaryPrivateIpsApi.Do(context.Background(), credential, &ctecs.EcsPortsAssignSecondaryPrivateIpsRequest{
		ClientToken:             &client_token,
		RegionID:                &region_id,
		NetworkInterfaceID:      &network_interface_id,
		SecondaryPrivateIpCount: &secondary_private_ip_count,
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
	client_token := "unassign-secondary-private-ips-01"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	network_interface_id := "port-webh3q806q"
	secondary_private_ips := []string{"192.168.0.3"}
	response, err := apis.EcsPortsUnassignSecondaryPrivateIpsApi.Do(context.Background(), credential, &ctecs.EcsPortsUnassignSecondaryPrivateIpsRequest{
		ClientToken:         &client_token,
		RegionID:            &region_id,
		NetworkInterfaceID:  &network_interface_id,
		SecondaryPrivateIps: &secondary_private_ips,
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
	client_token := "create-eip-test"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	project_id := "0"
	cycle_type := "month"
	cycle_count := 1
	name := "eip-name"
	band_width := 5
	band_width_id := "bandwidth-7hzv449r2j"
	demand_billing_type := "bandwidth"
	response, err := apis.EcsEipCreateApi.Do(context.Background(), credential, &ctecs.EcsEipCreateRequest{
		ClientToken:       &client_token,
		RegionID:          &region_id,
		ProjectID:         &project_id,
		CycleType:         &cycle_type,
		CycleCount:        &cycle_count,
		Name:              &name,
		Bandwidth:         &band_width,
		BandwidthID:       &band_width_id,
		DemandBillingType: &demand_billing_type,
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
	client_token := "delete-eip-test"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	project_id := "0"
	eip_id := "eip-dskebprsxl"
	response, err := apis.EcsEipDeleteApi.Do(context.Background(), credential, &ctecs.EcsEipDeleteRequest{
		ClientToken: &client_token,
		RegionID:    &region_id,
		ProjectID:   &project_id,
		EipID:       &eip_id,
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
	region_id := "bb9fdb42056f11eda1610242ac110002"
	instance_id := "63afb617-b8f5-d482-9ecd-6d8bb9124d4e"
	subnet_id := "subnet-3o8uvvp6h4"
	response, err := apis.EcsShareInterfaceAttachApi.Do(context.Background(), credential, &ctecs.EcsShareInterfaceAttachRequest{
		RegionID:   &region_id,
		InstanceID: &instance_id,
		SubnetID:   &subnet_id,
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
	nic_name := "nic-test-061901"
	nic_is_master := true
	subnet_id := "subnet-4c4333pc67"
	network_card := ctecs.EcsBackupCreateInstanceNetworkCardListRequest{
		NicName:  &nic_name,
		IsMaster: &nic_is_master,
		SubnetID: &subnet_id,
	}
	networkCardList = append(networkCardList, network_card)
	labelList := make([]ctecs.EcsBackupCreateInstanceLabelListRequest, 0)
	label_key := "label-key-test"
	label_value := "label-value-test"
	label := ctecs.EcsBackupCreateInstanceLabelListRequest{
		LabelKey:   &label_key,
		LabelValue: &label_value,
	}
	labelList = append(labelList, label)
	client_token := "ecs-backup-create-instance-test-061901"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	az_name := "cn-huadong1-jsnj1A-public-ctcloud"
	instance_name := "ecm-go-test-061901"
	display_name := "ecm-go-test-061901"
	instance_backup_id := "e718f4b2-0ff2-e486-b322-ea206fbce240"
	flavor_id := "34e1b6f6-e974-1575-20b2-172ba0e0bf83"
	vpc_id := "vpc-chz0ilszsp"
	on_demand := false
	ext_ip := "1"
	ip_version := ""
	band_width := 50
	project_id := ""
	ipv6_address_id := ""
	sec_group_list := []string{"sg-ku5edgbitc"}
	eip_id := ""
	affinity_group_id := ""
	keypair_id := ""
	user_password := ""
	cycle_count := 1
	cycle_type := "MONTH"
	auto_renew_status := 0
	user_data := "YmF0Y2hDcmVhdGVUZXN0MDgwMw=="
	monitor_service := true
	pay_voucher_price := 1819.50
	response, err := apis.EcsBackupCreateInstanceApi.Do(context.Background(), credential, &ctecs.EcsBackupCreateInstanceRequest{
		ClientToken:      &client_token,
		RegionID:         &region_id,
		AzName:           &az_name,
		InstanceName:     &instance_name,
		DisplayName:      &display_name,
		InstanceBackupID: &instance_backup_id,
		FlavorID:         &flavor_id,
		VpcID:            &vpc_id,
		OnDemand:         &on_demand,
		NetworkCardList:  networkCardList,
		ExtIP:            &ext_ip,
		IpVersion:        &ip_version,
		Bandwidth:        &band_width,
		ProjectID:        &project_id,
		Ipv6AddressID:    &ipv6_address_id,
		SecGroupList:     &sec_group_list,
		EipID:            &eip_id,
		AffinityGroupID:  &affinity_group_id,
		KeyPairID:        &keypair_id,
		UserPassword:     &user_password,
		CycleCount:       &cycle_count,
		CycleType:        &cycle_type,
		AutoRenewStatus:  &auto_renew_status,
		UserData:         &user_data,
		PayVoucherPrice:  &pay_voucher_price,
		LabelList:        labelList,
		MonitorService:   &monitor_service,
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}
	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))

}

func ecsAffinityGroupCreateExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "bb9fdb42056f11eda1610242ac110002"
	affinity_group_name := "agtest-01"
	policy_type := 2

	response, err := apis.EcsAffinityGroupCreateApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupCreateRequest{
		RegionID:          &region_id,
		AffinityGroupName: &affinity_group_name,
		PolicyType:        &policy_type,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupBindInstanceExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	instance_id := "74d7447a-1275-ec32-cede-8fd285e43f76"
	affinity_group_id := "f7b9c874-89b5-4fd6-977c-3e08b29b5872"
	response, err := apis.EcsAffinityGroupBindInstanceApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupBindInstanceRequest{
		RegionID:        &region_id,
		InstanceID:      &instance_id,
		AffinityGroupID: &affinity_group_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupUnbindInstanceExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	instance_id := "74d7447a-1275-ec32-cede-8fd285e43f76"
	affinity_group_id := "f7b9c874-89b5-4fd6-977c-3e08b29b5872"
	response, err := apis.EcsAffinityGroupUnbindInstanceApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupUnbindInstanceRequest{
		RegionID:        &region_id,
		InstanceID:      &instance_id,
		AffinityGroupID: &affinity_group_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupUnbindInstancesExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	instance_ids := "74d7447a-1275-ec32-cede-8fd285e43f76"
	affinity_group_id := "f7b9c874-89b5-4fd6-977c-3e08b29b5872"
	response, err := apis.EcsAffinityGroupUnbindInstancesApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupUnbindInstancesRequest{
		RegionID:        &region_id,
		InstanceIDs:     &instance_ids,
		AffinityGroupID: &affinity_group_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupListExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	affinity_group_id := "f7b9c874-89b5-4fd6-977c-3e08b29b5872"
	query_content := "ag"
	page_no := 1
	page_size := 10
	response, err := apis.EcsAffinityGroupListApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupListRequest{
		RegionID:        &region_id,
		AffinityGroupID: &affinity_group_id,
		QueryContent:    &query_content,
		PageNo:          &page_no,
		PageSize:        &page_size,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupListInstanceExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	affinity_group_id := "f7b9c874-89b5-4fd6-977c-3e08b29b5872"
	page_no := 1
	page_size := 10

	response, err := apis.EcsAffinityGroupListInstanceApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupListInstanceRequest{
		RegionID:        &region_id,
		AffinityGroupID: &affinity_group_id,
		PageNo:          &page_no,
		PageSize:        &page_size,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupUpdateExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	affinity_group_id := "f7b9c874-89b5-4fd6-977c-3e08b29b5872"
	affinity_group_name := "update-02"

	response, err := apis.EcsAffinityGroupUpdateApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupUpdateRequest{
		RegionID:          &region_id,
		AffinityGroupID:   &affinity_group_id,
		AffinityGroupName: &affinity_group_name,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupBindInstanceCheckExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	instance_id := "74d7447a-1275-ec32-cede-8fd285e43f76"
	affinity_group_id := "f7b9c874-89b5-4fd6-977c-3e08b29b5872"

	response, err := apis.EcsAffinityGroupBindInstanceCheckApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupBindInstanceCheckRequest{
		RegionID:        &region_id,
		InstanceID:      &instance_id,
		AffinityGroupID: &affinity_group_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupDeleteExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	affinity_group_id := "f7b9c874-89b5-4fd6-977c-3e08b29b5872"

	response, err := apis.EcsAffinityGroupDeleteApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupDeleteRequest{
		RegionID:        &region_id,
		AffinityGroupID: &affinity_group_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsAffinityGroupDetailsExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	instance_id := "74d7447a-1275-ec32-cede-8fd285e43f76"

	response, err := apis.EcsAffinityGroupDetailsApi.Do(context.Background(), credential, &ctecs.EcsAffinityGroupDetailsRequest{
		RegionID:   &region_id,
		InstanceID: &instance_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsBackupRepoCreateExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	project_id := "0"
	repo_name := "go-sdk-test"
	cycle_count := 6
	cycle_type := "MONTH"
	client_token := "create_repo_test_0001"
	size := 100
	auto_renew_status := 0
	pay_voucher_price := 100.01

	response, err := apis.EcsBackupRepoCreateApi.Do(context.Background(), credential, &ctecs.EcsBackupRepoCreateRequest{
		RegionID:        &region_id,
		ProjectID:       &project_id,
		RepositoryName:  &repo_name,
		CycleCount:      &cycle_count,
		CycleType:       &cycle_type,
		ClientToken:     &client_token,
		Size:            &size,
		AutoRenewStatus: &auto_renew_status,
		PayVoucherPrice: &pay_voucher_price,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsBackupRepoRenewExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	repo_id := "c6a56582-8dfb-41be-9223-f12bc19f2685"
	cycle_count := 1
	cycle_type := "MONTH"
	client_token := "renew-repo-test-0001"
	pay_voucher_price := 20.00

	response, err := apis.EcsBackupRepoRenewApi.Do(context.Background(), credential, &ctecs.EcsBackupRepoRenewRequest{
		RegionID:        &region_id,
		RepositoryID:    &repo_id,
		CycleCount:      &cycle_count,
		CycleType:       &cycle_type,
		ClientToken:     &client_token,
		PayVoucherPrice: &pay_voucher_price,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsBackupRepoUpgradeExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	repo_id := "c6a56582-8dfb-41be-9223-f12bc19f2685"
	client_token := "upgrade-repo-test-0001"
	size := 150
	pay_voucher_price := 40.55

	response, err := apis.EcsBackupRepoUpgradeApi.Do(context.Background(), credential, &ctecs.EcsBackupRepoUpgradeRequest{
		RegionID:        &region_id,
		RepositoryID:    &repo_id,
		ClientToken:     &client_token,
		Size:            &size,
		PayVoucherPrice: &pay_voucher_price,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsBackupRepoListExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	project_id := ""
	repo_name := ""
	repo_id := ""
	status := ""
	page_no := 1
	page_size := 5

	response, err := apis.EcsBackupRepoListApi.Do(context.Background(), credential, &ctecs.EcsBackupRepoListRequest{
		RegionID:       &region_id,
		ProjectID:      &project_id,
		RepositoryName: &repo_name,
		RepositoryID:   &repo_id,
		Status:         &status,
		PageNo:         &page_no,
		PageSize:       &page_size,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsBackupRepoDeleteExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	repo_id := "c6a56582-8dfb-41be-9223-f12bc19f2685"
	client_token := "delete-repo-test-0001"

	response, err := apis.EcsBackupRepoDeleteApi.Do(context.Background(), credential, &ctecs.EcsBackupRepoDeleteRequest{
		RegionID:     &region_id,
		RepositoryID: &repo_id,
		ClientToken:  &client_token,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVpcCreateSecurityGroupExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	client_token := "create-sg-test-01"
	region_id := "81f7728662dd11ec810800155d307d5b"
	project_id := "0"
	vpc_id := "vpc-riwxr5wpju"
	name := "go-sdk-test"
	description := "创建安全组测试"

	response, err := apis.EcsVpcCreateSecurityGroupApi.Do(context.Background(), credential, &ctecs.EcsVpcCreateSecurityGroupRequest{
		ClientToken: &client_token,
		RegionID:    &region_id,
		ProjectID:   &project_id,
		VpcID:       &vpc_id,
		Name:        &name,
		Description: &description,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVpcQuerySecurityGroupsExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	vpc_id := "vpc-riwxr5wpju"
	query_content := "test"
	instance_id := ""
	project_id := "0"
	page_no := 1
	page_size := 10

	response, err := apis.EcsVpcQuerySecurityGroupsApi.Do(context.Background(), credential, &ctecs.EcsVpcQuerySecurityGroupsRequest{
		RegionID:     &region_id,
		VpcID:        &vpc_id,
		QueryContent: &query_content,
		ProjectID:    &project_id,
		InstanceID:   &instance_id,
		PageNumber:   &page_no,
		PageNo:       &page_no,
		PageSize:     &page_size,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVpcDescribeSecurityGroupAttributeExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	sg_id := "sg-rq5u5gzpgx"
	project_id := "0"
	direction := "all"

	response, err := apis.EcsVpcDescribeSecurityGroupAttributeApi.Do(context.Background(), credential, &ctecs.EcsVpcDescribeSecurityGroupAttributeRequest{
		RegionID:        &region_id,
		SecurityGroupID: &sg_id,
		ProjectID:       &project_id,
		Direction:       &direction,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVpcDeleteSecurityGroupExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	client_token := "del-sg-test-01"
	region_id := "81f7728662dd11ec810800155d307d5b"
	project_id := "0"
	sg_id := "sg-rq5u5gzpgx"

	response, err := apis.EcsVpcDeleteSecurityGroupApi.Do(context.Background(), credential, &ctecs.EcsVpcDeleteSecurityGroupRequest{
		ClientToken:     &client_token,
		RegionID:        &region_id,
		ProjectID:       &project_id,
		SecurityGroupID: &sg_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVolumeCreateExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	disk_mode := "VBD"
	disk_type := "SATA"
	disk_name := "sdk-test"
	disk_size := 10
	client_token := "create_volume_test_001"
	az_name := "az1"
	multi_attach := false
	on_demand := true
	cycle_type := "YEAR"
	cycle_count := 2
	is_encrypt := false
	kms_uuid := ""
	project_id := "0"
	//image_id := ""

	response, err := apis.EcsVolumeCreateApi.Do(context.Background(), credential, &ctecs.EcsVolumeCreateRequest{
		RegionID:    &region_id,
		DiskMode:    &disk_mode,
		DiskType:    &disk_type,
		DiskName:    &disk_name,
		DiskSize:    &disk_size,
		ClientToken: &client_token,
		AzName:      &az_name,
		MultiAttach: &multi_attach,
		OnDemand:    &on_demand,
		CycleType:   &cycle_type,
		CycleCount:  &cycle_count,
		IsEncrypt:   &is_encrypt,
		KmsUUID:     &kms_uuid,
		ProjectID:   &project_id,
		//ImageID:     &image_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVolumeUpdateExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	disk_name := "sdk-test02"
	disk_id := "d9800a0e-b21b-4817-9b77-6283472c63c4"
	response, err := apis.EcsVolumeUpdateApi.Do(context.Background(), credential, &ctecs.EcsVolumeUpdateRequest{
		RegionID: &region_id,
		DiskName: &disk_name,
		DiskID:   &disk_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVolumeExtendExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	disk_id := "d9800a0e-b21b-4817-9b77-6283472c63c4"
	disk_size := 20
	client_token := "resize-volume-test-001"

	response, err := apis.EcsVolumeExtendApi.Do(context.Background(), credential, &ctecs.EcsVolumeExtendRequest{
		DiskSize:    &disk_size,
		DiskID:      &disk_id,
		RegionID:    &region_id,
		ClientToken: &client_token,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVolumeShowOldExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	resource_id := "1ca3987ee7404fd2a076d9968ef7d5ac"
	region_id := "81f7728662dd11ec810800155d307d5b"

	response, err := apis.EcsVolumeShowOldApi.Do(context.Background(), credential, &ctecs.EcsVolumeShowOldRequest{
		ResourceID: &resource_id,
		RegionID:   &region_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVolumeShowExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	disk_id := "d9800a0e-b21b-4817-9b77-6283472c63c4"

	response, err := apis.EcsVolumeShowApi.Do(context.Background(), credential, &ctecs.EcsVolumeShowRequest{
		DiskID:   &disk_id,
		RegionID: &region_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVolumeAttachExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	disk_id := "d9800a0e-b21b-4817-9b77-6283472c63c4"
	instance_id := "dba4eaed-fc14-607b-495c-347922ac96fe"

	response, err := apis.EcsVolumeAttachApi.Do(context.Background(), credential, &ctecs.EcsVolumeAttachRequest{
		DiskID:     &disk_id,
		RegionID:   &region_id,
		InstanceID: &instance_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVolumeDetachExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	disk_id := "d9800a0e-b21b-4817-9b77-6283472c63c4"
	instance_id := "dba4eaed-fc14-607b-495c-347922ac96fe"

	response, err := apis.EcsVolumeDetachApi.Do(context.Background(), credential, &ctecs.EcsVolumeDetachRequest{
		DiskID:     &disk_id,
		RegionID:   &region_id,
		InstanceID: &instance_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsVolumeDeleteExec(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)

	region_id := "81f7728662dd11ec810800155d307d5b"
	disk_id := "d9800a0e-b21b-4817-9b77-6283472c63c4"
	client_token := "del-vol-test-001"

	response, err := apis.EcsVolumeDeleteApi.Do(context.Background(), credential, &ctecs.EcsVolumeDeleteRequest{
		ClientToken: &client_token,
		DiskID:      &disk_id,
		RegionID:    &region_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyCreate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshot_policy_name := "api-create03"
	snapshot_time := "12,13"
	cycle_type := "day"
	cycle_day := 1
	cycle_week := "0,2,6"
	retention_type := "num"
	retention_day := 2
	retention_num := 3
	snapshot_policy_status := 1

	response, err := apis.EcsSnapshotPolicyCreateApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyCreateRequest{
		RegionID:             &region_id,
		SnapshotPolicyName:   &snapshot_policy_name,
		SnapshotTime:         &snapshot_time,
		CycleType:            &cycle_type,
		CycleDay:             &cycle_day,
		CycleWeek:            &cycle_week,
		RetentionType:        &retention_type,
		RetentionDay:         &retention_day,
		RetentionNum:         &retention_num,
		SnapshotPolicyStatus: &snapshot_policy_status,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyUpdate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	snapshotPolicy_name := "api-create03"
	snapshot_time := "12,14,15"
	cycle_type := "day"
	cycle_day := 1
	cycle_week := "0,2,6"
	retention_type := "num"
	retention_day := 2
	retention_num := 3
	response, err := apis.EcsSnapshotPolicyUpdateApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyUpdateRequest{
		RegionID:           &region_id,
		SnapshotPolicyID:   &snapshotPolicy_id,
		SnapshotPolicyName: &snapshotPolicy_name,
		SnapshotTime:       &snapshot_time,
		CycleType:          &cycle_type,
		CycleDay:           &cycle_day,
		CycleWeek:          &cycle_week,
		RetentionType:      &retention_type,
		RetentionDay:       &retention_day,
		RetentionNum:       &retention_num,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyEnable(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	response, err := apis.EcsSnapshotPolicyEnableApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyEnableRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyDisable(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	response, err := apis.EcsSnapshotPolicyDisableApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyDisableRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyExecute(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	response, err := apis.EcsSnapshotPolicyExecuteApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyExecuteRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyDetails(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	response, err := apis.EcsSnapshotPolicyDetailsApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyDetailsRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyDelete(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	response, err := apis.EcsSnapshotPolicyDeleteApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyDeleteRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyBindInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	instance_ids := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	response, err := apis.EcsSnapshotPolicyBindInstancesApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyBindInstancesRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
		InstanceIDs:      &instance_ids,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyUnbindInstances(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	instance_ids := "7f315d55-8ead-c470-3811-b9cad2dbecb4"
	response, err := apis.EcsSnapshotPolicyUnbindInstancesApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyUnbindInstancesRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
		InstanceIDs:      &instance_ids,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsVpcCreateSecurityGroupEgress(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	apis := ctecs.NewApis(client)
	direction := "egress"
	action := "accept"
	priority := 100
	protocol := "ANY"
	ether_type := "IPv4"
	destCidr_Ip := "0.0.0.0/0"
	description := "出方向"
	range_rule := "8000-9000"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	securityGroup_id := "sg-3lewm3bv1t"
	client_token := "create-security-group-egress-test082801"
	SecurityGroupRules := make([]ctecs.EcsVpcCreateSecurityGroupEgressSecurityGroupRulesRequest, 0)
	securityGroupRules := ctecs.EcsVpcCreateSecurityGroupEgressSecurityGroupRulesRequest{
		Direction:   &direction,
		Action:      &action,
		Priority:    &priority,
		Protocol:    &protocol,
		Ethertype:   &ether_type,
		DestCidrIp:  &destCidr_Ip,
		Description: &description,
		Range:       &range_rule,
	}

	SecurityGroupRules = append(SecurityGroupRules, securityGroupRules)

	response, err := apis.EcsVpcCreateSecurityGroupEgressApi.Do(context.Background(), credential, &ctecs.EcsVpcCreateSecurityGroupEgressRequest{
		RegionID:           &region_id,
		SecurityGroupID:    &securityGroup_id,
		SecurityGroupRules: SecurityGroupRules,
		ClientToken:        &client_token,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyTaskList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	page_no := 1
	page_size := 10
	response, err := apis.EcsSnapshotPolicyTaskListApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyTaskListRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
		PageNo:           &page_no,
		PageSize:         &page_size,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyInstanceList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	snapshotPolicy_id := "cdd169a6650411efbe670242ac110002"
	page_no := 1
	page_size := 10
	response, err := apis.EcsSnapshotPolicyInstanceListApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyInstanceListRequest{
		RegionID:         &region_id,
		SnapshotPolicyID: &snapshotPolicy_id,
		PageNo:           &page_no,
		PageSize:         &page_size,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}
func ecsSnapshotPolicyList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	region_id := "bb9fdb42056f11eda1610242ac110002"
	page_no := 1
	page_size := 10
	snapshotPolicy_status := 0
	queryContent := "test"
	response, err := apis.EcsSnapshotPolicyListApi.Do(context.Background(), credential, &ctecs.EcsSnapshotPolicyListRequest{
		RegionID:             &region_id,
		PageNo:               &page_no,
		PageSize:             &page_size,
		SnapshotPolicyStatus: &snapshotPolicy_status,
		QueryContent:         &queryContent,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsSnapshotCreateInstance(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	label_key := "test-key"
	label_value := "test-value"
	LabelList := make([]ctecs.EcsSnapshotCreateInstanceLabelListRequest, 0)
	labelList := ctecs.EcsSnapshotCreateInstanceLabelListRequest{
		LabelKey:   &label_key,
		LabelValue: &label_value,
	}

	LabelList = append(LabelList, labelList)
	nick_name := "nic-0701"
	fixed_ip := ""
	is_master := true
	subnet_id := "subnet-hazcjo9cm4"
	NetworkCardList := make([]ctecs.EcsSnapshotCreateInstanceNetworkCardListRequest, 0)
	networkCardList := ctecs.EcsSnapshotCreateInstanceNetworkCardListRequest{
		NicName:  &nick_name,
		FixedIP:  &fixed_ip,
		IsMaster: &is_master,
		SubnetID: &subnet_id,
	}

	NetworkCardList = append(NetworkCardList, networkCardList)
	client_token := "4cf2962d-e92c-4c00-9181-cfbb2218636ee082903"
	region_id := "bb9fdb42056f11eda1610242ac110002"
	project_id := ""
	instance_name := "go-sdk-082902"
	display_name := "go-sdk-082902"
	snapshot_id := "141ed492-b38e-8934-48b1-9186c8678a7c"
	vpc_id := "vpc-xq3tj5p30j"
	on_demand := false
	sec_group_list := []string{"sg-sn3ws4gwon"}
	ext_ip := "0"
	ip_version := ""
	bandwidth := 1
	ipv6_address_id := ""
	eip_id := ""
	affinity_group_id := ""
	key_pair_id := ""
	user_password := ""
	cycle_count := 1
	cycle_type := "MONTH"
	auto_renew_status := 0
	user_data := "ZWNobyBoZWxsbyBnb3N0YWNrIQ=="
	monitor_service := false

	response, err := apis.EcsSnapshotCreateInstanceApi.Do(context.Background(), credential, &ctecs.EcsSnapshotCreateInstanceRequest{
		ClientToken:     &client_token,
		RegionID:        &region_id,
		ProjectID:       &project_id,
		InstanceName:    &instance_name,
		DisplayName:     &display_name,
		SnapshotID:      &snapshot_id,
		VpcID:           &vpc_id,
		OnDemand:        &on_demand,
		SecGroupList:    &sec_group_list,
		NetworkCardList: NetworkCardList,
		ExtIP:           &ext_ip,
		IpVersion:       &ip_version,
		Bandwidth:       &bandwidth,
		Ipv6AddressID:   &ipv6_address_id,
		EipID:           &eip_id,
		AffinityGroupID: &affinity_group_id,
		KeyPairID:       &key_pair_id,
		UserPassword:    &user_password,
		CycleCount:      &cycle_count,
		CycleType:       &cycle_type,
		AutoRenewStatus: &auto_renew_status,
		UserData:        &user_data,
		LabelList:       LabelList,
		MonitorService:  &monitor_service,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsBackupPolicyCreate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	adv_day := 1
	adv_week := 1
	adv_month := 1
	adv_year := 1
	region_id := "bb9fdb42056f11eda1610242ac110002"
	policy_name := "test-bak-083002"
	cycle_type := "day"
	cycle_day := 1
	cycle_week := "0,2,6"
	time := "1,12"
	status := 1
	retention_type := "date"
	retention_day := 30
	retention_num := 20
	project_id := "0"
	total_backup := false
	adv_retention_status := true

	advRetention := ctecs.EcsBackupPolicyCreateAdvRetentionRequest{
		AdvDay:   &adv_day,
		AdvWeek:  &adv_week,
		AdvMonth: &adv_month,
		AdvYear:  &adv_year,
	}

	response, err := apis.EcsBackupPolicyCreateApi.Do(context.Background(), credential, &ctecs.EcsBackupPolicyCreateRequest{
		RegionID:           &region_id,
		PolicyName:         &policy_name,
		CycleType:          &cycle_type,
		CycleDay:           &cycle_day,
		CycleWeek:          &cycle_week,
		Time:               &time,
		Status:             &status,
		RetentionType:      &retention_type,
		RetentionDay:       &retention_day,
		RetentionNum:       &retention_num,
		ProjectID:          &project_id,
		TotalBackup:        &total_backup,
		AdvRetentionStatus: &adv_retention_status, // 假设启用了高级保留策略
		AdvRetention:       advRetention,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsBackupPolicyUpdate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)
	adv_day := 1
	adv_week := 1
	adv_month := 1
	adv_year := 1
	region_id := "bb9fdb42056f11eda1610242ac110002"
	policy_id := "3e251bce0d1411efb0a10242ac110002"
	policy_name := "test-bak-083003"
	cycle_type := "day"
	cycle_day := 1
	cycle_week := "0,2,6"
	time := "1,18"
	status := 1
	retention_type := "date"
	retention_day := 30
	retention_num := 20
	total_backup := false
	adv_retention_status := true

	advRetention := ctecs.EcsBackupPolicyUpdateAdvRetentionRequest{
		AdvDay:   &adv_day,
		AdvWeek:  &adv_week,
		AdvMonth: &adv_month,
		AdvYear:  &adv_year,
	}

	response, err := apis.EcsBackupPolicyUpdateApi.Do(context.Background(), credential, &ctecs.EcsBackupPolicyUpdateRequest{
		RegionID:           &region_id,
		PolicyID:           &policy_id,
		PolicyName:         &policy_name,
		CycleType:          &cycle_type,
		CycleDay:           &cycle_day,
		CycleWeek:          &cycle_week,
		Time:               &time,
		Status:             &status,
		RetentionType:      &retention_type,
		RetentionDay:       &retention_day,
		RetentionNum:       &retention_num,
		TotalBackup:        &total_backup,
		AdvRetentionStatus: &adv_retention_status, // 假设启用了高级保留策略
		AdvRetention:       advRetention,
	})

	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(response)
	fmt.Println(string(jsonstr))
}

func ecsOrderQueryUuid(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctecs.NewApis(client)

	master_order_id := "4fd3a1be605811efbff50242ac110005"
	response, err := apis.EcsOrderQueryUuidApi.Do(context.Background(), credential, &ctecs.EcsOrderQueryUuidRequest{
		MasterOrderId: &master_order_id,
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
	ecsOrderQueryUuid(*credential)
	ecsBackupPolicyUpdate(*credential)
	ecsBackupPolicyCreate(*credential)
	ecsSnapshotDelete(*credential)
	ecsSnapshotStatus(*credential)
	ecsSnapshotStatistics(*credential)
	ecsSnapshotDetails(*credential)
	ecsSnapshotList(*credential)
	ecsSnapshotBatchUpdate(*credential)
	ecsUpdateSnapshot(*credential)
	ecsRestoreSnapshot(*credential)
	ecsCreateSnapshot(*credential)
	ecsBackupPolicyUnbindRepo(*credential)
	ecsBackupPolicyBindRepo(*credential)
	ecsBackupPolicyUnbindInstances(*credential)
	ecsBackupPolicyBindInstances(*credential)
	ecsBackupPolicyListInstances(*credential)
	ecsBackupPolicyListApi(*credential)
	ecsBatchResetPassword(*credential)
	ecsVolumeStatistics(*credential)
	ecsVncDetails(*credential)
	ecsUpdateInstance(*credential)
	ecsSnapshotCreateInstance(*credential)
	ecsSnapshotPolicyList(*credential)
	ecsSnapshotPolicyInstanceList(*credential)
	ecsSnapshotPolicyTaskList(*credential)
	ecsVpcCreateSecurityGroupEgress(*credential)
	ecsSnapshotPolicyUnbindInstances(*credential)
	ecsSnapshotPolicyBindInstances(*credential)
	ecsSnapshotPolicyDelete(*credential)
	ecsSnapshotPolicyDetails(*credential)
	ecsSnapshotPolicyExecute(*credential)
	ecsSnapshotPolicyDisable(*credential)
	ecsSnapshotPolicyEnable(*credential)
	ecsSnapshotPolicyUpdate(*credential)
	ecsSnapshotPolicyCreate(*credential)
	ecsVolumeDeleteExec(*credential)
	ecsVolumeDetachExec(*credential)
	ecsVolumeAttachExec(*credential)
	ecsVolumeShowExec(*credential)
	ecsVolumeShowOldExec(*credential)
	ecsVolumeExtendExec(*credential)
	ecsVolumeUpdateExec(*credential)
	ecsVolumeCreateExec(*credential)
	ecsVpcDeleteSecurityGroupExec(*credential)
	ecsVpcDescribeSecurityGroupAttributeExec(*credential)
	ecsVpcQuerySecurityGroupsExec(*credential)
	ecsVpcCreateSecurityGroupExec(*credential)
	ecsBackupRepoDeleteExec(*credential)
	ecsBackupRepoListExec(*credential)
	ecsBackupRepoUpgradeExec(*credential)
	ecsBackupRepoRenewExec(*credential)
	ecsBackupRepoCreateExec(*credential)
	ecsAffinityGroupDetailsExec(*credential)
	ecsAffinityGroupDeleteExec(*credential)
	ecsAffinityGroupBindInstanceCheckExec(*credential)
	ecsAffinityGroupUpdateExec(*credential)
	ecsAffinityGroupListInstanceExec(*credential)
	ecsAffinityGroupListExec(*credential)
	ecsAffinityGroupUnbindInstancesExec(*credential)
	ecsAffinityGroupUnbindInstanceExec(*credential)
	ecsAffinityGroupBindInstanceExec(*credential)
	ecsAffinityGroupCreateExec(*credential)
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
