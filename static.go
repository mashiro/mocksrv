package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Static struct {
	Path      string
	ServePath string
}

// MarshalFlag marshals flag to string
func (s *Static) MarshalFlag() (string, error) {
	return fmt.Sprintf("%s:%s", s.Path, s.ServePath), nil
}

// UnmarshalFlag unmarshals string to flag
func (s *Static) UnmarshalFlag(value string) error {
	tokens := strings.Split(value, ":")
	if len(tokens) != 2 {
		return errors.New("invalid value")
	}

	s.Path = tokens[0]
	s.ServePath = tokens[1]
	return nil
}

type ServingOptions struct {
	Roots []Static `long:"root" description:"Serving static roots"`
	Files []Static `long:"file" description:"Serving static files"`
}

func init() {
	var options ServingOptions
	RegisterModule("Serving Options", &options, func(engine *gin.Engine) {
		for _, root := range options.Roots {
			engine.StaticFS(root.Path, http.Dir(root.ServePath))
		}
		for _, file := range options.Files {
			engine.StaticFile(file.Path, file.ServePath)
		}
	})
}
