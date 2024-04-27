package setting

//
//import (
//	"github.com/labstack/echo/v4"
//	"net/http"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//
//	"github.com/JiratTha/assessment-tax/admin/personal-struct-model"
//	"github.com/JiratTha/assessment-tax/admin/setting-function"
//)
//
//// MockDB is a mock implementation of the database interface for testing.
//type MockDB struct {
//	mock.Mock
//}
//
//func (m *MockDB) Exec(query string, args ...interface{}) (interface{}, error) {
//	argsList := m.Called(query, args)
//	return argsList.Get(0), argsList.Error(1)
//}
//
//func TestSettingKReceipt(t *testing.T) {
//	// Initialize mock DB and set it in the setting package.
//	mockDB := new(MockDB)
//	setting.SetDB(mockDB)
//
//	// Define test cases.
//	tests := []struct {
//		name           string
//		input          personal-struct-model.Admin
//		expectedOutput personal-struct-model.AdminResponse
//		mockExecResult interface{}
//		mockExecError  error
//	}{
//		{
//			name:           "Amount greater than limit",
//			input:          personal-struct-model.Admin{Amount: 100001.0},
//			expectedOutput: personal-struct-model.AdminResponse{PersonalDeduction: 100000.0},
//			mockExecResult: nil,
//			mockExecError:  nil,
//		},
//		{
//			name:           "Amount within limit",
//			input:          personal-struct-model.Admin{Amount: 80000.0},
//			expectedOutput: personal-struct-model.AdminResponse{PersonalDeduction: 80000.0},
//			mockExecResult: nil,
//			mockExecError:  nil,
//		},
//		{
//			name:           "Amount less than or equal 0",
//			input:          personal-struct-model.Admin{Amount: 0.0},
//			expectedOutput: echo.NewHTTPError(http.StatusBadRequest, "K-receipt must be greater than 0",
//			mockExecResult: nil,
//			mockExecError:  nil,
//		},
//		// Add more test cases as needed.
//	}
//
//	// Iterate through test cases.
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Mock DB.Exec method.
//			mockDB.On("Exec", mock.Anything, mock.Anything).Return(tt.mockExecResult, tt.mockExecError)
//
//			// Call the function under test.
//			output := setting.SettingPersonalDeduction(tt.input)
//
//			// Verify the result.
//			assert.Equal(t, tt.expectedOutput, output)
//
//			// Assert that the expected DB.Exec method was called.
//			mockDB.AssertCalled(t, "Exec", mock.Anything, mock.Anything)
//		})
//	}
//}
