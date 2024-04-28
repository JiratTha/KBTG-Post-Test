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

func TestSettingPersonalDeduction(t *testing.T) {

	dbMock, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbMock.Close()

	sqlxDB := sqlx.NewDb(dbMock, "sqlmock")

	db.SetDB(sqlxDB)

	mock.ExpectExec(`UPDATE project1."personal_deduction" SET amount=\$1 WHERE personal_deduction='personalDeduction'`).
		WithArgs(100000.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	t.Run("AmountGreaterThan100000", func(t *testing.T) {
		personalDeduction := model.Admin{Amount: 150000.0}
		expected := model.AdminResponse{PersonalDeduction: 100000.0}

		result := SettingPersonalDeduction(personalDeduction)
		assert.Equal(t, expected, result, "Expected and actual results should match")
	})

	t.Run("AmountLessThan100000", func(t *testing.T) {
		personalDeduction := model.Admin{Amount: 80000.0}
		expected := model.AdminResponse{PersonalDeduction: 80000.0}

		result := SettingPersonalDeduction(personalDeduction)
		assert.Equal(t, expected, result, "Expected and actual results should match")
	})

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
