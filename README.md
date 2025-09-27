[Slot Machine API.postman_collection.json](https://github.com/user-attachments/files/22573378/Slot.Machine.API.postman_collection.json)# Go Slot Machine Backend
<img width="714" height="684" alt="Screenshot 2025-09-27 at 2 55 06 PM" src="https://github.com/user-attachments/assets/2ae9e62d-3575-4431-b09a-7b706de4ac3e" />

[Uploading Slot Machine A{
	"info": {
		"_postman_id": "b283aaa1-ecdd-4342-8f38-5cc269331ed7",
		"name": "Slot Machine API",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "21479280"
	},
	"item": [
		{
			"name": "Create Game Session",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"name\": \"Ayman\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "{{host}}/api/v1/sessions",
					"host": [
						"{{host}}"
					],
					"path": [
						"api",
						"v1",
						"sessions"
					]
				}
			},
			"response": []
		},
		{
			"name": "Get Current Session",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"name\": \"Ayman\"\n}"
				},
				"url": {
					"raw": "{{host}}/api/v1/session",
					"host": [
						"{{host}}"
					],
					"path": [
						"api",
						"v1",
						"session"
					]
				}
			},
			"response": []
		},
		{
			"name": "Roll",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"name\": \"Ayman\"\n}"
				},
				"url": {
					"raw": "{{host}}/api/v1/roll",
					"host": [
						"{{host}}"
					],
					"path": [
						"api",
						"v1",
						"roll"
					]
				}
			},
			"response": []
		},
		{
			"name": "Cash Out",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"account_number\": \"123\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "{{host}}/api/v1/cash-out",
					"host": [
						"{{host}}"
					],
					"path": [
						"api",
						"v1",
						"cash-out"
					]
				}
			},
			"response": []
		}
	],
	"event": [
		{
			"listen": "prerequest",
			"script": {
				"type": "text/javascript",
				"packages": {},
				"requests": {},
				"exec": [
					""
				]
			}
		},
		{
			"listen": "test",
			"script": {
				"type": "text/javascript",
				"packages": {},
				"requests": {},
				"exec": [
					""
				]
			}
		}
	],
	"variable": [
		{
			"key": "host",
			"value": ""
		},
		{
			"key": "token",
			"value": ""
		}
	]
}PI.postman_collection.json…]()


A full-stack slot machine game built with Go backend and vanilla JavaScript frontend. Features session management, progressive slot reveals, credit tracking, and server-side cheating mechanics based on player credit levels.

## 🎮 Features

- **Dual UI Interfaces**: Beautiful emoji-based view and minimalistic table-based view
- **Session Management**: Cookie-based authentication with game sessions
- **Server-Side Cheating**: Dynamic reroll chances based on credit balance
- **Progressive Reveals**: Slots reveal sequentially (1s, 2s, 3s intervals)
- **Credit System**: Win rewards, roll costs, and cash-out functionality
- **Dodging Cash-Out Button**: 50% dodge chance, 40% unclickable chance (client-side)

## 🛠 Technologies Used

### Backend
- **Go 1.23.2** - Main language
- **Chi Router** - HTTP routing and middleware
- **SQLC** - Type-safe SQL code generation
- **MySQL 8.0** - Database
- **Templ** - Type-safe HTML templating
- **Zerolog** - Structured logging

### Frontend
- **Vanilla JavaScript** - Client-side logic
- **HTML5/CSS3** - UI and styling
- **Fetch API** - HTTP requests
- **CSS Animations** - Slot spinning effects

### Development Tools
- **Docker** - Database containerization
- **Templ CLI** - Template generation
- **SQLC CLI** - SQL code generation

## 📋 Prerequisites

- Go 1.23.2 or higher
- Docker and Docker Compose
- Make (optional)

## 🚀 Setup Instructions

### 1. Clone Repository
```bash
git clone <repository-url>
cd go-backend
```

### 2. Install Dependencies
```bash
go mod download
```

