package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, err := NewAccount(client)
	assert.NotNil(t, account)
	assert.Nil(t, err)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account, err := NewAccount(nil)
	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.Error(t, err, "client is required")
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}
