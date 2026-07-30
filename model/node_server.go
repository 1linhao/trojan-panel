package model

import "time"

type NodeServer struct {
	Id                *uint      `ddb:"id" json:"id"`
	Name              *string    `ddb:"name" json:"name"`
	Ip                *string    `ddb:"ip" json:"ip"`
	GrpcPort          *uint      `ddb:"grpc_port" json:"grpcPort"`
	GrpcTLSMode       *string    `ddb:"grpc_tls_mode" json:"grpcTlsMode"`
	GrpcTLSServerName *string    `ddb:"grpc_tls_server_name" json:"grpcTlsServerName"`
	CreateTime        *time.Time `ddb:"create_time" json:"createTime"`
	UpdateTime        *time.Time `ddb:"update_time" json:"updateTime"`
}
