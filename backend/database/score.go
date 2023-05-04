package database

import (
	"github.com/alicelerias/blog-golang/models"
)

func (s *PostgresDBRepository) Recomendations(uid string) (*[]models.User, error) {
	users := []models.User{}
	friendsOfFriends := []models.User{}
	recomendations := []models.User{}

	err := s.db.Raw(
		`SELECT recomended_user.id, recomended_user.user_name, recomended_user.bio, recomended_user.avatar
    FROM users AS me
    JOIN followings As my_friends
      ON me.id = my_friends.follower_id
    JOIN users AS friend 
      ON my_friends.following_id = friend.id
    JOIN followings AS recomendations
      ON friend.id = recomendations.follower_id 
        AND recomendations.following_id != me.id
    JOIN POPULARITY_SCORE AS pop
      ON recomendations.following_id = pop.id
    JOIN users AS recomended_user
      ON pop.id = recomended_user.id
    WHERE me.id = ?
    ORDER BY pop.score DESC
    LIMIT 5`, uid).
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
