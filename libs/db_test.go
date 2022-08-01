package libs

import (
	"goapi/models"
	"testing"

	"github.com/walkersumida/jaker"
)

func TestConnectDb(t *testing.T) {
	Connect()
}

func TestFakerData(t *testing.T) {
	for i := 0; i < 1; i++ {
		profile := jaker.Profile
		user := models.User{}
		user.UserName = profile.JaKanjiFullName
		db, _ := Connect()
		_ = db.Debug().Create(&user)
		favorite := models.Favorite{}
		favorite.UserId = user.ID
		favorite.UserId = user.ID
		_ = db.Debug().Create(&favorite)
	}
}
