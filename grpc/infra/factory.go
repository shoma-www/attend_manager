package infra

import "github.com/shoma-www/attend_manager/grpc/service"

// RepoFactory Repositoryのファクトリ
type RepoFactory struct {
}

// NewRepoFactory コンストラクタ
func NewRepoFactory() *RepoFactory {
	return &RepoFactory{}
}

// CreateUserRepository Repositoryつくるぞ
func (rf *RepoFactory) CreateUserRepository() service.UserRepository {
	return &userDAO{}
}
