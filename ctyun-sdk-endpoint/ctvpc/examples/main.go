package main

import (
	"context"
	"encoding/json"
	"fmt"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctvpc"
	"log"
)

func bandwidthOperation(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctvpc.NewApis(client)
	ctx := context.TODO()
	createRes, err := apis.BandwidthCreateApi.Do(ctx, credential, &ctvpc.BandwidthCreateRequest{
		RegionID:    "regionID",
		ClientToken: "abcedefghijklmn",
		Name:        "test",
		CycleType:   "on_demand",
		Bandwidth:   5,
	})

	if err != nil {
		panic(err)
	}

	for {
		if createRes.MasterResourceStatus == "in_progress" {
			createRes, err = apis.BandwidthCreateApi.Do(ctx, credential, &ctvpc.BandwidthCreateRequest{
				RegionID:    "regionID",
				ClientToken: "abcedefghijklmn",
				Name:        "test",
				CycleType:   "on_demand",
				Bandwidth:   5,
			})
			if err != nil {
				panic(err)
			}
		}

		if createRes.MasterResourceStatus != "in_progress" {
			break
		}
	}

	var bandwidthID string
	if createRes.MasterResourceStatus == "started" {
		bandwidthID = createRes.BandwidthId
	} else {
		panic(fmt.Errorf("order status %s", createRes.MasterResourceStatus))
	}

	changeNameHandler := ctvpc.NewApis(client).BandwidthChangeNameApi
	_, err = changeNameHandler.Do(ctx, credential, &ctvpc.BandwidthChangeNameRequest{
		ClientToken: "yyyyyy",
		RegionID:    "regionID",
		BandwidthID: bandwidthID,
		Name:        "wwwww",
	})
	if err != nil {
		panic(err)
	}

	addEipHandler := ctvpc.NewBandwidthAssociateEipApi(client)
	_, err = addEipHandler.Do(ctx, credential, &ctvpc.BandwidthAssociateEipRequest{
		ClientToken: "yyyyyy",
		RegionID:    "regionID",
		BandwidthID: bandwidthID,
		EipIDs:      []string{"eipID"},
	})
	if err != nil {
		panic(err)
	}

	removeEipHandler := ctvpc.NewApis(client).BandwidthDisassociateEipApi
	_, err = removeEipHandler.Do(ctx, credential, &ctvpc.BandwidthDisassociateEipRequest{
		ClientToken: "yyyyyy",
		RegionID:    "regionID",
		BandwidthID: bandwidthID,
		EipIDs:      []string{"eipID"},
	})
	if err != nil {
		panic(err)
	}

	showHandler := ctvpc.NewApis(client).BandwidthDescribeApi
	showRes, err := showHandler.Do(ctx, credential, &ctvpc.BandwidthDescribeRequest{
		BandwidthID: bandwidthID,
		RegionID:    "regionID",
	})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", showRes)

	modifySpecHandler := ctvpc.NewApis(client).BandwidthChangeSpecApi
	_, err = modifySpecHandler.Do(ctx, credential, &ctvpc.BandwidthChangeSpecRequest{
		ClientToken: "tttttttttttttt",
		RegionID:    "regionID",
		BandwidthID: bandwidthID,
		Bandwidth:   10,
	})
	if err != nil {
		panic(err)
	}

	deleteHandler := ctvpc.NewApis(client).BandwidthDeleteApi
	_, err = deleteHandler.Do(ctx, credential, &ctvpc.BandwidthDeleteRequest{
		ClientToken: "oooooooooooo",
		RegionID:    "regionID",
		BandwidthID: bandwidthID,
	})
	if err != nil {
		panic(err)
	}
}

