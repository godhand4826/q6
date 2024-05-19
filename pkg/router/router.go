package router

import (
	"net/http"
	"q6/lib/serve"
	"q6/pkg/entity"
	"q6/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

var _ serve.Router = (*matchingSystemRouter)(nil)

type matchingSystemRouter struct {
	m service.MatchingSystem
}

func NewMatchingSystemRouter(m service.MatchingSystem) serve.Router {
	return &matchingSystemRouter{
		m: m,
	}
}

func (r *matchingSystemRouter) BindOn(e *gin.Engine) {
	e.POST("/people", r.AddSinglePersonAndMatchHandler)
	e.GET("/people/male", r.QuerySingleMale)
	e.GET("/people/female", r.QuerySingleFemale)
	e.DELETE("/people/:id", r.RemoveSinglePerson)
}

// @Summary		AddSinglePersonAndMatchHandler
// @Router			/people [post]
// @Accept			json
// @Produce		json
// @Param			name	body		router.AddSinglePersonAndMatchHandler.Request	true	"user info"
// @Success		200		{object}	router.AddSinglePersonAndMatchHandler.Response
// @Failure		400		{object}	serve.HTTPError
// @Description	Add a new user to the matching system and find any possible matches for the new user.
// @Description	Gender must be "MALE" or "FEMALE"
// @Description	Dates must be greater than zero
// @Description	Returns the ID of the user.
func (r *matchingSystemRouter) AddSinglePersonAndMatchHandler(c *gin.Context) {
	type Request struct {
		Name   string `json:"name"`
		Height int    `json:"height"`
		Gender string `json:"gender"`
		Dates  int    `json:"dates"`
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		serve.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.m.AddSinglePersonAndMatch(req.Name, req.Height, entity.Gender(req.Gender), req.Dates)
	if err != nil {
		serve.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	type Response struct {
		ID int `json:"id"`
	}

	c.JSON(http.StatusCreated, Response{ID: id})
}

// @Summary		RemoveSinglePerson
// @Router			/people/{id} [delete]
// @Param			id	path	int	true	"user ID"
// @Success		204
// @Failure		400	{object}	serve.HTTPError
// @Failure		404	{object}	serve.HTTPError
// @Description	Remove a user from the matching system so that the user cannot be matched anymore.
func (r *matchingSystemRouter) RemoveSinglePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		serve.SendError(c, http.StatusBadRequest, "invalid id")
		return
	}
	if err := r.m.RemoveSinglePerson(id); err != nil {
		serve.SendError(c, http.StatusNotFound, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary		QuerySingleMale
// @Router			/people/male [get]
// @Param			n	query		int	true	"query size"
// @Success		200	{object}	[]entity.MatchRequest
// @Failure		400	{object}	serve.HTTPError
// @Description	Find the most N possible matched single male people, where N is a request parameter.
func (r *matchingSystemRouter) QuerySingleMale(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		serve.SendError(c, http.StatusBadRequest, "invalid n")
		return
	}

	result, err := r.m.QuerySinglePeople(n, entity.GenderMale)
	if err != nil {
		serve.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary		QuerySingleFemale
// @Description	Find the most N possible matched single female people, where N is a request parameter.
// @Param			n	query		int	true	"query size"
// @Success		200	{object}	[]entity.MatchRequest
// @Failure		400	{object}	serve.HTTPError
// @Router			/people/female [get]
func (r *matchingSystemRouter) QuerySingleFemale(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		serve.SendError(c, http.StatusBadRequest, "invalid n")
		return
	}

	users, err := r.m.QuerySinglePeople(n, entity.GenderFemale)
	if err != nil {
		serve.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}
