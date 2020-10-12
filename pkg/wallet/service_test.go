package wallet

import (
	"reflect"
	"testing"
	"github.com/Muhammadkhon0/wallet/pkg/types"
	"github.com/google/uuid"
)

var defaultFavorite = types.Favorite{
	ID:        uuid.New().String(),
	AccountID: 1,
	Name:      types.CategoryIt,
	Amount:    10,
	Category:  types.CategoryIt,
}

type testService struct {
	*Service
}

func newTestService() *testService {
	return &testService{
		Service: &Service{},
	}
}

func TestService_FindAccountByID_success(t *testing.T) {
	var service Service
	service.RegisterAccount("9127660305")

	account, err := service.FindAccountByID(1)

	if err != nil {
		t.Errorf("account => %v", account)
	}

}

func TestService_FindAccountByID_notFound(t *testing.T) {
	var service Service
	service.RegisterAccount("9127660305")

	account, err := service.FindAccountByID(2)

	if err == nil {
		t.Errorf("method returned nil error, account => %v", account)
	}

}

func TestService_Reject_success_user(t *testing.T) {
	var service Service
	service.RegisterAccount("9127660305")
	account, err := service.FindAccountByID(1)

	if err != nil {
		t.Errorf("error => %v", err)
	}

	err = service.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("error => %v", err)
	}

	payment, err := service.Pay(account.ID, 10_00, "Food")

	if err != nil {
		t.Errorf("error => %v", err)
	}

	pay, err := service.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("error => %v", err)
	}

	err = service.Reject(pay.ID)

	if err != nil {
		t.Errorf("error => %v", err)
	}

}

func TestService_Reject_fail_user(t *testing.T) {
	var service Service
	service.RegisterAccount("9127660305")
	account, err := service.FindAccountByID(1)

	if err != nil {
		t.Errorf("account => %v", account)
	}

	err = service.Deposit(account.ID, 100_00)
	if err != nil {
		t.Errorf("error => %v", err)
	}

	payment, err := service.Pay(account.ID, 10_00, "Food")

	if err != nil {
		t.Errorf("account => %v", account)
	}

	pay, err := service.FindPaymentByID(payment.ID)

	if err != nil {
		t.Errorf("payment => %v", payment)
	}

	err = service.Reject(pay.ID + "uu")

	if err == nil {
		t.Errorf("pay => %v", pay)
	}

}

