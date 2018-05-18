package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUsersTable_20180518_150752 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsersTable_20180518_150752{}
	m.Created = "20180518_150752"

	migration.Register("CreateUsersTable_20180518_150752", m)
}

// Run the migrations
func (m *CreateUsersTable_20180518_150752) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE IF NOT EXISTS account_go (
		id SERIAL PRIMARY KEY,
		nickname VARCHAR(255) NOT NULL,
		openid VARCHAR(255) NOT NULL,
		address VARCHAR(255) NOT NULL,
		create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)

}

// Reverse the migrations
func (m *CreateUsersTable_20180518_150752) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE  if exists account_go")

}
