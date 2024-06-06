package common

import ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"

const (
	EndpointNameCtecs = "ctecs"
	UrlProdCtecs      = "ctecs-global.ctapi.ctyun.cn"
	UrlTestCtecs      = "ctecs-global.ctapi-test.ctyun.cn:21443"
)

var EndpointCtecsProd = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtecs,
	Url:          UrlProdCtecs,
}

var EndpointCtecsTest = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtecs,
	Url:          UrlTestCtecs,
}
