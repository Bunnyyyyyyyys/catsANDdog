package main

import "github.com/gin-gonic/gin"

func catRouter(router gin.IRouter) {
	router.POST("/api/pets/add", AddCat)

	router.GET("/api/pets", GetAllCats)

	router.GET("/api/pets/:id", GetCat)

	router.DELETE("/api/pets/:id", DeleteCat)

	router.PUT("/api/pets/:id", EditCat)

}
