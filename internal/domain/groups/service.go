package groups

type GroupsService interface {
	Groups()
	Group()
	CreateGroup()
	InviteToGroup()
}

type groupsService struct {
	repo GroupsRepository
}

func (s *groupsService) Groups() {

}

func (s *groupsService) Group() {

}

func (s *groupsService) CreateGroup() {

}

func (s *groupsService) InviteToGroup() {

}

func NewGroupsService(repo GroupsRepository) GroupsService {
	return &groupsService{
		repo: repo,
	}
}
