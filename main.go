package main

import (
	"fmt"
	"os"
	"routerEze"
	"os/signal"
	"net/http"
    "github.com/gin-gonic/contrib/static"
    "github.com/gin-gonic/gin"
)

func main() {
// Making the router
router := gin.Default();

// Setting static files path
router.Use(static.Serve("/", static.LocalFile("./views", true)))

// setting router Group for API
api := router.Group("/api") 
// New scope 
{
api.GET("/", func (c *gin.Context) {
c.JSON(http.StatusOK, gin.H {
"message": "dude",
})
})
}
reciveData := make (chan string, 1)

sendToChannel := routerEze.MyChan{Channel: reciveData}

sendToChannel.SendData("Just Nothing")
fmt.Println(<-sendToChannel.Channel)

pepito := routerEze.Person{ Name: "Teresa", Age: 29}
fmt.Println(pepito.SayHello("Hello "))

router.Run(":3000")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println("Got signal:", s)
	os.Exit(0)

}

