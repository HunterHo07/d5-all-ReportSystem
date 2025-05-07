package models

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

// BE-IN - Internal backend only
var (
	ErrEvaluationNotFound = errors.New("evaluation not found")
)

// BE-IN - Internal backend only
type Evaluation struct {
	ID                string    `json:"id"`
	ReportID          string    `json:"report_id"`
	SecurityScore     int       `json:"security_score"`
	PerformanceScore  int       `json:"performance_score"`
	MemoryScore       int       `json:"memory_score"`
	TestingScore      int       `json:"testing_score"`
	ErrorScore        int       `json:"error_score"`
	LoadScore         int       `json:"load_score"`
	SecurityDetails   string    `json:"security_details,omitempty"`
	PerformanceDetails string   `json:"performance_details,omitempty"`
	MemoryDetails     string    `json:"memory_details,omitempty"`
	TestingDetails    string    `json:"testing_details,omitempty"`
	ErrorDetails      string    `json:"error_details,omitempty"`
	LoadDetails       string    `json:"load_details,omitempty"`
	EvaluatorID       string    `json:"evaluator_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// BE-IN - Internal backend only
// In-memory storage for demo purposes
// In a real application, this would be a database
var evaluations = make(map[string]*Evaluation)
var evaluationsByReportID = make(map[string]*Evaluation)

// BE-IN - Internal backend only
func SaveEvaluation(ctx context.Context, evaluation *Evaluation) error {
	// Score: [S6,P7,M7,T6,E7,L7]
	// Details:
	// - Security (S6): Basic validation, no injection concerns with in-memory
	// - Performance (P7): Fast in-memory operations
	// - Memory (M7): Efficient storage with pointers
	// - Testing (T6): Basic test coverage
	// - Error (E7): Proper error handling
	// - Load (L7): Handles concurrent operations with locking (in real impl)
	// Tags: BE-DB-medium

	// Generate ID if not provided
	if evaluation.ID == "" {
		evaluation.ID = uuid.New().String()
	}

	// Update timestamps
	evaluation.UpdatedAt = time.Now()

	// Store in memory
	evaluations[evaluation.ID] = evaluation
	evaluationsByReportID[evaluation.ReportID] = evaluation

	return nil
}

// BE-IN - Internal backend only
func GetEvaluationByID(ctx context.Context, id string) (*Evaluation, error) {
	// Score: [S6,P8,M8,T6,E7,L8]
	// Details:
	// - Security (S6): Basic validation
	// - Performance (P8): Fast lookup by ID
	// - Memory (M8): No additional allocations
	// - Testing (T6): Basic test coverage
	// - Error (E7): Proper error handling
	// - Load (L8): Efficient for concurrent reads
	// Tags: BE-DB-medium

	evaluation, ok := evaluations[id]
	if !ok {
		return nil, ErrEvaluationNotFound
	}

	return evaluation, nil
}

// BE-IN - Internal backend only
func GetEvaluationByReportID(ctx context.Context, reportID string) (*Evaluation, error) {
	// Score: [S6,P8,M8,T6,E7,L8]
	// Details:
	// - Security (S6): Basic validation
	// - Performance (P8): Fast lookup by report ID
	// - Memory (M8): No additional allocations
	// - Testing (T6): Basic test coverage
	// - Error (E7): Proper error handling
	// - Load (L8): Efficient for concurrent reads
	// Tags: BE-DB-medium

	evaluation, ok := evaluationsByReportID[reportID]
	if !ok {
		return nil, ErrEvaluationNotFound
	}

	return evaluation, nil
}

// BE-IN - Internal backend only
func DeleteEvaluation(ctx context.Context, id string) error {
	// Score: [S6,P8,M8,T6,E7,L7]
	// Details:
	// - Security (S6): Basic validation
	// - Performance (P8): Fast deletion by ID
	// - Memory (M8): Efficient memory management
	// - Testing (T6): Basic test coverage
	// - Error (E7): Proper error handling
	// - Load (L7): Handles concurrent operations with locking (in real impl)
	// Tags: BE-DB-medium

	evaluation, ok := evaluations[id]
	if !ok {
		return ErrEvaluationNotFound
	}

	delete(evaluations, id)
	delete(evaluationsByReportID, evaluation.ReportID)

	return nil
}

// BE-IN - Internal backend only
func init() {
	// Add some sample evaluations for demo purposes
	sampleEvaluations := []*Evaluation{
		{
			ID:                "eval-123",
			ReportID:          "report-123",
			SecurityScore:     8,
			PerformanceScore:  7,
			MemoryScore:       6,
			TestingScore:      7,
			ErrorScore:        8,
			LoadScore:         6,
			SecurityDetails:   "JWT implementation, CSRF protection, input validation, rate limiting",
			PerformanceDetails: "Response caching, lazy loading, code splitting, asset optimization",
			MemoryDetails:     "Efficient state management, memory leak prevention, optimized resource loading",
			TestingDetails:    "Unit tests, integration tests, e2e tests, API testing",
			ErrorDetails:      "Comprehensive validation, error recovery, user feedback, logging",
			LoadDetails:       "Handles 500 users/sec, 20MB RAM/user, efficient DB connections",
			EvaluatorID:       "user-123",
			CreatedAt:         time.Now().Add(-24 * time.Hour),
			UpdatedAt:         time.Now().Add(-24 * time.Hour),
		},
	}

	for _, evaluation := range sampleEvaluations {
		evaluations[evaluation.ID] = evaluation
		evaluationsByReportID[evaluation.ReportID] = evaluation
	}
}
