package pojo

type ListBucket struct {
	Name string `json:"name" from"name"`
}

type CreateBucket struct {
	Name   string `json:"name" form:"name" binding:"required"`
	BType  int    `json:"b_type" form:"b_type" binding:"required,min=1,max=3"`
	IsTemp bool   `json:"is_temp" form:"is_temp"`
}

type DeleteBucket struct {
	Name string `json:"name" form:"name" binding:"required"`
}
