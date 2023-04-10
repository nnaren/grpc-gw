package main
import (
    "context"
    "fmt"
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    pb "grpc-gw/gen/examplepb"
    "net"
    "net/http"
)
type BookService struct {
    pb.UnimplementedBookServiceServer
}

func (b *BookService) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
    resp := &pb.CreateBookResponse{}

    book := pb.Book{
        Name: req.GetName(),
        Id:   3,
    }

    fmt.Printf(" create a book")
    resp.Data = &book

    return resp, nil
}

func (b *BookService) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
    resp := &pb.GetBookResponse{}

    book := pb.Book{
        Name: "都很好",
        Id:   3,
    }

    resp.Data = &book
    return resp, nil
}

func main() {
    lis, err := net.Listen("tcp", ":8083")
    if err != nil {
        panic(err.Error())
    }
    grpcServer := grpc.NewServer()

    pb.RegisterBookServiceServer(grpcServer, &BookService{})
    go grpcServer.Serve(lis)

    mux := runtime.NewServeMux()
    err = pb.RegisterBookServiceHandlerFromEndpoint(context.Background(), mux, ":8083", []grpc.DialOption{grpc.WithInsecure()})
    if err != nil {
        panic(err.Error())
    }
    http.ListenAndServe(":8084", mux)
}
