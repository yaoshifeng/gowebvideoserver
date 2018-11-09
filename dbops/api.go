package dbops

import (
	"database/sql"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUE(?, ?)")
	if err != nil {
		return err
	}

	_, err := stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) error {
	stmtOut, err := connDB.Prepare("SELECT pwd from users WHERE login_name = ?")
	if err != nil {
		return err
	}

	var pwd string
	err := stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil {
		return err
	}
	defer stmtOut.Close()
	return nil
}

func AddVideoCredential(video_name string) error {

}
