package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

// User is the resolver for the user field.
func (r *creditTransactionResolver) User(ctx context.Context, obj *model.CreditTransaction) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Recipient is the resolver for the recipient field.
func (r *creditTransactionResolver) Recipient(ctx context.Context, obj *model.CreditTransaction) (*model.User, error) {
	return r.Service.GetUserByID(ctx, obj.RecipientID)
}

// Info is the resolver for the info field.
func (r *creditTransactionResolver) Info(ctx context.Context, obj *model.CreditTransaction) (*model.CreditTransactionInfo, error) {
	return r.Service.GetCreditTransactionInfoByID(ctx, &obj.CreditTransactionInfoID)
}

// RelevantTransaction is the resolver for the relevantTransaction field.
func (r *creditTransactionResolver) RelevantTransaction(ctx context.Context, obj *model.CreditTransaction) (*model.CreditTransaction, error) {
	return r.Service.GetCreditTransactionByID(ctx, obj.RelevantCreditTransactionID)
}

// Content is the resolver for the content field.
func (r *creditTransactionDescriptionVariableResolver) Content(ctx context.Context, obj *model.CreditTransactionDescriptionVariable) (string, error) {
	descriptionVariable, err := r.Service.GetDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Content, err
}

// URL is the resolver for the url field.
func (r *creditTransactionDescriptionVariableResolver) URL(ctx context.Context, obj *model.CreditTransactionDescriptionVariable) (*string, error) {
	descriptionVariable, err := r.Service.GetDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Url, err
}

// Image is the resolver for the image field.
func (r *creditTransactionDescriptionVariableResolver) Image(ctx context.Context, obj *model.CreditTransactionDescriptionVariable) (*string, error) {
	descriptionVariable, err := r.Service.GetDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Image, err
}

// DescriptionVariables is the resolver for the descriptionVariables field.
func (r *creditTransactionInfoResolver) DescriptionVariables(ctx context.Context, obj *model.CreditTransactionInfo) ([]*model.CreditTransactionDescriptionVariable, error) {
	return r.Service.GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx, &obj.ID)
}

// Template is the resolver for the template field.
func (r *creditTransactionInfoResolver) Template(ctx context.Context, obj *model.CreditTransactionInfo) (*model.CreditTransactionTemplate, error) {
	return r.Service.GetCreditTransactionTemplateByID(ctx, obj.CreditTransactionTemplateID)
}

// Cursor is the resolver for the cursor field.
func (r *creditTransactionsEdgeResolver) Cursor(ctx context.Context, obj *model.CreditTransactionsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(*obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *creditTransactionsEdgeResolver) Node(ctx context.Context, obj *model.CreditTransactionsEdge) (*model.CreditTransaction, error) {
	return r.Service.GetCreditTransactionByID(ctx, &obj.Node.ID)
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

// CreditTransactionInfo returns generated.CreditTransactionInfoResolver implementation.
func (r *Resolver) CreditTransactionInfo() generated.CreditTransactionInfoResolver {
	return &creditTransactionInfoResolver{r}
}

// CreditTransactionsEdge returns generated.CreditTransactionsEdgeResolver implementation.
func (r *Resolver) CreditTransactionsEdge() generated.CreditTransactionsEdgeResolver {
	return &creditTransactionsEdgeResolver{r}
}

type creditTransactionResolver struct{ *Resolver }
type creditTransactionDescriptionVariableResolver struct{ *Resolver }
type creditTransactionInfoResolver struct{ *Resolver }
type creditTransactionsEdgeResolver struct{ *Resolver }