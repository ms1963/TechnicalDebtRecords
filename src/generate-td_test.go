// generate-td_test.go
package main

import (
	"os"
	"strings"
	"testing"
)

// TestValidateTechnicalDebt checks the validation of the TechnicalDebt structure
func TestValidateTechnicalDebt(t *testing.T) {
	tests := []struct {
		name    string
		td      TechnicalDebt
		wantErr bool
	}{
		{
			name: "Valid TechnicalDebt",
			td: TechnicalDebt{
				Title:   "Outdated Library",
				Author:  "Jane Doe",
				Version: "1.0.0",
				Date:    "2024-04-15",
				State:   "Analyzed",
			},
			wantErr: false,
		},
		{
			name: "Missing Title",
			td: TechnicalDebt{
				Author:  "Jane Doe",
				Version: "1.0.0",
				Date:    "2024-04-15",
				State:   "Analyzed",
			},
			wantErr: true,
		},
		{
			name: "Missing Author",
			td: TechnicalDebt{
				Title:   "Outdated Library",
				Version: "1.0.0",
				Date:    "2024-04-15",
				State:   "Analyzed",
			},
			wantErr: true,
		},
		{
			name: "Missing Version",
			td: TechnicalDebt{
				Title:  "Outdated Library",
				Author: "Jane Doe",
				Date:   "2024-04-15",
				State:  "Analyzed",
			},
			wantErr: true,
		},
		{
			name: "Missing Date",
			td: TechnicalDebt{
				Title:   "Outdated Library",
				Author:  "Jane Doe",
				Version: "1.0.0",
				State:   "Analyzed",
			},
			wantErr: true,
		},
		{
			name: "Missing State",
			td: TechnicalDebt{
				Title:   "Outdated Library",
				Author:  "Jane Doe",
				Version: "1.0.0",
				Date:    "2024-04-15",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateTechnicalDebt(tt.td); (err != nil) != tt.wantErr {
				t.Errorf("validateTechnicalDebt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestGenerateMarkdown checks the generation of Markdown content
func TestGenerateMarkdown(t *testing.T) {
	td := TechnicalDebt{
		Title:          "Outdated Library",
		Author:         "Jane Doe",
		Version:        "1.0.0",
		Date:           "2024-04-15",
		State:          "Analyzed",
		Relations:      []string{"TDR-102", "TDR-103"},
		Summary:        "The library is outdated and causes security vulnerabilities.",
		Context:        "Originally chosen for quick implementation.",
		ImpactTech:     "Security risks and maintainability issues.",
		ImpactBus:      "Impact on customer satisfaction.",
		Symptoms:       "Error messages related to security protocols.",
		Severity:       "High",
		PotentialRisks: "Data breaches and legal consequences.",
		ProposedSol:    "Library replacement and implementation of 2FA.",
		CostDelay:      "Increased risk of security breaches.",
		Effort:         "4 weeks and €10,000.",
		Dependencies:   "Completion of the security audit.",
		Additional:     "Training for the development team.",
	}

	expected := `# Technical Debt Record

## Title
**Outdated Library**

## Author
Jane Doe

## Version
1.0.0

## Date
2024-04-15

## State
Analyzed

## Relations
- [TDR-102](#)
- [TDR-103](#)

## Summary
The library is outdated and causes security vulnerabilities.

## Context
Originally chosen for quick implementation.

## Impact
### Technical Impact
Security risks and maintainability issues.

### Business Impact
Impact on customer satisfaction.

## Symptoms
Error messages related to security protocols.

## Severity
High

## Potential Risks
Data breaches and legal consequences.

## Proposed Solution
Library replacement and implementation of 2FA.

## Cost of Delay
Increased risk of security breaches.

## Effort to Resolve
4 weeks and €10,000.

## Dependencies
Completion of the security audit.

## Additional Notes
Training for the development team.
`

	result := generateMarkdown(td)
	if strings.TrimSpace(result) != strings.TrimSpace(expected) {
		t.Errorf("generateMarkdown() = \n%s\n, want \n%s", result, expected)
	}
}

// Example excerpt from generate-td_test.go

func TestGenerateASCII(t *testing.T) {
	td := TechnicalDebt{
		Title:          "Outdated Library",
		Author:         "Jane Doe",
		Version:        "1.0.0",
		Date:           "2024-04-15",
		State:          "Analyzed",
		Relations:      []string{"TDR-102", "TDR-103"},
		Summary:        "The library is outdated and causes security vulnerabilities.",
		Context:        "Originally chosen for quick implementation.",
		ImpactTech:     "Security risks and maintainability issues.",
		ImpactBus:      "Impact on customer satisfaction.",
		Symptoms:       "Error messages related to security protocols.",
		Severity:       "High",
		PotentialRisks: "Data breaches and legal consequences.",
		ProposedSol:    "Library replacement and implementation of 2FA.",
		CostDelay:      "Increased risk of security breaches.",
		Effort:         "4 weeks and €10,000.",
		Dependencies:   "Completion of the security audit.",
		Additional:     "Training for the development team.",
	}

	expected := `Technical Debt Record
====================
    
Title:
------
Outdated Library
    
Author:
-------
Jane Doe
    
Version:
--------
1.0.0
    
Date:
-----
2024-04-15
    
State:
------
Analyzed
    
Relations:
----------
- TDR-102
- TDR-103
    
Summary:
--------
The library is outdated and causes security vulnerabilities.
    
Context:
--------
Originally chosen for quick implementation.
    
Impact:
-------
Technical Impact:
- Security risks and maintainability issues.
    
Business Impact:
- Impact on customer satisfaction.
    
Symptoms:
---------
Error messages related to security protocols.
    
Severity:
---------
High
    
Potential Risks:
----------------
Data breaches and legal consequences.
    
Proposed Solution:
-------------------
Library replacement and implementation of 2FA.
    
Cost of Delay:
---------------
Increased risk of security breaches.
    
Effort to Resolve:
-------------------
4 weeks and €10,000.
    
Dependencies:
-------------
Completion of the security audit.
    
Additional Notes:
-----------------
Training for the development team.
`

	result := generateASCII(td)
	if strings.TrimSpace(result) != strings.TrimSpace(expected) {
		t.Errorf("generateASCII() = \n%s\n, want \n%s", result, expected)
	}
}

// TestGeneratePDF checks the generation of a PDF file (basic existence check)
func TestGeneratePDF(t *testing.T) {
	td := TechnicalDebt{
		Title:          "Outdated Library",
		Author:         "Jane Doe",
		Version:        "1.0.0",
		Date:           "2024-04-15",
		State:          "Analyzed",
		Relations:      []string{"TDR-102", "TDR-103"},
		Summary:        "The library is outdated and causes security vulnerabilities.",
		Context:        "Originally chosen for quick implementation.",
		ImpactTech:     "Security risks and maintainability issues.",
		ImpactBus:      "Impact on customer satisfaction.",
		Symptoms:       "Error messages related to security protocols.",
		Severity:       "High",
		PotentialRisks: "Data breaches and legal consequences.",
		ProposedSol:    "Library replacement and implementation of 2FA.",
		CostDelay:      "Increased risk of security breaches.",
		Effort:         "4 weeks and €10,000.",
		Dependencies:   "Completion of the security audit.",
		Additional:     "Training for the development team.",
	}

	filename := "test_output.pdf"

	// Generate PDF
	err := generatePDF(td, filename)
	if err != nil {
		t.Fatalf("generatePDF() failed: %v", err)
	}

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("generatePDF() failed to create the file: %s", filename)
	}

	// Clean up
	os.Remove(filename)
}

// TestGenerateExcel checks the generation of an Excel file (basic existence check)
func TestGenerateExcel(t *testing.T) {
	td := TechnicalDebt{
		Title:          "Outdated Library",
		Author:         "Jane Doe",
		Version:        "1.0.0",
		Date:           "2024-04-15",
		State:          "Analyzed",
		Relations:      []string{"TDR-102", "TDR-103"},
		Summary:        "The library is outdated and causes security vulnerabilities.",
		Context:        "Originally chosen for quick implementation.",
		ImpactTech:     "Security risks and maintainability issues.",
		ImpactBus:      "Impact on customer satisfaction.",
		Symptoms:       "Error messages related to security protocols.",
		Severity:       "High",
		PotentialRisks: "Data breaches and legal consequences.",
		ProposedSol:    "Library replacement and implementation of 2FA.",
		CostDelay:      "Increased risk of security breaches.",
		Effort:         "4 weeks and €10,000.",
		Dependencies:   "Completion of the security audit.",
		Additional:     "Training for the development team.",
	}

	filename := "test_output.xlsx"

	// Generate Excel
	err := generateExcel(td, filename)
	if err != nil {
		t.Fatalf("generateExcel() failed: %v", err)
	}

	// Check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("generateExcel() failed to create the file: %s", filename)
	}

	// Clean up
	os.Remove(filename)
}
