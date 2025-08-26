package main

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ipResponse struct {
	IP       string `json:"ip"`
	Reversed string `json:"reversed"`
	Message  string `json:"message"`
}

func (app *application) ipReversedHandler(c *gin.Context) {

	clientIP := c.ClientIP()

	// res, err := http.Get("https://api64.ipify.org")
	// if err != nil {
	// 	log.Printf("ERROR: error fetching IP: %v", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// defer res.Body.Close()

	// clientIP, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	log.Printf("ERROR: error reading response body: %v", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	response := &ipResponse{
		IP:       string(clientIP),
		Reversed: reverseIP(string(clientIP)),
		Message:  "IP information retrieved successfully",
	}

	c.JSON(http.StatusOK, response)

	if err := app.Store.IP.SaveIP(c.Request.Context(), string(clientIP), reverseIP(string(clientIP))); err != nil {
		log.Printf("ERROR: error saving IP to database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
}

// reverseIP reverses IPv4 and IPv6 addresses
func reverseIP(ip string) string {

	var result string

	parsed := net.ParseIP(ip)
	if parsed == nil {
		return ip
	}

	if ipv4 := parsed.To4(); ipv4 != nil {
		parts := strings.Split(ipv4.String(), ".")

		for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
			parts[i], parts[j] = parts[j], parts[i]
		}

		result = strings.Join(parts, ".")
	}

	return result
}
