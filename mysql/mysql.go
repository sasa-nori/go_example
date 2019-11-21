package mysql

import (
    "fmt"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

func main() {
    // db接続
    db, err := sqlConnect()
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    result := []*Users{}
    db.Find(&result)
    for _, user := range result {
        fmt.Println(string(user.ID) + "_" + user.Name)
    }
    fmt.Println("Delete")
    // Deleteに構造体の配列をいれる
    error := db.Where("id = ?" , 1).Delete(Users{}).Error
    //// DELETE from users where id=1
    if error != nil {
        fmt.Println(error)
    }

    result = []*Users{}
    db.Find(&result)
    for _, user := range result {
        fmt.Println(string(user.ID) + "_" + user.Name)
    }
}

func update(db *gorm.DB, id int, name string) {
    result := []*Users{}
    db.Find(&result)
    for _, user := range result {
        fmt.Println(string(user.ID) + "_" + user.Name)
    }
    fmt.Println("update")
    // Modelに構造体の配列をいれる
    error := db.Model(Users{}).Where("id = ?", id).Update(&Users{
        Name:     name,
        UpdateAt: getDate(),
    }).Error
    //// UPDATE users SET name='ゴン太', update_at={現在日時};

    if error != nil {
        fmt.Println(error)
    }

    result = []*Users{}
    db.Find(&result)
    for _, user := range result {
        fmt.Println(string(user.ID) + "_" + user.Name)
    }
}

func findLike(db *gorm.DB, keyword string) {
    result := []*Users{}
    error := db.Where("name LIKE ?", "%"+keyword+"%").Find(&result).Error
    if error != nil || len(result) == 0 {
        return
    }
    for _, user := range result {
        fmt.Println(user.Name)
    }
}

func addUserData(db *gorm.DB) {
    error := db.Create(&Users{
        Name:     "テスト太郎",
        Age:      18,
        Address:  "東京都千代田区",
        UpdateAt: getDate(),
    }).Error
    if error != nil {
        fmt.Println(error)
    }
}

func getDate() string {
    const layout = "2006-01-02 15:04:05"
    now := time.Now()
    return now.Format(layout)
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
    DBMS := "mysql"
    USER := "go_example"
    PASS := "12345!"
    PROTOCOL := "tcp(127.0.0.1:3306)"
    DBNAME := "go_example"

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
    return gorm.Open(DBMS, CONNECT)
}

// Users ユーザー情報のテーブル情報
type Users struct {
    ID       int
    Name     string `json:"name"`
    Age      int    `json:"age"`
    Address  string `json:"address"`
    UpdateAt string `json:"updateAt" sql:"not null;type:date"`
}
