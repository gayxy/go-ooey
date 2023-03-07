// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	script "github.com/gayxy/go-ooey/app/ooey/internal/handler/script"
	"github.com/gayxy/go-ooey/app/ooey/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/scrpt/:id",
				Handler: script.GetScriptInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/ooey/v1"),
	)
}
