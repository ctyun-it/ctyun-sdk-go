package ctiam

import "github.com/ctyun/ctyun_sdk_core/ctyunsdk"

const (
	EndpointNameCtiam = "ctiam"
	UrlProdCtiam      = "ctiam-global.ctapi.ctyun.cn"
	UrlTestCtiam      = "ctiam-global.ctapi-test.ctyun.cn:21443"
)

var EndpointCtiamProd = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtiam,
	Url:          UrlProdCtiam,
}

var EndpointCtiamTest = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtiam,
	Url:          UrlTestCtiam,
}
