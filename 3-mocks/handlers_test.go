package handlers

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"testing/3-mocks/mocks"
	"testing/3-mocks/repository"
)

func TestGetComments(t *testing.T) {
	url := fmt.Sprintf("comments/%s", "some-product-id")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		t.Fatal(err)
	}

	repo := new(mocks.CommentsRepository)
	repo.On("GetComments", mock.Anything).Return([]repository.Comment{
		{
			ID:        1,
			ProductID: "4bb43201-be5d-4138-8031-fe583690b3de",
			Author:    "user",
			Content:   "golang awesome!",
		},
	}, nil)

	hc := headerComment{
		repository: repo,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hc.GetComments)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "[{\"id\":1,\"product_id\":\"4bb43201-be5d-4138-8031-fe583690b3de\",\"author\":\"user\",\"content\":\"golang awesome!\"}]\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestAddComment(t *testing.T) {
	url := fmt.Sprintf("comments/%s", "some product id")

	payload := `
			{
				"product_id": "de5ff4db-8196-4b22-b23b-2d44528ac725",
				"author": "user",
				"content": "awesome product!"
			}
		`
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	repo := new(mocks.CommentsRepository)
	repo.On("SaveComment", mock.Anything).Return(nil)

	hc := headerComment{
		repository: repo,
	}

	handler := http.HandlerFunc(hc.AddComment)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := "{\"id\":4,\"product_id\":\"\",\"author\":\"user\",\"content\":\"awesome product!\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