func eipOperation(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	createHandler := ctvpc.NewApis(client).EipCreateApi
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &ctvpc.EipCreateRequest{
		RegionId:    "regionID",
		ClientToken: "abcedefghij",
		Name:        "test",
		CycleType:   "on_demand",
	})

	if err != nil {
		panic(err)
	}

	for {
		if createRes.MasterResourceStatus == "in_progress" {
			createRes, err = createHandler.Do(ctx, credential, &ctvpc.EipCreateRequest{
				RegionId:    "regionID",
				ClientToken: "abcedefghij",
				Name:        "test",
				CycleType:   "on_demand",
			})
			if err != nil {
				panic(err)
			}
		}

		if createRes.MasterResourceStatus != "in_progress" {
			break
		}
	}

	var eipID string
	if createRes.MasterResourceStatus == "started" {
		eipID = createRes.EipId
	} else {
		panic(fmt.Errorf("order status %s", createRes.MasterResourceStatus))
	}

	changeNameHandler := ctvpc.NewApis(client).EipChangeNameApi

	_, err = changeNameHandler.Do(ctx, credential, &ctvpc.EipChangeNameRequest{
		ClientToken: "yyyyyy",
		RegionId:    "regionID",
		EipId:       eipID,
		Name:        "kkkkk",
	})
	if err != nil {
		panic(err)
	}

	associateHandler := ctvpc.NewApis(client).EipAssociateApi
	_, err = associateHandler.Do(ctx, credential, &ctvpc.EipAssociateRequest{
		ClientToken:     "yyyyyy",
		RegionId:        "regionID",
		EipId:           eipID,
		AssociationType: 1, // 1 vm; 2 bm; 3 vip
		AssociationId:   "vmID",
	})
	if err != nil {
		panic(err)
	}

	disassociationHandler := ctvpc.NewApis(client).EipDisassociateApi
	_, err = disassociationHandler.Do(ctx, credential, &ctvpc.EipDisassociateRequest{
		ClientToken: "yyyyyy",
		RegionId:    "regionID",
		EipId:       eipID,
	})
	if err != nil {
		panic(err)
	}

	changeSpecHandler := ctvpc.NewApis(client).EipModifySpecApi
	_, err = changeSpecHandler.Do(ctx, credential, &ctvpc.EipModifySpecRequest{
		ClientToken: "xxxxxxxxxxx",
		RegionId:    "regionID",
		EipId:       eipID,
		Bandwidth:   2,
	})
	if err != nil {
		panic(err)
	}

	deleteHandler := ctvpc.NewApis(client).EipDeleteApi
	_, err = deleteHandler.Do(ctx, credential, &ctvpc.EipDeleteRequest{
		ClientToken: "zzzzzzzzzzzz",
		RegionId:    "regionID",
		EipId:       eipID,
	})
	if err != nil {
		panic(err)
	}
}

func sgOperation(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	createHandler := ctvpc.NewApis(client).SecurityGroupCreateApi
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &ctvpc.SecurityGroupCreateRequest{
		RegionId:    "regionID",
		VpcId:       "vpcID",
		ClientToken: "xyz",
		Name:        "test",
		Description: "test",
	})

	if err != nil {
		panic(err)
	}

	sgID := createRes.SecurityGroupId

	descHandler := ctvpc.NewApis(client).SecurityGroupDescribeAttributeApi
	sg, err := descHandler.Do(ctx, credential, &ctvpc.SecurityGroupDescribeAttributeRequest{
		RegionId:        "regionID",
		SecurityGroupId: sgID,
	})

	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", sg)

	updateHandler := ctvpc.NewApis(client).SecurityGroupModifyAttributionApi
	_, err = updateHandler.Do(ctx, credential, &ctvpc.SecurityGroupModifyAttributionRequest{
		SecurityGroupId: sgID,
		RegionId:        "regionID",
		Name:            "test-test",
	})

	if err != nil {
		panic(err)
	}

	deleteHandler := ctvpc.NewApis(client).SecurityGroupDeleteApi
	_, err = deleteHandler.Do(ctx, credential, &ctvpc.SecurityGroupDeleteRequest{
		SecurityGroupId: sgID,
		RegionId:        "regionID",
	})

	if err != nil {
		panic(err)
	}
}

