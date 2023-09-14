package main

import (
	"net/http"

	inGrapghQL "github.com/iqbalnzls/graph-ql/internal/delivery/graph"
)

const defaultPort = "8080"

func main() {
	mux := http.NewServeMux()

	inGrapghQL.StartGraphQLServer(mux)
}
