package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rulyadhika/fga_digitalent_assignment_2/model/web"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			switch err.Err.(type) {

			case *NotFoundError:
				handlerError(ctx, http.StatusNotFound, err)
			case *BadRequestError:
				handlerError(ctx, http.StatusBadRequest, err)
			case *UnprocessableEntityError:
				handlerError(ctx, http.StatusUnprocessableEntity, err)
			default:
				handlerError(ctx, http.StatusInternalServerError, err)
			}

			ctx.Abort()
		}
	}
}

func handlerError(ctx *gin.Context, httpStatusCode int, err error) {
	response := &web.WebResponse{
		Status:  http.StatusText(httpStatusCode),
		Code:    httpStatusCode,
		Message: err.Error(),
		Data:    nil,
	}

	ctx.AbortWithStatusJSON(httpStatusCode, response)
}
