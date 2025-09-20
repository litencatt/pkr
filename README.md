# Poker CLI

[![CI](https://github.com/litencatt/pkr/actions/workflows/ci.yml/badge.svg)](https://github.com/litencatt/pkr/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/litencatt/pkr/branch/main/graph/badge.svg)](https://codecov.io/gh/litencatt/pkr)

Go言語で実装されたコマンドライン版ポーカーゲームです。

## 機能

- インタラクティブなポーカーゲーム
- 複数ラウンドのプレイ
- スコアとマルチプライヤーシステム
- カードの選択とアクション（Play/Discard/Cancel）

## 開発環境

本プロジェクトはDocker Composeを使用した開発環境をサポートしています。

### 必要なソフトウェア

- Docker
- Docker Compose
- Make

### セットアップ

```bash
# 開発環境の起動
docker compose up -d

# アプリケーションのビルド
docker compose exec app make build

# テストの実行
docker compose exec app go test ./...

# リンターの実行
docker compose exec app golangci-lint run

# アプリケーションの実行
docker compose exec app ./pkr
```

## CI/CD

GitHub Actionsを使用して以下の自動化を行っています：

- **テスト**: 単体テストの実行とカバレッジ測定
- **リント**: golangci-lintによるコード品質チェック
- **セキュリティスキャン**: gosecによる脆弱性検査
- **ビルド**: アプリケーションのビルド検証

## プロジェクト構成

```
.
├── cmd/pkr/          # メインアプリケーション
├── entity/           # ドメインエンティティ
├── service/          # ビジネスロジック
├── .github/workflows/ # CI/CD設定
├── docker-compose.yml # 開発環境設定
└── Makefile         # ビルドタスク
```

## ライセンス

MIT License
