package types

type ProfilePicReq struct {
	ProfilePictureURL string `json:"profilePicURL" binding:"required"`
}
