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
		PayImage:       req.PayImage,
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
	RegionID       *string `json:"regionID,omitempty"`
	InstanceID     *string `json:"instanceID,omitempty"`
	Password       *string `json:"password,omitempty"`
	ImageID        *string `json:"imageID,omitempty"`
	UserData       *string `json:"userData,omitempty"`
	InstanceName   *string `json:"instanceName,omitempty"`
	MonitorService *bool   `json:"monitorService,omitempty"`
	PayImage       *bool   `json:"payImage,omitempty"`
}

type EcsRebuildInstanceRequest struct {
	RegionID       *string
	InstanceID     *string
	Password       *string
	ImageID        *string
	UserData       *string
	InstanceName   *string
	MonitorService *bool
	PayImage       *bool
}

type EcsRebuildInstanceRealResponse struct {
	JobID string `json:"jobID"`
}

type EcsRebuildInstanceResponse struct {
	JobId string
}
