select id, user_name from (
  select * from (
    SELECT 0 priority, pop.id, pop.user_name, pop.bio, pop.avatar
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
    LIMIT 5
  )  friends_of_friends
  UNION
  select * from (
    SELECT 1 priority, profile.id, profile.user_name, profile.bio, profile.avatar
		FROM popularity_score AS pop
			JOIN users AS profile
				ON pop.id = profile.id
		ORDER BY pop.score DESC
		LIMIT 200
  ) most_popular
) result
order by priority
LIMIT 10
;
