package tools

import "time"

type mockDB struct {
}

// SetupDatabse implements DatabaseInterface.
func (d *mockDB) SetupDatabse() error {
	panic("unimplemented")
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"john": {
		AuthToken: "142ABC",
		Username:  "john",
	},
	"theo": {
		AuthToken: "1653ABC",
		Username:  "theo",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"john": {
		Coins:    230,
		Username: "john",
	},
	"theo": {
		Coins:    400,
		Username: "theo",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//Simulate DB Call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}
func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//Simulate DB Call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
