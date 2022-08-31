package mysqlModel

import (
	"database/sql"
	"time"
)

type MysqlModel struct {
	Id                int            `json:"id"`
	UsersId           sql.NullString `json:"users_id"`
	AnswerDate        time.Time      `json:"answer_date"`
	AnswerName        sql.NullString `json:"answer_name"`
	AnswerSummaryName sql.NullString `json:"answer_summary_name"`
	ServiceId         int            `json:"service_id"`
	Type47Result      sql.NullInt16  `json:"type_47_result"`
	AvailFlg          string         `json:"avail_flg"`
	CreatedDate       time.Time      `json:"created_date"`
	ModifiedDate      time.Time      `json:"modified_date"`
}
