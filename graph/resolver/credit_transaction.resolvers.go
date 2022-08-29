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

// User is the resolver for the user field.
func (r *creditTransactionResolver) User(ctx context.Context, obj *model.CreditTransaction) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Recipient is the resolver for the recipient field.
func (r *creditTransactionResolver) Recipient(ctx context.Context, obj *model.CreditTransaction) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Recipient - recipient"))
}

// Info is the resolver for the info field.
func (r *creditTransactionResolver) Info(ctx context.Context, obj *model.CreditTransaction) (*model.CreditTransactionInfo, error) {
	panic(fmt.Errorf("not implemented: Info - info"))
}

// RelevantTransaction is the resolver for the relevantTransaction field.
func (r *creditTransactionResolver) RelevantTransaction(ctx context.Context, obj *model.CreditTransaction) (*model.CreditTransaction, error) {
	panic(fmt.Errorf("not implemented: RelevantTransaction - relevantTransaction"))
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

// DescriptionVariables is the resolver for the descriptionVariables field.
func (r *creditTransactionInfoResolver) DescriptionVariables(ctx context.Context, obj *model.CreditTransactionInfo) ([]*model.CreditTransactionDescriptionVariable, error) {
	panic(fmt.Errorf("not implemented: DescriptionVariables - descriptionVariables"))
}

// Template is the resolver for the template field.
func (r *creditTransactionInfoResolver) Template(ctx context.Context, obj *model.CreditTransactionInfo) (*model.CreditTransactionTemplate, error) {
	panic(fmt.Errorf("not implemented: Template - template"))
}

// Content is the resolver for the content field.
func (r *creditTransactionTemplateResolver) Content(ctx context.Context, obj *model.CreditTransactionTemplate) (string, error) {
	panic(fmt.Errorf("not implemented: Content - content"))
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

// CreditTransactionInfo returns generated.CreditTransactionInfoResolver implementation.
func (r *Resolver) CreditTransactionInfo() generated.CreditTransactionInfoResolver {
	return &creditTransactionInfoResolver{r}
}

// CreditTransactionTemplate returns generated.CreditTransactionTemplateResolver implementation.
func (r *Resolver) CreditTransactionTemplate() generated.CreditTransactionTemplateResolver {
	return &creditTransactionTemplateResolver{r}
}

// CreditTransactionsEdge returns generated.CreditTransactionsEdgeResolver implementation.
func (r *Resolver) CreditTransactionsEdge() generated.CreditTransactionsEdgeResolver {
	return &creditTransactionsEdgeResolver{r}
}

type creditTransactionResolver struct{ *Resolver }
type creditTransactionDescriptionVariableResolver struct{ *Resolver }
type creditTransactionInfoResolver struct{ *Resolver }
type creditTransactionTemplateResolver struct{ *Resolver }
type creditTransactionsEdgeResolver struct{ *Resolver }
