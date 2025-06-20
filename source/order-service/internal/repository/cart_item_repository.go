package repository

import (
	"context"
	"fmt"
	"thanhldt060802/infrastructure"
	"thanhldt060802/internal/model"
	"thanhldt060802/utils"

	"github.com/google/uuid"
)

type cartItemRepository struct {
}

type CartItemRepository interface {
	// Main features
	Get(ctx context.Context, offset int, limit int, sortFields []utils.SortField) ([]model.CartItem, error)
	GetByUserId(ctx context.Context, userId string, offset int, limit int, sortFields []utils.SortField) ([]model.CartItem, error)
	GetByIdAndUserId(ctx context.Context, id string, userId string) (*model.CartItem, error)
	Create(ctx context.Context, newCartItem *model.CartItem) error
	Update(ctx context.Context, updatedCartItem *model.CartItem) error
	DeleteByIdAndUserId(ctx context.Context, id string, userId string) error
}

func NewCartItemRepository() CartItemRepository {
	return &cartItemRepository{}
}

//
//
// Main features
// ######################################################################################

func (cartItemRepository *cartItemRepository) Get(ctx context.Context, offset int, limit int, sortFields []utils.SortField) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	query := infrastructure.PostgresDB.NewSelect().Model(&cartItems).
		Offset(offset).
		Limit(limit)

	for _, sortField := range sortFields {
		query = query.Order(fmt.Sprintf("%s %s", sortField.Field, sortField.Direction))
	}

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (cartItemRepository *cartItemRepository) GetByUserId(ctx context.Context, userId string, offset int, limit int, sortFields []utils.SortField) ([]model.CartItem, error) {
	var cartItems []model.CartItem

	query := infrastructure.PostgresDB.NewSelect().Model(&cartItems).Where("user_id = ?", userId).
		Offset(offset).
		Limit(limit)

	for _, sortField := range sortFields {
		query = query.Order(fmt.Sprintf("%s %s", sortField.Field, sortField.Direction))
	}

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (cartItemRepository *cartItemRepository) GetByIdAndUserId(ctx context.Context, id string, userId string) (*model.CartItem, error) {
	var cartItem model.CartItem

	if err := infrastructure.PostgresDB.NewSelect().Model(&cartItem).Where("id = ? AND user_id = ?", id, userId).Scan(ctx); err != nil {
		return nil, err
	}

	return &cartItem, nil
}

func (cartItemRepository *cartItemRepository) Create(ctx context.Context, newCartItem *model.CartItem) error {
	newCartItem.Id = uuid.New().String()
	_, err := infrastructure.PostgresDB.NewInsert().Model(newCartItem).Exec(ctx)
	return err
}

func (cartItemRepository *cartItemRepository) Update(ctx context.Context, updatedCartItem *model.CartItem) error {
	_, err := infrastructure.PostgresDB.NewUpdate().Model(updatedCartItem).Where("id = ?", updatedCartItem.Id).Exec(ctx)
	return err
}

func (cartItemRepository *cartItemRepository) DeleteByIdAndUserId(ctx context.Context, id string, userId string) error {
	_, err := infrastructure.PostgresDB.NewDelete().Model(&model.CartItem{}).Where("id = ? AND user_id = ?", id, userId).Exec(ctx)
	return err
}
