package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/pocketgrok/art-gallery/migrations"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/template"
)

func main() {
	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		htmlTemplateRegistry := template.NewRegistry()

		// Redirect WebFinger requests to Mastodon subdomain.
		e.Router.GET("/.well-known/webfinger", func(c echo.Context) error {
			request := c.Request()
			url := request.URL
			url.Host = "mastodon." + request.Host
			return c.Redirect(302, url.String())
		})

		// Home page: Art Gallery
		e.Router.GET("/", func(c echo.Context) error {
			commissions, err := app.Dao().FindRecordsByExpr("commissions")
			if err != nil {
				return apis.NewApiError(
					http.StatusInternalServerError,
					"Unable to load commissions.",
					err,
				)
			}
			errs := app.Dao().ExpandRecords(commissions, []string{"artists", "images(commission)"}, nil)
			if len(errs) > 0 {
				return apis.NewApiError(500, "Commission artists not found.", nil)
			}
			html, err := htmlTemplateRegistry.LoadFiles(
				"views/layout.go.html",
				"views/gallery.go.html",
			).Render(map[string]any{
				"Commissions": commissions,
			})
			if err != nil {
				return apis.NewApiError(
					http.StatusInternalServerError,
					"Unable to render HTML.",
					err,
				)
			}
			return c.HTML(http.StatusOK, html)
		})

		// View commission page
		e.Router.GET("/commissions/:id", func(c echo.Context) error {
			id := c.PathParam("id")
			commission, err := app.Dao().FindRecordById("commissions", id)
			if err != nil {
				return apis.NewNotFoundError("Commission not found.", err)
			}
			errs := app.Dao().ExpandRecord(commission, []string{"artists", "images(commission)"}, nil)
			if len(errs) > 0 {
				return apis.NewApiError(500, "Commission artists not found.", nil)
			}
			html, err := htmlTemplateRegistry.LoadFiles(
				"views/layout.go.html",
				"views/commissions.go.html",
			).Render(map[string]any{
				"Id":         id,
				"Commission": commission,
			})
			if err != nil {
				// or redirect to a dedicated 404 HTML page
				return apis.NewNotFoundError("Could not render record.", err)
			}
			return c.HTML(http.StatusOK, html)
		})

		// Serve static files from the provided public dir (if exists).
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
