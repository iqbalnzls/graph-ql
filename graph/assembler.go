package graph

import "github.com/iqbalnzls/graph-ql/graph/model"

func toCharacters(m map[string]*model.CharacterResponse) []*model.CharacterResponse {
	characters := make([]*model.CharacterResponse, 0)

	for _, v := range m {
		characters = append(characters, v)
	}

	return characters
}
