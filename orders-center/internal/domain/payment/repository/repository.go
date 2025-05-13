package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"orders-center/internal/domain/payment/entity"
	"orders-center/internal/infrastructure/db/pgxtx"
)

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) GetOrderPaymentsByOrderID(ctx context.Context, orderID uuid.UUID) ([]*entity.OrderPayment, error) {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return nil, pgxtx.ErrNoTx
	}

	query := `
		SELECT id, order_id, type, sum, payed, info, contract_number, external_id
		FROM order_payments
		WHERE order_id = $1
	`

	rows, err := q.Query(ctx, query, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*entity.OrderPayment
	for rows.Next() {
		var payment entity.OrderPayment
		err := rows.Scan(
			&payment.ID,
			&payment.OrderID,
			&payment.Type,
			&payment.Sum,
			&payment.Payed,
			&payment.Info,
			&payment.ContractNumber,
			&payment.ExternalID,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return payments, nil
}

func (r *Repo) CreateOrderPayments(ctx context.Context, payments []*entity.OrderPayment) error {
	q, ok := pgxtx.GetTx(ctx)
	if !ok {
		return pgxtx.ErrNoTx
	}

	paymentQuery := `
		INSERT INTO order_payments (id, order_id, type, sum, payed, info, contract_number, external_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	creditDataQuery := `
		INSERT INTO credit_data (order_payment_id, bank, type, number_of_months, pay_sum_per_month, broker_id, iin)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	cardPaymentDataQuery := `
		INSERT INTO card_payment_data (order_payment_id, provider, transaction_id)
		VALUES ($1, $2, $3)
	`

	for _, payment := range payments {
		_, err := q.Exec(ctx, paymentQuery,
			payment.ID,
			payment.OrderID,
			payment.Type,
			payment.Sum,
			payment.Payed,
			payment.Info,
			payment.ContractNumber,
			payment.ExternalID,
		)
		if err != nil {
			return err
		}

		if payment.Type == entity.PaymentTypeCredit && payment.CreditData != nil {
			_, err := q.Exec(ctx, creditDataQuery,
				payment.ID,
				payment.CreditData.Bank,
				payment.CreditData.Type,
				payment.CreditData.NumberOfMonths,
				payment.CreditData.PaySumPerMonth,
				payment.CreditData.BrokerID,
				payment.CreditData.IIN,
			)
			if err != nil {
				return err
			}
		}

		if (payment.Type == entity.PaymentTypeCard || payment.Type == entity.PaymentTypeCardOnline) && payment.CardPaymentData != nil {
			_, err := q.Exec(ctx, cardPaymentDataQuery,
				payment.ID,
				payment.CardPaymentData.Provider,
				payment.CardPaymentData.TransactionId,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
