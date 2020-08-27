package main

import (

    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "gopkg.in/go-playground/validator.v9"
)

type Form struct {
    NameSei string `validate:"required"`
    NameMei string `validate:"required,max=5"`
    EMail string `validate:"required,email"`
}

func main() {
    router := gin.Default()

    router.Static("/static", "./static")
    router.LoadHTMLGlob("view/*")

    router.GET("/", func(c *gin.Context){
        c.HTML(http.StatusOK, "input.html",nil)
    })

    router.POST("/", postForm)

    router.Run(":8080")
}

func postForm (c *gin.Context) {
    form := Form{
        NameSei: c.PostForm("name_sei"),
        NameMei: c.PostForm("name_mei"),
        EMail: c.PostForm("email"),
    }
    log.Print("SEI : " , form.NameSei)
    log.Print("MEI : " , form.NameMei)
    log.Print("EMAIL : " , form.EMail)
    
    if ok, errors := form.Validate(); !ok {
        c.HTML(http.StatusOK, "input.html", gin.H{"Form": form,"errors": errors})
    } else {
        c.HTML(http.StatusOK, "confirm.html",gin.H{"Form": form})
    }

}

func (form *Form) Validate() (ok bool, result map[string]string) {
    result = make(map[string]string)
    // 構造体のデータをタグで定義した検証方法でチェック
    // err := validator.New().Struct(*form)
    validate := validator.New()
    err := validate.Struct(*form)
    if err != nil {
        errors := err.(validator.ValidationErrors)
        if len(errors) != 0 {
            for i := range errors {
                // フィールドごとに、検証
                switch errors[i].StructField() {
                case "NameSei":
                    switch errors[i].Tag() {
                    case "required":
                        result["NameSei"] = "お名前（姓）を入力してください。"
                    }
                case "NameMei":
                    switch errors[i].Tag() {
                    case "required":
                        result["NameMei"] = "お名前（名）を入力してください。"
                    case "max":
                        result["NameMei"] = "お名前（名）は5文字以内で入力してください。"
                    }
                case "EMail":
                    switch errors[i].Tag() {
                    case "required":
                        result["EMail"] = "メールアドレスを入力してください。"
                    case "email":
                        result["EMail"] = "メールアドレスを正しく入力してください。"
                    }
                 }
            }
        }
        return false, result
    }
    return true, result
}