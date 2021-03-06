package API

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"todorokvaja1/DataStructures"
)

func (a *Controller) GetOpravilo(c *gin.Context) {
	//[]DataStructure Opravilo - Array vseh opravil
	fmt.Println(c.Get("user_id"))
	opravilo, err := a.c.GetOpravilo(c.Request.Context())
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
	}

	// Vse ok - Status 200
	c.JSON(http.StatusOK, opravilo)
}
func (a *Controller) GetOpraviloById(c *gin.Context) {

	opravilo, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	OpraviloId, err := a.c.GetOpraviloById(c.Request.Context(), opravilo)
	if err != nil {
		//Vrne error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}
	//JSON serializacija
	c.JSON(http.StatusOK, OpraviloId)
}

func (a *Controller) InsertOpravilo(c *gin.Context) {

	var opravilo DataStructures.Opravilo

	err := c.BindJSON(&opravilo)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	err = a.c.InsertOpravilo(c.Request.Context(), opravilo)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	c.String(http.StatusOK, "vstavljam novo opravilo")
}

func (a *Controller) RemoveOpravilo(c *gin.Context) {

	opravilo, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		//log.Fatalf("sentry.Init: %s", err)
		sentry.CaptureException(err)      // Lovljenje errorjev na Sentry.io
		log.Printf("Sentry.init %s", err) //console log, ni fatal, da ne terminata vsega skupaj
		return
	}

	err = a.c.RemoveOpravilo(c.Request.Context(), opravilo)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}
	c.String(http.StatusOK, "odstranjujem opravilo")
}

func (a *Controller) UpdateOpravilo(c *gin.Context) {

	opraviloID, err := primitive.ObjectIDFromHex(c.Param("todo_id"))
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	var opraviloUpdate DataStructures.Opravilo
	err = c.BindJSON(&opraviloUpdate)
	if err != nil {
		//Vrne error 409 - Conflict
		c.String(http.StatusConflict, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	err = a.c.UpdateOpravilo(c.Request.Context(), opraviloID, opraviloUpdate)
	if err != nil {
		//Vrne error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		sentry.CaptureException(err)
		log.Printf("Sentry.init %s", err)
		return
	}

	c.String(http.StatusOK, "popravljam novo opravilo")
}
