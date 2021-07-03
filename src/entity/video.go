package entity

type Person struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Age       int8   `json:"age" validate:"gte=1,lte=130"`
	Email     string `json:"email" validate:"required,email"`
}

type Video struct {
	Title       string `json:"title" `
	Description string `json:"description" validate:"max=20"`
	URL         string `json:"url" validate:"required"`
	Author      Person `json:"author" validate:"required"`
}
