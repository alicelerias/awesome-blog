package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Recomendations(uid string) (*[]models.User, error) {
	// report := []struct {
	// 	column_a string `gorm:columa`
	// }
	users := []models.User{}

	recomendations := []models.User{}

	// rever

	err := s.db.Raw(
		`select distinct recomended_user.id, recomended_user.user_name, recomended_user.avatar, recomended_user.bio, pop.score
		from users as recomended_user
		join popularity_score as pop on recomended_user.id = pop.id
		join followings as recomendations on pop.id = recomendations.following_id
		join users as friend on recomendations.follower_id = friend.id
		join followings as my_friends on friend.id = my_friends.following_id
		join users as me on my_friends.follower_id = me.id and me.id = ?
		where recomendations.following_id != me.id and recomendations.follower_id != me.id
		order by pop.score DESC
		limit 5;
		`, uid).
		Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}

	if len(users) < 5 {
		s.db.Raw(
			`select * from popularity_score as pop
			join users as u on pop.id = u.id
			order by pop.score DESC
			`).Limit(5 - len(users)).
			Find(&recomendations)

		users = append(users, recomendations...)
	}

	return &users, nil
}
