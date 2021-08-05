package serializers

import "github.com/kousuketk/crudGo/app/models"

type Result struct {
	Id               int
	Name             string
	SelfIntroduction string
}

func UserSliceSerialize(user []models.User) []Result {
	var result []Result
	for i := 0; i < len(user); i++ {
		result = append(result, Result{
			Id:               int(user[i].ID),
			Name:             user[i].Name,
			SelfIntroduction: user[i].SelfIntroduction,
		})
	}
	return result
}

func UserSerialize(user models.User) Result {
	result := Result{
		Id:               int(user.ID),
		Name:             user.Name,
		SelfIntroduction: user.SelfIntroduction,
	}
	return result
}
