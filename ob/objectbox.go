package ob

import (
    "fmt"
    "go_example/model"
    "time"

    "github.com/objectbox/objectbox-go/objectbox"
)

func InitObjectBox() *objectbox.ObjectBox {
    objectBox, _ := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()
    return objectBox
}

// PrintVersion バージョン出力
func PrintVersion() {
    fmt.Println(objectbox.VersionInfo())
}

// AddTask データ追加
func AddTask(text string) uint64 {
    // initialize
    ob := InitObjectBox()

    defer ob.Close()
    // テーブル呼び出し
    box := model.BoxForTask(ob)

    // データ追加
    id, _ := box.Put(&model.Task{
        Text:        text,
        DateCreated: time.Now().Unix(),
    })

    return id
}

// UpdateTask データ更新
func UpdateTask(id uint64, updateText string) {
    // initialize
    ob := InitObjectBox()

    defer ob.Close()

    // テーブル呼び出し
    box := model.BoxForTask(ob)

    task, _ := box.Get(id)
    task.Text = updateText
    box.Put(task)
}

// ReadTaskAll 全件出力
func ReadTaskAll() {
    // initialize
    ob := InitObjectBox()

    defer ob.Close()

    // テーブル呼び出し
    box := model.BoxForTask(ob)

    // データ全件取得
    var list []*model.Task
    list, _ = box.GetAll()
    for _, task := range list {
        fmt.Println(task.Text)
        fmt.Println(task.Id)
        fmt.Println(time.Unix(task.DateCreated, 0))
    }
}

// RemoveTask レコード削除
func RemoveTask(id uint64) {
    // initialize
    ob := InitObjectBox()

    defer ob.Close()

    // テーブル呼び出し
    box := model.BoxForTask(ob)

    // 数を確認
    preCount, _ := box.Count()
    fmt.Println("削除前: " + fmt.Sprint(preCount))

    // IDが1のデータを取得
    task, _ := box.Get(id)
    box.Remove(task)

    // 数を確認
    postCount, _ := box.Count()
    fmt.Println("削除後: " + fmt.Sprint(postCount))
}

// QueryTask 検索
func QueryTask(keyword string) {
    // initialize
    ob := InitObjectBox()

    defer ob.Close()
    // テーブル呼び出し
    box := model.BoxForTask(ob)

    query := box.Query(model.Task_.Text.Equals(keyword, true))
    tasks, _ := query.Find()
    for _, task := range tasks {
        fmt.Println(task.Text)
        fmt.Println(task.Id)
        fmt.Println(time.Unix(task.DateCreated, 0))
    }

}

// RemoveAll 全件削除
func RemoveAll() {
    // initialize
    ob := InitObjectBox()

    defer ob.Close()
    // テーブル呼び出し
    box := model.BoxForTask(ob)

    box.RemoveAll()
}