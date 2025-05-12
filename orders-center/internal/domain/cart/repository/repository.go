package repository

import (
	"context"
	"errors"
	"orders-center/internal/domain/cart/entity"
	"orders-center/internal/infrastructure/db/pgxtx"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) FindItemsByOrderID(ctx context.Context, orderID string) ([]*entity.OrderItem, error) {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return nil, pgxtx.ErrNoTx
	}

	rows, err := q.Query(ctx, `
		SELECT 
			product_id,
			external_id,
			status,
			base_price,
			price,
			earned_bonuses,
			spent_bonuses,
			gift,
			owner_id,
			delivery_id,
			shop_assistant,
			warehouse,
			order_id
		FROM order_items 
		WHERE order_id = $1`, orderID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, entity.ErrOrderItemsNotFound
		}
		return nil, err
	}

	defer rows.Close()
	var items []*entity.OrderItem
	for rows.Next() {
		var item entity.OrderItem
		if err := rows.Scan(
			&item.ProductID,
			&item.ExternalID,
			&item.Status,
			&item.BasePrice,
			&item.Price,
			&item.EarnedBonuses,
			&item.SpentBonuses,
			&item.Gift,
			&item.OwnerID,
			&item.DeliveryID,
			&item.ShopAssistant,
			&item.Warehouse,
			&item.OrderID,
		); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repo) AddItemsToOrder(ctx context.Context, items []*entity.OrderItem) error {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return pgxtx.ErrNoTx
	}

	for _, item := range items {
		_, err := q.Exec(ctx, `
			INSERT INTO order_items (
				product_id,
				external_id,
				status,
				base_price,
				price,
				earned_bonuses,
				spent_bonuses,
				gift,
				owner_id,
				delivery_id,
				shop_assistant,
				warehouse,
				order_id
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
			item.ProductID,
			item.ExternalID,
			item.Status,
			item.BasePrice,
			item.Price,
			item.EarnedBonuses,
			item.SpentBonuses,
			item.Gift,
			item.OwnerID,
			item.DeliveryID,
			item.ShopAssistant,
			item.Warehouse,
			item.OrderID,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
