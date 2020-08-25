package infra

import (
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/ent"
	"github.com/shoma-www/attend_manager/grpc/service"
)

// RepoFactory Repositoryのファクトリ
type RepoFactory struct {
	logger   core.Logger
	dbClient *ent.Client
}

// NewRepoFactory コンストラクタ
func NewRepoFactory(l core.Logger, db *ent.Client) *RepoFactory {
	return &RepoFactory{logger: l, dbClient: db}
}

// CreateTransaction is created transaction
func (rf *RepoFactory) CreateTransaction() service.Transaction {
	return &transaction{l: rf.logger, cl: rf.dbClient}
}

// CreateUserRepository Repositoryつくるぞ
func (rf *RepoFactory) CreateUserRepository() service.UserRepository {
	return &userDAO{logger: rf.logger, cl: rf.dbClient}
}
