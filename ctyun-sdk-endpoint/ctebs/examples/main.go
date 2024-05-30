package main

import (
	"context"
	"encoding/json"
	"fmt"

	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctebs"
)

func ebsShow(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctebs.NewApis(client)
	ctx := context.TODO()
	createRes, err := apis.EbsShowApi.Do(ctx, credential, &ctebs.EbsShowRequest{
		RegionId: "bb9fdb42056f11eda1610242ac110002",
		DiskId:   "21cb3822-9083-482d-95a0-fa5d33892ad8",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func ebsChangeName(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctebs.NewApis(client)
	ctx := context.TODO()
	createRes, err := apis.EbsChangeNameApi.Do(ctx, credential, &ctebs.EbsChangeNameRequest{
		RegionId: "regionID",
		DiskId:   "diskID",
		DiskName: "diskName",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func ebsAssociate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctebs.NewApis(client)
	ctx := context.TODO()
	createRes, err := apis.EbsAssociateApi.Do(ctx, credential, &ctebs.EbsAssociateRequest{
		RegionId:   "regionID",
		DiskId:     "diskID",
		InstanceID: "instanceID",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func ebsDisassociate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctebs.NewApis(client)
	ctx := context.TODO()
	res, err := apis.EbsDisassociateApi.Do(ctx, credential, &ctebs.EbsDisassociateRequest{
		RegionId: "regionID",
		DiskId:   "diskID",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func ebsCreate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctebs.NewApis(client)
	ctx := context.TODO()
	res, err := apis.EbsCreateApi.Do(ctx, credential, &ctebs.EbsCreateRequest{
		ClientToken: "clientToken",
		DiskName:    "sdktest",
		DiskMode:    "VBD",
		DiskType:    "SATA",
		DiskSize:    20,
		RegionID:    "regionID",
		AzName:      "azName",
		OnDemand:    true,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func ebsChangeSize(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctebs.NewApis(client)
	ctx := context.TODO()
	res, err := apis.EbsChangeSizeApi.Do(ctx, credential, &ctebs.EbsChangeSizeRequest{
		ClientToken: "clientToken",
		DiskID:      "diskID",
		DiskSize:    30,
		RegionID:    "regionID",
	})
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func ebsDelete(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctebs.NewApis(client)
	ctx := context.TODO()
	res, err := apis.EbsDeleteApi.Do(ctx, credential, &ctebs.EbsDeleteRequest{
		ClientToken: "clientToken",
		DiskID:      "diskID",
		RegionID:    "regionID",
	})
	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func ebsList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctebs.NewApis(client)
	ctx := context.TODO()
	createRes, err := apis.EbsListApi.Do(ctx, credential, &ctebs.EbsListRequest{
		RegionID: "bb9fdb42056f11eda1610242ac110002",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func main() {
	credential, _ := ctyunsdk.NewCredential("3dd8e284f31e4a01ba24b700a049edcb", "c5b3891475b34e3a902f1676124802e0")
	ebsCreate(*credential)
	ebsList(*credential)
	ebsShow(*credential)
	ebsChangeName(*credential)
	ebsChangeSize(*credential)
	ebsAssociate(*credential)
	ebsDisassociate(*credential)
	ebsDelete(*credential)
}
