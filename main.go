package main

import (
	"net/http"

	inGrapghQL "github.com/iqbalnzls/graph-ql/internal/delivery/graph"
)

func main() {
	mux := http.NewServeMux()

	inGrapghQL.StartGraphQLServer(mux)
}
