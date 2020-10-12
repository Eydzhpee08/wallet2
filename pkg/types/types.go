package types


// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.).
type Money int64


// PaymentCategory представляет собой категория в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
type PaymentCategory string

// PaymentStatus представляет собой статус платежа.
type PaymentStatus string


// Предопределённые статусы платежей
const (
	PaymentStatusOK         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

const (
	CategoryFood = "food"
	CategoryIt   = "it"
	CategoryShop = "shop"
)


// Payment представляет информацию о платеже. 
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

// Phone представляет информацию о номере телефона пользователя.
type Phone string


// Account представляет информацию о счёте пользователя.
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

type Favorite struct {
	ID        string
	AccountID int64
	Name      string
	Amount    Money
	Category  PaymentCategory
}