package ctebm

import "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"

const (
	EndpointNameEbm = "ebm"
	UrlProdEbm      = "ebm-global.ctapi.ctyun.cn"
	UrlTestEbm      = "ebm-global.ctapi-test.ctyun.cn:21443"
)

var EndpointEbmProd = ctyunsdk.Endpoint{
	EndpointName: EndpointNameEbm,
	Url:          UrlProdEbm,
}

var EndpointEbmTest = ctyunsdk.Endpoint{
	EndpointName: EndpointNameEbm,
	Url:          UrlTestEbm,
}
