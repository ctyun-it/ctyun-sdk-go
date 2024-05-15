package ctvpc

import "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"

const (
	EndpointNameCtvpc = "ctvpc"
	UrlProdCtvpc      = "ctvpc-global.ctapi.ctyun.cn"
	UrlTestCtvpc      = "ctvpc-global.ctapi-test.ctyun.cn"
)

var EndpointCtvpcProd = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtvpc,
	Url:          UrlProdCtvpc,
}

var EndpointCtvpcTest = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtvpc,
	Url:          UrlTestCtvpc,
}
