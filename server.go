package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Item{})

	handler := newHandler(db)

	r := gin.Default()

	r.GET("/allUsers", handler.listUsersHandler)
	r.GET("/user/:id",/*TODO*/ handler.listUserByID)
	r.GET("/updateitems",/*TODO*/ handler.updateUserItems)
	r.GET("/login", handler.loginUsersHandler)
	r.POST("/register", handler.createUserHandler)
	r.DELETE("/del/:id", handler.deleteUserHandler)

	r.Run()
}

type Handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) *Handler {
	return &Handler{db}
}

type User struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Email     *string `json:"email"`
	Password  string  `json:"password"`
	Age       uint8   `json:"age"`
	Item      Item
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

/*func (h *Handler) listUserItems(tx *gin.Context) {
	var items []Item

	if result := h.db.Find(&items); result.Error != nil {
		tx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	tx.JSON(http.StatusOK, &items)
}

func (h *Handler) createUserItems(tx *gin.Context) {
	var item Item

	if err := tx.ShouldBindJSON(&item); err != nil {
		tx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&item); result.Error != nil {
		tx.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	tx.JSON(http.StatusCreated, gin.H{
		"Message": "Successful Added Item",
		"Status":  "201",
	})
	tx.JSON(http.StatusCreated, &item)
}

func (h *Handler) deleteUserItems(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&Item{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully Deleted User",
		"Status":  "200",
	})
}*/

func (h *Handler) listUsersHandler(c *gin.Context) {
	var users []User

	if result := h.db.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &users)
}

func (h *Handler) loginUsersHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.First(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message": "User Existence Confirmed",
		"Status":  "200",
	})
	c.JSON(http.StatusCreated, &user)
}

func (h *Handler) createUserHandler(c *gin.Context) {
	var user User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := h.db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message": "Successful Added User",
		"Status":  "201",
	})
	c.JSON(http.StatusCreated, &user)
}

func (h *Handler) deleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Successfully Deleted User",
		"Status":  "200",
	})
}

func (h *Handler) updateUserItems(c *gin.Context) {
	
}
