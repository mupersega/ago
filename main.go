package main

import (
	"ago/comps"
	"ago/helper"
	"context"
	"embed"
	"fmt"
	"log"

	// "log"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	// "github.com/joho/godotenv"
	"github.com/mavolin/go-htmx"
)

//go:embed static
var StaticFS embed.FS

const (
	maxAltitude = 10
)

func main() {
	store := NewStore()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(UserCookieMiddleware)

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
		userId, err := GetUserId(r)
		if err != nil {
			http.Error(w, "user_id not found", http.StatusInternalServerError)
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

		tm := comps.NewTileMap(maxAltitude, size, size, ConfigFromRequest(r))

		user := store.GetUser(userId)
		user = user.SetTileMap(tm)
		store.SetUser(userId, user)

		templ.Handler(comps.TileMapComponent(tm)).ServeHTTP(w, r)
	})

	r.Get("/display", func(w http.ResponseWriter, r *http.Request) {
		userId, err := GetUserId(r)
		if err != nil {
			http.Error(w, "user_id not found", http.StatusInternalServerError)
		}
		user := store.GetUser(userId)
		if user.TileMap.Width != 0 {
			templ.Handler(comps.TileMapComponent(user.TileMap)).ServeHTTP(w, r)
		}
	})

	r.Post("/shape/{id}", func(w http.ResponseWriter, r *http.Request) {
		// get id param
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			fmt.Println("Error converting id to int")
			panic(err)
		}
		userId, err := GetUserId(r)
		if err != nil {
			fmt.Println("Error getting user id")
			panic(err)
		}
		r.ParseForm()
		magnitude := r.Form.Get("magnitude")
		prescribedMagnitude := r.Form.Get("prescribedMagnitude")
		// get user
		user := store.GetUser(userId)
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
		store.SetTileMap(userId, tm)
		// return just the tile
		templ.Handler(comps.TileComponent(tm.Tiles[y][x])).ServeHTTP(w, r)
		store.AddEditedPoint(userId, comps.Coord{X: x, Y: y})
	})

	r.Get("/options", func(w http.ResponseWriter, r *http.Request) {
		templ.Handler(comps.OptionsComponent()).ServeHTTP(w, r)
	})

	r.Get("/smooth", func(w http.ResponseWriter, r *http.Request) {
		userId, err := GetUserId(r)
		if err != nil {
			http.Error(w, "user_id not found", http.StatusInternalServerError)
		}
		user := store.GetUser(userId)
		if user.TileMap.Width != 0 {
			user.TileMap.SeedData = user.TileMap.Smooth(1)
			user.TileMap.Tiles = user.TileMap.GenerateTiles()
			store.SetTileMap(userId, user.TileMap)
			templ.Handler(comps.TileMapComponent(user.TileMap)).ServeHTTP(w, r)
			store.ClearEditedPoints(userId)
		}
	})

	r.Get("/smoothedited", func(w http.ResponseWriter, r *http.Request) {
		userId, err := GetUserId(r)
		if err != nil {
			http.Error(w, "user_id not found", http.StatusInternalServerError)
		}
		user := store.GetUser(userId)
		if user.TileMap.Width != 0 {
			points := store.GetEditedPoints(user.Id)
			user.TileMap.SeedData = user.TileMap.SmoothPointsAndNeighbours(points, 1)
			user.TileMap.Tiles = user.TileMap.GenerateTiles()
			store.SetTileMap(user.Id, user.TileMap)
			templ.Handler(comps.TileMapComponent(user.TileMap)).ServeHTTP(w, r)
			store.ClearEditedPoints(user.Id)
		}
	})

	http.ListenAndServe(":8080", r)

}

func deduceXYFromId(id int, tileMap comps.TileMap) (x, y int, err error) {
	if tileMap.Width == 0 {
		// Return an error as dividing by zero is not allowed
		return 0, 0, fmt.Errorf("tileMap.Width is zero, division by zero is not allowed")
	}
	x = id % tileMap.Width
	y = id / tileMap.Width
	return x, y, nil
}

func UserCookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the cookie
		cookie, err := r.Cookie("user_id")
		if err != nil {
			log.Println("No user_id cookie found, generating new one:", err)
			// Generate a new cookie
			userID := uuid.New().String()
			http.SetCookie(w, &http.Cookie{
				Name:   "user_id",
				Value:  userID,
				Path:   "/",
				MaxAge: 60 * 60 * 24 * 365, // 1 year
			})
		}
		// Check if next handler is valid
		if next == nil {
			log.Println("Next handler is nil!")
		}

		// Add the cookie to the request context
		ctx := context.WithValue(r.Context(), "user_id", cookie.Value)
		r = r.WithContext(ctx)

		// Pass request to next handler
		next.ServeHTTP(w, r)
	})
}

func GetUserId(r *http.Request) (string, error) {
	userId, ok := r.Context().Value("user_id").(string)
	if !ok {
		return "", fmt.Errorf("user_id not found")
	}
	return userId, nil
}

func ConfigFromRequest(r *http.Request) comps.MapConfig {
	// Parse the form data
	r.ParseForm()

	// Load the default config
	config := comps.DefaultConfig()

	// Update config fields only if they exist in the form
	if val := r.Form.Get("SelectiveDistance"); val != "" {
		config.SelectiveDistance = helper.Atoi(val)
	}
	if val := r.Form.Get("WidthModifier"); val != "" {
		config.WidthModifier = helper.Atoi(val)
	}
	if val := r.Form.Get("PostSmoothDistance"); val != "" {
		config.PostSmoothDistance = helper.Atoi(val)
	}
	if val := r.Form.Get("InitialAltitude"); val != "" {
		config.InitialAltitude = comps.InitialAltitudeModifier(helper.Atoi(val))
	}
	if val := r.Form.Get("Mountains"); val != "" {
		config.Mountains = helper.Atoi(val)
	}
	if val := r.Form.Get("MountainAltitude"); val != "" {
		config.MountainAltitude = helper.Atoi(val)
	}
	if val := r.Form.Get("MountainAltitudeWindow"); val != "" {
		config.MountainAltitudeWindow = helper.Atoi(val)
	}
	if val := r.Form.Get("MountainRadius"); val != "" {
		config.MountainRadius = helper.Atoi(val)
	}
	if val := r.Form.Get("MountainRadiusWindow"); val != "" {
		config.MountainRadiusWindow = helper.Atoi(val)
	}
	if val := r.Form.Get("MountainRanges"); val != "" {
		config.MountainRanges = helper.Atoi(val)
	}
	if val := r.Form.Get("MountainRangeSize"); val != "" {
		config.MountainRangeSize = helper.Atoi(val)
	}
	if val := r.Form.Get("RangeSpread"); val != "" {
		config.RangeSpread = helper.Atoi(val)
	}
	if val := r.Form.Get("DefaultRunners"); val != "" {
		config.DefaultRunners = helper.Atoi(val)
	}
	if val := r.Form.Get("DefaultRunnerMinlength"); val != "" {
		config.DefaultRunnerMinlength = helper.Atoi(val)
	}
	if val := r.Form.Get("DefaultRunnerMaxlength"); val != "" {
		config.DefaultRunnerMaxlength = helper.Atoi(val)
	}

	return config
}
