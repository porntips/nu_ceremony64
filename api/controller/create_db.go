package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"nu_ceremony/connected"
	"nu_ceremony/model"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AddDefaultDb(c *gin.Context) {
	var ceremonyies []model.Ceremony
	if err := c.ShouldBindJSON(&ceremonyies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DropData()
	query := "INSERT INTO ceremonyDB(studentcode,sname,degreecertificate,facultyname,hornor,ceremonygroup,ceremonysequence,subsequence,ceremonydate,ceremonypack,ceremonypackno,ceremonysex,ceremonyprefix) VALUES "
	var inserts []string
	var params []interface{}
	for _, v := range ceremonyies {
		inserts = append(inserts, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		params = append(params, v.Studentcode, v.Sname, v.Degreecertificate, v.Facultyname, v.Hornor, v.Ceremonygroup, v.Ceremonysequence, v.Subsequence, v.Ceremonydate, v.Ceremonypack, v.Ceremonypackno, v.Ceremonysex, v.Ceremonyprefix)
	}

	queryVals := strings.Join(inserts, ",")
	query = query + queryVals

	ctx, cancelfunc := context.WithTimeout(context.Background(), 15*time.Hour)
	defer cancelfunc()
	stmt, err := connected.DB.PrepareContext(ctx, query)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"Error": fmt.Sprintf("Error %s when preparing SQL statement", err),
		})
		return
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Error": fmt.Sprintf("Error %s when inserting row into table", err),
		})
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return
	}
	log.Printf("%d created simulatneously", rows)
	c.JSON(http.StatusOK, gin.H{
		"created": rows,
	})
}

func DropData() {
	query := `DELETE FROM ceremonyDB`
	stmt, err := connected.DB.Prepare(query)
	if err != nil {
		log.Printf("Error %s when prepare", err)
		return
	}
	res, err := stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return
	}
	log.Println("rows affected is", rows)
}
