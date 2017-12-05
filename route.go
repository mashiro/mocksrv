package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route is mock route definition
type Route struct {
	Path   string
	Status int
}

// MarshalFlag marshals flag to string
func (r *Route) MarshalFlag() (string, error) {
	return fmt.Sprintf("%s=%d", r.Path, r.Status), nil
}

// UnmarshalFlag unmarshals string to flag
func (r *Route) UnmarshalFlag(value string) error {
	tokens := strings.Split(value, "=")
	if len(tokens) != 2 {
		return errors.New("invalid value")
	}

	r.Path = tokens[0]
	status, err := strconv.Atoi(tokens[1])
	if err != nil {
		return err
	}
	r.Status = status
	return nil
}

// Respond the route request
func (r *Route) Respond(c *gin.Context) {
	c.Status(r.Status)
}

var anyRoute = Route{Path: "/*paths", Status: 200}

type routeMap map[string][]Route

// RouteOptions is a route options
type RouteOptions struct {
	AnyRoutes     []Route `long:"any" description:"Any routes"`
	GetRoutes     []Route `long:"get" description:"GET routes"`
	PostRoutes    []Route `long:"post" description:"POST routes"`
	PutRoutes     []Route `long:"put" description:"PUT routes"`
	PatchRoutes   []Route `long:"patch" description:"PATCH routes"`
	HeadRoutes    []Route `long:"head" description:"HEAD routes"`
	OptionsRoutes []Route `long:"options" description:"OPTIONS routes"`
	DeleteRoutes  []Route `long:"delete" description:"DELETE routes"`
	ConnectRoutes []Route `long:"connect" description:"CONNECT routes"`
	TraceRoutes   []Route `long:"trace" description:"TRACE routes"`
}

// RouteMap returns method with routes map
func (o *RouteOptions) RouteMap() routeMap {
	return routeMap{
		"GET":     o.GetRoutes,
		"POST":    o.PostRoutes,
		"PUT":     o.PutRoutes,
		"PATCH":   o.PatchRoutes,
		"HEAD":    o.HeadRoutes,
		"OPTIONS": o.OptionsRoutes,
		"DELETE":  o.DeleteRoutes,
		"CONNECT": o.ConnectRoutes,
		"TRACE":   o.TraceRoutes,
	}
}

func init() {
	var options RouteOptions
	RegisterModule("Route Options", &options, func(engine *gin.Engine) {
		var routeCount int

		for _, r := range options.AnyRoutes {
			engine.Any(r.Path, r.Respond)
			routeCount++
		}

		for method, routes := range options.RouteMap() {
			for _, r := range routes {
				engine.Handle(method, r.Path, r.Respond)
				routeCount++
			}
		}

		if routeCount == 0 {
			engine.Any(anyRoute.Path, anyRoute.Respond)
		}
	})
}
