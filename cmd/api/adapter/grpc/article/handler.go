package article

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tanimutomo/go-grpc-api/pkg/data"
	pb "github.com/tanimutomo/go-grpc-api/pkg/proto/article"
)

type articleHandler struct {
	pb.UnimplementedArticleServer
}

func (s articleHandler) Get(ctx context.Context, req *pb.GetRequest) (res *pb.GetResponse, err error) {
	art := data.Articles[uint64(req.GetId())]

	cat, err := ptypes.TimestampProto(art.CreatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
	}
	uat, err := ptypes.TimestampProto(art.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "updated_at cannot be parsed to timestamp")
	}

	res = &pb.GetResponse{
		Article: &pb.ArticleResponse{
			Id:        art.ID,
			Title:     art.Title,
			CreatedAt: cat,
			UpdatedAt: uat,
		},
	}
	return res, nil
}

func (s articleHandler) List(ctx context.Context, req *pb.ListRequest) (res *pb.ListResponse, err error) {
	arts := data.Articles
	var resArts []*pb.ArticleResponse

	for _, art := range arts {
		cat, err := ptypes.TimestampProto(art.CreatedAt)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "created_at cannot be parsed to timestamp")
		}
		uat, err := ptypes.TimestampProto(art.UpdatedAt)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "updated_at cannot be parsed to timestamp")
		}

		r := &pb.ArticleResponse{
			Id:        art.ID,
			Title:     art.Title,
			CreatedAt: cat,
			UpdatedAt: uat,
		}
		resArts = append(resArts, r)
	}

	res = &pb.ListResponse{
		Articles: resArts,
	}
	return res, nil
}

func (s articleHandler) Create(ctx context.Context, req *pb.CreateRequest) (res *pb.CreateResponse, err error) {
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

	res = &pb.CreateResponse{
		Article: &pb.ArticleResponse{
			Id:        art.ID,
			Title:     art.Title,
			CreatedAt: cat,
			UpdatedAt: uat,
		},
	}
	return res, nil
}
