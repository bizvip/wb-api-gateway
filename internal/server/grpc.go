/**********************************************************************************************************************
 * Copyright (c) 2024. Archer++. All rights reserved.   Author ORCID: https://orcid.org/0009-0003-8150-367X           *
 **********************************************************************************************************************/

package server

import (
	"github.com/go-kratos/kratos/v2/middleware/validate"

	v1 "api_gateway_service/api/http/v1"
	"api_gateway_service/internal/conf"
	"api_gateway_service/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, userServ *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}

	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	srv := grpc.NewServer(opts...)
	v1.RegisterAccountServer(srv, userServ)
	v1.RegisterAddressesServer(srv, userServ)
	v1.RegisterReceivingMethodsServer(srv, userServ)

	return srv
}
