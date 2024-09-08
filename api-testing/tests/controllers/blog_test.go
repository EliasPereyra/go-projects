package controllers_test

import (
	"api-testing/controllers"
	"api-testing/models"
	model_test "api-testing/tests/models"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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
	req, err := http.NewRequest(http.MethodGet, "/posts", nil)
	if err != nil {
		log.Println("THere was an error with the request", req)
	}
	w := httptest.NewRecorder()

	post1 := models.Blog{Title: "First post", Body: "This is the first post"}
	post2 := models.Blog{Title: "Second post", Body: "This is the secon post"}
	models.DB.Create(&post1)
	models.DB.Create(&post2)

	controllers.GetAllPosts(w, req)
	res := w.Result()

	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, http.StatusOK, "API should return 200 status code")
}
