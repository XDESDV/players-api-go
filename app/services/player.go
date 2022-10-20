package services

import (
	"context"
	"errors"

	"github.com/xdesdv/cac-sdk/connectors/mongodb"
	"github.com/xdesdv/cac-sdk/functions"
	"github.com/xdesdv/cac-sdk/queries"
	"github.com/xdesdv/players-api-go/app/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(p types.Player) error {
	var err error

	p.TagID = functions.NewUUID()

	db := mongodb.GetInstance()
	cCollection := db.Collection(p.Collection())
	if err == nil {
		_, err = cCollection.InsertOne(context.TODO(), p)
	}
	return err

}

func Read(id string) (types.Player, error) {
	var (
		err         error
		p           types.Player
		queryParams queries.QueryParams
	)

	db := mongodb.GetInstance()
	c := db.Collection(p.Collection())

	queryParams.FilterClause = append(queryParams.FilterClause, "customID,"+id)
	filter := mongodb.SelectConstructeur(queryParams)
	err = c.FindOne(context.TODO(), filter).Decode(&p)
	if err == nil {
		if err == mongo.ErrNoDocuments {
			return p, err
		}

	}

	return p, err
}

func Find(queryParams queries.QueryParams) (types.Players, error) {
	var (
		err     error
		players types.Players
		p       types.Player
		coll    *mongo.Collection
		cursor  *mongo.Cursor
	)
	db := mongodb.GetInstance()
	coll = db.Collection(p.Collection())

	filter := mongodb.SelectConstructeur(queryParams)
	cursor, err = coll.Find(context.TODO(), filter)
	defer cursor.Close(context.TODO())
	if err == nil {
		for cursor.Next(context.TODO()) {
			// A new result variable should be declared for each document.
			err = cursor.Decode(&p)
			if err == nil {
				players = append(players, p)
			}
		}
	}
	err = cursor.Err()

	return players, err
}

func update(id string, p *types.Player) error {
	var (
		doc         interface{}
		result      *mongo.UpdateResult
		err         error
		queryParams queries.QueryParams
	)

	db := mongodb.GetInstance()
	c := db.Collection(p.Collection())

	queryParams.FilterClause = append(queryParams.FilterClause, "tagID,"+id)
	filter := mongodb.SelectConstructeur(queryParams)
	if doc, err = mongodb.ToDoc(p); err == nil {
		update := bson.M{"$set": doc}
		result, err = c.UpdateOne(context.TODO(), filter, update)
		if result.MatchedCount == 0 {
			err = errors.New("Tile to be modified was not found")
		}
		if err == nil && result.ModifiedCount == 0 {
			err = errors.New("Tile could not be updated")
		}

	}
	return err

}

// func Delete() {

// }
