package solar

import "github.com/gin-gonic/gin"

func RunApp(app *gin.Engine) {
	r := app.Group("/solar")
	{
		r.POST("", process)
	}
}
