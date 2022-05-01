package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/renatoviolin/go-node/dto"
	"github.com/renatoviolin/go-node/usecase"
)

type GinHandler struct {
	gin *gin.Engine
	uc  *usecase.FindAll
}

func NewHandler(uc *usecase.FindAll) *GinHandler {
	gin.SetMode(gin.ReleaseMode)
	return &GinHandler{gin: gin.New(), uc: uc}
}

func (h *GinHandler) SetupAndRun(address string) {
	h.gin.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	h.gin.GET("/json", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dto.NewPayload())
	})

	h.gin.GET("/mongo", func(ctx *gin.Context) {
		result, err := h.uc.Execute()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		ctx.JSON(http.StatusOK, result)
	})
	log.Printf("Gin running at %s\n", address)
	h.gin.Run(address)
}
