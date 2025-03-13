
![weaver_Banner](https://github.com/user-attachments/assets/bd3dec3d-bd8e-4d2e-bf15-fadefb8ad7ed)
# Weaver

Remember when you start a new backend project and you have to re-write the same stuff again and again everytime T_T, well not anymore :D,
Weaver is a powerful CLI tool designed to rapidly scaffold backend applications with CRUD operations and authentication, eliminating repetitive boilerplate code.

## Features

- **Multiple Framework Support**: Generate projects using popular Go web frameworks:
    - Chi
    - Echo
    - Fiber
    - Standard HTTP package

- **Authentication Integration**: Easily add OAuth providers to your project:
    - Google
    - GitHub
    - Discord

- **MongoDB Integration**: All generated projects come with MongoDB database connectivity out of the box (yeah don't worry I'll add others...... later (¬⤙¬ ) )

- **Complete CRUD Operations**: Each project includes ready-to-use endpoints for Create, Read, Update, and Delete operations

- **Interactive CLI**: User-friendly command-line interface with interactive prompts

## Installation

```bash
go install github.com/PrathamX595/weaver@latest
```

## Usage

### Basic Command

```bash
weaver create
```

This will launch an interactive CLI that guides you through project creation.

### Command Line Options

```bash
weaver create --name "MyProject" --framework "fiber" --auth "google,github"
```

Available options:
- `--name`: Project name
- `--framework`: Framework selection (chi, echo, fiber, http)
- `--auth`: Authentication providers (google, github, discord)

## Project Structure

Generated projects follow a clean, maintainable structure:

```
MyProject/
├── auth/           # Authentication handlers
├── config/         # Database and environment config
├── controller/     # Route handlers
├── models/         # Data models
├── routes/         # Route definitions
├── utils/          # Helper functions
├── .env.example    # Environment variables template
├── .gitignore      # Git ignore file
├── go.mod          # Go module file
└── server.go       # Main application entry
```

## Environment Variables

After creating your project, you'll need to set up your environment variables:

1. Copy `.env.example` to `.env`
2. Fill in your MongoDB connection string
3. Add authentication credentials if you selected auth providers

## License

This project is licensed under the Apache License 2.0 - see the LICENSE file for details.

## Personal Notes

yeah I made the mascot myself (don't judge me alright I had some free time (¬_¬") )
