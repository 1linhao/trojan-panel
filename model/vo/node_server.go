package vo

import "time"

type NodeServerVo struct {
	Id                uint      `json:"id"`
	Name              string    `json:"name"`
	Ip                string    `json:"ip"`
	GrpcPort          uint      `json:"grpcPort"`
	GrpcTLSMode       string    `json:"grpcTlsMode"`
	GrpcTLSServerName string    `json:"grpcTlsServerName"`
	CreateTime        time.Time `json:"createTime"`

	Status                 int    `json:"status"`
	TrojanPanelCoreVersion string `json:"trojanPanelCoreVersion"`
	KernelSummary          string `json:"kernelSummary"`
}

type NodeServerPageVo struct {
	NodeServers []NodeServerVo `json:"nodeServers"`
	BaseVoPage
}

type NodeServerOneVo struct {
	Id                uint      `json:"id"`
	Name              string    `json:"name"`
	Ip                string    `json:"ip"`
	GrpcPort          uint      `json:"grpcPort"`
	GrpcTLSMode       string    `json:"grpcTlsMode"`
	GrpcTLSServerName string    `json:"grpcTlsServerName"`
	CreateTime        time.Time `json:"createTime"`
}

type NodeServerListVo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type NodeServerInfoVo struct {
	CpuUsed  float32 `json:"cpuUsed"`
	MemUsed  float32 `json:"memUsed"`
	DiskUsed float32 `json:"diskUsed"`
}

type NodeServerExportVo struct {
	Name              string    `json:"name" ddb:"name"`
	Ip                string    `json:"ip" ddb:"ip"`
	GrpcPort          uint      `json:"grpcPort" ddb:"grpc_port"`
	GrpcTLSMode       string    `json:"grpcTlsMode" ddb:"grpc_tls_mode"`
	GrpcTLSServerName string    `json:"grpcTlsServerName" ddb:"grpc_tls_server_name"`
	CreateTime        time.Time `json:"createTime" ddb:"create_time"`
}
