package dto

type NotifRequest struct {
	UserID  []int             `json:"user_id" binding:"required"`
	Message map[string]string `json:"message" binding:"required"`
}
