package infra

import (
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/ent"
	"github.com/shoma-www/attend_manager/grpc/service"
)

// Factory Repositoryのファクトリ
type Factory struct {
	logger   core.Logger
	dbClient *ent.Client
}

// NewFactory コンストラクタ
func NewFactory(l core.Logger, db *ent.Client) *Factory {
	return &Factory{logger: l, dbClient: db}
}

// CreateTransaction is created transaction
func (rf *Factory) CreateTransaction() service.Transaction {
	return &transaction{l: rf.logger, cl: rf.dbClient}
}

// CreateUserRepository Repositoryつくるぞ
func (rf *Factory) CreateUserRepository() service.UserRepository {
	return &userDAO{logger: rf.logger, cl: rf.dbClient}
}
