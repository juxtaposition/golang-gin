package main

import (
	"fmt"
	"os"
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

router.Run(":3000")

// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)

	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	signal.Notify(c, os.Interrupt)
	// Block until any signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
	os.Exit(0)

}

