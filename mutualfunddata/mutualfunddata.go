package mutualfunddata

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type MutualFundData struct {
	ID             string `json:"id"`
	FundHouse      string `json:"fund_house"`
	SchemeType     string `json:"scheme_type"`
	SchemeCategory string `json:"scheme_category"`
	SchemeCode     string `json:"scheme_code"`
	SchemeName     string `json:"scheme_name"`
	Date           string `json:"date"`
	Nav            string `json:"nav"`
}

func Create(router *gin.Engine, database *sql.DB) {
	db = database
	router.GET("/mf_data", GetMutualFundData)
	router.POST("/mf_data", CreateMutualFundData)
}

func SetDB(database *sql.DB) {
	db = database
}

func GetMutualFundData(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, fund_house, scheme_type, scheme_category, scheme_code, scheme_name, date, nav FROM mutual_fund_data")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var mutualFundData []MutualFundData
	for rows.Next() {
		var m MutualFundData
		err := rows.Scan(&m.ID, &m.FundHouse, &m.SchemeType, &m.SchemeCategory, &m.SchemeCode, &m.SchemeName, &m.Date, &m.Nav)
		if err != nil {
			log.Fatal(err)
		}
		mutualFundData = append(mutualFundData, m)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, mutualFundData)
}

func CreateMutualFundData(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var m MutualFundData
	if err := c.BindJSON(&m); err != nil {
		log.Fatal(err)
	}

	_, err := db.Exec("INSERT INTO mutual_fund_data (fund_house, scheme_type, scheme_category, scheme_code, scheme_name, date, nav) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		m.FundHouse, m.SchemeType, m.SchemeCategory, m.SchemeCode, m.SchemeName, m.Date, m.Nav)
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, m)
}

func CheckData(date string, schemeCode string) (string, error) {
	rows, err := db.Query("SELECT id FROM mutual_fund_data WHERE date = $1 AND scheme_code = $2", date, schemeCode)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var id string
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return "", err
		}
	}
	err = rows.Err()
	if err != nil {
		return "", err
	}

	log.Default().Println("Record present ID: ", id)
	return id, nil
}

func CreateMutualFund(m MutualFundData) (sql.Result, error) {
	res, err := db.Exec("INSERT INTO mutual_fund_data (fund_house, scheme_type, scheme_category, scheme_code, scheme_name, date, nav) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		m.FundHouse, m.SchemeType, m.SchemeCategory, m.SchemeCode, m.SchemeName, m.Date, m.Nav)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateMutualFund(m MutualFundData) (sql.Result, error) {
	res, err := db.Exec("UPDATE mutual_fund_data SET fund_house = $1, scheme_type = $2, scheme_category = $3, scheme_name = $4, date = $5, nav = $6, scheme_code = $7 WHERE id = $8",
		m.FundHouse, m.SchemeType, m.SchemeCategory, m.SchemeName, m.Date, m.Nav, m.SchemeCode, m.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpsertMutualFund(m MutualFundData) (sql.Result, error) {
	mf_id, err := CheckData(m.Date, m.SchemeCode)
	if err != nil {
		return nil, err
	}

	var res sql.Result
	if mf_id != "" {
		m.ID = mf_id
		res, err = UpdateMutualFund(m)
	} else {
		res, err = CreateMutualFund(m)
	}
	if err != nil {
		return nil, err
	}
	return res, nil
}
