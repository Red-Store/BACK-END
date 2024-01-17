package handler

import "MyEcommerce/features/user"

type UserResponse struct {
	Name         string `json:"name" form:"name"`
	UserName     string `json:"user_name" form:"user_name"`
	Email        string `json:"email" form:"email"`
	PhotoProfile string `json:"photo_profile" form:"photo_profile"`
}

func CoreToResponse(data *user.Core) UserResponse {
	var result = UserResponse{
		Name:         data.Name,
		UserName:     data.UserName,
		Email:        data.Email,
		PhotoProfile: data.PhotoProfile,
	}
	return result
}

func CoreToResponseList(data []user.Core) []UserResponse {
	var results []UserResponse
	for _, v := range data {
		var result = UserResponse{
			Name:         v.Name,
			UserName:     v.UserName,
			Email:        v.Email,
			PhotoProfile: v.PhotoProfile,
		}
		results = append(results, result)
	}
	return results
}
