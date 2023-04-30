package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicelerias/blog-golang/api/controllers"
	"github.com/alicelerias/blog-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"golang.org/x/net/context"
)

type MockRepository struct {
}

type MockCache struct {
}

func newMockRepository() *MockRepository {
	return &MockRepository{}
}

func newMockCache() *MockCache {
	return &MockCache{}
}

func performRequest(method, path string, router *gin.Engine, payload io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, payload)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
func TestServer(t *testing.T) {
	mockRepository := newMockRepository()
	mockCache := newMockCache()
	server := controllers.NewServer(mockRepository, mockCache)
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("uid", "1")
	})

	router.GET("/users", server.GetUsers)
	router.GET("/users/:id", server.GetUser)
	router.POST("/users", server.CreateUser)
	router.PUT("/users/:id", server.UpdateUser)
	router.DELETE("/users/:id", server.DeleteUser)

	router.GET("/posts", server.GetPosts)
	router.GET("/posts/:id", server.GetPost)
	router.POST("/posts", server.CreatePost)
	router.PUT("/posts/:id", server.UpdatePost)
	router.DELETE("/posts/:id", server.DeletePost)

	router.POST("/favorite/:id", server.Favorite)
	router.DELETE("/favorite/:id", server.Unfavorite)
	router.GET("/favorites", server.GetFavoritesPosts)
	router.GET("/favorites/:id", server.GetFavoritesByPost)

	router.POST("/comment/:id", server.CreateComment)
	router.DELETE("/comment/:id", server.DeleteComment)
	router.GET("/comment/:id", server.GetPostComments)

	router.POST("/follow/:id", server.CreateFollow)
	router.DELETE("/follow/:id", server.Unfollow)

	router.GET("/feed", server.Feed)

	gin.SetMode(gin.TestMode)

	// Test case 1: test 404
	res := performRequest("GET", "/undefined", router, nil)
	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.Equal(t, "404 page not found", res.Body.String())

	// Test case 2: create user
	payload, _ := json.Marshal(&models.User{
		UserName: "LukeSkywalker",
		Email:    "luke@email.com",
		Password: "luke123",
	})
	res = performRequest("POST", "/users", router, bytes.NewReader(payload))
	assert.Equal(t, http.StatusCreated, res.Code)

	// invalid user

	invalidPayload, _ := json.Marshal(&models.User{
		UserName: "",
	})
	res = performRequest("POST", "/users", router, bytes.NewReader(invalidPayload))
	assert.Equal(t, http.StatusBadRequest, res.Code)

	// Test case 3: get users
	res = performRequest("GET", "/users", router, nil)
	assert.Equal(t, http.StatusOK, res.Code)

	var data struct {
		Users []*models.User `json:"users"`
	}
	json.Unmarshal([]byte(res.Body.String()), &data)
	assert.Equal(t, "Leia Ogana", data.Users[0].UserName)

	// wrong user

	assert.NotEqual(t, "Chewbaca", data.Users[0].UserName)

	// Test case 4: get user
	res = performRequest("GET", "/users/1", router, nil)
	assert.Equal(t, http.StatusOK, res.Code)

	var userData *models.User
	json.Unmarshal([]byte(res.Body.String()), &userData)
	assert.Equal(t, "Leia Ogana", userData.UserName)

	// wrong user
	assert.NotEqual(t, "Chewbaca", userData.UserName)

	// Test case 6: delete user

	res = performRequest("DELETE", "/users/1", router, nil)
	assert.Equal(t, http.StatusNoContent, res.Code)

	// Test case 7: create post
	payload, _ = json.Marshal(&models.Post{
		Title:   "titleshausahusauashusauh",
		Content: "content",
	})
	res = performRequest("POST", "/posts", router, bytes.NewReader(payload))

	assert.Equal(t, http.StatusCreated, res.Code)

	// Test case 8: get posts
	res = performRequest("GET", "/posts", router, nil)
	assert.Equal(t, http.StatusOK, res.Code)

	var postsData struct {
		Posts []*models.Post `json:"content"`
	}
	json.Unmarshal([]byte(res.Body.String()), &postsData)

	assert.Equal(t, "title", postsData.Posts[0].Title)

	// wrong title
	assert.NotEqual(t, "wrong title", postsData.Posts[0].Title)

	// Test case 9: get post
	res = performRequest("GET", "/posts/1", router, nil)
	assert.Equal(t, http.StatusOK, res.Code)

	var postData *models.Post
	json.Unmarshal([]byte(res.Body.String()), &postData)
	assert.Equal(t, "title", postData.Title)

	// wrong title
	assert.NotEqual(t, "wrong title", postData.Title)

	// Test case 10: update post
	payload, _ = json.Marshal(map[string]string{
		"title": "new title",
	})
	res = performRequest("PUT", "/posts/1", router, bytes.NewReader(payload))
	assert.Equal(t, http.StatusOK, res.Code)

	var expectedPostData = &models.Post{}
	json.Unmarshal(res.Body.Bytes(), &expectedPostData)
	assert.Equal(t, "new title", expectedPostData.Title)

	// Test case 11: delete post

	res = performRequest("DELETE", "/posts/1", router, nil)
	assert.Equal(t, http.StatusNoContent, res.Code)

	// Test case 12: follow

	res = performRequest("POST", "/follow/1", router, nil)
	assert.Equal(t, http.StatusCreated, res.Code)

	// Test case 13: unfollow

	res = performRequest("DELETE", "/follow/1", router, nil)
	assert.Equal(t, http.StatusNoContent, res.Code)

	// Test case 14: feed

	res = performRequest("GET", "/feed", router, nil)
	assert.Equal(t, http.StatusOK, res.Code)

	var feedData struct {
		Feed []*models.Post `json:"content"`
	}

	json.Unmarshal([]byte(res.Body.String()), &feedData)

	assert.Equal(t, "post", feedData.Feed[1].Title)

	// post not found

	assert.NotEqual(t, "wrong post", feedData.Feed[1].Title)

	// Test case 15: favorite post

	res = performRequest("POST", "/favorite/1", router, nil)

	assert.Equal(t, http.StatusCreated, res.Code)

	// Test case 16: unfavorite post

	res = performRequest("DELETE", "/favorite/1", router, nil)

	assert.Equal(t, http.StatusNoContent, res.Code)

	// Test case 17: get favorites posts

	res = performRequest("GET", "/favorites", router, nil)

	assert.Equal(t, http.StatusOK, res.Code)

	var favorites struct {
		Favorites []*models.Post `json:"content"`
	}

	json.Unmarshal([]byte(res.Body.String()), &favorites)
	assert.Equal(t, "title", favorites.Favorites[1].Title)

	// post not found
	assert.NotEqual(t, "wrong title", favorites.Favorites[1].Title)

	// Test case 18: get favorites by post
	res = performRequest("GET", "/favorites/1", router, nil)

	assert.Equal(t, http.StatusOK, res.Code)

	var favoritesByPost struct {
		Favorites []*models.Favorite `json:"content"`
	}

	json.Unmarshal([]byte(res.Body.String()), &favoritesByPost)
	assert.Equal(t, uint32(12), favoritesByPost.Favorites[0].PostId)

	// Test case 19: create comment
	payload, _ = json.Marshal(&models.Comment{
		Content: "comment",
	})
	res = performRequest("POST", "/comment/1", router, bytes.NewReader(payload))

	assert.Equal(t, http.StatusCreated, res.Code)

	// Test case 20: delete comment

	res = performRequest("DELETE", "/comment/1", router, nil)

	assert.Equal(t, http.StatusNoContent, res.Code)

	// Test cas 21: get post comments

	res = performRequest("GET", "/comment/1", router, nil)

	assert.Equal(t, http.StatusOK, res.Code)

	var comments struct {
		Comments []*models.Comment `json:"comments"`
	}

	json.Unmarshal([]byte(res.Body.String()), &comments)
	assert.Equal(t, "comment", comments.Comments[0].Content)

	// wrong comment
	assert.NotEqual(t, "wrong comment", comments.Comments[0].Content)

}

