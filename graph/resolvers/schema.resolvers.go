package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gormql/graph/generated"
	"gormql/graph/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.CreateUserInput) (*models.User, error) {
	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	result := r.DB.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, result.Error
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id uint, input models.UpdateUserInput) (*models.User, error) {
	user := models.User{}

	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if input.Username != nil {
		user.Username = *input.Username
	}

	if input.Password != nil {
		user.Password = *input.Password
	}

	result = r.DB.Save(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id uint) (*models.User, error) {
	user := models.User{}

	result := r.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.DB.Delete(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input models.CreateTodoInput) (*models.Todo, error) {
	user := models.User{}

	result := r.DB.First(&user, input.UserID)

	if result.Error != nil {
		return nil, result.Error
	}

	todo := models.Todo{
		Text:   input.Text,
		Done:   input.Done,
		User:   user,
		UserID: input.UserID,
	}

	result = r.DB.Create(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id uint, input models.UpdateTodoInput) (*models.Todo, error) {
	user := models.User{}

	result := r.DB.First(&user, *input.UserID)

	if result.Error != nil {
		return nil, result.Error
	}

	todo := models.Todo{}

	result = r.DB.Create(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, id uint) (*models.Todo, error) {
	todo := models.Todo{}

	result := r.DB.First(&todo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.DB.Delete(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id uint) (*models.User, error) {
	user := models.User{}
	result := r.DB.First(&user, id)
	return &user, result.Error
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	users := []*models.User{}
	result := r.DB.Find(&users)
	return users, result.Error
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id uint) (*models.Todo, error) {
	todo := models.Todo{}
	result := r.DB.First(&todo, id)
	return &todo, result.Error
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	todos := []*models.Todo{}
	result := r.DB.Preload("User").Find(&todos)
	return todos, result.Error
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
