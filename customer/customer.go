package customer

import (
	"fmt"
	"net/http"

	"github.com/DemionK/finalexam/database"
	"github.com/gin-gonic/gin"
)

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func InitDatabase() {
	err := database.CreateTB()
	fmt.Println(err)
}

func PostHandler(c *gin.Context) {
	ct := Customer{}
	if err := c.ShouldBindJSON(&ct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": http.StatusText(http.StatusBadRequest)})
		return
	}
	row, err := database.InsertRow(ct.Name, ct.Email, ct.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": http.StatusText(http.StatusInternalServerError)})
		fmt.Println(err.Error())
		return
	}
	err = row.Scan(&ct.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusCreated, ct)
	return
}

func GetByIDHandler(c *gin.Context) {
	ID := c.Param("id")
	row, err := database.SelectByID(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	ct := Customer{}
	err = row.Scan(&ct.ID, &ct.Name, &ct.Email, &ct.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": http.StatusText(http.StatusBadRequest)})
		return
	}
	c.JSON(http.StatusOK, ct)
	return
}

func GetHandler(c *gin.Context) {
	rows, err := database.SelectAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	ctm := []Customer{}
	for rows.Next() {
		ct := Customer{}
		err := rows.Scan(&ct.ID, &ct.Name, &ct.Email, &ct.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": http.StatusText(http.StatusBadRequest)})
			return
		}
		ctm = append(ctm, ct)
	}
	c.JSON(200, ctm)
	return
}

func PutHandler(c *gin.Context) {
	ID := c.Param("id")
	ct := Customer{}
	if err := c.ShouldBindJSON(&ct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": http.StatusText(http.StatusBadRequest)})
		return
	}
	row, err := database.UpdateRow(ID, ct.Name, ct.Email, ct.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	err = row.Scan(&ct.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": http.StatusText(http.StatusBadRequest)})
		return
	}
	c.JSON(http.StatusOK, ct)
}

func DeleteByIDHandler(c *gin.Context) {
	ID := c.Param("id")
	err := database.DeleteRow(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
	return
}
