package domain

// CoffeeRepository = สัญญาของ "ตู้เก็บข้อมูลกาแฟ"
type CoffeeRepository interface {
	FindByID(id string) (*Coffee, error)
	FindAll() ([]Coffee, error)
}

// OrderRepository = สัญญาของ "ตู้เก็บออเดอร์"
type OrderRepository interface {
	Save(order *Order) error
	FindByID(id string) (*Order, error)
	FindAll() ([]Order, error)
}