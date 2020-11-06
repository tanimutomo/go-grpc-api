package article

import (
	"google.golang.org/grpc"

	pb "github.com/tanimutomo/grpcapi-go-server/pkg/proto/article"
)

func SetHandler(s *grpc.Server) {
	pb.RegisterArticleServer(s, &articleHandler{})
}
