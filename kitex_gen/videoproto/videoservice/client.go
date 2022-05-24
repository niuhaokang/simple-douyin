// Code generated by Kitex v0.3.1. DO NOT EDIT.

package videoservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"simple-douyin/kitex_gen/videoproto"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateVideo(ctx context.Context, Req *videoproto.CreateVideoReq, callOptions ...callopt.Option) (r *videoproto.CreateVideoResp, err error)
	GetVideosByUserId(ctx context.Context, Req *videoproto.GetVideosByUserIdReq, callOptions ...callopt.Option) (r *videoproto.GetVideosByUserIdResp, err error)
	GetVideosByTime(ctx context.Context, Req *videoproto.GetVideosByTimeReq, callOptions ...callopt.Option) (r *videoproto.GetVideosByTimeResp, err error)
	LikeVideo(ctx context.Context, Req *videoproto.LikeVideoReq, callOptions ...callopt.Option) (r *videoproto.LikeVideoResp, err error)
	UnLikeVideo(ctx context.Context, Req *videoproto.UnLikeVideoReq, callOptions ...callopt.Option) (r *videoproto.UnLikeVideoResp, err error)
	CreateComment(ctx context.Context, Req *videoproto.CreateCommentReq, callOptions ...callopt.Option) (r *videoproto.CreateCommentResp, err error)
	DeleteComment(ctx context.Context, Req *videoproto.DeleteCommentReq, callOptions ...callopt.Option) (r *videoproto.DeleteCommentResp, err error)
	GetComments(ctx context.Context, Req *videoproto.GetCommentsReq, callOptions ...callopt.Option) (r *videoproto.GetCommentsResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kVideoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) CreateVideo(ctx context.Context, Req *videoproto.CreateVideoReq, callOptions ...callopt.Option) (r *videoproto.CreateVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateVideo(ctx, Req)
}

func (p *kVideoServiceClient) GetVideosByUserId(ctx context.Context, Req *videoproto.GetVideosByUserIdReq, callOptions ...callopt.Option) (r *videoproto.GetVideosByUserIdResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideosByUserId(ctx, Req)
}

func (p *kVideoServiceClient) GetVideosByTime(ctx context.Context, Req *videoproto.GetVideosByTimeReq, callOptions ...callopt.Option) (r *videoproto.GetVideosByTimeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetVideosByTime(ctx, Req)
}

func (p *kVideoServiceClient) LikeVideo(ctx context.Context, Req *videoproto.LikeVideoReq, callOptions ...callopt.Option) (r *videoproto.LikeVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LikeVideo(ctx, Req)
}

func (p *kVideoServiceClient) UnLikeVideo(ctx context.Context, Req *videoproto.UnLikeVideoReq, callOptions ...callopt.Option) (r *videoproto.UnLikeVideoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnLikeVideo(ctx, Req)
}

func (p *kVideoServiceClient) CreateComment(ctx context.Context, Req *videoproto.CreateCommentReq, callOptions ...callopt.Option) (r *videoproto.CreateCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateComment(ctx, Req)
}

func (p *kVideoServiceClient) DeleteComment(ctx context.Context, Req *videoproto.DeleteCommentReq, callOptions ...callopt.Option) (r *videoproto.DeleteCommentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteComment(ctx, Req)
}

func (p *kVideoServiceClient) GetComments(ctx context.Context, Req *videoproto.GetCommentsReq, callOptions ...callopt.Option) (r *videoproto.GetCommentsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetComments(ctx, Req)
}
