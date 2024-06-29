/******************************************************************************
 * Copyright (c) 2024. Archer++. All rights reserved.                         *
 * Author ORCID: https://orcid.org/0009-0003-8150-367X                        *
 ******************************************************************************/

package server

import (
	"github.com/go-kratos/kratos/v2/middleware/validate"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "api_gateway_service/api/http/v1"
	"api_gateway_service/internal/conf"
	"api_gateway_service/internal/service"
	"api_gateway_service/pkg/kratos/custom"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, userService *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			// middleware.Chain(
			recovery.Recovery(),
			validate.Validator(),
			// ),
		),
		http.ResponseEncoder(custom.ResponseEncoder),
		http.ErrorEncoder(custom.ErrorEncoder),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	v1.RegisterAccountHTTPServer(srv, userService)
	v1.RegisterAddressesHTTPServer(srv, userService)
	v1.RegisterReceivingMethodsHTTPServer(srv, userService)

	return srv
}
