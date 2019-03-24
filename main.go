package main

import (
	"log"
	"net/http"

	"./auth"
)

func main() {

	/* ルートへのアクセスを処理する。認証状態によって遷移先を変更 */
	http.HandleFunc("/", moveHandler)

	/* チャットルームの生成と開始 */
	chatroom := newRoom()
	go chatroom.run()

	/* 認証モデルの生成 */
	auth.SetAuthInfo()

	/* メッセージを処理するハンドラ */
	http.Handle("/room", chatroom)

	// 認証プロバイダ選択ページへの遷移を行うハンドラ
	http.Handle("/login", &templateHandler{filename: "/login.html"})

	// 選択したプロバイダによる認証を行うハンドラ
	http.Handle("/auth/", &auth.AuthHandler{Path: "/chat.html"})

	// 認証情報を削除し再度認証ページへアクセスするハンドラ
	http.HandleFunc("/logout", logoutHandler)

	// チャットページへのハンドラ
	http.Handle("/chat", &templateHandler{filename: "/chat.html"})

	/* webサーバを開始する */
	log.Println("webサーバを開始します。")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("webサーバの起動に失敗しました。:", err)
	}
}
