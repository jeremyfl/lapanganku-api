package field

import (
	"net/http"

	"github.com/jeremylombogia/lapanganku-api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func IndexField(c echo.Context) error {
	var fields, err = models.FetchField()

	if err != nil {
		return c.JSON(500, Presenter{500, "Something went wrong and it's not your fault!", nil})
	}

	return c.JSON(200, Presenter{200, nil, fields})
}

func PostField(c echo.Context) (err error) {
	payload := new(Payload)
	if err = c.Bind(payload); err != nil {
		return
	}

	var field models.Field
	field.ID = bson.NewObjectId()
	field.Name = payload.Data.Name
	field.Price = payload.Data.Price
	field.Availability = payload.Data.Availability

	if err = models.StoreField(&field); err != nil {
		return c.JSON(http.StatusBadRequest, Presenter{500, "Something went wrong", err.Error()})
	}

	return c.JSON(201, Presenter{201, "Record created", field})
}
