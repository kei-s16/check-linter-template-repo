package main

import (
	"context"
	"os"

	"dagger.io/dagger"
	"github.com/joho/godotenv"
)

func main() {
	// projectRoot/dagger/.env を読む
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	linterTargetDir := os.Getenv("LINTER_TARGET_DIR")

	if err := lint(context.Background(), linterTargetDir); err != nil {
		os.Exit(1)
	}
}

func lint(ctx context.Context, targetDir string) error {
	// daggerクライアントの初期化
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout), dagger.WithWorkdir("../"))
	if err != nil {
		return err
	}
	defer client.Close()

	// workdirのパスを取得
	projectRoot := client.Host().Workdir()

	// `node` コンテナイメージを持ってくる
	node := client.Container().From("node:18.11.0-alpine3.15")

	// `node` コンテナの/srcにworkdirをマウントする
	node = node.WithMountedDirectory("/src", projectRoot).WithWorkdir("/src")

	// `node` コンテナ内で必要なnpmパッケージをインストールする
	node = node.
		Exec(dagger.ContainerExecOpts{
			Args: []string{"npm", "ci"},
		})

	// `node` コンテナ内でtextlintを流す
	_, err = node.
		Exec(dagger.ContainerExecOpts{
			Args: []string{"npx", "textlint", targetDir},
		}).
		ExitCode(ctx)

	// NOTE: ExitCodeが常に0を返してくるので、workaroundでerrの有無で判定する
	if err != nil {
		return err
	}

	return nil
}
