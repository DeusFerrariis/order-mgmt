package internal

type (
	ItemData struct {
		Sku         string `json:"sku"`
		Description string `json:"description"`
	}
	ItemPricing struct {
		Cost  int64 `json:"cost"`
		Price int64 `json:"price"`
	}
	LineItemData struct {
		ItemData
		ItemPricing
		LineNumber int64 `json:"line_number"`
		Quantity   int64 `json:"quantity"`
		OrderId    int64 `json:"order_id"`
	}
	LineItemRecord struct {
		LineItemData
		Id int64 `json:"id"`
	}
	OrderRecord struct {
		Id         int64 `json:"id"`
		CustomerId int64 `json:"customer_id"`
	}
)

const (
	OrderTable  = "orders"
	OrderSchema = `
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER NOT NULL UNIQUE PRIMARY KEY,
			customer_id INTEGER NOT NULL,
		);
	`
	LineItemTable  = "lineItems"
	LineItemSchema = `
		CREATE TABLE IF NOT EXISTS lineItems (
			id INTEGER NOT NULL UNIQUE PRIMARY KEY,
			order_id INTEGER NOT NULL,
			line_number INTEGER NOT NULL,
			sku TEXT NOT NULL,
			description TEXT NOT NULL,
			cost INTEGER NOT NULL,
			price INTEGER NOT NULL,
			quantity INTEGER NOT NULL
		);
	`
)
