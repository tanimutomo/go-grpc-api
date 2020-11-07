package article

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tanimutomo/grpcapi-go-server/pkg/db"
	pb "github.com/tanimutomo/grpcapi-go-server/pkg/grpcs/article"
)

type handler struct {
	articleDB db.ArticleHandler
	pb.UnimplementedArticleServiceServer
}

func SetHandler(s *grpc.Server) {
	h := &handler{
		articleDB: db.NewArticleHandler(),
	}
	pb.RegisterArticleServiceServer(s, h)
}

func (s handler) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (res *pb.GetArticleResponse, err error) {
	art, err := s.articleDB.Find(uint64(req.GetId()))
	if err != nil {
		return res, status.Errorf(codes.NotFound, err.Error())
	}

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

func (s handler) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (res *pb.ListArticlesResponse, err error) {
	arts, _ := s.articleDB.FindAll()
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

func (s handler) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (res *pb.CreateArticleResponse, err error) {
	art := db.Article{
		Title: req.GetTitle(),
	}
	art, _ = s.articleDB.Create(art)

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
