package graph

import (
	"github.com/iqbalnzls/graph-ql/internal/delivery/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CharacterStore map[string]*model.CharacterResponse
}
