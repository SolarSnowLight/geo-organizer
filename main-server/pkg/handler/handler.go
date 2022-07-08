package handler

import (
	"main-server/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	_ "main-server/docs"

	route "main-server/pkg/constants/route"

	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

/* Инициализация маршрутов */
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.MaxMultipartMemory = 50 << 20 // 50 MiB
	router.Static("/public", "./public")

	router.LoadHTMLGlob("pkg/templates/*")

	// Настройка CORS-политики
	router.Use(cors.New(cors.Config{
		//AllowAllOrigins: true, // для тестов
		AllowOrigins:     []string{viper.GetString("client_url")},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group(route.AUTH_MAIN_ROUTE)
	{
		auth.POST(route.AUTH_SIGN_UP_ROUTE, h.signUp)
		auth.POST(route.AUTH_SIGN_IN_ROUTE, h.signIn)
		auth.POST(route.AUTH_SIGN_IN_GOOGLE_ROUTE, h.signInOAuth2)
		auth.GET(route.AUTH_ACTIVATE_ROUTE, h.activate)

		// With middlewares (for get data from access token)
		auth.POST(route.AUTH_REFRESH_TOKEN_ROUTE, h.userIdentityLogout, h.refresh)
		auth.POST(route.AUTH_LOGOUT_ROUTE, h.userIdentity, h.logout)
	}

	user := router.Group(route.USER_MAIN_ROUTE, h.userIdentity)
	{
		user.POST(route.USER_CREATE_ARTICLE_ROUTE, h.createArticle)
		user.POST(route.USER_DELETE_ARTICLE_ROUTE, h.deleteArticle)
		user.POST(route.USER_GET_ARTICLE_ROUTE, h.getArticle)
		user.POST(route.USER_GET_ARTICLES_ROUTE, h.getArticles)
	}

	/*api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}*/

	return router
}
