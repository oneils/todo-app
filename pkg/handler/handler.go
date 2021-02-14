package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oneils/todo-app/pkg/service"
)

type Handler struct {
	services *service.Service
}

// New creates a new Handler with service.Service specified
func New(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes create a new router for te app
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.POST("/", h.createList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := api.Group("/items")
		{
			items.GET("/", h.getAllItems)
			items.GET("/:id", h.getItemById)
			items.POST("/", h.createItem)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
