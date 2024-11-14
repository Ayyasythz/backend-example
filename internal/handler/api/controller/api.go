package controller

import (
	"github.com/gofiber/swagger"
	_ "sagara_backend_test/docs"
	"sagara_backend_test/internal/usecases"
	"sagara_backend_test/lib/router"
	"time"
)

type API struct {
	prefix         string
	port           uint
	readTimeout    time.Duration
	writeTimeout   time.Duration
	requestTimeout time.Duration
	enableSwagger  bool
	wardrobeUc     usecases.WardrobeUseCases
}

type Options struct {
	Prefix         string
	Port           uint
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	RequestTimeout time.Duration
	EnableSwagger  bool
	WardrobeUc     usecases.WardrobeUseCases
}

func New(opts *Options) *API {
	return &API{
		prefix:         opts.Prefix,
		port:           opts.Port,
		readTimeout:    opts.ReadTimeout,
		writeTimeout:   opts.WriteTimeout,
		requestTimeout: opts.RequestTimeout,
		enableSwagger:  opts.EnableSwagger,
		wardrobeUc:     opts.WardrobeUc,
	}
}

func (api *API) RegisterRoute() *router.FastRouter {
	myRouter := router.New(&router.Options{
		Prefix:         api.prefix,
		Port:           api.port,
		ReadTimeout:    api.readTimeout,
		WriteTimeout:   api.writeTimeout,
		RequestTimeout: api.requestTimeout,
	})

	if api.enableSwagger {
		myRouter.CustomHandler("GET", "/docs/*", swagger.HandlerDefault, router.MustAuthorized(false))
	}

	myRouter.GET("/health", api.Ping, router.MustAuthorized(false))

	myRouter.Group("/v1", func(v1 *router.FastRouter) {
		v1.Group("/wardrobe", func(wardrobe *router.FastRouter) {
			wardrobe.GET("/search", api.Search, router.MustAuthorized(false))
			wardrobe.GET("/ready", api.GetAvailable, router.MustAuthorized(false))
			wardrobe.GET("/out", api.GetUnavailable, router.MustAuthorized(false))
			wardrobe.GET("/less", api.GetLessThan, router.MustAuthorized(false))
			wardrobe.PUT("/:id", api.Update, router.MustAuthorized(false))
			wardrobe.GET("/:id", api.GetById, router.MustAuthorized(false))
			wardrobe.DELETE("/:id", api.Delete, router.MustAuthorized(false))
			wardrobe.PUT("/:id/addStock", api.AddStock, router.MustAuthorized(false))
			wardrobe.PUT("/:id/subStock", api.SubStock, router.MustAuthorized(false))
			wardrobe.GET("", api.GetAll, router.MustAuthorized(false))
			wardrobe.POST("", api.Insert, router.MustAuthorized(false))
		})
	})

	//myRouter.Group("/v1", func(v1 *router.FastRouter) {
	//	v1.Group("/accounts", func(account *router.FastRouter) {
	//		account.POST("/register", api.RegisterAccount, router.MustAuthorized(false))
	//		account.GET("/me", api.GetAccount)
	//		account.PATCH("/update", api.UpdateAccount)
	//	})
	//	v1.Group("/users", func(user *router.FastRouter) {
	//		user.GET("/me", api.GetUser)
	//		user.PATCH("/update", api.UpdateUser)
	//		user.PATCH("/password", api.ChangePassword)
	//		user.POST("/login", api.LoginUser, router.MustAuthorized(false))
	//	})
	//
	//})

	return myRouter
}
