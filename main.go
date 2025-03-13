package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

type Message struct {
	Org  string `json:"org"`
	Repo string `json:"repo"`
}

func main() {
	for {
		// メッセージを標準入力で読み取り
		message, err := readMessage(os.Stdin)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Fprintf(os.Stderr, "エラー: メッセージの読み取りに失敗: %v\n", err)
			return
		}

		var msg Message
		if err := json.Unmarshal(message, &msg); err != nil {
			fmt.Fprintf(os.Stderr, "エラー: JSONのパースに失敗: %v\n", err)
			continue
		}

		// リポジトリのパスを構築
		repoPath := filepath.Join(os.Getenv("HOME"), "ghq", "github.com", msg.Org, msg.Repo)

		// パスが存在するか確認
		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			sendResponse(fmt.Sprintf("エラー: リポジトリが見つかりません: %s", repoPath))
			continue
		}

		// cursorコマンドを実行
		cmd := exec.Command("/usr/local/bin/cursor", repoPath)
		if err := cmd.Run(); err != nil {
			sendResponse(fmt.Sprintf("エラー: cursorの実行に失敗しました: %v", err))
			continue
		}

		sendResponse("成功: リポジトリを開きました")
	}
}

func readMessage(r io.Reader) ([]byte, error) {
	var length uint32
	if err := binary.Read(r, binary.LittleEndian, &length); err != nil {
		return nil, err
	}
	message := make([]byte, length)
	if _, err := io.ReadFull(r, message); err != nil {
		return nil, err
	}
	return message, nil
}

func sendResponse(message string) {
	response, _ := json.Marshal(map[string]string{"response": message})
	binary.Write(os.Stdout, binary.LittleEndian, uint32(len(response)))
	os.Stdout.Write(response)
}
