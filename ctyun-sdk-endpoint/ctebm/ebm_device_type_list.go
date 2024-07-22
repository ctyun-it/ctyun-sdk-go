package ctebm

import (
	"context"
	"github.com/ctyun-it/ctyun-sdk-go/ctyun-sdk-core"
	"net/http"
)

// EbmDeviceTypeListApi D:\Project\go-sdk-auto-write\docs\查询资源池内物理机套餐
// https://eop.ctyun.cn/ebp/ctapiDocument/search?sid=16&api=4574&data=97&isNormal=1
type EbmDeviceTypeListApi struct {
	ctyunsdk.CtyunRequestBuilder
	client *ctyunsdk.CtyunClient
}

func NewEbmDeviceTypeListApi(client *ctyunsdk.CtyunClient) *EbmDeviceTypeListApi {
	return &EbmDeviceTypeListApi{
		client: client,
		CtyunRequestBuilder: ctyunsdk.CtyunRequestBuilder{
			Method:  http.MethodGet,
			UrlPath: "/v4/ebm/device-type-list",
		},
	}
}

func (this *EbmDeviceTypeListApi) Do(ctx context.Context, credential ctyunsdk.Credential, req *EbmDeviceTypeListRequest) (*EbmDeviceTypeListResponse, ctyunsdk.CtyunRequestError) {
	builder := this.WithCredential(&credential)

	builder.
		AddParam("regionID", req.RegionID).
		AddParam("azName", req.AzName).
		AddParam("deviceType", req.DeviceType)

	response, err := this.client.RequestToEndpoint(ctx, EndpointNameEbm, builder)
	if err != nil {
		return nil, err
	}

	var realResponse EbmDeviceTypeListRealResponse
	err = response.ParseByStandardModelWithCheck(&realResponse)
	if err != nil {
		return nil, err
	}

	var results []EbmDeviceTypeListResultsResponse
	for _, res := range realResponse.Results {
		results = append(results, EbmDeviceTypeListResultsResponse{
			Id:                      res.Id,
			Region:                  res.Region,
			AzName:                  res.AzName,
			DeviceType:              res.DeviceType,
			NameZh:                  res.NameZh,
			NameEn:                  res.NameEn,
			CpuSockets:              res.CpuSockets,
			NumaNodeAmount:          res.NumaNodeAmount,
			CpuAmount:               res.CpuAmount,
			CpuThreadAmount:         res.CpuThreadAmount,
			CpuManufacturer:         res.CpuManufacturer,
			CpuModel:                res.CpuModel,
			CpuFrequency:            res.CpuFrequency,
			MemAmount:               res.MemAmount,
			MemSize:                 res.MemSize,
			MemFrequency:            res.MemFrequency,
			NicAmount:               res.NicAmount,
			NicRate:                 res.NicRate,
			SystemVolumeAmount:      res.SystemVolumeAmount,
			SystemVolumeSize:        res.SystemVolumeSize,
			SystemVolumeType:        res.SystemVolumeType,
			SystemVolumeInterface:   res.SystemVolumeInterface,
			SystemVolumeDescription: res.SystemVolumeDescription,
			DataVolumeAmount:        res.DataVolumeAmount,
			DataVolumeSize:          res.DataVolumeSize,
			DataVolumeInterface:     res.DataVolumeInterface,
			DataVolumeType:          res.DataVolumeType,
			DataVolumeDescription:   res.DataVolumeDescription,
			SmartNicExist:           res.SmartNicExist,
			NvmeVolumeAmount:        res.NvmeVolumeAmount,
			NvmeVolumeSize:          res.NvmeVolumeSize,
			NvmeVolumeType:          res.NvmeVolumeType,
			NvmeVolumeInterface:     res.NvmeVolumeInterface,
			GpuAmount:               res.GpuAmount,
			GpuSize:                 res.GpuSize,
			GpuManufacturer:         res.GpuManufacturer,
			GpuModel:                res.GpuModel,
			ComputeIBAmount:         res.ComputeIBAmount,
			ComputeIBRate:           res.ComputeIBRate,
			StorageIBAmount:         res.StorageIBAmount,
			StorageIBRate:           res.StorageIBRate,
			ComputeRoCEAmount:       res.ComputeRoCEAmount,
			ComputeRoCERate:         res.ComputeRoCERate,
			StorageRoCEAmount:       res.StorageRoCEAmount,
			StorageRoCERate:         res.StorageRoCERate,
			SupportCloud:            res.SupportCloud,
			CloudBoot:               res.CloudBoot,
			CreateTime:              res.CreateTime,
			UpdateTime:              res.UpdateTime,
		})
	}

	return &EbmDeviceTypeListResponse{
		TotalCount: realResponse.TotalCount,
		Results:    results,
	}, nil
}

type EbmDeviceTypeListRealRequest struct {
	RegionID   string `json:"regionID,omitempty"`
	AzName     string `json:"azName,omitempty"`
	DeviceType string `json:"deviceType,omitempty"`
}

type EbmDeviceTypeListRequest struct {
	RegionID   string
	AzName     string
	DeviceType string
}

