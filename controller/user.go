package controller

import (
	"net/http"
	"strconv"

	"github.com/himu62/echo-testing/model"
	"github.com/labstack/echo"
)

type (
	User struct {
		store []*model.User
	}
)

func AddRoutes(e *echo.Echo) {
	ctl := &User{
		store: make([]*model.User, 0, 30),
	}

	g := e.Group("/users")

	g.GET("/:id", ctl.fetch)
	g.GET("", ctl.list)
	g.POST("", ctl.create)
	g.PUT("/:id", ctl.update)
	g.DELETE("/:id", ctl.remove)
}

func (ctl *User) list(c echo.Context) error {
	return c.JSON(http.StatusOK, ctl.store)
}

func (ctl *User) fetch(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "invalid user id"})
	}
	if id < 0 || id > len(ctl.store) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "invalid user id"})
	}
	return c.JSON(http.StatusOK, ctl.store[id])
}

func (ctl *User) create(c echo.Context) error {
	user := &model.User{
		Name:  c.FormValue("name"),
		Email: c.FormValue("email"),
	}

	if valid, errs := user.Valid(); !valid {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errs})
	}

	ctl.store = append(ctl.store, user)

	return c.JSON(http.StatusOK, user)
}

func (ctl *User) update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "invalid user id"})
	}
	if id < 0 || id > len(ctl.store) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "invalid user id"})
	}

	user := &model.User{
		Name:  c.FormValue("name"),
		Email: c.FormValue("email"),
	}

	if valid, errs := user.Valid(); !valid {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errs})
	}

	ctl.store[id] = user

	return c.JSON(http.StatusOK, user)
}

func (ctl *User) remove(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "invalid user id"})
	}
	if id < 0 || id > len(ctl.store) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "invalid user id"})
	}

	ctl.store = append(ctl.store[:id], ctl.store[id+1:]...)

	return nil
}
