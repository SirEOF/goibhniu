package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/emperorcow/goibhniu/common"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	data       *common.FermentationData
	router     *gin.Engine
	listenPort int
	listenHost string
}

/* Get a new API server to allow the main routing to pass information to us
 */
func NewAPIServer(data *common.FermentationData, port int, host string) *APIServer {
	return &APIServer{
		data:       data,
		router:     gin.New(),
		listenPort: port,
		listenHost: host,
	}
}

func (this *APIServer) sendError(c *gin.Context, code int, msg string, err error) {
	c.JSON(code, gin.H{
		"message": msg,
		"error":   err.Error(),
	})
	log.WithFields(log.Fields{
		"error": err.Error(),
	}).Error(msg)
}

func (this *APIServer) sendResponseData(c *gin.Context, data interface{}) {
	dataJSON, err := json.Marshal(data)

	if err != nil {
		this.sendError(c, 500, "Unable to parse output data.", err)
	} else {
		c.Data(200, "application/json", dataJSON)
	}
}

func (this *APIServer) sendSuccess(c *gin.Context, msg string) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": msg,
	})
}

func (this *APIServer) QueryVessel(c *gin.Context) {
	log.Debug("Getting all vessels for API.")
	this.sendResponseData(c, this.data.Fermenters)
}

func (this *APIServer) CreateVessel(c *gin.Context) {
	log.Debug("Creating new vessel.")

	var newVessel common.Vessel

	err := c.BindJSON(&newVessel)
	if err != nil {
		this.data.Fermenters.Add(newVessel)
		this.sendSuccess(c, "Vessel added.")
	} else {
		this.sendError(c, 500, "Unable to process input vessel.", err)
	}
}

func (this *APIServer) ReadVessel(c *gin.Context) {
	id := c.Param("id")
	log.WithFields(log.Fields{
		"id": id,
	}).Debug("Getting individual vessel.")

	vessel, err := this.data.Fermenters.Get(id)

	if err != nil {
		this.sendError(c, 404, "That fermenter does not exist.", err)
	} else {
		this.sendResponseData(c, vessel)
	}
}

func (this *APIServer) UpdateVessel(c *gin.Context) {
	var newVessel common.Vessel

	err := c.BindJSON(&newVessel)
	if err != nil {
		this.sendError(c, 500, "Unable to proccess input vessel.", err)
	}

	err = this.data.Fermenters.Update(newVessel)
	if err != nil {
		this.sendError(c, 404, "Unable to update vessel.", err)
	} else {
		this.sendSuccess(c, "Vessel updated.")
	}
}

func (this *APIServer) DeleteVessel(c *gin.Context) {
	id := c.Param("id")
	log.WithFields(log.Fields{
		"id": id,
	}).Debug("Deleting individual vessel.")

	err := this.data.Fermenters.Delete(id)
	if err != nil {
		this.sendError(c, 404, "Unable to delete vessel.", err)
	} else {
		this.sendSuccess(c, "Vessel deleted.")
	}
}

func (this *APIServer) QueryFermentable(c *gin.Context) {
	log.Debug("Getting all fermentables.")
	this.sendResponseData(c, this.data.Brews.GetAll())
}

func (this *APIServer) CreateFermentable(c *gin.Context) {
	log.Debug("Creating new fermentable.")

	var newBrew common.Fermentable

	err := c.BindJSON(&newBrew)
	if err != nil {
		this.data.Brews.Add(newBrew)
		this.sendSuccess(c, "Fermentable added.")
	} else {
		this.sendError(c, 500, "Unable to process input data.", err)
	}
}

func (this *APIServer) ReadFermentable(c *gin.Context) {
	idStr := c.Param("id")
	log.WithFields(log.Fields{
		"id": idStr,
	}).Debug("Getting individual fermentable.")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.sendError(c, 400, "Input ID was not a number.", err)
		return
	}

	brew, err := this.data.Brews.Get(id)

	if err != nil {
		this.sendError(c, 404, "That fermentable does not exist.", err)
	} else {
		this.sendResponseData(c, brew)
	}
}

func (this *APIServer) UpdateFermentable(c *gin.Context) {
	var newBrew common.Fermentable

	err := c.BindJSON(&newBrew)
	if err != nil {
		this.sendError(c, 500, "Unable to proccess input data.", err)
		return
	}

	err = this.data.Brews.Update(newBrew)
	if err != nil {
		this.sendError(c, 404, "Unable to update fermentable.", err)
	} else {
		this.sendSuccess(c, "Fermentable updated successfully.")
	}
}

func (this *APIServer) DeleteFermentable(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.sendError(c, 400, "Input ID was not a number.", err)
		return
	}

	this.data.Brews.Delete(id)
	this.sendSuccess(c, "Fermentable deleted.")
}

func (this *APIServer) AddFermentableReading(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.sendError(c, 400, "Input ID was not a number.", err)
		return
	}

	brew, err := this.data.Brews.Get(id)
	if err != nil {
		this.sendError(c, 404, "Unable to find that fermentable", err)
		return
	}

	var newReading common.Reading

	err = c.BindJSON(&newReading)
	if err != nil {
		this.sendError(c, 500, "Unable to parse input data.", err)
		return
	}

	brew.AddReading(newReading)
	this.sendSuccess(c, "Reading taken.")
}

/* Serves up our API to whomever is calling it.
 */
func (this *APIServer) Serve(webroot string) {
	// Setup our logger
	this.router.Use(ginrus.Ginrus(log.StandardLogger(), time.RFC3339, true))

	// QCRUD for Vessels
	this.router.GET("/vessel/", this.QueryVessel)
	this.router.GET("/vessel/:id", this.ReadVessel)
	this.router.PUT("/vessel/", this.UpdateVessel)
	this.router.POST("/vessel/", this.CreateVessel)
	this.router.DELETE("/vessel/", this.DeleteVessel)

	// QCRUD for Fermentables
	this.router.GET("/fermentable/", this.QueryFermentable)
	this.router.GET("/fermentable/:id", this.ReadFermentable)
	this.router.PUT("/fermentable/", this.UpdateFermentable)
	this.router.POST("/fermentable/", this.CreateFermentable)
	this.router.DELETE("/fermentable/", this.DeleteFermentable)

	// Reading Updates
	this.router.POST("/fermentable/:id/reading", this.AddFermentableReading)

	// Static hosting of our web servers
	this.router.Static("/", webroot)

	// Setup a listen string of host:port and then start our listener
	listenString := fmt.Sprintf("%s:%d", this.listenHost, this.listenPort)
	for {
		log.Error(this.router.Run(listenString))
	}
}
