package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.27

import (
	"context"

	"github.com/plogto/core/db"
	"github.com/plogto/core/graph/generated"
	"github.com/plogto/core/graph/model"
	"github.com/plogto/core/util"
)

// User is the resolver for the user field.
func (r *creditTransactionResolver) User(ctx context.Context, obj *db.CreditTransaction) (*db.User, error) {
	return r.Service.GetUserByID(ctx, obj.UserID)
}

// Recipient is the resolver for the recipient field.
func (r *creditTransactionResolver) Recipient(ctx context.Context, obj *db.CreditTransaction) (*db.User, error) {
	return r.Service.GetUserByID(ctx, obj.RecipientID)
}

// Amount is the resolver for the amount field.
func (r *creditTransactionResolver) Amount(ctx context.Context, obj *db.CreditTransaction) (float64, error) {
	return float64(obj.Amount), nil
}

// Type is the resolver for the type field.
func (r *creditTransactionResolver) Type(ctx context.Context, obj *db.CreditTransaction) (*model.CreditTransactionType, error) {
	return (*model.CreditTransactionType)(&obj.Type), nil
}

// Info is the resolver for the info field.
func (r *creditTransactionResolver) Info(ctx context.Context, obj *db.CreditTransaction) (*db.CreditTransactionInfo, error) {
	if &obj.CreditTransactionInfoID == nil {
		return nil, nil
	} else {
		return r.Service.GetCreditTransactionInfoByID(ctx, obj.CreditTransactionInfoID)
	}
}

// RelevantTransaction is the resolver for the relevantTransaction field.
func (r *creditTransactionResolver) RelevantTransaction(ctx context.Context, obj *db.CreditTransaction) (*db.CreditTransaction, error) {
	if &obj.RelevantCreditTransactionID == nil {
		return nil, nil
	} else {
		return r.Service.GetCreditTransactionByID(ctx, obj.RelevantCreditTransactionID.UUID)
	}
}

// Type is the resolver for the type field.
func (r *creditTransactionDescriptionVariableResolver) Type(ctx context.Context, obj *db.CreditTransactionDescriptionVariable) (model.CreditTransactionDescriptionVariableType, error) {
	return model.CreditTransactionDescriptionVariableType(obj.Type), nil
}

// Key is the resolver for the key field.
func (r *creditTransactionDescriptionVariableResolver) Key(ctx context.Context, obj *db.CreditTransactionDescriptionVariable) (model.CreditTransactionDescriptionVariableKey, error) {
	return model.CreditTransactionDescriptionVariableKey(obj.Key), nil
}

// Content is the resolver for the content field.
func (r *creditTransactionDescriptionVariableResolver) Content(ctx context.Context, obj *db.CreditTransactionDescriptionVariable) (string, error) {
	descriptionVariable, err := r.Service.GetDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Content, err
}

// URL is the resolver for the url field.
func (r *creditTransactionDescriptionVariableResolver) URL(ctx context.Context, obj *db.CreditTransactionDescriptionVariable) (*string, error) {
	descriptionVariable, err := r.Service.GetDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Url, err
}

// Image is the resolver for the image field.
func (r *creditTransactionDescriptionVariableResolver) Image(ctx context.Context, obj *db.CreditTransactionDescriptionVariable) (*string, error) {
	descriptionVariable, err := r.Service.GetDescriptionVariableContentByTypeAndContentID(ctx, obj.Type, obj.ContentID)
	return descriptionVariable.Image, err
}

// Description is the resolver for the description field.
func (r *creditTransactionInfoResolver) Description(ctx context.Context, obj *db.CreditTransactionInfo) (*string, error) {
	return &obj.Description.String, nil
}

// DescriptionVariables is the resolver for the descriptionVariables field.
func (r *creditTransactionInfoResolver) DescriptionVariables(ctx context.Context, obj *db.CreditTransactionInfo) ([]*db.CreditTransactionDescriptionVariable, error) {
	return r.Service.GetCreditTransactionDescriptionVariablesByCreditTransactionInfoID(ctx, obj.ID)
}

// Status is the resolver for the status field.
func (r *creditTransactionInfoResolver) Status(ctx context.Context, obj *db.CreditTransactionInfo) (model.CreditTransactionStatus, error) {
	return model.CreditTransactionStatus(obj.Status), nil
}

// Template is the resolver for the template field.
func (r *creditTransactionInfoResolver) Template(ctx context.Context, obj *db.CreditTransactionInfo) (*db.CreditTransactionTemplate, error) {
	if !obj.CreditTransactionTemplateID.Valid {
		return nil, nil
	}

	return r.Service.GetCreditTransactionTemplateByID(ctx, obj.CreditTransactionTemplateID)
}

// Name is the resolver for the name field.
func (r *creditTransactionTemplateResolver) Name(ctx context.Context, obj *db.CreditTransactionTemplate) (model.CreditTransactionTemplateName, error) {
	return model.CreditTransactionTemplateName(obj.Name), nil
}

// Cursor is the resolver for the cursor field.
func (r *creditTransactionsEdgeResolver) Cursor(ctx context.Context, obj *model.CreditTransactionsEdge) (string, error) {
	return util.ConvertCreateAtToCursor(obj.Node.CreatedAt), nil
}

// Node is the resolver for the node field.
func (r *creditTransactionsEdgeResolver) Node(ctx context.Context, obj *model.CreditTransactionsEdge) (*db.CreditTransaction, error) {
	if &obj.Node.ID == nil {
		return nil, nil
	} else {
		return r.Service.GetCreditTransactionByID(ctx, obj.Node.ID)
	}
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
