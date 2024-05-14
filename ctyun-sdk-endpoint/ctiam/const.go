package ctiam

import "github.com/ctyun-it/ctyun-sdk-go/ctyunsdk"

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
