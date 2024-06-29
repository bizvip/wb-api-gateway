/******************************************************************************
 * Copyright (c) 2024. Archer++. All rights reserved.                         *
 * Author ORCID: https://orcid.org/0009-0003-8150-367X                        *
 ******************************************************************************/

package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/anypb"

	v1 "api_gateway_service/api/http/v1"
)

type UserService struct {
	v1.UnimplementedAccountServer
	v1.UnimplementedAddressesServer
	v1.UnimplementedReceivingMethodsServer
	log *log.Helper
}

func NewUserService(logger log.Logger) *UserService {
	return &UserService{
		log: log.NewHelper(logger),
	}
}

// RegisterAccount 注册账户
func (s *UserService) RegisterAccount(ctx context.Context, req *v1.RegisterAccountRequest) (*v1.Reply, error) {
	// 创建 UserAccount 实例
	account := &v1.UserAccount{
		Id:              0,
		Usr:             "",
		Pwd:             "",
		SecPwd:          "",
		Mobile:          "",
		Email:           "",
		Lang:            "",
		Timezone:        "",
		RealName:        "",
		IdNo:            "",
		Currency:        "",
		IsGoogleAuth:    false,
		GoogleAuth:      "",
		DeviceId:        "",
		IsAllowedMsg:    false,
		IsAllowedEmail:  false,
		MsgType:         0,
		IsRealAuth:      false,
		IsAddedBankcard: false,
		SecurityLevel:   0,
		IsUnreadNotice:  false,
		IsCoupon:        false,
		Status:          0,
		CreatedAt:       "",
		UpdatedAt:       "",
		DeletedAt:       "",
	}

	anyData, err := anypb.New(account)
	if err != nil {
		return nil, err
	}

	return &v1.Reply{Code: 0, Msg: "", Data: anyData}, v1.ErrorUserNotFound("user %s not found", "kratos", "withmetadata")
}

// GetAccount 读取账户
func (s *UserService) GetAccount(ctx context.Context, req *v1.GetAccountRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// UpdateAccount 更新账户
func (s *UserService) UpdateAccount(ctx context.Context, req *v1.UpdateAccountRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// ListAddresses 地址簿列表
func (s *UserService) ListAddresses(ctx context.Context, req *empty.Empty) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// GetAddress 读取单条地址
func (s *UserService) GetAddress(ctx context.Context, req *v1.GetAddressRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// AddAddress 添加地址
func (s *UserService) AddAddress(ctx context.Context, req *v1.AddAddressRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// UpdateAddress 更新地址
func (s *UserService) UpdateAddress(ctx context.Context, req *v1.UpdateAddressRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// DeleteAddress 删除地址
func (s *UserService) DeleteAddress(ctx context.Context, req *v1.DeleteAddressRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// ListReceivingMethods 收款方式列表
func (s *UserService) ListReceivingMethods(ctx context.Context, req *empty.Empty) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// GetReceivingMethod 获取单个收款方式
func (s *UserService) GetReceivingMethod(ctx context.Context, req *v1.GetReceivingMethodRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// AddReceivingMethod 添加收款方式
func (s *UserService) AddReceivingMethod(ctx context.Context, req *v1.AddReceivingMethodRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// UpdateReceivingMethod 更新收款方式
func (s *UserService) UpdateReceivingMethod(ctx context.Context, req *v1.UpdateReceivingMethodRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}

// DeleteReceivingMethod 删除收款方式
func (s *UserService) DeleteReceivingMethod(ctx context.Context, req *v1.DeleteReceivingMethodRequest) (*v1.Reply, error) {
	// 调用 gRPC 服务，转发请求
	// 这里填充转发逻辑
	return &v1.Reply{Code: 0, Msg: "", Data: nil}, nil
}
