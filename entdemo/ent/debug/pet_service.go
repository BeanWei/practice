package entrest

// Pet is the model entity for the Pet schema.
type Pet struct {
	Name string `json:"name,omitempty"`
}

// CreatePetRequest .
type CreatePetRequest struct {
	Name string `json:"name,omitempty" v:"required#请输入宠物的名字"`
}

// UpdatePetRequest .
type UpdatePetRequest struct {
	Name string `json:"name,omitempty" v:"required#请输入宠物的名字"`
}

// DeletePetRequest .
type DeletePetRequest struct {
	ID int `json:"id,omitempty"`
}

// GetPetRequest .
type GetPetRequest struct {
	ID int `json:"id,omitempty"`
}

// ListPetRequest .
type ListPetRequest struct {
	PageSize int    `json:"pageSize,omitempty"`
	PageNum  int    `json:"pageNum,omitempty"`
	Sort     string `json:"sort,omitempty"`
}
