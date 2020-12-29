package rpcs

import (
	"context"
	"github.com/galehuang/message-service-project/api/business"
)

var _ business.MessageServiceServer = (*MessageServiceServerImpl)(nil)

type MessageServiceServerImpl struct {
}

func (m MessageServiceServerImpl) SenderCreate(ctx context.Context, req *business.Email_SenderCreateReq) (*business.CommonRsp, error) {
	panic("implement me")
}

func (m MessageServiceServerImpl) ReceiverCreate(ctx context.Context, req *business.Email_ReceiverCreateReq) (*business.CommonRsp, error) {
	panic("implement me")
}

func (m MessageServiceServerImpl) EmailCreate(ctx context.Context, req *business.Email_EmailCreateReq) (*business.CommonRsp, error) {
	panic("implement me")
}

func (m MessageServiceServerImpl) EmailUpdate(ctx context.Context, req *business.Email_EmailUpdateReq) (*business.CommonRsp, error) {
	panic("implement me")
}

func (m MessageServiceServerImpl) EmailDelete(ctx context.Context, req *business.Email_EmailDeleteReq) (*business.CommonRsp, error) {
	panic("implement me")
}

func (m MessageServiceServerImpl) EmailList(ctx context.Context, req *business.Email_EmailListReq) (*business.Email_EmailListRsp, error) {
	panic("implement me")
}

func (m MessageServiceServerImpl) EmailSend(ctx context.Context, req *business.Email_EmailSendReq) (*business.CommonRsp, error) {
	panic("implement me")
}

