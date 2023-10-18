package responses

import "main/models"

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type RegisterResponse = models.User
