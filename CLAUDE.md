# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.
すべて回答は日本語で行って。

## Project Overview

This is a CLI poker game application written in Go, using the Cobra framework for CLI commands and the Survey library for interactive prompts.

## Architecture

### Main Components

1. **CLI Layer** (`poker_cli.go`, `cmd/`): Interactive terminal UI using Survey prompts
2. **Service Layer** (`service/poker_service.go`): Game logic and state management
3. **Entity Layer** (`entity/`): Core domain models (Deck, Trump cards, Poker hands, Rounds)

### Key Design Patterns

- Service interface pattern for game logic abstraction
- Entity-based domain modeling for poker game elements
- Cobra command structure for CLI extensibility

## Development Environment

### Docker Compose Setup

基本的にすべてのコード実行はDocker Composeで起動したコンテナ内で行います。

```bash
# コンテナの起動
docker compose up -d

# コンテナ内でシェルを起動
docker compose exec app bash

# コンテナの停止
docker compose down
```

## Development Commands

**Note: 以下のコマンドはすべてDockerコンテナ内で実行してください。**

### Build

```bash
# コンテナ内で実行
make build
# or directly:
go build -ldflags="$(BUILD_LDFLAGS)" -o ./pkr cmd/pkr/main.go
```

### Run the game

```bash
# コンテナ内で実行
./pkr run        # Normal mode
./pkr run -d     # Debug mode (shows detailed card information)
```

### Code Quality

```bash
# コンテナ内で実行
go fmt ./...     # Format code
go vet ./...     # Run static analysis
go mod tidy      # Clean up dependencies
```

### Release Preparation

```bash
# コンテナ内で実行
make deps        # Install required tools (cobra-cli, ghch, gocredits)
make prerelease  # Prepare release (update changelog, credits)
make release     # Create and push git tag, run goreleaser
```

## Key Implementation Details

- Game state is managed through `PokerRound` which tracks hands, discards, score requirements
- Card deck uses standard 52-card Trump entities with suits and ranks
- Poker hand evaluation supports standard hands from High Card to Royal Flush
- Interactive UI uses Survey library for card selection and action prompts
- Terminal clearing is OS-specific (Linux/Darwin supported)

## Agentic Coding Development Guidelines

### Git Commit Strategy
- 変更毎に意味のあるコミットを作成する
- 各機能追加や改善は独立したコミットとして管理
- コミットメッセージは変更内容を明確に記述
- 段階的な開発でも各ステップをコミットして進捗を追跡可能にする

### Development Process
- 大きな機能は小さな単位に分割して実装
- 各段階で動作確認とコミットを実行
- リファクタリングと機能追加は分離してコミット
- エラーや問題が発生した場合は前のコミットに戻りやすくする

### Code Quality Assurance
- **コード修正完了後は必ずテストを実行する**
  ```bash
  # すべてのテストを実行
  docker compose exec app go test ./...
  
  # 詳細なテスト結果を確認
  docker compose exec app go test -v ./...
  ```
- **ビルドが成功することを確認する**
  ```bash
  docker compose exec app make build
  ```
- **コード品質チェックを実行する**
  ```bash
  docker compose exec app go fmt ./...  # フォーマット
  docker compose exec app go vet ./...  # 静的解析
  ```
- テスト失敗やビルドエラーが発生した場合は修正してからコミット
- フォーマット変更があった場合は追加のコミットとして記録

## Important Notes

- Docker Compose環境（Go 1.22-bullseye）を使用してすべての開発作業を行う
- Entity層とService層にテストファイルが存在（*_test.go）
- Debug mode controlled via `-d` flag shows card details during gameplay
- Config file support via Viper (default: `$HOME/.pkr.yaml`)
- Version information embedded at build time via ldflags
