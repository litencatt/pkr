# Poker CLI

[![CI](https://github.com/litencatt/pkr/actions/workflows/ci.yml/badge.svg)](https://github.com/litencatt/pkr/actions/workflows/ci.yml)

A command-line poker game implemented in Go.

## Features

- Interactive poker game
- Multiple rounds of play
- Score and multiplier system
- Card selection and actions (Play/Discard/Cancel)

## How to Play

### Starting the Game

```bash
# Normal mode
./pkr run

# Debug mode (shows detailed card information)
./pkr run -d
```

### Game Flow

1. **Game Start**: The game begins with 5 cards dealt to you
2. **Card Selection**: Review your cards and choose which cards to play
3. **Action Selection**: For each card, choose from the following actions:

   - **Play**: Play the card and evaluate it as part of a poker hand
   - **Discard**: Discard the card (no effect on score)
   - **Cancel**: Cancel the selection

4. **Hand Evaluation**: Played cards are evaluated as poker hands and score is added
5. **Next Round**: After the round ends, proceed to the next round

### Poker Hands

The following poker hands are recognized (in order of strength):

- **Royal Flush**: A-K-Q-J-10 of the same suit
- **Straight Flush**: Five consecutive cards of the same suit
- **Four of a Kind**: Four cards of the same rank
- **Full House**: Three of a kind plus a pair
- **Flush**: Five cards of the same suit
- **Straight**: Five consecutive cards (any suit)
- **Three of a Kind**: Three cards of the same rank
- **Two Pair**: Two pairs of cards
- **One Pair**: Two cards of the same rank
- **High Card**: Any other combination

### Scoring

- Each poker hand has a unique score value
- Stronger hands yield higher scores
- The multiplier system provides bonus scores for consecutive hands

### Game End

- Play the set number of rounds or manually end the game
- Final score is displayed when the game ends

## Development Environment

This project supports a development environment using Docker Compose.

### Required Software

- Docker
- Docker Compose
- Make

### Setup

```bash
# Start development environment
docker compose up -d

# Build application
docker compose exec app make build

# Run tests
docker compose exec app go test ./...

# Run linter
docker compose exec app golangci-lint run

# Run application
docker compose exec app ./pkr
```

## CI/CD

The following automation is performed using GitHub Actions:

- **Testing**: Unit test execution and coverage measurement
- **Linting**: Code quality checks with golangci-lint
- **Security Scanning**: Vulnerability detection with gosec
- **Build**: Application build verification

## Project Structure

```
.
├── cmd/pkr/          # Main application
├── entity/           # Domain entities
├── service/          # Business logic
├── .github/workflows/ # CI/CD configuration
├── docker-compose.yml # Development environment configuration
└── Makefile         # Build tasks
```

## License

MIT License
