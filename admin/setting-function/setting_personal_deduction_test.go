package setting

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/JiratTha/assessment-tax/admin/model"
	"github.com/JiratTha/assessment-tax/admin/setting-function"
)

// MockDB is a mock implementation of the database interface for testing.
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Exec(query string, args ...interface{}) (interface{}, error) {
	argsList := m.Called(query, args)
	return argsList.Get(0), argsList.Error(1)
}

func TestSettingPersonnelDeduction(t *testing.T) {
	// Initialize mock DB and set it in the setting package.
	mockDB := new(MockDB)
	setting.SetDB(mockDB)

	// Define test cases.
	tests := []struct {
		name           string
		input          model.Admin
		expectedOutput model.AdminResponse
		mockExecResult interface{}
		mockExecError  error
	}{
		{
			name:           "Amount greater than limit",
			input:          model.Admin{Amount: 120000.0},
			expectedOutput: model.AdminResponse{PersonalDeduction: 100000.0},
			mockExecResult: nil,
			mockExecError:  nil,
		},
		{
			name:           "Amount within limit",
			input:          model.Admin{Amount: 80000.0},
			expectedOutput: model.AdminResponse{PersonalDeduction: 80000.0},
			mockExecResult: nil,
			mockExecError:  nil,
		},
		{
			name:           "Amount less than or equal 10000",
			input:          model.Admin{Amount: 10000.0},
			expectedOutput: echo.NewHTTPError(http.StatusBadRequest, "Personal Deduction must be greater than 10000",
			mockExecResult: nil,
			mockExecError:  nil,
		},
		// Add more test cases as needed.
	}

	// Iterate through test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock DB.Exec method.
			mockDB.On("Exec", mock.Anything, mock.Anything).Return(tt.mockExecResult, tt.mockExecError)

			// Call the function under test.
			output := setting.SettingPersonnelDeduction(tt.input)

			// Verify the result.
			assert.Equal(t, tt.expectedOutput, output)

			// Assert that the expected DB.Exec method was called.
			mockDB.AssertCalled(t, "Exec", mock.Anything, mock.Anything)
		})
	}
}
