package testing

import (
	"github.com/cloudfoundry-incubator/notifications/application"
	"github.com/cloudfoundry-incubator/notifications/db"
	"github.com/cloudfoundry-incubator/notifications/models"
)

func TruncateTables(database *db.DB) {
	env := application.NewEnvironment()
	dbMigrator := models.DatabaseMigrator{}
	dbMigrator.Migrate(database.RawConnection(), env.ModelMigrationsPath)
	models.Setup(database)

	connection := database.Connection().(*db.Connection)
	err := connection.TruncateTables()
	if err != nil {
		panic(err)
	}
}
