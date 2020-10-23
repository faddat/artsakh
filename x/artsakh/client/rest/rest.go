package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers artsakh-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/artsakh/add", createAddHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/artsakh/add", listAddHandler(cliCtx, "artsakh")).Methods("GET")
		r.HandleFunc("/artsakh/add/{key}", getAddHandler(cliCtx, "artsakh")).Methods("GET")
		r.HandleFunc("/artsakh/add", setAddHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/artsakh/add", deleteAddHandler(cliCtx)).Methods("DELETE")

		
}
