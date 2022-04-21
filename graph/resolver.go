package graph

import "LearningGoGraphQL/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	FilmStore map[string]model.Film
}