type EbmDeviceTypeListResultsRealResponse struct {
	Id                      int    `json:"id,omitempty"`
	Region                  string `json:"region,omitempty"`
	AzName                  string `json:"azName,omitempty"`
	DeviceType              string `json:"deviceType,omitempty"`
	NameZh                  string `json:"nameZh,omitempty"`
	NameEn                  string `json:"nameEn,omitempty"`
	CpuSockets              int    `json:"cpuSockets,omitempty"`
	NumaNodeAmount          int    `json:"numaNodeAmount,omitempty"`
	CpuAmount               int    `json:"cpuAmount,omitempty"`
	CpuThreadAmount         int    `json:"cpuThreadAmount,omitempty"`
	CpuManufacturer         string `json:"cpuManufacturer,omitempty"`
	CpuModel                string `json:"cpuModel,omitempty"`
	CpuFrequency            string `json:"cpuFrequency,omitempty"`
	MemAmount               int    `json:"memAmount,omitempty"`
	MemSize                 int    `json:"memSize,omitempty"`
	MemFrequency            int    `json:"memFrequency,omitempty"`
	NicAmount               int    `json:"nicAmount,omitempty"`
	NicRate                 int    `json:"nicRate,omitempty"`
	SystemVolumeAmount      int    `json:"systemVolumeAmount,omitempty"`
	SystemVolumeSize        int    `json:"systemVolumeSize,omitempty"`
	SystemVolumeType        string `json:"systemVolumeType,omitempty"`
	SystemVolumeInterface   string `json:"systemVolumeInterface,omitempty"`
	SystemVolumeDescription string `json:"systemVolumeDescription,omitempty"`
	DataVolumeAmount        int    `json:"dataVolumeAmount,omitempty"`
	DataVolumeSize          int    `json:"dataVolumeSize,omitempty"`
	DataVolumeInterface     string `json:"dataVolumeInterface,omitempty"`
	DataVolumeType          string `json:"dataVolumeType,omitempty"`
	DataVolumeDescription   string `json:"dataVolumeDescription,omitempty"`
	SmartNicExist           bool   `json:"smartNicExist,omitempty"`
	NvmeVolumeAmount        int    `json:"nvmeVolumeAmount,omitempty"`
	NvmeVolumeSize          int    `json:"nvmeVolumeSize,omitempty"`
	NvmeVolumeType          string `json:"nvmeVolumeType,omitempty"`
	NvmeVolumeInterface     string `json:"nvmeVolumeInterface,omitempty"`
	GpuAmount               int    `json:"gpuAmount,omitempty"`
	GpuSize                 int    `json:"gpuSize,omitempty"`
	GpuManufacturer         string `json:"gpuManufacturer,omitempty"`
	GpuModel                string `json:"gpuModel,omitempty"`
	ComputeIBAmount         int    `json:"computeIBAmount,omitempty"`
	ComputeIBRate           int    `json:"computeIBRate,omitempty"`
	StorageIBAmount         int    `json:"storageIBAmount,omitempty"`
	StorageIBRate           int    `json:"storageIBRate,omitempty"`
	ComputeRoCEAmount       int    `json:"computeRoCEAmount,omitempty"`
	ComputeRoCERate         int    `json:"computeRoCERate,omitempty"`
	StorageRoCEAmount       int    `json:"storageRoCEAmount,omitempty"`
	StorageRoCERate         int    `json:"storageRoCERate,omitempty"`
	SupportCloud            bool   `json:"supportCloud,omitempty"`
	CloudBoot               bool   `json:"cloudBoot,omitempty"`
	CreateTime              string `json:"createTime,omitempty"`
	UpdateTime              string `json:"updateTime,omitempty"`
}

type EbmDeviceTypeListRealResponse struct {
	TotalCount int                                    `json:"totalCount,omitempty"`
	Results    []EbmDeviceTypeListResultsRealResponse `json:"results,omitempty"`
}

type EbmDeviceTypeListResultsResponse struct {
	Id                      int
	Region                  string
	AzName                  string
	DeviceType              string
	NameZh                  string
	NameEn                  string
	CpuSockets              int
	NumaNodeAmount          int
	CpuAmount               int
	CpuThreadAmount         int
	CpuManufacturer         string
	CpuModel                string
	CpuFrequency            string
	MemAmount               int
	MemSize                 int
	MemFrequency            int
	NicAmount               int
	NicRate                 int
	SystemVolumeAmount      int
	SystemVolumeSize        int
	SystemVolumeType        string
	SystemVolumeInterface   string
	SystemVolumeDescription string
	DataVolumeAmount        int
	DataVolumeSize          int
	DataVolumeInterface     string
	DataVolumeType          string
	DataVolumeDescription   string
	SmartNicExist           bool
	NvmeVolumeAmount        int
	NvmeVolumeSize          int
	NvmeVolumeType          string
	NvmeVolumeInterface     string
	GpuAmount               int
	GpuSize                 int
	GpuManufacturer         string
	GpuModel                string
	ComputeIBAmount         int
	ComputeIBRate           int
	StorageIBAmount         int
	StorageIBRate           int
	ComputeRoCEAmount       int
	ComputeRoCERate         int
	StorageRoCEAmount       int
	StorageRoCERate         int
	SupportCloud            bool
	CloudBoot               bool
	CreateTime              string
	UpdateTime              string
}

type EbmDeviceTypeListResponse struct {
	TotalCount int
	Results    []EbmDeviceTypeListResultsResponse
}
