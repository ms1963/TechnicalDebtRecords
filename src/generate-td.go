//////////////////////////////////////////////////////////////////////////
//
// generate-md, (c) 2024, Michael Stal
// Published unter MIT license
// and written in Go (Golang)
//
// This tool allows to create Technical Debt Records in different formats.
// Call tool with command line parameter -help to get more information 
// about usage.
//
/////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/phpdave11/gofpdf"
	"github.com/xuri/excelize/v2"
)

// TechnicalDebt represents a technical debt record
type TechnicalDebt struct {
	Title          string
	Author         string
	Version        string
	Date           string
	State          string // New field added
	Relations      []string
	Summary        string
	Context        string
	ImpactTech     string
	ImpactBus      string
	Symptoms       string
	Severity       string
	PotentialRisks string
	ProposedSol    string
	CostDelay      string
	Effort         string
	Dependencies   string
	Additional     string
	Empty          bool
}

// AllowedStates defines the possible states of a technical debt record
var AllowedStates = []string{
	"Identified",
	"Analyzed",
	"Approved",
	"In Progress",
	"Resolved",
	"Closed",
	"Rejected",
}

// getInput prompts the user for input and returns the entered value
func getInput(prompt string, required bool) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		input = strings.TrimSpace(input)
		if required && input == "" {
			fmt.Println("This field is required. Please enter a value.")
			continue
		}
		return input, nil
	}
}

// getRelations prompts the user to enter related Technical Debt IDs
func getRelations() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	var relations []string
	fmt.Println("\nEnter related Technical Debt IDs (leave blank to finish):")
	for {
		fmt.Print(" - Related TD ID: ")
		rel, err := reader.ReadString('\n')
		if err != nil {
			return relations, err
		}
		rel = strings.TrimSpace(rel)
		if rel == "" {
			break
		}
		relations = append(relations, rel)
	}
	return relations, nil
}

// getState prompts the user to select a state from predefined options
func getState() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nSelect the State of the Technical Debt:")
		for i, state := range AllowedStates {
			fmt.Printf("  %d) %s\n", i+1, state)
		}
		fmt.Print("Enter the number corresponding to the state: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		input = strings.TrimSpace(input)
		// Convert input to integer
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(AllowedStates) {
			fmt.Println("Invalid selection. Please enter a valid number corresponding to the state.")
			continue
		}
		return AllowedStates[choice-1], nil
	}
}

// validateTechnicalDebt checks if the required fields are filled
func validateTechnicalDebt(td TechnicalDebt) error {
	if td.Title == "" {
		return errors.New("Title is required")
	}
	if td.Author == "" {
		return errors.New("Author is required")
	}
	if td.Version == "" {
		return errors.New("Version is required")
	}
	if td.Date == "" {
		return errors.New("Date is required")
	}
	if td.State == "" {
		return errors.New("State is required")
	}
	return nil
}

// generateMarkdown generates the Markdown content
func generateMarkdown(td TechnicalDebt) string {
	relationsFormatted := "None"
	if len(td.Relations) > 0 {
		var rels []string
		for _, rel := range td.Relations {
			rels = append(rels, fmt.Sprintf("- [%s](#)", rel))
		}
		relationsFormatted = strings.Join(rels, "\n")
	}

	if td.Empty {
		return fmt.Sprintf(`# Technical Debt Record

## Title
**[Enter Title Here]**

## Author
[Enter Author Here]

## Version
[Enter Version Here]

## Date
[Enter Date Here]

## State
[Enter State Here]

## Summary
*A brief overview of the technical debt, explaining the problem in one or two sentences.*

## Context
*Provide the historical context and reasons why this technical debt exists.*

## Impact
### Technical Impact
*Describe how the debt affects the system’s performance, scalability, or maintainability.*

### Business Impact
*Explain how the debt affects the business, such as increased risk, customer dissatisfaction, or slower feature delivery.*

## Symptoms
*List specific signs that indicate the presence of technical debt.*

## Risk Assessment
- **Severity**: *[Enter Severity Here: Critical / High / Medium / Low]*
- **Potential Risks**: *Potential security vulnerabilities leading to data breaches.*

## Proposed Solution
*Describe how to resolve the technical debt.*

## Cost of Delay
*Explain the consequences of delaying the resolution of the technical debt.*

## Effort to Resolve
*Estimate the time, resources, and effort needed to address the debt.*

## Dependencies
*List any dependencies or blockers that need to be resolved before tackling the debt.*

## Relations with Other Technical Debt Records
%s

## Additional Notes
*Any other relevant information or considerations.*
`, relationsFormatted)
	}

	// Normal Markdown generation
	return fmt.Sprintf(`# Technical Debt Record

## Title
**%s**

## Author
%s

## Version
%s

## Date
%s

## State
%s

## Summary
%s

## Context
%s

## Impact
### Technical Impact
%s

### Business Impact
%s

## Symptoms
%s

## Risk Assessment
- **Severity**: %s
- **Potential Risks**: %s

## Proposed Solution
%s

## Cost of Delay
%s

## Effort to Resolve
%s

## Dependencies
%s

## Relations with Other Technical Debt Records
%s

## Additional Notes
%s
`, td.Title, td.Author, td.Version, td.Date, td.State, td.Summary, td.Context, td.ImpactTech, td.ImpactBus,
		td.Symptoms, td.Severity, td.PotentialRisks, td.ProposedSol, td.CostDelay, td.Effort,
		td.Dependencies, relationsFormatted, td.Additional)
}

