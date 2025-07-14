package v1

import (
	"log/slog"
	"net/http"
	"strconv"
	"test/internal/common"
	"test/internal/dto"
	"test/internal/usecase"

	"github.com/gin-gonic/gin"
)

type subRoutes struct {
	log *slog.Logger
	s   *usecase.SubService
}

func NewSubscriptionRoutes(log *slog.Logger, handler *gin.RouterGroup, s *usecase.SubService) {
	r := &subRoutes{
		log: log,
		s:   s,
	}

	g := handler.Group("/subscriptions")
	{
		g.POST("", r.create)
		g.GET("", r.get)
		g.PUT("/:id", r.update)
		g.GET("/:id", r.getByID)
		g.DELETE("/:id", r.delete)
		g.POST("/sum", r.sum)
	}
}

func handlErr(c *gin.Context, log *slog.Logger, err error) {
	log.Error(err.Error())
	status, err := common.ParseErr(err)
	c.JSON(status, gin.H{"error": err.Error()})
}

func getIDFromURL(c *gin.Context) (int, error) {
	urlParam, ok := c.Params.Get("id")
	if !ok {
		return -1, common.ErrBadURL
	}

	id, err := strconv.Atoi(urlParam)
	if err != nil {
		return -1, common.ErrBadType
	}

	return id, nil
}

// @Summary     Create subscription
// @Description Create subscription
// @ID          CreateSubscription
// @Tags  	    Subscritpion
// @Accept      json
// @Param 		subscription body dto.SubscritpionRequest false "Subscritpion creation data"
// @Produce     json
// @Success     200 {object} dto.SubscritpionResponse
// @Failure     400
// @Failure     500
// @Router      /subscriptions [post]
func (r *subRoutes) create(c *gin.Context) {
	const op = "subRoutes.create"
	log := r.log.With(
		slog.String("op", op),
	)

	var req *dto.SubscritpionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handlErr(c, log, err)
		return
	}

	ent, err := req.ToEntity()
	if err != nil {
		handlErr(c, log, err)
		return
	}

	sub, err := r.s.Create(c.Request.Context(), *ent)
	if err != nil {
		handlErr(c, log, err)
		return
	}

	resp := dto.ToSubscriptionResponse(sub)
	c.JSON(http.StatusOK, resp)
}

// @Summary     Update subscription
// @Description Update subscription
// @ID          UpdateSubscription
// @Tags  	    Subscritpion
// @Accept      json
// @Param 		subscription body dto.SubscritpionRequest false "Subscritpion update data"
// @Produce     json
// @Success     200 {object} dto.SubscritpionResponse
// @Failure     400
// @Failure     404
// @Failure     500
// @Router      /subscriptions/{id} [put]
func (r *subRoutes) update(c *gin.Context) {
	const op = "subRoutes.update"
	log := r.log.With(
		slog.String("op", op),
	)

	id, err := getIDFromURL(c)
	if err != nil {
		handlErr(c, log, err)
		return
	}

	_, err = r.s.GetByID(c.Request.Context(), id)
	if err != nil {
		handlErr(c, log, err)
		return
	}

	var req *dto.SubscritpionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handlErr(c, log, err)
		return
	}

	ent, err := req.ToEntity()
	if err != nil {
		handlErr(c, log, err)
		return
	}
	ent.ID = id

	sub, err := r.s.Update(c.Request.Context(), *ent)
	if err != nil {
		handlErr(c, log, err)
		return
	}

	resp := dto.ToSubscriptionResponse(sub)
	c.JSON(http.StatusOK, resp)
}

// @Summary     Get subscription by id
// @Description Get subscription by id
// @ID          GetSubscriptionByID
// @Tags  	    Subscritpion
// @Produce     json
// @Success     200 {object} dto.SubscritpionResponse
// @Failure     400
// @Failure     404
// @Failure     500
// @Router      /subscriptions/{id} [get]
func (r *subRoutes) getByID(c *gin.Context) {
	const op = "subRoutes.getByID"
	log := r.log.With(
		slog.String("op", op),
	)

	id, err := getIDFromURL(c)
	if err != nil {
		handlErr(c, log, err)
		return
	}

	sub, err := r.s.GetByID(c.Request.Context(), id)
	if err != nil {
		handlErr(c, r.log, err)
		return
	}

	resp := dto.ToSubscriptionResponse(sub)
	c.JSON(http.StatusOK, resp)
}

// @Summary     Delete subscription
// @Description Delete subscription
// @ID          DeleteSubscritpion
// @Tags  	    Subscritpion
// @Success     200
// @Failure     400
// @Failure     500
// @Router      /subscriptions/{id} [delete]
func (r *subRoutes) delete(c *gin.Context) {
	const op = "subRoutes.delete"
	log := r.log.With(
		slog.String("op", op),
	)

	id, err := getIDFromURL(c)
	if err != nil {
		handlErr(c, log, err)
		return
	}

	err = r.s.Delete(c.Request.Context(), id)
	if err != nil {
		handlErr(c, r.log, err)
		return
	}

	c.JSON(http.StatusOK, "")
}

// @Summary     Get subscriptions
// @Description Get subscriptions
// @ID          GetSubscriptions
// @Tags  	    Subscritpion
// @Produce     json
// @Success     200 {object} dto.SubscritpionsResponse
// @Failure     500
// @Router      /subscriptions [get]
func (r *subRoutes) get(c *gin.Context) {
	const op = "subRoutes.get"
	log := r.log.With(
		slog.String("op", op),
	)

	sub, err := r.s.GetAll(c.Request.Context())
	if err != nil {
		handlErr(c, log, err)
		return
	}

	resp := dto.ConvertToSubscriptionsResponse(sub)
	c.JSON(http.StatusOK, resp)
}

// @Summary     Get subscriptions price sum
// @Description Get subscriptions price sum
// @ID          GetSubscriptionsPriceSum
// @Tags  	    Subscritpion
// @Accept      json
// @Param 		filter body dto.FilterRequest false "Filter"
// @Produce     json
// @Success     200 {object} dto.SumResponse
// @Failure     400
// @Failure     500
// @Router      /subscriptions/sum [post]
func (r *subRoutes) sum(c *gin.Context) {
	const op = "subRoutes.sum"
	log := r.log.With(
		slog.String("op", op),
	)

	var req *dto.FilterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handlErr(c, log, err)
		return
	}

	ent, err := req.ToEntity()
	if err != nil {
		handlErr(c, log, err)
		return
	}

	resp, err := r.s.GetSum(c.Request.Context(), *ent)
	if err != nil {
		handlErr(c, log, err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
