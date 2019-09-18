package firestore

import (
    "context"
    "log"

    "cloud.google.com/go/firestore"
    firebase "firebase.google.com/go"
    "google.golang.org/api/iterator"
    "google.golang.org/api/option"
)

func FirebaseInit(ctx context.Context) (*firestore.Client, error) {
    // Use a service account
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
        return nil, err
    }

    return client, nil
}

func dataAdd() {
    // 初期化
    ctx := context.Background()
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    // データ追加
    _, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
        "first":  "Ada",
        "middle": "Mathison",
        "last":   "Lovelace",
        "born":   1815,
    })
    if err != nil {
        log.Fatalf("Failed adding alovelace: %v", err)
    }

    // 切断
    defer client.Close()
}

func dataSet() {
    // 初期化
    ctx := context.Background()
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    // データ追加
    _, err = client.Collection("users").Doc("user2").Set(ctx, map[string]interface{}{
        "first":  "Ada",
        "middle": "Mathison",
        "last":   "Lovelace",
        "born":   1815,
    })
    if err != nil {
        log.Fatalf("Failed adding alovelace: %v", err)
    }

    // 切断
    defer client.Close()
}

func dataRead() {
    // 初期化
    ctx := context.Background()
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    // データ読み取り
    iter := client.Collection("users").Documents(ctx)
    for {
        doc, err := iter.Next()
        if err == iterator.Done {
            break
        }
        if err != nil {
            log.Fatalf("Failed to iterate: %v", err)
        }

        log.Println(doc.Data())
    }

    // 切断
    defer client.Close()
}
func dataUpdate() {
    // 初期化
    ctx := context.Background()
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    // データ更新
    _, updateError := client.Collection("users").Doc("user2").Set(ctx, map[string]interface{}{
        "first": "Yeah",
    }, firestore.MergeAll)
    if updateError != nil {
        // Handle any errors in an appropriate way, such as returning them.
        log.Printf("An error has occurred: %s", err)
    }

    // 切断
    defer client.Close()
}

func filedDelete() {
    // 初期化
    ctx := context.Background()
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    // フィールド削除
    _, errorDelete := client.Collection("users").Doc("user2").Update(ctx, []firestore.Update{
        {
            Path:  "middle",
            Value: firestore.Delete,
        },
    })
    if errorDelete != nil {
        // Handle any errors in an appropriate way, such as returning them.
        log.Printf("An error has occurred: %s", err)
    }

    // 切断
    defer client.Close()
}

func docDelete() {
    // 初期化
    ctx := context.Background()
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }

    // フィールド削除
    _, errorDelete := client.Collection("users").Doc("uesr2").Delete(ctx)
    if errorDelete != nil {
        // Handle any errors in an appropriate way, such as returning them.
        log.Printf("An error has occurred: %s", err)
    }

    // 切断
    defer client.Close()
}

func deleteCollection() error {
    ctx := context.Background()
    sa := option.WithCredentialsFile("path/to/serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }
    ref := client.Collection("users")

    for {
        // Get a batch of documents
        iter := ref.Limit(10).Documents(ctx)
        numDeleted := 0

        // Iterate through the documents, adding
        // a delete operation for each one to a
        // WriteBatch.
        batch := client.Batch()
        for {
            doc, err := iter.Next()
            if err == iterator.Done {
                break
            }
            if err != nil {
                return err
            }

            batch.Delete(doc.Ref)
            numDeleted++
        }

        // If there are no documents to delete,
        // the process is over.
        if numDeleted == 0 {
            return nil
        }

        _, err := batch.Commit(ctx)
        if err != nil {
            return err
        }
    }
}