// generateASCII generates the Plain ASCII content
func generateASCII(td TechnicalDebt) string {
	relationsFormatted := "None"
	if len(td.Relations) > 0 {
		var rels []string
		for _, rel := range td.Relations {
			rels = append(rels, fmt.Sprintf("- %s", rel))
		}
		relationsFormatted = strings.Join(rels, "\n")
	}

	if td.Empty {
		return fmt.Sprintf(`Technical Debt Record
====================
    
Title:
------
[Enter Title Here]
    
Author:
-------
[Enter Author Here]
    
Version:
--------
[Enter Version Here]
    
Date:
-----
[Enter Date Here]
    
State:
------
[Enter State Here]
    
Summary:
--------
*A brief overview of the technical debt, explaining the problem in one or two sentences.*
    
Context:
--------
*Provide the historical context and reasons why this technical debt exists.*
    
Impact:
-------
Technical Impact:
- *Describe the technical impact.*
    
Business Impact:
- *Describe the business impact.*
    
Symptoms:
---------
*List symptoms of technical debt.*
    
Risk Assessment:
----------------
- **Severity**: *[Enter Severity Here: Critical / High / Medium / Low]*
- **Potential Risks**: *Potential security vulnerabilities leading to data breaches.*
    
Proposed Solution:
-------------------
*Describe how to resolve the technical debt.*
    
Cost of Delay:
--------------
*Explain the consequences of delaying the resolution of the technical debt.*
    
Effort to Resolve:
-------------------
*Estimate the time, resources, and effort needed to address the debt.*
    
Dependencies:
-------------
*List any dependencies.*
    
Relations with Other Technical Debt Records:
--------------------------------------------
%s
    
Additional Notes:
-----------------
*Any other relevant information or considerations.*
`, relationsFormatted)
	}

	// Normal ASCII generation
	return fmt.Sprintf(`Technical Debt Record
====================

Title:
------
%s

Author:
-------
%s

Version:
--------
%s

Date:
-----
%s

State:
------
%s

Summary:
--------
%s

Context:
--------
%s

Impact:
-------
Technical Impact:
- %s

Business Impact:
- %s

Symptoms:
---------
%s

Risk Assessment:
----------------
- **Severity**: %s
- **Potential Risks**: %s

Proposed Solution:
-------------------
%s

Cost of Delay:
--------------
%s

Effort to Resolve:
-------------------
%s

Dependencies:
-------------
%s

Relations with Other Technical Debt Records:
--------------------------------------------
%s

Additional Notes:
-----------------
%s
`, td.Title, td.Author, td.Version, td.Date, td.State, td.Summary, td.Context, td.ImpactTech, td.ImpactBus,
		td.Symptoms, td.Severity, td.PotentialRisks, td.ProposedSol, td.CostDelay, td.Effort,
		td.Dependencies, relationsFormatted, td.Additional)
}

// generatePDF generates a PDF file using the gofpdf library
func generatePDF(td TechnicalDebt, filename string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Set font and size for the title
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Technical Debt Record")
	pdf.Ln(12)

	// Add sections to PDF
	addPDFSection(pdf, "Title", td.Title)
	addPDFSection(pdf, "Author", td.Author)
	addPDFSection(pdf, "Version", td.Version)
	addPDFSection(pdf, "Date", td.Date)
	addPDFSection(pdf, "State", td.State)
	addPDFSection(pdf, "Summary", td.Summary)
	addPDFSection(pdf, "Context", td.Context)
	addPDFSection(pdf, "Technical Impact", td.ImpactTech)
	addPDFSection(pdf, "Business Impact", td.ImpactBus)
	addPDFSection(pdf, "Symptoms", td.Symptoms)
	addPDFSection(pdf, "Severity", td.Severity)
	addPDFSection(pdf, "Potential Risks", td.PotentialRisks)
	addPDFSection(pdf, "Proposed Solution", td.ProposedSol)
	addPDFSection(pdf, "Cost of Delay", td.CostDelay)
	addPDFSection(pdf, "Effort to Resolve", td.Effort)
	addPDFSection(pdf, "Dependencies", td.Dependencies)

	// Output the PDF
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return fmt.Errorf("error saving PDF file: %w", err)
	}
	return nil
}

