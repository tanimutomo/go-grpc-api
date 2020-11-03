package article

import (
	"google.golang.org/grpc"

	pb "github.com/tanimutomo/go-grpc-api/pkg/proto/article"
)

func SetHandler(s *grpc.Server) {
	pb.RegisterArticleServer(s, &articleHandler{})
}
