package article

import (
	"google.golang.org/grpc"

	pb "github.com/tanimutomo/grpcapi-go-server/pkg/grpcs/article"
)

func SetHandler(s *grpc.Server) {
	pb.RegisterArticleServiceServer(s, &articleHandler{})
}
