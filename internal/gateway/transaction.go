package gateway

import "github.com/flaviomilan/fullcycle-walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
