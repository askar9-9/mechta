package entity

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestOrderItem_Validate(t *testing.T) {
	validUUID := uuid.New()

	tests := []struct {
		name    string
		item    OrderItem
		wantErr error
	}{
		{
			name: "valid item",
			item: OrderItem{
				ProductID:     "123",
				ExternalID:    "ext-456",
				Status:        "confirmed",
				BasePrice:     100.0,
				Price:         90.0,
				EarnedBonuses: 10.0,
				SpentBonuses:  5.0,
				Gift:          false,
				OwnerID:       "owner-1",
				DeliveryID:    "del-1",
				ShopAssistant: "assistant-1",
				Warehouse:     "wh-1",
				OrderID:       validUUID,
			},
			wantErr: nil,
		},
		{
			name:    "missing ProductID",
			item:    OrderItem{ExternalID: "ext", Status: "ok", BasePrice: 1, Price: 1, OwnerID: "owner", Warehouse: "wh", OrderID: validUUID},
			wantErr: ErrProductIDRequired,
		},
		{
			name:    "missing ExternalID",
			item:    OrderItem{ProductID: "123", Status: "ok", BasePrice: 1, Price: 1, OwnerID: "owner", Warehouse: "wh", OrderID: validUUID},
			wantErr: ErrExternalIDRequired,
		},
		{
			name:    "missing Status",
			item:    OrderItem{ProductID: "123", ExternalID: "ext", BasePrice: 1, Price: 1, OwnerID: "owner", Warehouse: "wh", OrderID: validUUID},
			wantErr: ErrInvalidStatus,
		},
		{
			name:    "negative BasePrice",
			item:    OrderItem{ProductID: "123", ExternalID: "ext", Status: "ok", BasePrice: -1, Price: 1, OwnerID: "owner", Warehouse: "wh", OrderID: validUUID},
			wantErr: ErrBasePriceInvalid,
		},
		{
			name:    "negative Price",
			item:    OrderItem{ProductID: "123", ExternalID: "ext", Status: "ok", BasePrice: 1, Price: -1, OwnerID: "owner", Warehouse: "wh", OrderID: validUUID},
			wantErr: ErrPriceInvalid,
		},
		{
			name:    "negative EarnedBonuses",
			item:    OrderItem{ProductID: "123", ExternalID: "ext", Status: "ok", BasePrice: 1, Price: 1, EarnedBonuses: -1, OwnerID: "owner", Warehouse: "wh", OrderID: validUUID},
			wantErr: ErrBonusesInvalid,
		},
		{
			name:    "negative SpentBonuses",
			item:    OrderItem{ProductID: "123", ExternalID: "ext", Status: "ok", BasePrice: 1, Price: 1, SpentBonuses: -1, OwnerID: "owner", Warehouse: "wh", OrderID: validUUID},
			wantErr: ErrBonusesInvalid,
		},
		{
			name:    "missing OwnerID",
			item:    OrderItem{ProductID: "123", ExternalID: "ext", Status: "ok", BasePrice: 1, Price: 1, Warehouse: "wh", OrderID: validUUID},
			wantErr: ErrOwnerIDRequired,
		},
		{
			name:    "missing Warehouse",
			item:    OrderItem{ProductID: "123", ExternalID: "ext", Status: "ok", BasePrice: 1, Price: 1, OwnerID: "owner", OrderID: validUUID},
			wantErr: ErrWarehouseRequired,
		},
		{
			name:    "invalid OrderID",
			item:    OrderItem{ProductID: "123", ExternalID: "ext", Status: "ok", BasePrice: 1, Price: 1, OwnerID: "owner", Warehouse: "wh", OrderID: uuid.Nil},
			wantErr: ErrOrderIDInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.item.Validate()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}
		})
	}
}
