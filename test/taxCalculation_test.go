package test

import (
	"encoding/json"
	Personnel_model "github.com/JiratTha/assessment-tax/Personnel/model"
	"github.com/JiratTha/assessment-tax/tax/model"
	tax_calculation "github.com/JiratTha/assessment-tax/tax/tax-calculation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxCalculation(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "Scenario 1: Total income less than or equal to 150000",
			input: `{
				"totalIncome": 150000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 0.0
					},
					{
						"allowanceType": "donation",
						"amount": 0.0
					}
				]
			}`,
			expected: `{
				"tax": 0.0,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 0.0},
					{"level": "500,001-1,000,000", "tax": 0.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		{
			name: "Scenario 2: Total income less than or equal to 500000",
			input: `{
				"totalIncome": 500000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 0.0
					},
					{
						"allowanceType": "donation",
						"amount": 0.0
					}
				]
			}`,
			expected: `{
				"tax": 15000.0,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 15000.0},
					{"level": "500,001-1,000,000", "tax": 0.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		{
			name: "Scenario 3: Total income between 500001 and 1000000",
			input: `{
				"totalIncome": 800000.0,
				"wht": 0.0,
				"allowances": [
					{
						"allowanceType": "k-receipt",
						"amount": 50000.0
					},
					{
						"allowanceType": "donation",
						"amount": 20000.0
					}
				]
			}`,
			expected: `{
				"tax": 85000.0,
				"taxLevel": [
					{"level": "0-150,000", "tax": 0.0},
					{"level": "150,001-500,000", "tax": 15000.0},
					{"level": "500,001-1,000,000", "tax": 70000.0},
					{"level": "1,000,001-2,000,000", "tax": 0.0},
					{"level": "2,000,001 ขึ้นไป", "tax": 0.0}
				]
			}`,
		},
		// Add more test cases for other scenarios as needed
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var requestBody Personnel_model.Personnel
			if err := json.Unmarshal([]byte(tc.input), &requestBody); err != nil {
				t.Fatalf("failed to unmarshal request body: %v", err)
			}

			expectedResponse := model.TaxResponse{}
			if err := json.Unmarshal([]byte(tc.expected), &expectedResponse); err != nil {
				t.Fatalf("failed to unmarshal expected response: %v", err)
			}

			actualResponse := tax_calculation.TaxCalculation(requestBody.TotalIncome, requestBody.Wht)

			assert.Equal(t, expectedResponse, actualResponse, "Response does not match expected")
		})
	}
}
