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
	"github.com/mavolin/go-htmx"
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
		host = "172.105.161.248:443"
	} else {
		host = "localhost:" + strconv.Itoa(*FlagPort)
	}

	store := NewStore()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(comps.IndexComponent()).ServeHTTP(w, r)
	})

	r.Handle("/static/*", http.FileServer(http.FS(StaticFS)))

	r.Post("/new", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			htmx.Reswap(r, htmx.SwapNone)
			w.WriteHeader(http.StatusBadRequest)
		}
		formSize := r.Form.Get("size")
		var size int
		switch formSize {
		case "s":
			size = 30
		case "m":
			size = 50
		case "l":
			size = 70
		default:
			size = 30
		}

		tm := comps.NewTileMap(10, size, size)

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
		prescribedMagnitude := r.Form.Get("prescribedMagnitude")
		if err != nil {
			panic(err)
		}
		println("magnitude:", magnitude)
		println("prescribedMagnitude:", prescribedMagnitude)
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
		altInt := 0
		if prescribedMagnitude != "undefined" && helper.Atoi(prescribedMagnitude) != -1 {
			altInt = helper.Atoi(prescribedMagnitude)
		} else {
			altInt = tm.AltAt(x, y) + helper.Atoi(magnitude)
		}
		// checks
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
	if *FlagPort == 0 {
		http.ListenAndServeTLS(":443", "cert.pem", "privkey.pem", r)
		http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.URL.String(), http.StatusMovedPermanently)
		}))
	} else {
		http.ListenAndServe(host, r)
	}

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
