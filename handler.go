package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    // "strconv"
    "fmt"
)

func ScoreLASKupedes(c *gin.Context) {
    var las_t_scoring       Las_t_kupedes_scoring
    var las_t_scoring_clean Las_t_kupedes_scoring_clean
    var las_scored          Las_kupedes_scored
    if err := c.ShouldBindJSON(&las_t_scoring); err == nil {
        las_t_scoring_clean = las_t_scoring.ToClean()
        las_scored = las_t_scoring_clean.Score()
        // ProduceKafka(las_scored) // Produce data to Kafka topic
        fmt.Printf("Credit scoring result: %v\n",las_scored)
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": las_scored})
    } else {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }
}
