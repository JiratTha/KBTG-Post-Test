package allowance_calculation

import (
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/model"
	"testing"
)

// Mock DB implementation for testing purposes
type MockDB struct{}

func (m *MockDB) Get(dest interface{}, query string, args ...interface{}) error {
	// Mock implementation for database query
	// You can return different results based on the query and args if needed
	return nil
}

func TestAllowanceCalculation(t *testing.T) {
	// Mocking input data for testing
	db := &MockDB{}
	if db != nil {
		t.Errorf("db is nil")
	}
	mockAllowance := Personnel_model.Personnel{
		Allowance: []Personnel_model.Allowance{
			{AllowanceType: "donation", Amount: 100001.0},
			{AllowanceType: "k-receipt", Amount: 50001.0},
		},
	}

	// Creating a mock DB instance

	// Testing AllowanceCalculation function
	totalAllowance, donationAmount, kReceiptAmount := AllowanceCalculation(mockAllowance)

	// Define expected results based on your mock data
	expectedTotalAllowance := 150000.0
	expectedDonationAmount := 100001.0 // Limited by database max allowance for donation
	expectedKReceiptAmount := 50001.0  // Limited by database max allowance for k-receipt

	// Check if the actual results match the expected results
	if totalAllowance != expectedTotalAllowance {
		t.Errorf("Total allowance calculation incorrect. Got: %f, Expected: %f", totalAllowance, expectedTotalAllowance)
	}
	if donationAmount != expectedDonationAmount {
		t.Errorf("Donation amount calculation incorrect. Got: %f, Expected: %f", donationAmount, expectedDonationAmount)
	}
	if kReceiptAmount != expectedKReceiptAmount {
		t.Errorf("K-Receipt amount calculation incorrect. Got: %f, Expected: %f", kReceiptAmount, expectedKReceiptAmount)
	}
}
