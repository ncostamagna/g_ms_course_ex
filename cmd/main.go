package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/ncostamagna/g_ms_course_ex/internal/course"
	"github.com/ncostamagna/g_ms_course_ex/pkg/bootstrap"
	"github.com/ncostamagna/g_ms_course_ex/pkg/handler"
)

func main() {

	_ = godotenv.Load()
	l := bootstrap.InitLogger()
	db, err := bootstrap.DBConnection()
	if err != nil {
		l.Fatal(err)
	}

	ctx := context.Background()
	courseRepo := course.NewRepo(db, l)
	courseSrv := course.NewService(l, courseRepo)
	h := handler.NewCourseHTTPServer(ctx, course.MakeEndpoints(courseSrv))
	port := os.Getenv("PORT")

	srv := &http.Server{
		Handler:      accessControl(h),
		Addr:         fmt.Sprintf("127.0.0.1:%s", port),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  4 * time.Second,
	}

	errCh := make(chan error)

	go func() {
		errCh <- srv.ListenAndServe()
	}()

	err = <-errCh
	if err != nil {
		log.Fatal(err)
	}

}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Accept,Authorization,Cache-Control,Content-Type,DNT,If-Modified-Since,Keep-Alive,Origin,User-Agent,X-Requested-With")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
