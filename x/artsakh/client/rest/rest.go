package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
    // this line is used by starport scaffolding # 1
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers artsakh-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 3
    r.HandleFunc("/artsakh/posts/{id}", getPostHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/artsakh/posts", listPostHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 4
    r.HandleFunc("/artsakh/posts", createPostHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/artsakh/posts/{id}", updatePostHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/artsakh/posts/{id}", deletePostHandler(clientCtx)).Methods("POST")

}

