package mysqlController

import (
	"github.com/gin-gonic/gin"
	"golang-api-demo/initMySQL"
	mysqlModel "golang-api-demo/modules/mySQLDemo/model"
	"log"
	"net/http"
)

func HandleGetData(c *gin.Context) {
	response, err := GetData()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error bad request": err})
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetData() ([]mysqlModel.MysqlModel, error) {
	DB, _ := initMySQL.InitMySQL()
	rows, err := DB.Query("SELECT * FROM user_question_answers where users_id = ? order by created_date desc", "d2d9438a-a8d1-4dac-9aa1-741346584529")
	defer DB.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var datas []mysqlModel.MysqlModel
	for rows.Next() {
		var data mysqlModel.MysqlModel
		if err := rows.Scan(&data.Id, &data.UsersId, &data.AnswerDate, &data.AnswerName, &data.AnswerSummaryName, &data.ServiceId, &data.Type47Result, &data.AvailFlg, &data.CreatedDate, &data.ModifiedDate); err != nil {
			log.Fatal(err)
			return datas, err
		}
		datas = append(datas, data)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return datas, err
	}

	return datas, nil
}
