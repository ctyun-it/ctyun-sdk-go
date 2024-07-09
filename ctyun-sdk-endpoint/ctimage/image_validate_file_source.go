package ctimage

import (
	"context"
	"net/http"

	ctyunsdk "github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
)

// ImageValidateFileSource 校验镜像文件地址
// https://www.ctyun.cn/document/10027726/10042097

type ImageValidateApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewImageValidateApi(client *ctyunsdk.CtyunClient) *ImageValidateApi {
	return &ImageValidateApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/image/validate-file-source",
		},
	}
}

func (this *ImageValidateApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *ImageValidateRequest) (*ImageValidateResponse, ctyunsdk.CtyunRequestError) {
	request := this.WithCredential(&credential).
		AddParam("regionID", req.RegionID).
		AddParam("imageFileSource", req.ImageFileSource)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtimage, request)
	if err != nil {
		return nil, err
	}

	imageValidateResponse := &ImageValidateResponse{}
	if err = response.ParseByStandardModelWithCheck(imageValidateResponse); err != nil {
		return nil, err
	}

	return imageValidateResponse, nil
}

type ImageValidateRequest struct {
	ImageFileSource string // 镜像文件地址，格式应为 {internetEndpoint}/{bucket}/{key}
	RegionID        string // 资源池 ID
}

type ImageValidateResponse struct {
	Archive    bool `json:"archive"`
	Bucket     bool `json:"bucket"`
	Encryption bool `json:"encryption"`
	Object     bool `json:"object"`
}
