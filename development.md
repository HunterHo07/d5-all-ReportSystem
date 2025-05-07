# Development Guidelines

This document outlines the development standards, architecture, and workflow for the Project Reporting System.

## Project Structure

```
project-reporting-system/
├── research.md                # Market research and analysis
├── readme.md                  # Project overview and documentation
├── development.md             # Development guidelines (this file)
├── demo/                      # Demo implementation
│   ├── index.html
│   ├── styles/
│   ├── scripts/
│   └── assets/
├── fe/                        # SvelteKit frontend implementation
│   ├── src/
│   │   ├── routes/
│   │   ├── lib/
│   │   ├── components/
│   │   └── stores/
│   ├── static/
│   └── tests/
├── be/                        # EncoreGo backend implementation
│   ├── api/
│   ├── auth/
│   ├── services/
│   ├── models/
│   └── tests/
├── db/                        # Database schema and migrations
│   ├── schema/
│   ├── migrations/
│   └── seeds/
└── deployment/                # Deployment configurations
    ├── cloudflare/
    ├── github-actions/
    └── monitoring/
```

## Architecture Overview

### System Architecture

The Project Reporting System follows a modern web application architecture:

1. **Frontend Layer**: SvelteKit application providing the user interface
2. **API Layer**: EncoreGo backend exposing gRPC endpoints
3. **Service Layer**: Business logic implementation
4. **Data Layer**: Pocketbase database with structured schemas
5. **Integration Layer**: Connectors to external systems

### Component Relationships

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  Frontend   │────▶│    API      │────▶│  Services   │
└─────────────┘     └─────────────┘     └─────────────┘
                                              │
                                              ▼
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│ Integrations│◀────│  Database   │◀────│   Models    │
└─────────────┘     └─────────────┘     └─────────────┘
```

### API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/reports` | GET | List all reports |
| `/api/reports` | POST | Create new report |
| `/api/reports/:id` | GET | Get report details |
| `/api/reports/:id` | PUT | Update report |
| `/api/reports/:id/submit` | POST | Submit report to department |
| `/api/reports/:id/evaluate` | POST | Add/update evaluation scores |
| `/api/departments` | GET | List available departments |
| `/api/users` | GET | List users (admin only) |
| `/api/analytics` | GET | Get reporting analytics |

### Database Schema

#### Reports Table
- `id`: UUID (primary key)
- `title`: String
- `description`: Text
- `project_id`: UUID (foreign key)
- `author_id`: UUID (foreign key)
- `department_id`: UUID (foreign key)
- `status`: Enum (draft, submitted, approved, rejected)
- `created_at`: Timestamp
- `updated_at`: Timestamp
- `submitted_at`: Timestamp
- `metadata`: JSON

#### Evaluations Table
- `id`: UUID (primary key)
- `report_id`: UUID (foreign key)
- `security_score`: Integer (0-10)
- `performance_score`: Integer (0-10)
- `memory_score`: Integer (0-10)
- `testing_score`: Integer (0-10)
- `error_score`: Integer (0-10)
- `load_score`: Integer (0-10)
- `security_details`: Text
- `performance_details`: Text
- `memory_details`: Text
- `testing_details`: Text
- `error_details`: Text
- `load_details`: Text
- `evaluator_id`: UUID (foreign key)
- `created_at`: Timestamp
- `updated_at`: Timestamp

## Development Standards

### Code Organization Rules

All code should follow our scope and direction rules:

```javascript
// BE-IN - Internal backend only
function processInternalData() {
  // Implementation
}

// FE-OUT - Frontend with external data
function fetchAndDisplayData() {
  // Implementation
}
```

### Scoring System

Evaluate each component using the following format:

```
// Score: [S{s},P{p},M{m},T{t},E{e},L{l}]
// Details:
// - Security (S{s}): [details about security implementation]
// - Performance (P{p}): [details about performance optimizations]
// - Memory (M{m}): [details about memory management]
// - Testing (T{t}): [details about testing coverage]
// - Error (E{e}): [details about error handling]
// - Load (L{l}): [details about load impact]
// Tags: [type]-[scope]-[impact]
```

### Coding Standards

#### Frontend (SvelteKit)
- Use TypeScript for all components
- Follow component-based architecture
- Implement proper state management with stores
- Use Zod for form validation
- Document all components with JSDoc

#### Backend (EncoreGo)
- Follow clean architecture principles
- Implement proper error handling
- Use dependency injection
- Write comprehensive tests
- Document all endpoints

## Development Workflow

1. **Feature Planning**
   - Create feature specification
   - Define acceptance criteria
   - Design component architecture

2. **Implementation**
   - Develop backend services
   - Implement frontend components
   - Write tests for all functionality

3. **Testing**
   - Run unit tests
   - Perform integration testing
   - Conduct UI/UX testing

4. **Review & Refinement**
   - Code review
   - Performance optimization
   - Security audit

5. **Deployment**
   - Stage for testing
   - Verify in staging environment
   - Deploy to production

## Getting Started

### Prerequisites
- Node.js 18+
- Go 1.20+
- Docker
- Pocketbase CLI

### Setup Instructions
1. Clone the repository
2. Run `npm install` in the `fe` directory
3. Run `go mod download` in the `be` directory
4. Start Pocketbase with `pocketbase serve` in the `db` directory
5. Start the frontend with `npm run dev` in the `fe` directory
6. Start the backend with `encore run` in the `be` directory

## Testing

- Frontend tests: `npm test` in the `fe` directory
- Backend tests: `go test ./...` in the `be` directory
- E2E tests: `npm run test:e2e` in the root directory
