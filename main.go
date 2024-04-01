package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"rest_postgres/albums"
	"rest_postgres/mutualfund"
	"rest_postgres/mutualfunddata"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s timezone=utc",
		"postgres",
		"root",
		"0.0.0.0",
		6543,
		"albums",
		"disable",
	))
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	albums.Create(router, db)
	mutualfunddata.Create(router, db)

	router.GET("/mf_latest/:key", GetMutualFundLatest)
	router.GET("/mf_history/:key", GetHistoryNav)

	router.GET("/insert_mf_latest/:key", UpsertMutualFundLatest)

	router.Run("localhost:8080")
}

func GetHistoryNav(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	symbol := c.Param("key")
	navData, err := mutualfund.NewHandler().GetHistoryNavData(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, navData)
}

func GetMutualFundLatest(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	symbol := c.Param("key")
	navData, err := mutualfund.NewHandler().GetLatestNavData(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, navData)
}

func UpsertMutualFundLatest(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	symbol := c.Param("key")
	navData, err := mutualfund.NewHandler().GetLatestNavData(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mf_date, _ := time.Parse("02-01-2006", navData.Data[0].Date)
	mf_dateStr := fmt.Sprintln(mf_date.Format("2006-01-02"))

	mutualfunddata.SetDB(db)

	var mf_data = mutualfunddata.MutualFundData{
		FundHouse:      navData.Meta.FuncdHouse,
		SchemeType:     navData.Meta.SchemeType,
		SchemeCategory: navData.Meta.SchemaCategory,
		SchemeCode:     strconv.Itoa(navData.Meta.SchemeCode),
		SchemeName:     navData.Meta.SchemeName,
		Date:           mf_dateStr,
		Nav:            navData.Data[0].Nav,
	}

	res, err := mutualfunddata.UpsertMutualFund(mf_data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
