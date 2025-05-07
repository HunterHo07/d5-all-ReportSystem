# Project Reporting System Frontend

This directory contains the SvelteKit frontend implementation of the Project Reporting System.

## Technology Stack

- **Framework**: SvelteKit
- **UI Components**: Shadcn-Svelte
- **Form Management**: Superform with Zod validation
- **Animations**: GSAP, Svelte Motion
- **Styling**: TailwindCSS

## Project Structure

```
fe/
├── src/
│   ├── routes/                 # SvelteKit routes
│   │   ├── +layout.svelte      # Main layout
│   │   ├── +page.svelte        # Dashboard page
│   │   └── reports/            # Report-related routes
│   │       ├── +page.svelte    # Reports list
│   │       └── new/            # New report form
│   ├── lib/                    # Shared code
│   │   ├── components/         # Reusable components
│   │   │   └── ReportForm.svelte  # Report form component
│   │   └── stores/             # Svelte stores
│   └── app.css                 # Global styles
├── static/                     # Static assets
├── tests/                      # Test files
├── package.json                # Dependencies
└── svelte.config.js            # SvelteKit configuration
```

## Features

- **Interactive Dashboard**: Overview of reports and metrics
- **Report Creation**: Form with validation for creating new reports
- **Standardized Scoring**: Implementation of the scoring system
- **Real-time Validation**: Immediate feedback on form inputs
- **Responsive Design**: Works on all device sizes

## Scoring System Implementation

The frontend implements the standardized scoring system with the following features:

- Visual sliders for each score dimension (S, P, M, T, E, L)
- Color-coded indicators based on score values
- Detailed text fields for justification
- Automatic calculation of overall score
- Formatted score display in the standard format: [S8,P7,M6,T7,E8,L6]

## Development

### Prerequisites

- Node.js 18+
- npm or yarn

### Setup

1. Install dependencies:
   ```bash
   npm install
   ```

2. Start the development server:
   ```bash
   npm run dev
   ```

3. Open your browser at `http://localhost:5173`

### Building for Production

```bash
npm run build
```

The built application will be in the `build` directory.

## Testing

### Running Tests

```bash
# Run unit tests
npm test

# Run with coverage
npm test -- --coverage

# Run end-to-end tests
npm run test:e2e
```

## Evaluation

The frontend implementation has been evaluated using our scoring system:

```
// Score: [S8,P8,M7,T8,E8,L7]
// Details:
// - Security (S8): Input validation, CSRF protection, secure forms
// - Performance (P8): Code splitting, lazy loading, optimized assets
// - Memory (M7): Efficient state management, minimal re-renders
// - Testing (T8): Unit tests, component tests, e2e tests
// - Error (E8): Comprehensive validation, user feedback, error boundaries
// - Load (L7): Handles large datasets, efficient rendering
// Tags: FE-module-high
```
