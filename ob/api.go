package ob

import (
    "go_example/model"
    "net/http"
    "time"

    "github.com/labstack/echo"
    "github.com/objectbox/objectbox-go/objectbox"
)

// RunObjectBoxApi ObjectBox+echoのAPI実行
func RunObjectBoxApi() {
    e := echo.New()
    e.Router()
    e.POST("/api/add", addFavorite)
    e.POST("/api/find", find)
    e.POST("/api/update", update)
    e.GET("/api/get/all", getAll)
    e.POST("/api/remove", remove)
    e.Logger.Fatal(e.Start(":3022"))
}

func remove(c echo.Context) error {
    name := c.FormValue("name")
    if name == "" {
        return c.JSON(http.StatusNotFound, "param is not found")
    }
    box := getBox()
    // 検索
    list, error := findName(box, name)
    if error != nil {
        return c.JSON(http.StatusNotFound, "item not found")
    }
    // 削除
    for _, item := range list {
        box.Remove(item)
    }
    // 削除結果取得
    list, error = box.GetAll()
    if error != nil {
        return error
    }
    return c.JSON(http.StatusOK, list)
}

func getAll(c echo.Context) error {
    // テーブル呼び出し
    box := getBox()

    var list []*model.Favorite
    list, err := box.GetAll()
    if err != nil {
        return err
    }

    return c.JSON(http.StatusOK, list)
}

func update(c echo.Context) error {
    name := c.FormValue("name")
    description := c.FormValue("description")
    // パラメーターなしの場合はエラーで返す
    if name == "" && description == "" {
        return c.JSON(http.StatusNotFound, "param is not found")
    }
    box := getBox()
    // 検索
    list, error := findName(box, name)
    if error != nil {
        return c.JSON(http.StatusNotFound, "item not found")
    }
    // 更新
    for _, item := range list {
        item.Description = description
        box.Put(item)
    }
    // 更新結果取得
    list, error = findName(box, name)
    if error != nil {
        return error
    }
    return c.JSON(http.StatusOK, list)
}

func find(c echo.Context) error {
    name := c.FormValue("name")
    description := c.FormValue("description")
    keyword := c.FormValue("keyword")
    // パラメーターなしの場合はエラーで返す
    if name == "" && description == "" && keyword == "" {
        return c.JSON(http.StatusNotFound, "param is not found")
    }

    // テーブル呼び出し
    box := getBox()

    // NameでLIKE検索
    if name != "" {
        result, error := findName(box, name)
        if error != nil {
            return error
        }
        return c.JSON(http.StatusOK, result)
    }

    // DescriptionでLIKE検索
    if description != "" {
        result, error := findDescription(box, description)
        if error != nil {
            return error
        }
        return c.JSON(http.StatusOK, result)
    }

    // Name,DescriptionでのLIKE検索
    result, error := findKeyword(box, keyword)
    if error != nil {
        return error
    }

    return c.JSON(http.StatusOK, result)
}

func findName(box *model.FavoriteBox, name string) ([]*model.Favorite, error) {
    return box.Query(model.Favorite_.Name.Contains(name, true)).Find()
}

func findDescription(box *model.FavoriteBox, description string) ([]*model.Favorite, error) {
    return box.Query(model.Favorite_.Description.Contains(description, true)).Find()
}

func findKeyword(box *model.FavoriteBox, keyword string) ([]*model.Favorite, error) {
    return box.Query(objectbox.Any(model.Favorite_.Name.Contains(keyword, true),
        model.Favorite_.Description.Contains(keyword, true))).Find()
}

func addFavorite(c echo.Context) error {
    favorite := new(AddFavoriteRequest)
    // データ受け取り
    if error := c.Bind(favorite); error != nil {
        return error
    }

    // テーブル呼び出し
    box := getBox()

    const dateLayout = "2006/01/02 15:04:05"
    // データ追加
    id, error := box.Put(&model.Favorite{
        Name:        favorite.Name,
        Description: favorite.Description,
        CreatedAt:   time.Now().Format(dateLayout),
    })
    if error != nil {
        return error
    }

    // 追加したデータを取得
    result, error := box.Get(id)
    if error != nil {
        return error
    }

    // 追加したデータをレスポンスに渡す
    return c.JSON(http.StatusOK, result)
}

// AddFavoriteRequest データ追加リクエストのパース用構造体
type AddFavoriteRequest struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}

func getBox() *model.FavoriteBox {
    return model.BoxForFavorite(InitObjectBox())
}

// RemoveAll 全件削除
func removeAll() {
    // テーブル呼び出し
    box := getBox()

    box.RemoveAll()
}
