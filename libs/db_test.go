package libs

import (
	"github.com/walkersumida/jaker"
	"goapi/models"
	"testing"
)

func TestConnectDb(t *testing.T) {
	Connect()
}

func TestFakerData(t *testing.T) {

	for i := 0; i < 10; i++ {
		profile := jaker.Profile
		user := models.User{}
		user.UserName = profile.JaKanjiFullName
		db, _ := Connect()
		_ = db.Create(&user)
		favorite := models.Favorite{}
		favorite.UserId = user.ID
		favorite.UserId = user.ID
		_ = db.Create(&favorite)
	}
}
