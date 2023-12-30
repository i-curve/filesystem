package pojo

type CreateBucket struct {
	Name  string `json:"name" form:"name" binding:"required"`
	BType int    `json:"b_type" form:"b_type" binding:"required,min=1,max=3"`
}

type DeleteBucket struct {
	Name string `json:"name" form:"name" binding:"required"`
}
