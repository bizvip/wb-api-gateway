package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"

	"api_gateway_service/api/common"
	v1 "api_gateway_service/api/http/v1"
)

type GatewayService struct {
	v1.UnimplementedAccountServer
	log *log.Helper
}

func NewGatewayService(logger log.Logger) *GatewayService {
	return &GatewayService{
		log: log.NewHelper(logger),
	}
}

// RegisterAccount 注册账户
func (s *GatewayService) RegisterAccount(ctx context.Context, req *v1.RegisterAccountRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// GetAccount 读取账户
func (s *GatewayService) GetAccount(ctx context.Context, req *v1.GetAccountRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// UpdateAccount 更新账户
func (s *GatewayService) UpdateAccount(ctx context.Context, req *v1.UpdateAccountRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// ListAddresses 地址簿列表
func (s *GatewayService) ListAddresses(ctx context.Context, req *empty.Empty) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// GetAddress 读取单条地址
func (s *GatewayService) GetAddress(ctx context.Context, req *v1.GetAddressRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// AddAddress 添加地址
func (s *GatewayService) AddAddress(ctx context.Context, req *v1.AddAddressRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// UpdateAddress 更新地址
func (s *GatewayService) UpdateAddress(ctx context.Context, req *v1.UpdateAddressRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// DeleteAddress 删除地址
func (s *GatewayService) DeleteAddress(ctx context.Context, req *v1.DeleteAddressRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// ListReceivingMethods 收款方式列表
func (s *GatewayService) ListReceivingMethods(ctx context.Context, req *empty.Empty) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// GetReceivingMethod 获取单个收款方式
func (s *GatewayService) GetReceivingMethod(ctx context.Context, req *v1.GetReceivingMethodRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// AddReceivingMethod 添加收款方式
func (s *GatewayService) AddReceivingMethod(ctx context.Context, req *v1.AddReceivingMethodRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// UpdateReceivingMethod 更新收款方式
func (s *GatewayService) UpdateReceivingMethod(ctx context.Context, req *v1.UpdateReceivingMethodRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// DeleteReceivingMethod 删除收款方式
func (s *GatewayService) DeleteReceivingMethod(ctx context.Context, req *v1.DeleteReceivingMethodRequest) (*common.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &common.Reply{Code: 0, Msg: "", Data: nil}, nil
}
