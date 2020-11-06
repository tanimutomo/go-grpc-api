package article

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tanimutomo/grpcapi-go-server/pkg/data"
	pb "github.com/tanimutomo/grpcapi-go-server/pkg/grpcs/article"
)

type articleHandler struct {
	pb.UnimplementedArticleServiceServer
}

func (s articleHandler) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (res *pb.GetArticleResponse, err error) {
	art := data.Articles[uint64(req.GetId())]

	cat, err := ptypes.TimestampProto(art.CreatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
	}
	uat, err := ptypes.TimestampProto(art.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "updated_at cannot be parsed to timestamp")
	}

	res = &pb.GetArticleResponse{
		Article: &pb.Article{
			Id:        art.ID,
			Title:     art.Title,
			CreatedAt: cat,
			UpdatedAt: uat,
		},
	}
	return res, nil
}

func (s articleHandler) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (res *pb.ListArticlesResponse, err error) {
	arts := data.Articles
	var resArts []*pb.Article

	for _, art := range arts {
		cat, err := ptypes.TimestampProto(art.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
		}
		uat, err := ptypes.TimestampProto(art.UpdatedAt)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "updated_at cannot be parsed to timestamp")
		}

		r := &pb.Article{
			Id:        art.ID,
			Title:     art.Title,
			CreatedAt: cat,
			UpdatedAt: uat,
		}
		resArts = append(resArts, r)
	}

	res = &pb.ListArticlesResponse{
		Articles: resArts,
	}
	return res, nil
}

func (s articleHandler) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (res *pb.CreateArticleResponse, err error) {
	id := uint64(len(data.Articles) + 1)
	art := data.Article{
		ID:        id,
		Title:     req.GetTitle(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data.Articles[id] = art

	cat, err := ptypes.TimestampProto(art.CreatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
	}
	uat, err := ptypes.TimestampProto(art.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "updated_at cannot be parsed to timestamp")
	}

	res = &pb.CreateArticleResponse{
		Article: &pb.Article{
			Id:        art.ID,
			Title:     art.Title,
			CreatedAt: cat,
			UpdatedAt: uat,
		},
	}
	return res, nil
}
