package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Recomendations(uid string) (*[]models.User, error) {
	users := []models.User{}

	recomendations := []models.User{}

	err := s.db.Raw(
		`select distinct recomended_user.id, recomended_user.user_name, recomended_user.avatar, recomended_user.bio, pop.score
		from users as recomended_user
		join popularity_score as pop on recomended_user.id = pop.id
		join followings as recomendations on pop.id = recomendations.following_id
		join users as friend on recomendations.follower_id = friend.id
		join followings as my_friends on friend.id = my_friends.following_id
		join users as me on my_friends.follower_id = me.id and me.id = ?
		where recomendations.following_id != me.id and recomendations.follower_id != me.id and recomendations.following_id != my_friends.following_id
		order by pop.score DESC
		limit 5;
		`, uid).
		Find(&users).Error
	if err != nil {
		return &[]models.User{}, err
	}

	if len(users) < 5 {
		s.db.Raw(
			`SELECT u.id, u.user_name, u.bio, u.avatar, ps.score
			FROM public.users u
			INNER JOIN public.popularity_score ps ON u.id = ps.id
			LEFT JOIN (
				SELECT following_id
				FROM public.followings
				WHERE follower_id = ?
			) f ON u.id = f.following_id
			WHERE f.following_id IS NULL
			ORDER BY ps.score DESC
			`, uid).Limit(5).
			Find(&recomendations)

		users = recomendations
	}

	return &users, nil
}
