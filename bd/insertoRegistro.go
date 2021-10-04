package bd

import (
	"context"
	"time"

	"github.com/EstenoesJavi/go_tuiter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistro es la parada final con la db para registrar los datos del usuario*/
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("gotuiter")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
