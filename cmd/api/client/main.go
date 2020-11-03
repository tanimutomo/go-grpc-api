package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	article "github.com/tanimutomo/go-grpc-api/pkg/proto/article"
)

const host = "localhost:50051"

func main() {
	doArticleRequests()
}

func doArticleRequests() {
	fmt.Println("do articles")

	conn, err := grpc.Dial(
		host,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("error in connection")
		return
	}
	defer conn.Close()
	c := article.NewArticleClient(conn)

	if err := articleGetRequest(c, uint64(1)); err != nil {
		log.Fatalf("error in Get: %v\n", err)
		return
	}
	if err := articleListRequest(c); err != nil {
		log.Fatalf("error in List: %v\n", err)
		return
	}
	if err := articleCreateRequest(c, "title"); err != nil {
		log.Fatalf("error in Create: %v\n", err)
		return
	}

	fmt.Println("\nend articles")
}

func articleGetRequest(client article.ArticleClient, id uint64) error {
	fmt.Println("\ndo article/Get")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	req := article.GetRequest{
		Id: id,
	}
	res, err := client.Get(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "failed to receive response")
	}
	log.Printf("response: %+v\n", res.GetArticle())
	return nil
}

func articleListRequest(client article.ArticleClient) error {
	fmt.Println("\ndo article/List")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	req := article.ListRequest{}
	res, err := client.List(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "failed to receive response")
	}
	log.Printf("response: %+v\n", res.GetArticles())
	return nil
}

func articleCreateRequest(client article.ArticleClient, title string) error {
	fmt.Println("\ndo article/Create")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)
	defer cancel()

	req := article.CreateRequest{
		Title: title,
	}
	res, err := client.Create(ctx, &req)
	if err != nil {
		return errors.Wrap(err, "failed to receive response")
	}
	log.Printf("response: %+v\n", res.GetArticle())
	return nil
}
