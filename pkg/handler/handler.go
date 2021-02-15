package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oneils/todo-app/pkg/service"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	_ "github.com/oneils/todo-app/docs" // init generated docs
)

type Handler struct {
	services *service.Service
}

// NewHandler creates a new Handler with service.Service specified
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes create a new router for te app
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// add swagger to route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
