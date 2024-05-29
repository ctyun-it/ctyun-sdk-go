package ctecs

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EcsRebuildInstanceApi 重装一台云主机
// https://www.ctyun.cn/document/10026730/10106439
type EcsRebuildInstanceApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEcsRebuildInstanceApi(client *ctyunsdk.CtyunClient) *EcsRebuildInstanceApi {
	return &EcsRebuildInstanceApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodPost,
			UrlPath: "/v4/ecs/rebuild-instance",
		},
	}
}

func (this *EcsRebuildInstanceApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EcsRebuildInstanceRequest) (*EcsRebuildInstanceResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)
	_, err := builder.WriteJson(&ecsRebuildInstanceRealRequest{
		RegionID:       req.RegionID,
		InstanceID:     req.InstanceID,
		Password:       req.Password,
		ImageID:        req.ImageID,
		UserData:       req.UserData,
		InstanceName:   req.InstanceName,
		MonitorService: req.MonitorService,
	})
	if err != nil {
		return nil, err
	}

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameCtecs, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EcsRebuildInstanceRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}
	return &EcsRebuildInstanceResponse{
		JobId: realResponse.JobID,
	}, nil
}

type ecsRebuildInstanceRealRequest struct {
	RegionID       string `json:"regionID"`
	InstanceID     string `json:"instanceID"`
	Password       string `json:"password"`
	ImageID        string `json:"imageID"`
	UserData       string `json:"userData"`
	InstanceName   string `json:"instanceName"`
	MonitorService bool   `json:"monitorService"`
}

type EcsRebuildInstanceRequest struct {
	RegionID       string
	InstanceID     string
	Password       string
	ImageID        string
	UserData       string
	InstanceName   string
	MonitorService bool
}

type EcsRebuildInstanceRealResponse struct {
	JobID string `json:"jobID"`
}

type EcsRebuildInstanceResponse struct {
	JobId string
}
