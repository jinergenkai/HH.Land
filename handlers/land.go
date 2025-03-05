package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"land_service/database"
	"land_service/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateLand godoc
// @Summary Thêm vùng đất mới
// @Description Tạo một vùng đất mới với tọa độ GeoJSON
// @Tags Lands
// @Accept json
// @Produce json
// @Param land body models.Land true "Dữ liệu vùng đất"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /land [post]
func CreateLand(c *gin.Context) {
	var newLand models.Land

	// Bind JSON vào struct
	if err := c.ShouldBindJSON(&newLand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lấy collection trong từng request (tránh nil pointer)
	landCollection := database.GetCollection("lands")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := landCollection.InsertOne(ctx, newLand)
	if err != nil {
		log.Println("❌ Lỗi khi thêm vùng đất:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể thêm vùng đất"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Thêm vùng đất thành công!", "id": result.InsertedID})
}

// @Summary Lấy danh sách vùng đất
// @Description Trả về danh sách tất cả các vùng đất
// @Tags Lands
// @Produce json
// @Success 200 {array} models.Land
// @Router /land [get]
func GetLands(c *gin.Context) {
	var lands []models.Land

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	landCollection := database.GetCollection("lands")

	cursor, err := landCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách vùng đất"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var land models.Land
		if err := cursor.Decode(&land); err != nil {
			log.Println("❌ Lỗi khi decode vùng đất:", err)
			continue
		}
		lands = append(lands, land)
	}

	c.JSON(http.StatusOK, lands)
}

// @Summary Xóa vùng đất theo ID
// @Description Xóa vùng đất dựa trên ID được cung cấp
// @Tags Lands
// @Param id path string true "ID của vùng đất cần xóa"
// @Router /land/{id} [delete]
func DeleteLand(c *gin.Context) {
	landID := c.Param("id")

	// Kiểm tra ID có hợp lệ không (MongoDB ObjectID)
	objID, err := primitive.ObjectIDFromHex(landID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	landCollection := database.GetCollection("lands")

	// Xóa vùng đất theo ID
	result, err := landCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi khi xóa vùng đất"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy vùng đất"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xóa thành công"})
}