func securityGroupRuleOperation(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	ctx := context.TODO()

	ingressCreateHandler := ctvpc.NewApis(client).SecurityGroupRuleIngressCreateApi

	createRes, err := ingressCreateHandler.Do(ctx, credential, &ctvpc.SecurityGroupRuleIngressCreateRequest{
		RegionId:        "regionID",
		ClientToken:     "xyz",
		SecurityGroupId: "securityGroupID",
		SecurityGroupRules: []ctvpc.SecurityGroupRuleIngressCreateSecurityGroupRulesRequest{{
			Direction:   "ingress",
			Action:      "accept",
			Priority:    1,
			Protocol:    "TCP",
			Ethertype:   "IPv4",
			DestCidrIp:  "0.0.0.0/0",
			Description: "dafgsdfd",
			Range:       "1-200",
		}},
	})

	if err != nil {
		panic(err)
	}

	ingressSgRuleID := createRes.SgRuleIds[0]

	egressCreateHandler := ctvpc.NewApis(client).SecurityGroupRuleEgressCreateApi

	egressCreateRes, err := egressCreateHandler.Do(ctx, credential, &ctvpc.SecurityGroupRuleEgressCreateRequest{
		RegionId:        "regionID",
		ClientToken:     "xyz",
		SecurityGroupId: "securityGroupID",
		SecurityGroupRules: []ctvpc.SecurityGroupRuleEgressCreateSecurityGroupRulesRequest{{
			Direction:   "gress",
			Action:      "accept",
			Priority:    1,
			Protocol:    "TCP",
			Ethertype:   "IPv4",
			DestCidrIp:  "0.0.0.0/0",
			Description: "dafgsdfd",
			Range:       "1-200",
		}},
	})

	if err != nil {
		panic(err)
	}

	egressSgRuleID := egressCreateRes.SgRuleIds[0]

	descHandler := ctvpc.NewApis(client).SecurityGroupRuleDescribeApi
	res, err := descHandler.Do(ctx, credential, &ctvpc.SecurityGroupRuleDescribeRequest{
		RegionId:            "regionID",
		SecurityGroupId:     "securityGroupID",
		SecurityGroupRuleId: ingressSgRuleID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", res)

	res, err = descHandler.Do(ctx, credential, &ctvpc.SecurityGroupRuleDescribeRequest{
		RegionId:            "regionID",
		SecurityGroupId:     "securityGroupID",
		SecurityGroupRuleId: egressSgRuleID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", res)

	ingressModifyHandler := ctvpc.NewApis(client).SecurityGroupRuleIngressModifyApi
	_, err = ingressModifyHandler.Do(ctx, credential, &ctvpc.SecurityGroupRuleIngressModifyRequest{
		RegionId:            "regionID",
		SecurityGroupId:     "securityGroupID",
		SecurityGroupRuleId: ingressSgRuleID,
		ClientToken:         "xyz",
		Description:         "xxxxx",
	})
	if err != nil {
		panic(err)
	}

	egressModifyHandler := ctvpc.NewApis(client).SecurityGroupRuleEgressModifyApi
	_, err = egressModifyHandler.Do(ctx, credential, &ctvpc.SecurityGroupRuleEgressModifyRequest{
		RegionId:            "regionID",
		SecurityGroupId:     "securityGroupID",
		SecurityGroupRuleId: egressSgRuleID,
		ClientToken:         "xyz",
		Description:         "xxxxx",
	})
	if err != nil {
		panic(err)
	}

	ingressRevokeHandler := ctvpc.NewApis(client).SecurityGroupRuleIngressRevokeApi
	_, err = ingressRevokeHandler.Do(ctx, credential, &ctvpc.SecurityGroupRuleIngressRevokeRequest{
		RegionId:            "regionID",
		SecurityGroupId:     "securityGroupID",
		SecurityGroupRuleId: ingressSgRuleID,
		ClientToken:         "xyz",
	})
	if err != nil {
		panic(err)
	}

	egressRevokeHandler := ctvpc.NewApis(client).SecurityGroupRuleEgressRevokeApi
	_, err = egressRevokeHandler.Do(ctx, credential, &ctvpc.SecurityGroupRuleEgressRevokeRequest{
		RegionId:            "regionID",
		SecurityGroupId:     "securityGroupID",
		SecurityGroupRuleId: ingressSgRuleID,
		ClientToken:         "xyz",
	})
	if err != nil {
		panic(err)
	}
}

func listSubnets(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	handler := ctvpc.NewApis(client).SubnetListApi
	res, err := handler.Do(context.TODO(), credential, &ctvpc.SubnetListRequest{RegionId: "regionID"})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", res)
}

func subnetOperation(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	createHandler := ctvpc.NewApis(client).SubnetCreateApi
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &ctvpc.SubnetCreateRequest{
		RegionId:    "regionID",
		VpcId:       "vpcID",
		ClientToken: "xyz",
		Name:        "test",
		Cidr:        "192.168.1.0/24",
		Description: "test",
	})

	if err != nil {
		panic(err)
	}

	updateHandler := ctvpc.NewApis(client).SubnetUpdateApi
	_, err = updateHandler.Do(ctx, credential, &ctvpc.SubnetUpdateRequest{
		SubnetId: createRes.SubnetId,
		RegionId: "regionID",
		Name:     "test-test",
	})

	if err != nil {
		panic(err)
	}

	deleteHandler := ctvpc.NewApis(client).SubnetDeleteApi
	_, err = deleteHandler.Do(ctx, credential, &ctvpc.SubnetDeleteRequest{
		SubnetId: createRes.SubnetId,
		RegionId: "regionID",
	})

	if err != nil {
		panic(err)
	}
}

func listVpcs(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	handler := ctvpc.NewApis(client).VpcListApi
	res, err := handler.Do(context.TODO(), credential, &ctvpc.VpcListRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		VpcIds:     []string{"vpc-chz0ilszsp"},
		VpcName:    "vpc-782c",
		PageNumber: 1,
		PageSize:   10,
		ProjectId:  "0",
	})
	if err != nil {
		fmt.Printf("错误信息为：%s", err)
		return
	}

	jsonstr, _ := json.Marshal(res)
	fmt.Println(string(jsonstr))
}

func vpcOperation(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)

	createHandler := ctvpc.NewApis(client).VpcCreateApi
	ctx := context.TODO()
	createRes, err := createHandler.Do(ctx, credential, &ctvpc.VpcCreateRequest{
		RegionId:    "regionID",
		ClientToken: "xyz",
		Name:        "test",
		Cidr:        "192.168.1.0/24",
		Description: "test",
		EnableIpv6:  false,
	})

	if err != nil {
		panic(err)
	}

	updateHandler := ctvpc.NewApis(client).VpcUpdateApi
	_, err = updateHandler.Do(ctx, credential, &ctvpc.VpcUpdateRequest{
		VpcId:       createRes.VpcId,
		ClientToken: "xyz",
		RegionId:    "regionID",
		Name:        "test-test",
	})

	if err != nil {
		panic(err)
	}

	deleteHandler := ctvpc.NewApis(client).VpcDeleteApi
	_, err = deleteHandler.Do(ctx, credential, &ctvpc.VpcDeleteRequest{
		VpcId:    createRes.VpcId,
		RegionId: "regionID",
	})

	if err != nil {
		panic(err)
	}
}

func main() {
	credential, _ := ctyunsdk.NewCredential("ak", "sk")
	bandwidthOperation(*credential)
	eipOperation(*credential)
	sgOperation(*credential)
	securityGroupRuleOperation(*credential)
	listSubnets(*credential)
	subnetOperation(*credential)
	listVpcs(*credential)
	vpcOperation(*credential)
}
