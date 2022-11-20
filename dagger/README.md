# techblog-linter-settings/dagger
## これはなにか？
[dagger.io](https://dagger.io/)で記述されたCIです。  
GitHub Actionsだけでなく、ローカルでも同じ処理を走らせることができます。

## 実行に必要なもの
- [docker](https://www.docker.com/)
  - 実行ユーザが`docker`グループに追加されている、もしくは`sudo`を実行できる必要があります
- Goの実行環境
  - [asdf-vm](https://asdf-vm.com/)を導入していれば、`.tool-versions`を見て適切なバージョンをインストールしてくれます。

## 実行手順
初回のみ`{REPO}/dagger/.env`内の`LINTER_TARGET_DIR`の値を設定してください。デフォルトは`{REPO}/contents`が指定されています。

```sh
$ go get -v
$ go run main.go
```

## GitHub Actions上での実行
`{REPO}/.github/workflows`にGitHub Actions上でdaggerによるCIを実行する定義を置いています。  
デフォルトの発火条件は`on: pull_request`になっています。必要に応じて`path`を指定するなど、追加の設定を行なってください。  
特に、Privateリポジトリで使用する場合、デフォルトの設定ではActionsの実行時間を消費しすぎてしまう可能性があります。
