package controllers

import (
	"net/http"

	"example.com/golang-restfulapi/models"
	"github.com/gin-gonic/gin"
)

// GET
func (db *DBController) GetCollection(c *gin.Context) {
	_type := c.Query("type")
	_where := map[string]interface{}{}

	if _type != "" {
		_where["type"] = _type
	}

	var collections []models.Collections
	db.Database.Where(_where).Find(&collections)

	for i, _ := range collections {
		db.Database.Model(collections[i]).Association("Groups").Find(&collections[i].Groups)
	}

	c.JSON(http.StatusOK, gin.H{"results": &collections})
}

// GET BY ID
func (db *DBController) GetCollectionById(c *gin.Context) {
	id := c.Param("id")
	var collections models.Collections

	db.Database.First(&collections, id)
	db.Database.Model(&collections).Association("Groups").Find(&collections.Groups)

	c.JSON(http.StatusOK, gin.H{"results": &collections})
}

// POST
func (db *DBController) CreateCollection(c *gin.Context) {
	var collection models.Collections
	err := c.ShouldBind(&collection)

	result := db.Database.Create(&collection)
	for i, _ := range collection.Groups {
		collection.Groups[i].CollectionId = collection.Id
		db.Database.Create(collection.Groups[i])
	}

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &collection})
	}
}

// PATCH
func (db *DBController) UpdateCollection(c *gin.Context) {

	var collection models.Collections
	err := c.ShouldBind(&collection)

	result := db.Database.Updates(collection)

	if result.Error != nil || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"meassage": "Bad request."})
	} else {
		c.JSON(http.StatusOK, gin.H{"results": &collection})
	}
}

// DELETE
func (db *DBController) DeleteCollection(c *gin.Context) {
	id := c.Param("id")
	var collections models.Collections
	db.Database.Delete(&collections, id)

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK})
}
