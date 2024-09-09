package controllers_test

import (
	"api-testing/controllers"
	"api-testing/models"
	model_test "api-testing/tests/models"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := model_test.SetupTest()
	if err != nil {
		log.Fatalf("Error setting up test db: %v", err)
	}

	code := m.Run()

	model_test.TeardownTestDB()

	os.Exit(code)
}

func TestGetAllPosts(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/posts", nil)
	resRec := httptest.NewRecorder()

	post1 := models.Blog{Title: "First post", Body: "This is the first post"}
	post2 := models.Blog{Title: "Second post", Body: "This is the secon post"}
	models.DB.Create(&post1)
	models.DB.Create(&post2)

	controllers.GetAllPosts(resRec, req)
	res := resRec.Result()

	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode, "API should return 200 status code for getting all posts")
}

func TestGetOnePost(t *testing.T) {
	postTest := models.Blog{Title: "Post Test", Body: "This is a test post"}
	models.DB.Create(&postTest)

	t.Run("Valid post ID", func(t *testing.T) {
		postId := strconv.FormatUint(uint64(postTest.ID), 10)
		req := httptest.NewRequest(http.MethodGet, "/posts/"+postId, nil)
		resRec := httptest.NewRecorder()

		controllers.GetOnePost(resRec, req)
		res := resRec.Result()
		defer res.Body.Close()

		// First make sure the post was created
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var postCreated models.Blog
		decodeErr := json.NewDecoder(res.Body).Decode(&postCreated)
		assert.NoError(t, decodeErr)

		// Make sure that the post has its correspondant data
		assert.Equal(t, postTest.Title, postCreated.Title, "The post should have a title")
		assert.Equal(t, postTest.Body, postCreated.Body, "The post should have a body content")
	})
}
