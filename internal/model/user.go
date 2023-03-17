package model

type User struct {
	First_name *string `bson:"first_name" validate:"required, min=5, max=100"`
	Last_name  *string `bson:"last_name" validate:"required, min=5, max=150"`
	Email      *string `bson:"email" validate:"email, required"`
	Username   *string `bson:"username" validate:"min=8"`
	Password   *string `bson:"password" validate:"required, min=8"`
}
