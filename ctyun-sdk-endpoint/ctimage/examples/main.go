package main

import (
	"context"
	"encoding/json"
	"fmt"
	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctimage"
)

func imageDetail(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctimage.NewApis(client)
	ctx := context.TODO()
	createRes, err := apis.ImageDetailApi.Do(ctx, credential, &ctimage.ImageDetailRequest{
		RegionId: "regionID",
		ImageId:  "imageID",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(createRes)
	fmt.Printf(string(marshal))
}

func imageDelete(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctimage.NewApis(client)
	ctx := context.TODO()
	res, err := apis.ImageDeleteApi.Do(ctx, credential, &ctimage.ImageDeleteRequest{
		RegionId: "regionID",
		ImageId:  "imageID",
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func imageList(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctimage.NewApis(client)
	ctx := context.TODO()
	res, err := apis.ImageListApi.Do(ctx, credential, &ctimage.ImageListRequest{
		RegionId:   "bb9fdb42056f11eda1610242ac110002",
		AzName:     "cn-huadong1-jsnj1A-public-ctcloud",
		Visibility: 1,
		PageNo:     1,
		PageSize:   10,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func imageUpdate(credential ctyunsdk.Credential) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	apis := ctimage.NewApis(client)
	ctx := context.TODO()
	res, err := apis.ImageUpdateApi.Do(ctx, credential, &ctimage.ImageUpdateRequest{
		RegionId:   "regionID",
		ImageId:    "imageID",
		ImageName:  "imageName",
		MaximumRam: 16,
		MinimumRam: 4,
	})

	if err != nil {
		panic(err)
	}
	marshal, _ := json.Marshal(res)
	fmt.Printf(string(marshal))
}

func main() {
	credential, _ := ctyunsdk.NewCredential("ak", "sk")
	imageList(*credential)
	imageDetail(*credential)
	imageUpdate(*credential)
	imageDelete(*credential)
}
