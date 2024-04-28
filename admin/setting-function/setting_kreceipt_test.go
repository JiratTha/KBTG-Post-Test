package setting

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/JiratTha/assessment-tax/admin/model"
	"github.com/JiratTha/assessment-tax/db"
)

func TestSettingKReceipt(t *testing.T) {

	dbMock, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	sqlxDB := sqlx.NewDb(dbMock, "sqlmock")

	db.SetDB(sqlxDB)

	mock.ExpectExec(`UPDATE project1."allowance" SET amount=\$1 WHERE allowance_type='k-receipt'`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Mock successful update

	t.Run("AmountGreaterThan100000", func(t *testing.T) {
		admin := model.Admin{Amount: 150000}
		expected := model.AdminResponse{KReceipt: 100000}

		result := SettingKReceipt(admin)
		assert.Equal(t, expected, result, "Expected and actual results should match")
	})

	t.Run("AmountLessThan100000", func(t *testing.T) {
		admin := model.Admin{Amount: 80000}
		expected := model.AdminResponse{KReceipt: 80000}

		result := SettingKReceipt(admin)
		assert.Equal(t, expected, result, "Expected and actual results should match")
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
