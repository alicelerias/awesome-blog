select distinct recomended_user.id, recomended_user.user_name, pop.score from users as recomended_user
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
where me.id = 1465 and recomendations.following_id != me.id
order by pop.score DESC
limit 5;
