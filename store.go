package main

import (
	"ago/comps"
	"net/http"
)

type Store struct {
	Users map[string]UserStore
}

type UserStore struct {
	IP             string
	KV             map[string]string
	TileMap        comps.TileMap
	PointsToSmooth []int
}

func NewStore() Store {
	return Store{
		Users: make(map[string]UserStore),
	}
}

func (s *Store) GetUser(ip string) UserStore {
	if _, ok := s.Users[ip]; !ok {
		s.Users[ip] = UserStore{
			IP: ip,
			KV: make(map[string]string),
		}
	}
	return s.Users[ip]
}

func (s *Store) SetUser(ip string, user UserStore) {
	s.Users[ip] = user
}

func (s *Store) GetKV(ip, key string) string {
	return s.GetUser(ip).KV[key]
}

func (s *Store) SetTileMap(ip string, tm comps.TileMap) {
	user := s.GetUser(ip)
	user.TileMap = tm
	s.SetUser(ip, user)
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
		println("ip: " + user.IP)
		for key, value := range user.KV {
			println(key + " : " + value)
		}
		user.TileMap.Display()
	}
}
