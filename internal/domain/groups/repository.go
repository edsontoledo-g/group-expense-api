package groups

import "github.com/jackc/pgx/v5/pgxpool"

type GroupsRepository interface {
	GroupByIDForUser(id string, userID string) (*Group, error)
	GroupsByUser(userID string) ([]Group, error)
}

type groupsRepository struct {
	pool *pgxpool.Pool
}

func (repo *groupsRepository) GroupByIDForUser(id string, userID string) (*Group, error) {
	return nil, nil
}

func (repo *groupsRepository) GroupsByUser(userID string) ([]Group, error) {
	return nil, nil
}

func NewGroupsRepository(db *pgxpool.Pool) GroupsRepository {
	return &groupsRepository{
		pool: db,
	}
}
