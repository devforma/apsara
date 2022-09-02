package ecs

import "github.com/devforma/apsara/core"

type RunInstanceRequest struct {
	core.EmbededRequest
}

func NewRunInstanceRequest() *RunInstanceRequest {
	return &RunInstanceRequest{
		core.EmbededRequest{
			Product:    "ecs",
			Version:    "2014-12-12",
			Action:     "RunInstance",
			Style:      core.RequestStyleRPC,
			BizQueries: map[string]string{},
		},
	}
}

type RunInstanceResponse struct {
	core.EmbededResponse
}

func (r *RunInstanceResponse) GetInfo() string {
	return r.Headers["adaw"]
}
