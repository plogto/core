package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

// Sender is the resolver for the sender field.
func (r *creditTransactionResolver) Sender(ctx context.Context, obj *model.CreditTransaction) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.SenderID)
}

// Receiver is the resolver for the receiver field.
func (r *creditTransactionResolver) Receiver(ctx context.Context, obj *model.CreditTransaction) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.ReceiverID)
}

// DescriptionVariables is the resolver for the descriptionVariables field.
func (r *creditTransactionResolver) DescriptionVariables(ctx context.Context, obj *model.CreditTransaction) ([]*model.CreditTransactionDescriptionVariable, error) {
	return r.Service.GetCreditTransactionDescriptionVariablesByCreditTransactionID(ctx, &obj.ID)
}

// Type is the resolver for the type field.
func (r *creditTransactionResolver) Type(ctx context.Context, obj *model.CreditTransaction) (*model.CreditTransactionType, error) {
	return r.Service.GetCreditTransactionTypeByID(ctx, obj.CreditTransactionTypeID)
}

// Content is the resolver for the content field.
func (r *creditTransactionDescriptionVariableResolver) Content(ctx context.Context, obj *model.CreditTransactionDescriptionVariable) (string, error) {
	descriptionVariable, err := r.Service.GeDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Content, err
}

// URL is the resolver for the url field.
func (r *creditTransactionDescriptionVariableResolver) URL(ctx context.Context, obj *model.CreditTransactionDescriptionVariable) (*string, error) {
	descriptionVariable, err := r.Service.GeDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Url, err
}

// Image is the resolver for the image field.
func (r *creditTransactionDescriptionVariableResolver) Image(ctx context.Context, obj *model.CreditTransactionDescriptionVariable) (*string, error) {
	descriptionVariable, err := r.Service.GeDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Image, err
}

// Cursor is the resolver for the cursor field.
func (r *creditTransactionsEdgeResolver) Cursor(ctx context.Context, obj *model.CreditTransactionsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(*obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *creditTransactionsEdgeResolver) Node(ctx context.Context, obj *model.CreditTransactionsEdge) (*model.CreditTransaction, error) {
	return r.Service.GetCreditTransactionByID(ctx, &obj.Node.ID)
}

// CreateCreditTransaction is the resolver for the createCreditTransaction field.
func (r *mutationResolver) CreateCreditTransaction(ctx context.Context, input model.CreateCreditTransactionInput) (*model.CreditTransaction, error) {
	panic(fmt.Errorf("not implemented: CreateCreditTransaction - createCreditTransaction"))
}

// GetCreditTransactions is the resolver for the getCreditTransactions field.
func (r *queryResolver) GetCreditTransactions(ctx context.Context, pageInfoInput *model.PageInfoInput) (*model.CreditTransactions, error) {
	return r.Service.GetCreditTransactions(ctx, pageInfoInput)
}

// CreditTransaction returns generated.CreditTransactionResolver implementation.
func (r *Resolver) CreditTransaction() generated.CreditTransactionResolver {
	return &creditTransactionResolver{r}
}

// CreditTransactionDescriptionVariable returns generated.CreditTransactionDescriptionVariableResolver implementation.
func (r *Resolver) CreditTransactionDescriptionVariable() generated.CreditTransactionDescriptionVariableResolver {
	return &creditTransactionDescriptionVariableResolver{r}
}

// CreditTransactionsEdge returns generated.CreditTransactionsEdgeResolver implementation.
func (r *Resolver) CreditTransactionsEdge() generated.CreditTransactionsEdgeResolver {
	return &creditTransactionsEdgeResolver{r}
}

type creditTransactionResolver struct{ *Resolver }
type creditTransactionDescriptionVariableResolver struct{ *Resolver }
type creditTransactionsEdgeResolver struct{ *Resolver }
