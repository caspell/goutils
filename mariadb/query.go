package mariadb

import (
	"fmt"
	"strconv"

	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/wunicorns/goutils/config"
)

var DB sqlx.DB

func Init() error {

	fmt.Println("db init")

	id := config.DB_ID
	pw := config.DB_PW
	ip := config.DB_IP
	port, _ := strconv.Atoi(config.DB_PORT)
	dbName := config.DB_NAME

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", id, pw, ip, port, dbName)
	db, err := sqlx.Connect("mysql", connStr)
	if err != nil {
		return err
	}

	DB = *db
	return nil
}

func SelectServerList(ip string) (*[]Server, error) {
	servers := []Server{}
	q := `
    SELECT IFNULL(ts.status, '') as status,
	  IFNULL(vss1.severity_count, 0) as critical_count,
	  IFNULL(vss2.severity_count, 0) as warning_count,
	  (case when vss1.severity_count then 2 when vss2.severity_count then 1 else 0 end) as alarm_status
      FROM tb_server ts 
	  LEFT JOIN (SELECT tsg.group_id, tsg.group_name, tsgr.server_no FROM tb_server_group_rel tsgr LEFT JOIN tb_server_group tsg  ON tsgr.group_id = tsg.group_id ) mb ON ts.server_no = mb.server_no
	  LEFT JOIN (
		select ta.device_no AS device_no, ta.severity AS severity, count(0) AS severity_count
		from tb_alarms ta 
		where ta.category = 'E001' and ta.clear_flag = 'N' and ta.acknown_flag = 'N' group by ta.device_no, ta.severity
	  ) vss1 on (ts.server_no = vss1.device_no and vss1.severity = 'E101')
	  LEFT JOIN (
		select ta.device_no AS device_no, ta.severity AS severity, count(0) AS severity_count
		from tb_alarms ta 
		where ta.category = 'E001' and ta.clear_flag = 'N' and ta.acknown_flag = 'N' group by ta.device_no, ta.severity
	  ) vss2 on (ts.server_no = vss2.device_no and vss2.severity = 'E102')
		WHERE ts.del_yn = 'N'
    	%s
	`
	args := []interface{}{}

	if ip == "" {
		q = fmt.Sprintf(q, "")
	} else {
		q = fmt.Sprintf(q, "AND bmc_ip_addr = ?")
		args = append(args, ip)
	}

	stmt, _ := DB.Preparex(q)
	err := stmt.Select(&servers, args...)
	if err != nil {
		fmt.Println("Error occured while SelectServerList:", err)
		return &servers, err
	} else if len(servers) == 0 {
		return &servers, err
	}
	defer stmt.Close()
	return &servers, nil
}
