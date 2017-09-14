package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)
func main() {
    r := gin.Default()
    r.Use(Test())
    r.GET("/", func(c *gin.Context) {
        if cookie, err := c.Request.Cookie("test"); err == nil {
            c.String(http.StatusOK, cookie.Value)
        } else {
            cookie := &http.Cookie{
                Name:  "king",
                Value: "shane",
            }
            http.SetCookie(c.Writer, cookie)
            c.String(http.StatusOK, "100")
        }
    })

    r.GET("/get", func( c *gin.Context ){
	if cookie, err := c.Request.Cookie("king"); err == nil {
		value := cookie.Value
		fmt.Println( value )
	}
    })
    r.Run(":8010")
}
func Test() gin.HandlerFunc {
    return func(c *gin.Context) {
        if cookie, err := c.Request.Cookie("king"); err == nil {
            value := cookie.Value
            if len(value) == 0 {
                cookie.Value = "1"
            } else {
                if v, err := strconv.Atoi(value); err == nil {
                    i := v + 1
                    cookie.Value = fmt.Sprintf("%d", i)
                }
            }
            http.SetCookie(c.Writer, cookie)
            c.Next()
        }
    }
}
