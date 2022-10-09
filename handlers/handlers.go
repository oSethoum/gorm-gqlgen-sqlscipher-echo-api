package handlers

import (
	"gormql/db"
	"gormql/graph/generated"
	"gormql/graph/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func GqlHandler(c echo.Context) error {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{
		DB: db.DB,
	}}))
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}

func GqlPlaygroundHandler(c echo.Context) error {
	h := playground.Handler("GraphQL playground", "/query")
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}
