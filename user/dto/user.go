package dto

type CreateUserDto struct {
	Name        string `json:name`
	PhoneNumber string `json: phoneNumber`
}
