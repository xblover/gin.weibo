package helpers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path"
	"time"

	"gin.weibo/config"
)

var (
	//存储 mix-manifest.josn 解析出来的path map
	manifests = make(map[string]string)
)

// Mix 根据laravel-mix 生成静态文件path
func Mix(staticFilePath string) string {
	result := manifests[staticFilePath]

	if result == "" {
		filename := path.Join(config.ProjectConfig.PublicPath, "mix-manifest.json")
		file, err := os.Open(filename)
		if err != nil {
			return staticFilePath
		}
		defer file.Close()

		dec := json.NewDecoder(file)
		if err := dec.Decode(&manifests); err != nil {
			return staticFilePath
		}
	}

	if result == "" {
		return staticFilePath
	}

	return "/" + config.ProjectConfig.PublicPath + result
}

// FormatAsDate 格式化日期
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d/%02d/%02d", year, month, day)
}

// csrf input
func CsrfField() template.HTML {
	return template.HTML(fmt.Sprintf(`<input type="hidden" name="_token" value="%s">`, "asd"))
}

func HasSession(key string) bool {

}

func GetSession(key string) string {

}

func Old(key string, defaultValue string) string {

}