// addPDFSection adds a section to a PDF document
func addPDFSection(pdf *gofpdf.Fpdf, title, content string) {
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, title+":")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, content, "", "", false)
	pdf.Ln(6)
}

// generateExcel generates an Excel file using the excelize library
func generateExcel(td TechnicalDebt, filename string) error {
	f := excelize.NewFile()

	// Create a sheet
	sheet, err := f.NewSheet("TechnicalDebt")
	if err != nil {
		return err
	}

	// Set headers
	headers := []string{
		"Title",
		"Author",
		"Version",
		"Date",
		"State",
		"Summary",
		"Context",
		"Technical Impact",
		"Business Impact",
		"Symptoms",
		"Severity",
		"Potential Risks",
		"Proposed Solution",
		"Cost of Delay",
		"Effort to Resolve",
		"Dependencies",
		"Relations with Other Technical Debt Records",
		"Additional Notes",
	}

	for i, header := range headers {
		cell, _ := excelize.CoordinatesToCellName(1, i+1)
		f.SetCellValue("TechnicalDebt", cell, header)
	}

	// Set values
	values := []string{
		td.Title,
		td.Author,
		td.Version,
		td.Date,
		td.State,
		td.Summary,
		td.Context,
		td.ImpactTech,
		td.ImpactBus,
		td.Symptoms,
		td.Severity,
		td.PotentialRisks,
		td.ProposedSol,
		td.CostDelay,
		td.Effort,
		td.Dependencies,
	}

	// Relations
	if len(td.Relations) > 0 {
		values = append(values, strings.Join(td.Relations, ", "))
	} else {
		values = append(values, "None")
	}

	// Additional Notes
	values = append(values, td.Additional)

	for i, value := range values {
		cell, _ := excelize.CoordinatesToCellName(2, i+1)
		f.SetCellValue("TechnicalDebt", cell, value)
	}

	// Set active sheet and save the Excel file
	f.SetActiveSheet(sheet)
	if err := f.SaveAs(filename); err != nil {
		return fmt.Errorf("error saving Excel file: %w", err)
	}
	return nil
}

