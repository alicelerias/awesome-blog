package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Recomendations(uid string) (*[]models.User, error) {
	users := []models.User{}
	friendsOfFriends := []models.User{}
	recomendations := []models.User{}

	// rever

	err := s.db.Raw(
		`select distinct recomended_user.id, recomended_user.user_name, recomended_user.avatar, recomended_user.bio, pop.score from users as recomended_user
		join popularity_score as pop
		on recomended_user.id = pop.id
		join followings as recomendations
		on pop.id = recomendations.following_id
		join users as friend
		on recomendations.follower_id = friend.id
		join followings as my_friends
		on friend.id = my_friends.following_id
		join users as me 
		on my_friends.follower_id = me.id
		where me.id = ? and recomendations.following_id != me.id
		order by pop.score DESC
		limit 5;
		`, uid).
		Find(&friendsOfFriends).Error
	if err != nil {
		return &[]models.User{}, err
	}

	s.db.Raw(
		`SELECT profile.id, profile.user_name, profile.bio, profile.avatar
		FROM popularity_score AS pop
			JOIN users AS profile
				ON pop.id = profile.id
		ORDER BY pop.score DESC
		LIMIT 5`).
		Find(&recomendations)

	friendsOfFriends = append(friendsOfFriends, recomendations...)
	users = append(users, friendsOfFriends...)

	return &users, nil
}
