package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"user-api/config"
	"user-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("insert user success", func(mt *mtest.T) {
	

		user := models.User{
			Name:        "Test User",
			DOB:         "2000-01-01",
			Address:     "Test City",
			Description: "Test Description",
		}
		body, _ := json.Marshal(user)

		req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		config.DB = mt.Client
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		CreateUser(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Expected 201 Created, got %d", rr.Code)
		}
	})
}

func TestGetAllUsers(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("get all users success", func(mt *mtest.T) {
		

		config.DB = mt.Client

		userDoc := bson.D{
			{"_id", primitive.NewObjectID()},
			{"name", "Test User"},
			{"dob", "2000-01-01"},
			{"address", "Test Address"},
			{"description", "Test Description"},
			{"createdAt", time.Now().Format("2006-01-02 15:04:05")},
		}
		cursor := mtest.CreateCursorResponse(1, "userdb.users", mtest.FirstBatch, userDoc)
		mt.AddMockResponses(cursor)

		req, _ := http.NewRequest("GET", "/users", nil)
		rr := httptest.NewRecorder()

		GetAllUsers(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected 200 OK, got %d", rr.Code)
		}
	})
}

func TestGetUserByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("get user by id success", func(mt *mtest.T) {
		

		config.DB = mt.Client

		id := primitive.NewObjectID()
		userDoc := bson.D{
			{"_id", id},
			{"name", "User One"},
			{"dob", "2000-01-01"},
			{"address", "Address"},
			{"description", "Description"},
			{"createdAt", time.Now().Format("2006-01-02 15:04:05")},
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "userdb.users", mtest.FirstBatch, userDoc))

		req, _ := http.NewRequest("GET", "/users/"+id.Hex(), nil)
		rr := httptest.NewRecorder()

		ctx := req.Context()
		req = req.WithContext(context.WithValue(ctx, "id", id.Hex()))

		GetUserByID(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected 200 OK, got %d", rr.Code)
		}
	})
}

func TestUpdateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("update user success", func(mt *mtest.T) {
	

		config.DB = mt.Client

		id := primitive.NewObjectID()
		user := models.User{
			Name:        "Updated Name",
			DOB:         "1990-01-01",
			Address:     "New Address",
			Description: "Updated",
		}
		body, _ := json.Marshal(user)

		req, _ := http.NewRequest("PUT", "/users/"+id.Hex(), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		ctx := req.Context()
		req = req.WithContext(context.WithValue(ctx, "id", id.Hex()))

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		UpdateUser(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected 200 OK, got %d", rr.Code)
		}
	})
}

func TestDeleteUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("delete user success", func(mt *mtest.T) {
	

		config.DB = mt.Client

		id := primitive.NewObjectID()

		req, _ := http.NewRequest("DELETE", "/users/"+id.Hex(), nil)
		rr := httptest.NewRecorder()

		ctx := req.Context()
		req = req.WithContext(context.WithValue(ctx, "id", id.Hex()))

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		DeleteUser(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected 200 OK, got %d", rr.Code)
		}
	})
}