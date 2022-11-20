# techblog-linter
日本語の文章をlintしてくれるオレオレ設定

## 必要なもの
### 共通
nodejs 18系の実行環境が必要です。  
[asdf-vm](https://asdf-vm.com/)を導入していれば、`.tool-versions`を見て適切なバージョンをインストールしてくれます。

### vim/neovim
LSPの設定と、[mattn/efm-langserver](https://github.com/mattn/efm-langserver)が必要です。

### vscode
[vscode-textlint](https://marketplace.visualstudio.com/items?itemName=taichi.vscode-textlint)が必要です。

## 設定
### textlintもろもろのインストール
```sh
$ npm ci
```

### (vim/neovimのみ)efm-langserverの設定を置く
`~/.config/efm-langserver/config.yaml` を自分で定義するか、[kei-s16が使っている設定](https://raw.githubusercontent.com/kei-s16/dotfiles/master/.config/efm-langserver/config.yaml)を持ってきてください。

## lintの実行
```sh
$ npx textlint {対象ファイルorディレクトリ}
```
