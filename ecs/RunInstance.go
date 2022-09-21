package ecs

import "github.com/devforma/apsara/core"

type RunInstanceRequest struct {
	core.EmbededRequest
}

func NewRunInstanceRequest() *RunInstanceRequest {
	return &RunInstanceRequest{
		core.EmbededRequest{
			Product:    "Ecs",
			Version:    "2014-05-26",
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
