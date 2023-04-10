package main

import (
    "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
    "google.golang.org/grpc"
    pb "grpc-gw/gen/examplepb"
    "net/http"
    "context"
)

func main() {


    mux := runtime.NewServeMux()
    err := pb.RegisterBookServiceHandlerFromEndpoint(context.Background(), mux, ":8083", []grpc.DialOption{grpc.WithInsecure()})
    if err != nil {
        panic(err.Error())
    }
    http.ListenAndServe(":8085", mux)
}
