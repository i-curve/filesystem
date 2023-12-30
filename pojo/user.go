package pojo

type CreateUser struct {
	Name  string `json:"name" binding:"required,min=4"`
	UType int    `json:"u_type" binding:"required,min=1,max=3"`
}

type DeleteUser struct {
	Name string `json:"name" binding:"required"`
}

type UpdateUser struct {
	Name  string `json:"name" binding:"required"`
	Auth  string `json:"auth"`
	UType int    `json:"u_type"`
}
