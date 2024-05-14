package ctimage

import "github.com/ctyun/ctyun_sdk_core/ctyunsdk"

const (
	EndpointNameCtimage = "ctimage"
	UrlProdCtiamge      = "ctimage-global.ctapi.ctyun.cn"
	UrlTestCtiamge      = "ctimage-global.ctapi-test.ctyun.cn:21443"
)

var EndpointCtimageProd = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtimage,
	Url:          UrlProdCtiamge,
}

var EndpointCtimageTest = ctyunsdk.Endpoint{
	EndpointName: EndpointNameCtimage,
	Url:          UrlTestCtiamge,
}
