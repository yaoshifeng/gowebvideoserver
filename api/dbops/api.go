package dbops

import (
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil{
		return err
	}
	def stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd from ? where login_name = ?")
	if err != nil {
		return err
	}
	
	var pwd string
	_, err = stmtOut.Exec(loginName).Scan(&pwd)
	if err != nil {
		return err
	}
	
	def stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM user where login_name = ? and pwd = ?")
	if err != nil {
		return err
	}
	
	_, err = stmtOut.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	
	def stmtOut.Close()
	return err
}

func AddNewVideo(aid int, name string) (*def.VideoInfo, error) {
	
}

func GetVideoInfo(vid string) (*def.VideoInfo, error) {

}

func DeleteVideoInfo(aid string) (*def.VideoInfo, error) {

}

func AddNewComments(vid string, aid int, content string) error {

}

func DeleteListComments(vid string, from, to int) ([]*def.Comments, error) {

}
