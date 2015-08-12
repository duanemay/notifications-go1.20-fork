package db_test

import (
	"database/sql"
	"testing"

	"github.com/cloudfoundry-incubator/notifications/application"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDBSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Suite")
}

var sqlDB *sql.DB

var _ = BeforeEach(func() {
	env := application.NewEnvironment()

	var err error
	sqlDB, err = sql.Open("mysql", env.DatabaseURL)
	Expect(err).NotTo(HaveOccurred())
})
