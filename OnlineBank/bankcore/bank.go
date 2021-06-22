package bankcore

type Customer struct {
	Name string
	Address string
	Phone string
}

type Account struct {
	Customer
	Number int32
	Balance float64
}

func Hello() string {
	return "Hey! I'm working!"
}
