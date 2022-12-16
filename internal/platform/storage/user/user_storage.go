package user

type Response struct {
	ID       uint32 `json:"id" copier:"must"`
	FullName string `json:"full_name"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Nickname string `json:"nickname"`
	IdCard   string `json:"id_card"`
	Email    string `json:"email" copier:"must"`
}
