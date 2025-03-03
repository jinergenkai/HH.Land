package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"land_service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Lấy thông tin ranh giới từ bên thứ ba
// @Description Gọi API của Guland để lấy dữ liệu ranh giới theo tọa độ
// @Tags Lands
// @Accept  json
// @Produce  json
// @Param lat query string true "Latitude"
// @Param lng query string true "Longitude"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /get-bound [get]
func GetBoundFromThirdParty(c *gin.Context) {
	// Lấy `lat` và `lng` từ query string
	lat := c.Query("lat")
	lng := c.Query("lng")

	// Kiểm tra tham số bắt buộc
	if lat == "" || lng == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Thiếu lat hoặc lng"})
		return
	}
	apiURL := fmt.Sprintf("https://guland.vn/get-bound-2?marker_lat=%s&marker_lng=%s", lat, lng)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Lỗi tạo request:", err)
		return
	}

	// Giả lập request từ trình duyệt thật
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
	req.Header.Set("Referer", "https://guland.vn/")
	req.Header.Set("Origin", "https://guland.vn")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Lỗi khi gửi request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Lỗi đọc body:", err)
		return
	}

	// fmt.Println("Dữ liệu nhận được:", string(body))

	var rawData models.LandParcelRaw
	if err := json.Unmarshal(body, &rawData); err != nil {
		fmt.Println("Error:", err)
		return
	}

	geoData := models.ConvertToGeoJSON(rawData)

	c.JSON(http.StatusOK, geoData)
}