func TestService_RegisterAccount(t *testing.T) {
	type fields struct {
		nextAccountID int64
		accounts      []*types.Account
		payments      []*types.Payment
	}
	type args struct {
		phone types.Phone
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Account
		wantErr bool
	}{
		{
			name:   "user successfully registered.",
			fields: fields{},
			args:   args{phone: "9127660305"},
			want: &types.Account{
				ID:      1,
				Phone:   "9127660305",
				Balance: 0,
			},
			wantErr: false,
		},
		{
			name: "phone already registered",
			fields: fields{
				nextAccountID: 10,
				accounts:      Accounts(),
				payments:      nil,
			},
			args: args{
				phone: "9127660305",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				nextAccountID: tt.fields.nextAccountID,
				accounts:      tt.fields.accounts,
				payments:      tt.fields.payments,
			}
			got, err := s.RegisterAccount(tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RegisterAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Deposit(t *testing.T) {
	var accounts []*types.Account
	accounts = append(accounts,
		&types.Account{ID: 1, Phone: "9127660305", Balance: 0},
		&types.Account{ID: 3, Phone: "9127660307", Balance: 2},
		&types.Account{ID: 2, Phone: "9127660306", Balance: 1})

	type fields struct {
		nextAccountID int64
		accounts      []*types.Account
		payments      []*types.Payment
	}
	type args struct {
		accountID int64
		amount    types.Money
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "amount must be greater than zero",
			fields: fields{},
			args: args{
				accountID: 1,
				amount:    0,
			},
			wantErr: true,
		},
		{
			name: "account not found",
			fields: fields{
				nextAccountID: 0,
				accounts:      accounts,
				payments:      nil,
			},
			args: args{
				accountID: 4,
				amount:    10,
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				nextAccountID: 0,
				accounts:      accounts,
				payments:      nil,
			},
			args: args{
				accountID: 1,
				amount:    10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				nextAccountID: tt.fields.nextAccountID,
				accounts:      tt.fields.accounts,
				payments:      tt.fields.payments,
			}
			if err := s.Deposit(tt.args.accountID, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_FindAccountByID(t *testing.T) {
	type fields struct {
		nextAccountID int64
		accounts      []*types.Account
		payments      []*types.Payment
	}
	type args struct {
		accountID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				nextAccountID: tt.fields.nextAccountID,
				accounts:      tt.fields.accounts,
				payments:      tt.fields.payments,
			}
			got, err := s.FindAccountByID(tt.args.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAccountByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAccountByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_FindPaymentByID(t *testing.T) {
	type fields struct {
		nextAccountID int64
		accounts      []*types.Account
		payments      []*types.Payment
	}
	type args struct {
		paymentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Payment
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				nextAccountID: tt.fields.nextAccountID,
				accounts:      tt.fields.accounts,
				payments:      tt.fields.payments,
			}
			got, err := s.FindPaymentByID(tt.args.paymentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindPaymentByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPaymentByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Pay(t *testing.T) {
	type fields struct {
		nextAccountID int64
		accounts      []*types.Account
		payments      []*types.Payment
	}
	type args struct {
		accountID int64
		amount    types.Money
		category  types.PaymentCategory
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Payment
		wantErr bool
	}{
		{
			name:   "amount must be greater than zero",
			fields: fields{},
			args: args{
				accountID: 0,
				amount:    10,
				category:  types.CategoryFood,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "account not found",
			fields: fields{
				nextAccountID: 0,
				accounts:      Accounts(),
				payments:      nil,
			},
			args: args{
				accountID: 10,
				amount:    10,
				category:  types.CategoryIt,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "not enough balance in account",
			fields: fields{
				nextAccountID: 0,
				accounts:      Accounts(),
				payments:      nil,
			},
			args: args{
				accountID: 4,
				amount:    10,
				category:  types.CategoryFood,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				nextAccountID: tt.fields.nextAccountID,
				accounts:      tt.fields.accounts,
				payments:      tt.fields.payments,
			}
			got, err := s.Pay(tt.args.accountID, tt.args.amount, tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pay() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pay() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Reject(t *testing.T) {
	type fields struct {
		nextAccountID int64
		accounts      []*types.Account
		payments      []*types.Payment
	}
	type args struct {
		paymentID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				nextAccountID: tt.fields.nextAccountID,
				accounts:      tt.fields.accounts,
				payments:      tt.fields.payments,
			}
			if err := s.Reject(tt.args.paymentID); (err != nil) != tt.wantErr {
				t.Errorf("Reject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Accounts() []*types.Account {
	var accounts []*types.Account
	accounts = append(
		accounts,
		&types.Account{ID: 1, Phone: "9127660305", Balance: 0},
		&types.Account{ID: 2, Phone: "9127660306", Balance: 1},
		&types.Account{ID: 3, Phone: "9127660307", Balance: 2},
		&types.Account{ID: 4, Phone: "9127660307", Balance: 3})
	return accounts
}

func Payments() []*types.Payment {
	var payments []*types.Payment
	payments = append(
		payments,
		&types.Payment{
			ID:        "e1dceb29-6cc4-48c5-acd4-455530f9d50a",
			AccountID: 3,
			Amount:    1,
			Category:  types.CategoryFood,
			Status:    types.PaymentStatusInProgress,
		},
		&types.Payment{
			ID:        "5f7834e5-1c1e-42ff-bd78-9fd7f1b5da28",
			AccountID: 2,
			Amount:    1,
			Category:  types.CategoryIt,
			Status:    types.PaymentStatusInProgress,
		})
	return payments
}

func TestService_Repeat_success(t *testing.T) {
	s := newTestService()

	account, err := s.AddAccountWithBalance("9127660305", 100)
	if err != nil {
		t.Errorf("account => %v", account)
		return
	}

	payment, err := s.Pay(account.ID, 10, types.CategoryIt)
	if err != nil {
		t.Errorf("payment => %v", payment)
		return
	}

	newPayment, err := s.Repeat(payment.ID)
	if err != nil {
		t.Errorf("newPayment => %v", newPayment)
		return
	}
}

func TestService_FavoritePayment_success(t *testing.T) {
	s := newTestService()
	account, err := s.AddAccountWithBalance("9127660305", 100)
	if err != nil {
		t.Errorf("account => %v", account)
		return
	}

	payment, err := s.Pay(account.ID, 10, types.CategoryIt)
	if err != nil {
		t.Errorf("payment => %v", payment)
		return
	}

	favorite, err := s.FavoritePayment(payment.ID, types.CategoryIt)
	if err != nil {
		t.Errorf("favorite => %v", favorite)
		return
	}

	payFromFavorite, err := s.PayFromFavorite(favorite.ID)
	if err != nil {
		t.Errorf("payFromFavorite => %v", payFromFavorite)
		return
	}
}

func TestService_FindFavoriteByID(t *testing.T) {
	type fields struct {
		nextAccountID int64
		accounts      []*types.Account
		payments      []*types.Payment
		favorites     []*types.Favorite
	}
	type args struct {
		favoriteID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Favorite
		wantErr bool
	}{
		{
			name: "successFull find",
			fields: fields{
				nextAccountID: 0,
				accounts:      Accounts(),
				payments:      Payments(),
				favorites:     Favorites(),
			},
			args: args{
				favoriteID: defaultFavorite.ID,
			},
			want:    &defaultFavorite,
			wantErr: false,
		},
		{
			name: "successFull find",
			fields: fields{
				nextAccountID: 0,
				accounts:      Accounts(),
				payments:      Payments(),
				favorites:     Favorites(),
			},
			args: args{
				favoriteID: "nonExistingFavoriteID",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				nextAccountID: tt.fields.nextAccountID,
				accounts:      tt.fields.accounts,
				payments:      tt.fields.payments,
				favorites:     tt.fields.favorites,
			}
			got, err := s.FindFavoriteByID(tt.args.favoriteID)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindFavoriteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindFavoriteByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Favorites() []*types.Favorite {
	var favorites []*types.Favorite
	favorites = append(
		favorites,
		&types.Favorite{
			ID:        defaultFavorite.ID,
			AccountID: Accounts()[0].ID,
			Name:      types.CategoryIt,
			Amount:    10,
			Category:  types.CategoryIt,
		}, &types.Favorite{
			ID:        uuid.New().String(),
			AccountID: Accounts()[1].ID,
			Name:      types.CategoryIt,
			Amount:    10,
			Category:  types.CategoryIt,
		}, &types.Favorite{
			ID:        uuid.New().String(),
			AccountID: Accounts()[2].ID,
			Name:      types.CategoryIt,
			Amount:    10,
			Category:  types.CategoryIt,
		}, &types.Favorite{
			ID:        uuid.New().String(),
			AccountID: Accounts()[3].ID,
			Name:      types.CategoryIt,
			Amount:    10,
			Category:  types.CategoryIt,
		})
	return favorites
}