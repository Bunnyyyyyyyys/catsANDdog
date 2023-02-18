package main

import "github.com/gin-gonic/gin"

func dogRouter(router gin.IRouter) {
	router.POST("/api/pets/add", AddDog)

	router.GET("/api/pets", GetAllDogs)

	router.GET("/api/pets/:id", GetDog)

	router.DELETE("/api/pets/:id", DeleteDog)

	router.PUT("/api/pets/:id", EditDog)

}