### 3. Install Required Tools
```bash
# Install Templ for template generation
go install github.com/a-h/templ/cmd/templ@latest

# Install SQLC for SQL code generation
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### 4. Start Database
```bash
cd local-env
./up.sh
# Or manually: docker-compose up -d
```

This starts MySQL on port 4306 with:
- Database: `default`
- User: `dbuser`
- Password: `dbuserpassword`

### 5. Generate Code
```bash
# Generate database models from SQL
sqlc generate

# Generate HTML templates
templ generate
```

### 6. Build and Run
```bash
# Build the application
./build.sh
# Or manually:
# go build -o api cmd/api/main.go

# Run the server
./api
```

The server will start on the port specified in your configuration (check `internal/app/config/config.go`).

## 🎯 Usage

### Access the Game
- **Main Interface**: `http://localhost:<port>/` or `http://localhost:<port>/slot-machine`
- **Minimal Interface**: `http://localhost:<port>/minimal` or `http://localhost:<port>/minimal-slot`

### Game Flow
1. **Create Session**: Click "Create Game Session" / "Start Game"
2. **Roll Slots**: Each roll costs 1 credit
3. **Win Conditions**: Match all 3 symbols
   - Cherry (C/🍒): 10 credits
   - Lemon (L/🍋): 20 credits
   - Orange (O/🍊): 30 credits
   - Watermelon (W/🍉): 40 credits
4. **Cash Out**: End session and transfer credits to account

### Server Mechanics
- **< 40 credits**: Truly random rolls
- **40-60 credits**: 30% chance to reroll winning combinations
- **> 60 credits**: 60% chance to reroll winning combinations

## 🏗 Project Structure

```
go-backend/
├── cmd/
│   ├── api/main.go              # API server entry point
│   └── cli/                     # CLI commands
├── internal/app/
│   ├── api/
│   │   ├── apihandlers/         # JSON API handlers
│   │   ├── viewhandlers/        # HTML template handlers
│   │   ├── middleware/          # Custom middleware
│   │   ├── routes/              # Route definitions
│   │   └── server.go            # Server setup
│   ├── database/
│   │   ├── dbmodels/            # Generated SQLC models
│   │   ├── dbqueries/           # SQL query definitions
│   │   ├── dbschemas/           # Database schemas
│   │   ├── dbrepos/             # Repository pattern
│   │   └── client.go            # Database connection
│   ├── config/                  # Configuration management
│   └── utils/                   # Utility functions
├── views/                       # Templ templates
├── web/static/                  # Static assets (JS, CSS)
├── local-env/                   # Docker development environment
└── build.sh                     # Build script
```

## 🔧 Development

### Database Changes
1. Modify schemas in `internal/app/database/dbschemas/`
2. Add queries in `internal/app/database/dbqueries/`
3. Run `sqlc generate` to regenerate models
4. Update repository methods in `internal/app/database/dbrepos/`

### Template Changes
1. Modify `.templ` files in `views/`
2. Run `templ generate` to regenerate Go code
3. Rebuild and restart server

### Adding New Endpoints
1. Create handler in `internal/app/api/apihandlers/`
2. Add routes in `internal/app/api/routes/routes.go`
3. Register with appropriate HTTP method helper

## 🎨 AI Assistance Note

*Due to tight development timelines, AI was extensively used to generate the majority of the frontend JavaScript code, HTML templates, and CSS styling. The core backend architecture, database design, and business logic were manually designed and implemented. Some frontend refinements and integrations were also manually adjusted to ensure proper functionality and user experience.*

## 📝 Environment Configuration

Copy `.env.example` to `.env` and configure:
```bash
APP_PORT=8080
DATABASE_HOST=localhost
DATABASE_PORT=4306
DATABASE_NAME=default
DATABASE_USER=dbuser
DATABASE_PASS=dbuserpassword
```

## 🧪 Testing

```bash
# Run tests (if implemented)
go test ./...

# Test specific package
go test ./internal/app/database/dbrepos
```

## 🐳 Docker Support

The project includes Docker configuration for local development:

```bash
# Start all services
docker-compose up -d

# Stop services
docker-compose down
```

## 📄 License

[Add your license information here]

## 🤝 Contributing

[Add contribution guidelines here]
