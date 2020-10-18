package app

import "github.com/laur1s/bookstore_users-api/controllers"

func mapUrls()  {
	router.GET("/ping",controllers.Ping)
}