package routes

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/server/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	db  []domain.Ticket
}

func NewRouter(eng *gin.Engine, db []domain.Ticket) *router {
	return &router{eng, db}
}

func (r *router) MapRoutes() {
	rp := tickets.NewRepository(r.db)
	sv := tickets.NewService(rp)
	h := handler.NewService(sv)

	route := r.eng.Group("/ticket")

	route.GET("/ping", func(ctx *gin.Context) { ctx.String(200, "Pong") })

	route.GET("/getByCountry/:dest", h.GetTicketsByCountry())
	route.GET("getAverage/:dest", h.AverageDestination())
}
