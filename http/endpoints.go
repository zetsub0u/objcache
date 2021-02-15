package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetObject is the endpoint to retrieve objects
// @Summary get object
// @Description get an object from the store tha can be released later
// @Success 200 {object} int
// @Router /object [get]
func (s *Server) GetObject(c *gin.Context) {
	c.JSON(http.StatusOK, s.mgr.GetObject())
}

// ReturnObject is the endpoint to return objects into the store.
// @Summary return the object
// @Description return an object into the store
// @Param obj path int true "object to return"
// @Success 200
// @Router /object/{obj} [put]
func (s *Server) ReturnObject(c *gin.Context) {
	obj := c.Param("obj")
	if obj == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	intObj, err := strconv.Atoi(obj)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	s.mgr.FreeObject(&intObj)
	c.Status(http.StatusOK)
	return
}

type CreateReq struct {
	Obj int `json:"obj"`
}

// CreateObject is the endpoint to create objects, it receives them from a json body (took liberty here)
// @Summary create a new object in the store
// @Description create  an object into the store
// @Accept json
// @Success 200
// @Param req body CreateReq true "object to create"
// @Router /object [post]
func (s *Server) CreateObject(c *gin.Context) {
	var req CreateReq

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	s.mgr.FreeObject(&req.Obj)
	c.Status(http.StatusOK)
	return
}
