package database

import (
	"database/sql"
	"net"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenConnection() {
	datab, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(err)
	}
	datab.Exec(`
		CREATE TABLE IF NOT EXISTS addressbook (
			ip string PRIMARY KEY,
			port integer
		)
	`)
	db = datab
}

func CloseConn() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

func QueryAddressbook() []*net.TCPAddr {
	rows, err := db.Query("SELECT * from addressbook")
	if err != nil {
		return nil
	}
	var ipAddr string
	var port int

	addrs := []*net.TCPAddr{}
	for rows.Next() {
		err = rows.Scan(&ipAddr, &port)
		if err != nil {
			continue
		}
		ip := net.ParseIP(ipAddr)
		addr := &net.TCPAddr{
			IP:   ip,
			Port: port,
		}
		addrs = append(addrs, addr)
	}
	rows.Close()
	return addrs
}

func AddContact(addr *net.TCPAddr) error {
	template, err := db.Prepare("INSERT INTO addressbook(ip, port) values(?, ?)")
	if err != nil {
		return err
	}

	_, err = template.Exec(addr.IP.String(), addr.Port)
	if err != nil {
		return err
	}
	return nil
}
