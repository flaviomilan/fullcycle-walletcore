package gateway

import "github.com/flaviomilan/fullcycle-walletcore/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
