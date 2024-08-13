package cart

import (
	"fmt"

	"github.com/ahenla/go-ecom/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIds := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %d", item.ProductID)
		}

		productIds[i] = item.ProductID
	}

	return productIds, nil
}
