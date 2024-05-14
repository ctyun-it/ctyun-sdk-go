package ctebs

import "github.com/ctyun/ctyun_sdk_core/ctyunsdk"

const (
	EndpointNameCtebs = "ctebs"
	UrlProdCtebs      = "ebs-global.ctapi.ctyun.cn"
	UrlTestCtebs      = "ebs-global.ctapi-test.ctyun.cn"
)

var EndpointCtebsProd = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtebs,
	Url:          UrlProdCtebs,
}

var EndpointCtebsTest = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtebs,
	Url:          UrlTestCtebs,
}
