package dto

type GroupRequest struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	UserIDs     []string `json:"userIds"`
}

type GroupInviteRequest struct {
	Token string `json:"token"`
}
