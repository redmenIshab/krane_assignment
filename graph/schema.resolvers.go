package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"eventManagemntSystem/model"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, user *model.UserInput) (*model.User, error) {
	res, err := r.UserRepo.CreateUser(*user)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id *string, user *model.UserInput) (*model.User, error) {
	res, err := r.UserRepo.UpdateUser(id, *user)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, userID string, event *model.EventInput) (*model.Event, error) {
	_, err := r.UserRepo.GetUserByID(userID)
	if err != nil {
		panic(fmt.Errorf("Invalid User id"))
	}
	res, err := r.EventRepo.CreateEvent(userID, *event)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateEvent is the resolver for the updateEvent field.
func (r *mutationResolver) UpdateEvent(ctx context.Context, id *string, userID *string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented: UpdateEvent - updateEvent"))
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, id *string, userID *string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteEvent - deleteEvent"))
}

// AddUserToEvent is the resolver for the addUserToEvent field.
func (r *mutationResolver) AddUserToEvent(ctx context.Context, managerID string, userEvent model.UserEventInput) (*model.UserEvent, error) {
	_, err := r.UserRepo.GetUserByID(userEvent.UserID)
	if err != nil {
		panic(fmt.Errorf("invalid User id"))
	}
	fmt.Println("verified user id")
	// Verify for sufficient permission
	userEventRes, err := r.EventRepo.GetUserEventByUserAndEventID(managerID, userEvent.EventID)
	fmt.Printf("userEventRes %+v", userEventRes)
	if userEventRes.Role != model.UserRoleAdmin && userEventRes.Role != model.UserRoleManager {
		panic(fmt.Errorf("role insuffiecient"))
	}
	if err != nil {
		panic(fmt.Errorf("invalid user or event id"))
	}
	fmt.Println("user verified ---", userEventRes)
	res, err := r.EventRepo.AddUserToEvent(managerID, userEvent)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UpdateUserEvent is the resolver for the updateUserEvent field.
func (r *mutationResolver) UpdateUserEvent(ctx context.Context, managerID string, userEvent model.UserEventInput) (*model.UserEvent, error) {
	_, err := r.UserRepo.GetUserByID(userEvent.UserID)
	if err != nil {
		panic(fmt.Errorf("invalid User id"))
	}
	fmt.Println("verified user id")
	// Verify for sufficient permission
	userEventRes, err := r.EventRepo.GetUserEventByUserAndEventID(managerID, userEvent.EventID)
	fmt.Printf("userEventRes %+v", userEventRes)
	if userEventRes.Role != model.UserRoleAdmin && userEventRes.Role != model.UserRoleManager {
		panic(fmt.Errorf("role insuffiecient"))
	}
	if err != nil {
		panic(fmt.Errorf("invalid user or event id"))
	}
	fmt.Println("user verified ---", userEventRes)
	res, err := r.EventRepo.UpdateUserEvent(managerID, userEvent)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// CreateExpense is the resolver for the createExpense field.
func (r *mutationResolver) CreateExpense(ctx context.Context, adminID string, expense model.ExpenseInput) (*model.Expense, error) {
	_, err := r.UserRepo.GetUserByID(adminID)
	if err != nil {
		panic(fmt.Errorf("invalid User id"))
	}
	fmt.Println("verified user id")
	// Verify for sufficient permission
	userEventRes, err := r.EventRepo.GetUserEventByUserAndEventID(adminID, expense.EventID)
	fmt.Printf("userEventRes %+v", userEventRes)
	if userEventRes.Role != model.UserRoleAdmin {
		panic(fmt.Errorf("role insuffiecient"))
	}
	if err != nil {
		panic(fmt.Errorf("invalid user or event id"))
	}
	fmt.Println("user verified ---", userEventRes)
	res, err := r.ExpenseRepo.CreateExpense(adminID, expense)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UserRepo.GetUserByID(id)
}

// Event is the resolver for the event field.
func (r *queryResolver) Event(ctx context.Context, id string) (*model.Event, error) {
	return r.EventRepo.GetEventById(id)
}

// Expense is the resolver for the expense field.
func (r *queryResolver) Expense(ctx context.Context, eventID string, userID string, expensetype model.ExpenseType) ([]*model.Expense, error) {
	panic(fmt.Errorf("not implemented: Expense - expense"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
