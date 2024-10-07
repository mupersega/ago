package main

import (
	"ago/factory"
	"ago/vector"
	"fmt"
	"net/http"
)

type Store struct {
	Users map[string]UserStore
}

type UserStore struct {
	Id           string
	KV           map[string]string
	TileMap      factory.TileMap
	EditedPoints []vector.Vec2
}

func NewStore() Store {
	return Store{
		Users: make(map[string]UserStore),
	}
}

func (s *Store) GetUser(id string) UserStore {
	if _, ok := s.Users[id]; !ok {
		fmt.Println("Creating new user with ip: " + id)
		s.Users[id] = UserStore{
			Id: id,
			KV: make(map[string]string),
		}
	}
	userStore := s.Users[id]
	return userStore
}

func (s *Store) SetUser(ip string, user UserStore) {
	s.Users[ip] = user
}

func (s *Store) GetKV(ip, key string) string {
	return s.GetUser(ip).KV[key]
}

func (s *Store) SetTileMap(ip string, tm factory.TileMap) {
	user := s.GetUser(ip)
	user.TileMap = tm
	s.SetUser(ip, user)
}

func (us UserStore) SetTileMap(tm factory.TileMap) UserStore {
	us.TileMap = tm
	return us
}

func (s *Store) SetKV(ip, key, value string) {
	user := s.GetUser(ip)
	user.KV[key] = value
	s.SetUser(ip, user)
}

func (s *Store) GetIP(r *http.Request) string {
	return r.RemoteAddr
}

func (s *Store) DisplayStore() {
	for _, user := range s.Users {
		println("ip: " + user.Id)
		for key, value := range user.KV {
			println(key + " : " + value)
		}
		user.TileMap.String()
	}
}

func (s *Store) GetEditedPoints(ip string) []vector.Vec2 {
	return s.GetUser(ip).EditedPoints
}

func (s *Store) SetEditedPoints(ip string, points []vector.Vec2) {
	user := s.GetUser(ip)
	user.EditedPoints = points
	s.SetUser(ip, user)
}

func (s *Store) AddEditedPoint(ip string, point vector.Vec2) {
	// add point only if it doesn't already exist
	user := s.GetUser(ip)
	for _, editedPoint := range user.EditedPoints {
		if editedPoint == point {
			return
		}
	}
	user.EditedPoints = append(user.EditedPoints, point)
	s.SetUser(ip, user)
}

func (s *Store) ClearEditedPoints(ip string) {
	user := s.GetUser(ip)
	user.EditedPoints = make([]vector.Vec2, 0)
	s.SetUser(ip, user)
}