func main() {
	fmt.Println("generate-td  (c) Michael Stal, 2024")
	// Define command-line flags
	formatPtr := flag.String("format", "markdown", "Output format: markdown, ascii, pdf, excel")
	filenamePtr := flag.String("output", "", "Output filename (optional)")
	emptyPtr := flag.Bool("empty", false, "Generate an empty template")

	// Custom usage message
	flag.Usage = func() {
		usageText := `Usage: generate_td [OPTIONS]

Generates a technical debt record in the specified format.

Options:
  -format string
        Output format: markdown, ascii, pdf, excel (default "markdown")
  -output string
        Output filename (optional). If not provided, a default filename with the appropriate extension is generated.
  -empty
        Generate an empty template with placeholders without prompting for input.
  -h, --help
        Show this help message and exit.

Examples:
  Generate a Markdown file:
        generate_td -format markdown

  Generate a PDF file with a custom filename:
        generate_td -format pdf -output my_debt_record.pdf

  Generate an Excel file with an empty template:
        generate_td -format excel -empty

  Show help:
        generate_td --help
`
		fmt.Fprintf(flag.CommandLine.Output(), usageText)
	}

	// Parse flags
	flag.Parse()

	// Check if help is requested
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			if arg == "-h" || arg == "--help" {
				flag.Usage()
				return
			}
		}
	}

	// Supported formats
	supportedFormats := map[string]bool{
		"markdown": true,
		"ascii":    true,
		"pdf":      true,
		"excel":    true,
	}

	// Validate format
	format := strings.ToLower(*formatPtr)
	if !supportedFormats[format] {
		fmt.Println("Unsupported format. Supported formats are: markdown, ascii, pdf, excel")
		return
	}

	// Map formats to file extensions
	formatExtensions := map[string]string{
		"markdown": ".md",
		"ascii":    ".txt",
		"pdf":      ".pdf",
		"excel":    ".xlsx",
	}

	// Determine output filename
	var filename string
	if *filenamePtr != "" {
		ext := filepath.Ext(*filenamePtr)
		if ext == "" {
			// Append the correct extension
			filename = *filenamePtr + formatExtensions[format]
			fmt.Printf("No file extension provided. Appending '%s' to the filename.\n", formatExtensions[format])
		} else {
			// Check if the extension matches the format
			expectedExt := formatExtensions[format]
			if strings.ToLower(ext) != expectedExt {
				fmt.Printf("Warning: The provided filename extension '%s' does not match the format '%s'. Expected '%s'.\n",
					ext, format, expectedExt)
			}
			filename = *filenamePtr
		}
	} else {
		// Generate default filename
		filename = "technical_debt_record" + formatExtensions[format]
	}

	// Create an empty technical debt record if the -empty flag is set
	td := TechnicalDebt{Empty: *emptyPtr}

	// If not generating an empty file, prompt user for inputs
	if !*emptyPtr {
		var err error

		td.Title, err = getInput("Enter the Title of the Technical Debt: ", true)
		if err != nil {
			fmt.Println("Error reading title:", err)
			return
		}

		td.Author, err = getInput("Enter the Author of the Document: ", true)
		if err != nil {
			fmt.Println("Error reading author:", err)
			return
		}

		td.Version, err = getInput("Enter the Version (e.g., 1.0.0): ", true)
		if err != nil {
			fmt.Println("Error reading version:", err)
			return
		}

		// Prompt for Date with default as today
		dateInput, err := getInput("Enter the Date (YYYY-MM-DD) [Leave blank for today]: ", false)
		if err != nil {
			fmt.Println("Error reading date:", err)
			return
		}
		if dateInput == "" {
			td.Date = time.Now().Format("2006-01-02")
		} else {
			// Validate date format
			_, err := time.Parse("2006-01-02", dateInput)
			if err != nil {
				fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
				return
			}
			td.Date = dateInput
		}

		// Get State
		td.State, err = getState()
		if err != nil {
			fmt.Println("Error reading state:", err)
			return
		}

		td.Relations, err = getRelations()
		if err != nil {
			fmt.Println("Error reading relations:", err)
			return
		}

		// Additional fields
		td.Summary, err = getInput("Enter Summary: ", false)
		if err != nil {
			fmt.Println("Error reading summary:", err)
			return
		}

		td.Context, err = getInput("Enter Context: ", false)
		if err != nil {
			fmt.Println("Error reading context:", err)
			return
		}

		td.ImpactTech, err = getInput("Enter Technical Impact: ", false)
		if err != nil {
			fmt.Println("Error reading technical impact:", err)
			return
		}

		td.ImpactBus, err = getInput("Enter Business Impact: ", false)
		if err != nil {
			fmt.Println("Error reading business impact:", err)
			return
		}

		td.Symptoms, err = getInput("Enter Symptoms: ", false)
		if err != nil {
			fmt.Println("Error reading symptoms:", err)
			return
		}

		td.Severity, err = getInput("Enter Severity (Critical / High / Medium / Low): ", false)
		if err != nil {
			fmt.Println("Error reading severity:", err)
			return
		}

		td.PotentialRisks, err = getInput("Enter Potential Risks: ", false)
		if err != nil {
			fmt.Println("Error reading potential risks:", err)
			return
		}

		td.ProposedSol, err = getInput("Enter Proposed Solution: ", false)
		if err != nil {
			fmt.Println("Error reading proposed solution:", err)
			return
		}

		td.CostDelay, err = getInput("Enter Cost of Delay: ", false)
		if err != nil {
			fmt.Println("Error reading cost of delay:", err)
			return
		}

		td.Effort, err = getInput("Enter Effort to Resolve: ", false)
		if err != nil {
			fmt.Println("Error reading effort:", err)
			return
		}

		td.Dependencies, err = getInput("Enter Dependencies: ", false)
		if err != nil {
			fmt.Println("Error reading dependencies:", err)
			return
		}

		td.Additional, err = getInput("Enter Additional Notes: ", false)
		if err != nil {
			fmt.Println("Error reading additional notes:", err)
			return
		}

		// Validate inputs
		if err := validateTechnicalDebt(td); err != nil {
			fmt.Println("Validation error:", err)
			return
		}
	}

	// Generate content based on format
	switch format {
	case "markdown":
		content := generateMarkdown(td)
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error generating Markdown file:", err)
			return
		}
	case "ascii":
		content := generateASCII(td)
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error generating ASCII file:", err)
			return
		}
	case "pdf":
		err := generatePDF(td, filename)
		if err != nil {
			fmt.Println("Error generating PDF file:", err)
			return
		}
	case "excel":
		err := generateExcel(td, filename)
		if err != nil {
			fmt.Println("Error generating Excel file:", err)
			return
		}
	default:
		fmt.Println("Unsupported format.")
		return
	}

	fmt.Printf("\nTechnical Debt record has been saved to '%s'.\n", filename)
}

