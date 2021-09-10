package server

import (
	"context"
	"fmt"
	"net/http/httptest"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func strptr(v string) *string { return &v }
func intptr(v int) *int       { return &v }

var users = []*User{
	{
		ID: "1", Name: "Yours Truly", LuckyNumber: intptr(17),
		Hair: &Hair{Color: strptr("Black")},
	},
	{ID: "2", Name: "Raven", LuckyNumber: intptr(-1), Hair: nil},
}

var animals = []*Animal{
	{
		ID: "3", Name: "Fido", Species: SpeciesDog, Owner: userByID("1"),
		Hair: &BeingsHair{HasHair: true},
	},
	{
		ID: "4", Name: "Old One", Species: SpeciesCoelacanth, Owner: nil,
		Hair: &BeingsHair{HasHair: false},
	},
}

func userByID(id string) *User {
	for _, user := range users {
		if id == user.ID {
			return user
		}
	}
	return nil
}

func beingByID(id string) Being {
	for _, user := range users {
		if id == user.ID {
			return user
		}
	}
	for _, animal := range animals {
		if id == animal.ID {
			return animal
		}
	}
	return nil
}

func (r *queryResolver) Me(ctx context.Context) (*User, error) {
	return userByID("1"), nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	return userByID(id), nil
}

func (r *queryResolver) Being(ctx context.Context, id string) (Being, error) {
	return beingByID(id), nil
}

func (r *queryResolver) Beings(ctx context.Context, ids []string) ([]Being, error) {
	ret := make([]Being, len(ids))
	for i, id := range ids {
		ret[i] = beingByID(id)
	}
	return ret, nil
}

func (r *queryResolver) LotteryWinner(ctx context.Context, number int) (Lucky, error) {
	for _, user := range users {
		if user.LuckyNumber != nil && *user.LuckyNumber == number {
			return user, nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Fail(ctx context.Context) (*bool, error) {
	f := true
	return &f, fmt.Errorf("oh no")
}

func RunServer() *httptest.Server {
	gqlgenServer := handler.New(NewExecutableSchema(Config{Resolvers: &resolver{}}))
	gqlgenServer.AddTransport(transport.POST{})
	return httptest.NewServer(gqlgenServer)
}

type (
	resolver      struct{}
	queryResolver struct{}
)

func (r *resolver) Query() QueryResolver { return &queryResolver{} }

//go:generate go run github.com/99designs/gqlgen