func (c *MockCache) SetKey(key string, id string, value interface{}, expiration time.Duration) error {
	return nil
}

func (c *MockCache) GetKey(key string, id string, model interface{}) error {
	return nil
}

func (c *MockCache) DeleteKey(key string, id string) error {
	return nil
}
func (s *MockRepository) GetHome() error {
	return nil
}

func (s *MockRepository) CreateUser(context.Context, *models.User) error {
	return nil
}

func (s *MockRepository) FindAllUsers(context.Context, *models.User) (*[]models.User, error) {
	return &[]models.User{
		{
			ID:       1,
			UserName: "Leia Ogana",
			Password: "leia123",
			Email:    "leia@email.com",
		},
	}, nil
}

func (s *MockRepository) FindUserByID(context.Context, string) (*models.User, error) {
	return &models.User{
		UserName: "Leia Ogana",
	}, nil
}

func (s *MockRepository) GetUser(context.Context, string) (*models.User, error) {
	return &models.User{
		UserName: "Leia Ogana",
	}, nil
}

func (s *MockRepository) UpdateUser(ctx context.Context, value interface{}, id string) (*models.User, error) {
	data := value.(map[string]interface{})
	return &models.User{
		UserName: data["user_name"].(string),
	}, nil
}

func (s *MockRepository) DeleteUser(context.Context, string) error {
	return nil
}

