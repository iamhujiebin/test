package utils

import (
	"../config"
	"gopkg.in/mgo.v2"
)


func WithinFEWebDB(s func(*mgo.Database) error) error {
	return WithinDatabase(config.GetGlobalStringValue("feweb_db_url", ""),
		config.GetGlobalStringValue("feweb_db", ""), s)
}

func WithinInteractionDB(s func(*mgo.Database) error) error {
	return WithinDatabase(config.GetGlobalStringValue("interaction_db_url", ""),
		config.GetGlobalStringValue("interaction_db", ""), s)
}