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
		DiskId:   "83783d88-76fa-48dc-9bd4-00e1994b361d",
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

func main() {
	credential, _ := ctyunsdk.NewCredential("ak", "sk")
	ebsCreate(*credential)
	ebsShow(*credential)
	ebsChangeName(*credential)
	ebsChangeSize(*credential)
	ebsAssociate(*credential)
	ebsDisassociate(*credential)
	ebsDelete(*credential)
}
