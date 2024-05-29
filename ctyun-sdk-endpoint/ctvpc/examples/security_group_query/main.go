package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-endpoint/ctvpc"
)

func sgOperation(ak, sk, regionID string) {
	client := ctyunsdk.EnvOf(ctyunsdk.EnvironmentProd)
	sg := ctvpc.NewApis(client).SecurityGroupQueryApi

	credential, _ := ctyunsdk.NewCredential(ak, sk)

	req := ctvpc.SecurityGroupQueryRequest{
		RegionId: regionID,
	}
	res, err := sg.Do(context.TODO(), *credential, &req)
	if err != nil {
		panic(err)
	}

	for _, item := range res.SecurityGroups {
		fmt.Println(item)
	}
}

func main() {
	var ak string
	var sk string
	var regionID string
	flag.StringVar(&ak, "ak", "", "access key")
	flag.StringVar(&sk, "sk", "", "secret key")
	flag.StringVar(&regionID, "region-id", "", "region id")
	flag.Parse()

	if len(ak) == 0 || len(sk) == 0 || len(regionID) == 0 {
		log.Print("ak or sk or region-id or vpc-id is required")
		return
	}

	sgOperation(ak, sk, regionID)
}