func (s *MockRepository) Favorite(context.Context, *models.Favorite) error {
	return nil
}

func (s *MockRepository) Unfavorite(ctx context.Context, postId uint32, userId uint32) error {
	return nil
}

func (s *MockRepository) GetFavorite(ctx context.Context, postId string, userId string) bool {
	return true
}

func (s *MockRepository) GetFavoritesPostsByUser(ctx context.Context, cursor string, userId uint32) ([]models.Post, error) {
	return []models.Post{
		{
			Title:    "title",
			Content:  "content",
			AuthorID: 1,
		},
		{
			Title:    "title",
			Content:  "content",
			AuthorID: 1,
		},
	}, nil
}

func (s *MockRepository) GetFavoritesByPost(ctx context.Context, postId uint32) (*[]models.Favorite, error) {
	return &[]models.Favorite{
		{
			PostId: 12,
			UserId: 2,
		},
	}, nil
}
func (s *MockRepository) CreatePost(context.Context, *models.Post) error {
	return nil
}

func (s *MockRepository) GetPosts(context.Context, string, *models.Post) ([]models.Post, error) {
	return []models.Post{
		{
			Title:    "title",
			Content:  "content",
			AuthorID: 1,
		},
	}, nil
}

func (s *MockRepository) GetPost(context.Context, string) (*models.Post, error) {
	return &models.Post{
		Title:    "title",
		Content:  "content",
		AuthorID: 1,
	}, nil
}

func (s *MockRepository) UpdatePost(ctx context.Context, value interface{}, id string) (*models.Post, error) {
	data := value.(map[string]interface{})
	return &models.Post{
		Title: data["title"].(string),
	}, nil
}

func (s *MockRepository) DeletePost(context.Context, string) error {
	return nil
}

func (s *MockRepository) CreateComment(ctx context.Context, comment *models.Comment) error {
	return nil
}

func (s *MockRepository) DeleteComment(ctx context.Context, id uint32, authorId uint32) error {
	return nil
}

func (s *MockRepository) GetPostComments(ctx context.Context, cursor string, postId uint32) ([]models.Comment, error) {
	return []models.Comment{
		{
			Content: "comment",
		},
	}, nil
}

func (s *MockRepository) Follow(context.Context, *models.Following) error {
	return nil
}

func (s *MockRepository) GetFollows(context.Context, *models.Following) (*[]models.Following, error) {
	return &[]models.Following{}, nil
}

func (s *MockRepository) IsFollowing(ctx context.Context, followerId string, followingId string) bool {
	return true
}

func (s *MockRepository) Unfollow(context.Context, string, string) error {
	return nil
}

func (s *MockRepository) Feed(context.Context, string, string) ([]models.Post, error) {
	return []models.Post{
		{
			Title:    "post",
			Content:  "content",
			AuthorID: 1,
		},
		{
			Title:    "post",
			Content:  "content",
			AuthorID: 1,
		},
		{
			Title:    "post",
			Content:  "content",
			AuthorID: 1,
		},
	}, nil
}

func (s *MockRepository) GetPostsByUser(ctx context.Context, post *models.Post, cursor string, uid string) ([]models.Post, error) {
	return []models.Post{}, nil
}
