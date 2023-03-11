package tests

import (
	"app/config"
	"app/internal/services"
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func Test_news(t *testing.T) {
	config.NewConfig()
	//repo.NewDb()
	//news := repo.GetNews()
	//fmt.Println(news)
	fmt.Println(viper.Get("APP_URL"))
	t.Log("ok")
}

func Test_token(t *testing.T) {
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["id"] = 12
	strToken, _ := services.CreateToken(data)
	fmt.Println(strToken)
	t.Log("ok")
}
func Test_parseToken(t *testing.T) {
	strToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImlkIjoxMn0sIm5iZiI6MTY3MjE0MzEzMn0.h8PaL4mlhLUbvbt9_aPbA_7Fm9I1kKTFn3X0xtG1yyI"
	res := services.ParseToken(strToken)
	fmt.Println(res["data"])
}
