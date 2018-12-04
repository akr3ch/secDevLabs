package db

import (
	"github.com/rafaveira3/secDevLabs/a5-ecommerce-api/pass"
	"github.com/rafaveira3/secDevLabs/a5-ecommerce-api/types"
	"gopkg.in/mgo.v2/bson"
)

// GetUserData queries MongoDB and returns user's data based on its username.
func GetUserData(mapParams map[string]interface{}) (types.UserData, error) {
	userDataResponse := types.UserData{}
	session, err := Connect()
	if err != nil {
		return userDataResponse, err
	}
	userDataQuery := []bson.M{}
	for k, v := range mapParams {
		userDataQuery = append(userDataQuery, bson.M{k: v})
	}
	userDataFinalQuery := bson.M{"$and": userDataQuery}
	err = session.SearchOne(userDataFinalQuery, nil, UserCollection, &userDataResponse)
	return userDataResponse, err
}

// RegisterUser regisiter into MongoDB a new user and returns an error.
func RegisterUser(userData types.UserData) error {
	session, err := Connect()
	if err != nil {
		return err
	}

	userData.HashedPassword, err = pass.BcrpytPassword(userData.RawPassword)
	if err != nil {
		return err
	}

	newUserData := bson.M{
		"username":       userData.Username,
		"hashedPassword": userData.HashedPassword,
		"userID":         userData.UserID,
		"ticket":         userData.Ticket,
	}
	err = session.Insert(newUserData, UserCollection)
	return err

}
