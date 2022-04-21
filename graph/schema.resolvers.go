package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"LearningGoGraphQL/graph/generated"
	"LearningGoGraphQL/graph/model"
	"context"
	"fmt"
	"strconv"
)

func (r *mutationResolver) UpsertFilm(ctx context.Context, input model.FilmInput) (*model.Film, error) {
	id := input.ID
	film := model.Film{Name: input.Name, Genre: input.Genre, Year: input.Year, Director: &model.Director{Name: input.Name}}
	length := len(r.FilmStore)
	if length == 0 {
		r.FilmStore = make(map[string]model.Film)
	}
	if id != nil {
		_, ok := r.FilmStore[*id]
		if !ok {
			return nil, fmt.Errorf("Film is not found")
		}
		r.FilmStore[*id] = film
	} else {
		nextId := strconv.Itoa(length + 1)
		film.ID = nextId
		r.FilmStore[nextId] = film
	}
	return &film, nil
}

func (r *queryResolver) Film(ctx context.Context, id string) (*model.Film, error) {
	if film, ok := r.FilmStore[id]; ok {
		return &film, nil
	} else {
		return nil, fmt.Errorf("Film is not found")
	}
}

func (r *queryResolver) Films(ctx context.Context, genre model.Genre) ([]*model.Film, error) {
	films := make([]*model.Film, 0)
	for idx := range r.FilmStore {
		film := r.FilmStore[idx]
		if film.Genre == genre {
			films = append(films, &film)
		}
	}
	return films, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
