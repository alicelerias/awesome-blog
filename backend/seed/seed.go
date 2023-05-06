package main

import (
	"math/rand"
	"strconv"

	"github.com/alicelerias/blog-golang/api/auth"
	"github.com/alicelerias/blog-golang/config"
	"github.com/alicelerias/blog-golang/database"
	"github.com/alicelerias/blog-golang/models"
	"github.com/brianvoe/gofakeit/v6"
)

func GenUser(r database.Repository, handler func(*models.User)) (err error) {

	model := &models.User{
		UserName: gofakeit.Username(),
		Password: "123456",
		Email:    gofakeit.Email(),
		Bio:      gofakeit.Phrase(),
		Avatar:   gofakeit.Person().Image,
	}
	if err = auth.CreateUser(r, model); err != nil {
		return
	}
	handler(model)
	return
}

func genUsers(
	quantity int,
	r database.Repository,
	handler func(m *models.User)) (err error) {
	for i := 0; i < quantity; i++ {
		err = GenUser(r, handler)
		if err != nil {
			return
		}
	}
	return
}

func genFollowing(r database.Repository, followerId uint32, followingId uint32) (err error) {
	model := &models.Following{
		FollowerID:  followerId,
		FollowingID: followingId,
	}
	if err = r.Follow(model); err != nil {
		return
	}
	return
}

func followingPosts(r database.Repository, cursor string, followingId uint32) ([]models.Post, error) {
	post := &models.Post{}
	uidStr := strconv.FormatUint(uint64(followingId), 10)
	posts, err := r.GetPostsByUser(post, cursor, uidStr)
	if err != nil {
		return []models.Post{}, err
	}
	return posts, nil
}

func genPost(r database.Repository, handler func(*models.Post), authorId uint32) (err error) {
	model := &models.Post{
		Title:    gofakeit.Phrase(),
		Content:  gofakeit.LoremIpsumSentence(30),
		Img:      gofakeit.ImageURL(50, 100),
		AuthorID: authorId,
	}

	if err = r.CreatePost(model); err != nil {
		return
	}
	handler(model)
	return
}

func genPosts(
	quantity int,
	r database.Repository,
	authorId uint32,
	handler func(m *models.Post)) (err error) {
	for i := 0; i < quantity; i++ {
		err = genPost(r, handler, authorId)
		if err != nil {
			return
		}
	}
	return
}

func genComment(r database.Repository, postId uint32, authorId uint32) (err error) {
	model := &models.Comment{
		PostId:   postId,
		AuthorId: authorId,
		Content:  gofakeit.LoremIpsumSentence(15),
	}
	if err = r.CreateComment(model); err != nil {
		return
	}
	return
}

func genComments(
	quantity int,
	r database.Repository,
	postId uint32,
	authorId uint32,
	handler func(m *models.Comment)) (err error) {
	for i := 0; i < quantity; i++ {
		err = genComment(r, postId, authorId)
		if err != nil {
			return
		}
	}
	return
}

func genFavorite(r database.Repository, postId uint32, userId uint32) (err error) {
	model := &models.Favorite{
		PostId: postId,
		UserId: userId,
	}
	if err = r.Favorite(model); err != nil {
		return
	}
	return
}

func chance(perc float64) bool {
	if rand.Float64() < perc {
		return true
	}
	return false
}

func cleanDatabase(r database.Repository) (err error) {
	if err := r.DeleteUsersTable(); err != nil {
		return err
	}

	if err := r.DeletePostsTable(); err != nil {
		return err
	}

	if err := r.DeleteFollowingsTable(); err != nil {
		return err
	}

	if err := r.DeleteCommentsTable(); err != nil {
		return err
	}

	if err := r.DeleteFavoritesTable(); err != nil {
		return err
	}

	return nil
}

func main() {
	configs := config.GetConfig()
	connection, err := database.GetConnection(configs)
	defer connection.Close()
	if err != nil {
		panic(err)
	}
	database.MigrateDB(connection)
	r := database.NewPostgresDBRepository(connection)

	// delete database

	if err := cleanDatabase(r); err != nil {
		return
	}

	// continue

	var seed int64 = 0
	gofakeit.Seed(seed)
	rand.Seed(seed)

	users := []*models.User{}
	nnu := 10000
	genUsers(nnu, r, func(user *models.User) {
		randomCreatePosts := rand.Intn(30)
		genPosts(randomCreatePosts, r, user.ID, func(post *models.Post) {})
		users = append(users, user)
	})

	for uidx, user := range users {
		num_followees := rand.Intn(int(float64(nnu)*0.04)) + 100
		for i := 0; i <= num_followees; i++ {
			randomIndex := rand.Intn(len(users))
			for randomIndex == uidx {
				randomIndex = rand.Intn(len(users))
			}
			randomFollowing := users[randomIndex]

			if err := genFollowing(r, user.ID, randomFollowing.ID); err != nil {
				continue
			}

			fPosts, err := followingPosts(r, "", randomFollowing.ID)
			if err != nil {
				return
			}
			for _, post := range fPosts {

				if chance(0.5) {
					if err := genComment(r, uint32(post.ID), user.ID); err != nil {
						return
					}
				}

				if chance(0.5) {
					if err := genFavorite(r, uint32(post.ID), user.ID); err != nil {
						return
					}
				}
			}
		}
	}
}
