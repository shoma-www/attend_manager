package infra

import (
	"github.com/shoma-www/attend_manager/core"
	"github.com/shoma-www/attend_manager/grpc/service"
)

// RepoFactory Repositoryのファクトリ
type RepoFactory struct {
	logger core.Logger
}

// NewRepoFactory コンストラクタ
func NewRepoFactory(l core.Logger) *RepoFactory {
	return &RepoFactory{logger: l}
}

// CreateUserRepository Repositoryつくるぞ
func (rf *RepoFactory) CreateUserRepository() service.UserRepository {
	return &userDAO{logger: rf.logger}
}
