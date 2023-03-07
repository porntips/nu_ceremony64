package controller

import (
	"fmt"
	"log"
	"net/http"
	"nu_ceremony/connected"
	"nu_ceremony/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllGrad(c *gin.Context) {
	all_grad, err := GetAll()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Error": fmt.Sprintf("%s", err),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"all_result": all_grad.Ceremony,
		"all_count":  all_grad.Count,
	})
}

func GetAllCeremony(c *gin.Context, pack int) model.ReturnGrad {
	pack_all, err := GetPack(pack)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Error": fmt.Sprintf("%s", err),
		})
	}
	pack_receive, err := GetPackReceive(pack)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Error": fmt.Sprintf("%s", err),
		})
	}

	remain, err := GetRemain()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Error": fmt.Sprintf("%s", err),
		})
	}

	receive, err := GetReceive()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Error": fmt.Sprintf("%s", err),
		})
	}

	return model.ReturnGrad{
		Pack_count:   pack_all.Count,
		Pack_remain:  pack_all.Count - pack_receive.Count,
		Remain_result:  remain.Ceremony,
		Receive_result:   receive.Ceremony,
		// Receive_count:    receive.Count, //ยอดรับแล้วทั้งหมด
		Receive_count:    pack_receive.Count, //ยอดรับแล้วของแต่ละช่วง
		Ceremonypack: pack,
	}

}

func RunningCeremony(c *gin.Context) {
	studentcode := c.Param("studentcode")
	ceremony, _ := strconv.ParseBool(c.Param("ceremony"))

	query := `UPDATE ceremonyDB 
	SET ceremony = ? 
	WHERE studentcode = ?`
	stmt, err := connected.DB.Prepare(query)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"Error": fmt.Sprintf("Error %s when prepare", err),
		})
		return
	}
	res, err := stmt.Exec(ceremony, studentcode)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated": rows,
	})
}

func GetAll() (model.ReturnCeremony, error) {
	var (
		grad  model.Ceremony
		grads []model.Ceremony
	)

	stmt, err := connected.DB.Prepare("SELECT * FROM ceremonyDB ORDER BY ceremonygroup,ceremonysequence")
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	defer stmt.Close()

	theRows, err := stmt.Query()
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	for theRows.Next() {
		err := theRows.Scan(&grad.Studentcode, &grad.Sname, &grad.Degreecertificate, &grad.Facultyname, &grad.Hornor, &grad.Ceremonygroup,
			&grad.Ceremonysequence, &grad.Ceremonydate, &grad.Ceremonypack, &grad.Ceremonypackno, &grad.Ceremonysex, &grad.Ceremonyprefix, &grad.Ceremony)
		grads = append(grads, grad)

		if err != nil {
			return model.ReturnCeremony{}, err
		}
	}
	err = theRows.Err()
	if err != nil {
		return model.ReturnCeremony{}, err
	}

	return model.ReturnCeremony{
		Ceremony: grads,
		Count:    len(grads),
	}, nil
}

func GetPack(pack int) (model.ReturnCeremony, error) {
	var (
		grad  model.Ceremony
		grads []model.Ceremony
	)

	stmt, err := connected.DB.Prepare("SELECT * FROM ceremonyDB WHERE ceremonypack = ? ORDER BY ceremonygroup,ceremonysequence")
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	defer stmt.Close()

	theRows, err := stmt.Query(pack)
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	for theRows.Next() {
		err := theRows.Scan(&grad.Studentcode, &grad.Sname, &grad.Degreecertificate, &grad.Facultyname, &grad.Hornor, &grad.Ceremonygroup,
			&grad.Ceremonysequence, &grad.Ceremonydate, &grad.Ceremonypack, &grad.Ceremonypackno, &grad.Ceremonysex, &grad.Ceremonyprefix, &grad.Ceremony)
		grads = append(grads, grad)

		if err != nil {
			return model.ReturnCeremony{}, err
		}
	}
	err = theRows.Err()
	if err != nil {
		return model.ReturnCeremony{}, err
	}

	return model.ReturnCeremony{
		Ceremony: grads,
		Count:    len(grads),
	}, nil
}
func GetPackReceive(pack int) (model.ReturnCeremony, error) {
	var (
		grad  model.Ceremony
		grads []model.Ceremony
	)

	stmt, err := connected.DB.Prepare("SELECT * FROM ceremonyDB WHERE ceremonypack = ? AND ceremony = true ORDER BY ceremonygroup,ceremonysequence")
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	defer stmt.Close()

	theRows, err := stmt.Query(pack)
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	for theRows.Next() {
		err := theRows.Scan(&grad.Studentcode, &grad.Sname, &grad.Degreecertificate, &grad.Facultyname, &grad.Hornor, &grad.Ceremonygroup,
			&grad.Ceremonysequence, &grad.Ceremonydate, &grad.Ceremonypack, &grad.Ceremonypackno, &grad.Ceremonysex, &grad.Ceremonyprefix, &grad.Ceremony)
		grads = append(grads, grad)

		if err != nil {
			return model.ReturnCeremony{}, err
		}
	}
	err = theRows.Err()
	if err != nil {
		return model.ReturnCeremony{}, err
	}

	return model.ReturnCeremony{
		Ceremony: grads,
		Count:    len(grads),
	}, nil
}

func GetRemain() (model.ReturnCeremony, error) {
	var (
		grad  model.Ceremony
		grads []model.Ceremony
	)

	stmt, err := connected.DB.Prepare("SELECT * FROM ceremonyDB WHERE ceremony = false ORDER BY ceremonygroup,ceremonysequence LIMIT 20")
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	defer stmt.Close()

	theRows, err := stmt.Query()
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	for theRows.Next() {
		err := theRows.Scan(&grad.Studentcode, &grad.Sname, &grad.Degreecertificate, &grad.Facultyname, &grad.Hornor, &grad.Ceremonygroup,
			&grad.Ceremonysequence, &grad.Ceremonydate, &grad.Ceremonypack, &grad.Ceremonypackno, &grad.Ceremonysex, &grad.Ceremonyprefix, &grad.Ceremony)
		grads = append(grads, grad)

		if err != nil {
			return model.ReturnCeremony{}, err
		}
	}
	err = theRows.Err()
	if err != nil {
		return model.ReturnCeremony{}, err
	}

	return model.ReturnCeremony{
		Ceremony: grads,
		Count:    len(grads),
	}, nil
}

func GetReceive() (model.ReturnCeremony, error) {
	var (
		grad  model.Ceremony
		grads []model.Ceremony
	)

	stmt, err := connected.DB.Prepare("SELECT * FROM ceremonyDB WHERE ceremony = true ORDER BY ceremonysequence DESC")
	if err != nil {
		return model.ReturnCeremony{}, err
	}
	defer stmt.Close()

	theRows, err := stmt.Query()
	if err != nil {
		return model.ReturnCeremony{}, err
	}

	for theRows.Next() {
		err := theRows.Scan(&grad.Studentcode, &grad.Sname, &grad.Degreecertificate, &grad.Facultyname, &grad.Hornor, &grad.Ceremonygroup,
			&grad.Ceremonysequence, &grad.Ceremonydate, &grad.Ceremonypack, &grad.Ceremonypackno, &grad.Ceremonysex, &grad.Ceremonyprefix, &grad.Ceremony)
		grads = append(grads, grad)

		if err != nil {
			return model.ReturnCeremony{}, err
		}
	}
	err = theRows.Err()
	if err != nil {
		return model.ReturnCeremony{}, err
	}

	return model.ReturnCeremony{
		Ceremony: grads,
		Count:    len(grads),
	}, nil
}
