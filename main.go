package main

import (
	"ago/comps"
	"ago/helper"
	"embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

//go:embed static
var StaticFS embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	FlagPort := flag.Int("port", 0, "Port to listen on.")
	flag.Parse()

	host := ""

	if *FlagPort == 0 {
		host = "172.105.161.248:80"
	} else {
		host = "localhost:" + strconv.Itoa(*FlagPort)
	}

	// hostIP := "172.105.161.248:" + strconv.Itoa(port)
	store := NewStore()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(comps.IndexComponent()).ServeHTTP(w, r)
	})

	// vv this is the old way of doing it with http route vv
	// r.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	r.Handle("/static/*", http.FileServer(http.FS(StaticFS)))

	r.Get("/new", func(w http.ResponseWriter, r *http.Request) {
		tm := comps.NewTileMap(10, 30, 30)

		user := store.GetUser(helper.GetIpFromRequest(r))
		store.SetTileMap(user.IP, tm)

		templ.Handler(comps.TileMapComponent(tm)).ServeHTTP(w, r)
	})

	r.Get("/display", func(w http.ResponseWriter, r *http.Request) {
		user := store.GetUser(helper.GetIpFromRequest(r))
		if user.TileMap.Width != 0 {
			templ.Handler(comps.TileMapComponent(user.TileMap)).ServeHTTP(w, r)
		}
	})

	r.Post("/shape/{id}", func(w http.ResponseWriter, r *http.Request) {
		// get id param
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			panic(err)
		}
		r.ParseForm()
		magnitude := r.Form.Get("magnitude")
		if err != nil {
			panic(err)
		}
		println("magnitude:", magnitude)
		// println("id:", id)
		// get user
		user := store.GetUser(helper.GetIpFromRequest(r))
		// get tilemap
		tm := user.TileMap
		// get x and y from id
		x, y, err := deduceXYFromId(id, tm)

		if err != nil {
			panic(err)
		}
		// get altitude from tilemap
		altInt := tm.AltAt(x, y)
		// increment altitude
		altInt += helper.Atoi(magnitude)
		if altInt > tm.MaxAltitude {
			altInt = tm.MaxAltitude
		}
		if altInt < 0 {
			altInt = 0
		}
		// set altitude in tilemap
		tm.Set(x, y, altInt)
		// regenerate tiles and set user tilemap
		tm.Tiles = tm.GenerateTiles()
		store.SetTileMap(user.IP, tm)
		// return just the tile
		templ.Handler(comps.TileComponent(tm.Tiles[y][x])).ServeHTTP(w, r)
		store.AddEditedPoint(user.IP, comps.Coord{X: x, Y: y})
	})

	r.Get("/options", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(comps.OptionsComponent()).ServeHTTP(w, r)
	})

	r.Get("/smooth", func(w http.ResponseWriter, r *http.Request) {
		user := store.GetUser(helper.GetIpFromRequest(r))
		if user.TileMap.Width != 0 {
			user.TileMap.SeedData = user.TileMap.Smooth(1)
			user.TileMap.Tiles = user.TileMap.GenerateTiles()
			store.SetTileMap(user.IP, user.TileMap)
			templ.Handler(comps.TileMapComponent(user.TileMap)).ServeHTTP(w, r)
			store.ClearEditedPoints(user.IP)
		}
	})

	r.Get("/smoothedited", func(w http.ResponseWriter, r *http.Request) {
		user := store.GetUser(helper.GetIpFromRequest(r))
		if user.TileMap.Width != 0 {
			points := store.GetEditedPoints(user.IP)
			user.TileMap.SeedData = user.TileMap.SmoothPointsAndNeighbours(points, 1)
			user.TileMap.Tiles = user.TileMap.GenerateTiles()
			store.SetTileMap(user.IP, user.TileMap)
			templ.Handler(comps.TileMapComponent(user.TileMap)).ServeHTTP(w, r)
			store.ClearEditedPoints(user.IP)
		}
	})

	// announce server
	fmt.Println("Server running on " + host)
	http.ListenAndServe(host, r)

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
