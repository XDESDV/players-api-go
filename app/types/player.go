package types

import "errors"

type Player struct {
	TagID    string `bson:"tagID" json:"tagID"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
}

// Collection Mongodb collection
func (p Player) Collection() string {
	return "players"
}

// Controls : control fields requires
func (p Player) Controls() error {
	var lstErr string
	rc := "\n"
	lstErr = ""

	if p.Username == "" {
		lstErr += "Username error : must not be empty" + rc
	}

	if lstErr == "" {
		return nil
	}

	return errors.New(lstErr)

}

// Players list of player
type Players []Player
