package responses

import "main/models"

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterResponse = models.User
