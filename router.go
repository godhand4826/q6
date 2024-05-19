package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Router interface {
	BindOn(r *gin.Engine)
}

type matchingSystemRouter struct {
	m MatchingSystem
}

func NewMatchingSystemRouter(m MatchingSystem) Router {
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

func (r *matchingSystemRouter) AddSinglePersonAndMatchHandler(c *gin.Context) {
	type Request struct {
		Name   string `json:"name"`
		Height int    `json:"height"`
		Gender string `json:"gender"`
		Dates  int    `json:"dates"`
	}

	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}

	id, err := r.m.AddSinglePersonAndMatch(req.Name, req.Height, Gender(req.Gender), req.Dates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (r *matchingSystemRouter) RemoveSinglePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := r.m.RemoveSinglePerson(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (r *matchingSystemRouter) QuerySingleMale(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid n"})
		return
	}

	r.m.QuerySinglePeople(n, GENDER_MALE)
}

func (r *matchingSystemRouter) QuerySingleFemale(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid n"})
		return
	}

	users, err := r.m.QuerySinglePeople(n, GENDER_FEMALE)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
