package main

import (
	"embed"
	"flag"
	"log"
	"net/http"
	"time"
	"github.com/zsmatrix62/templ-goat/bootstrap"
	"github.com/zsmatrix62/templ-goat/config"
	pkgConfig "github.com/zsmatrix62/templ-goat/pkg/config"
	"github.com/zsmatrix62/templ-goat/pkg/logger"

	"github.com/getsentry/sentry-go"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

//go:embed public/*
var publicFS embed.FS

func init() {
	config.Initialize()
}

func setupSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://848d371a4a19a0ae3aaee1040908106d@o249661.ingest.us.sentry.io/4508351896879104",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

//go:generate pnpm run build
func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()

	pkgConfig.InitConfig(env)

	r := bootstrap.SetupRoutes(publicFS)

	bootstrap.SetupLogger()
	bootstrap.SetupDB()

	setupSentry()

	g = errgroup.Group{}
	g.Go(func() error {
		logger.Info("Server started at " + pkgConfig.Get("app.url"))
		srv := &http.Server{
			Handler:      r,
			Addr:         pkgConfig.Get("app.url"),
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		}
		return srv.ListenAndServe()
	})
	logger.LogIf(g.Wait())
}
