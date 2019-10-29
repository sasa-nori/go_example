package fcm

import (
    "context"
    "log"

    firebase "firebase.google.com/go"
    "firebase.google.com/go/messaging"
    "google.golang.org/api/option"
)

func sendMessage() {
    ctx := context.Background()
    // Firebase初期化
    client, err := firebaseInit(ctx)
    if err != nil {
        log.Fatal(err)
    }

    message := createMessage("test", "テスト送信", "これはテスト送信です。", "test1")

    // Send a message to the devices subscribed to the provided topic.
    response, err := client.Send(ctx, message)
    if err != nil {
        log.Fatalln(err)
    }
    // Response is a message ID string.
    log.Println("Successfully sent message:", response)
}

func createMessage(topic string, title string, body string, tag string) *messaging.Message {
    // android用の設定初期化
    android := new(messaging.AndroidConfig)
    // 通知優先度設定
    android.Priority = "high"
    // android用の通知設定初期化
    androirNotification := new(messaging.AndroidNotification)
    // チャンネル設定(Android8以降は必須。受信する側の設定と合わせる)
    androirNotification.ChannelID = "channel_1"
    // タグ設定(あってもなくてもいい)
    androirNotification.Tag = tag
    android.Notification = androirNotification
    // 大本の通知設定の初期化
    notification := new(messaging.Notification)
    // タイトル
    notification.Title = title
    // 本文
    notification.Body = body
    // メッセージ構造体の初期化
    message := &messaging.Message{
        // データの設定(通知を出す前にデータだけ受け取りたいときはこっちに設定する)
        Data: map[string]string{
            "title": title,
            "body":  body,
        },
        // Android用の設定
        Android: android,
        // 通知設定
        Notification: notification,
        // 配信先(トピック)
        Topic: topic,
    }
    return message
}

// firebaseInit Firebaseの初期化
func firebaseInit(ctx context.Context) (*messaging.Client, error) {
    // Use a service account
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    client, err := app.Messaging(ctx)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    return client, nil
}
