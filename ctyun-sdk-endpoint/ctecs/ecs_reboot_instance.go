package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsRebootInstanceApi 重启一台云主机
// https://www.ctyun.cn/document/10026730/10106399
type EcsRebootInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsRebootInstanceApi(client *ctyunsdk.CtyunClient) *EcsRebootInstanceApi {
	return &EcsRebootInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/reboot-instance",
		},
	}
}

func (this *EcsRebootInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsRebootInstanceRequest) (*EcsRebootInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	_, err := builder.WriteJson(&ecsRebootInstanceRealRequest{
		RegionID:   req.RegionID,
		InstanceID: req.InstanceID,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsRebootInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsRebootInstanceResponse{
		JobId: realResponse.JobID,
	}, nil
}

type ecsRebootInstanceRealRequest struct {
	RegionID   string `json:"regionID"`
	InstanceID string `json:"instanceID"`
}

type EcsRebootInstanceRequest struct {
	RegionID   string
	InstanceID string
}

type EcsRebootInstanceRealResponse struct {
	JobID string `json:"jobID"`
}

type EcsRebootInstanceResponse struct {
	JobId string
}
