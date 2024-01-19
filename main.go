package main

import (
	"ago/comps"
	"ago/helper"
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	store := NewStore()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(comps.IndexComponent()).ServeHTTP(w, r)
	})

	r.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		tm := comps.NewTileMap(11, 10, 10)

		user := store.GetUser(helper.GetIpFromRequest(r))
		store.SetTileMap(user.IP, tm)

		// store.DisplayStore()

		templ.Handler(comps.TileMapComponent(tm)).ServeHTTP(w, r)
	})

	r.Get("/display", func(w http.ResponseWriter, r *http.Request) {
		user := store.GetUser(helper.GetIpFromRequest(r))
		if user.TileMap.Width != 0 {
			templ.Handler(comps.TileMapComponent(user.TileMap)).ServeHTTP(w, r)
		}
	})

	r.Get("/lift/{id}", func(w http.ResponseWriter, r *http.Request) {
		// get id param
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			panic(err)
		}
		// println("id:", id)
		// get user
		user := store.GetUser(helper.GetIpFromRequest(r))
		// get tilemap
		tileMap := user.TileMap
		// get x and y from id
		x, y, err := deduceXYFromId(id, tileMap)

		if err != nil {
			panic(err)
		}
		// get altitude from tilemap
		altInt := tileMap.AltAt(x, y)
		// increment altitude
		if altInt < tileMap.MaxAltitude {
			altInt++
		}
		// set altitude in tilemap
		tileMap.Set(x, y, altInt)
		// // store tilemap
		tileMap.Tiles = tileMap.GenerateTiles()
		// store.SetTileMap(user.IP, tileMap)
		// tileMap.Display()
		// return just the tile
		templ.Handler(comps.TileComponent(tileMap.Tiles[y][x])).ServeHTTP(w, r)
	})

	http.ListenAndServe("127.0.0.1:3000", r)

}

func deduceXYFromId(id int, tileMap comps.TileMap) (x, y int, err error) {
	if tileMap.Width == 0 {
		// Return an error as dividing by zero is not allowed
		return 0, 0, fmt.Errorf("tileMap.Width is zero, division by zero is not allowed")
	}
	fmt.Println("width: ", tileMap.Width)
	x = id % tileMap.Width
	y = id / tileMap.Width
	return x, y, nil
}
