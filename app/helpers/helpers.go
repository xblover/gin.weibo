package helpers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"gin.weibo/config"
)

// Mix 根据laravel-mix 生成静态文件path
func Mix(staticFilePath string) string {
	manifests := make(map[string]string)

	filename := path.Join(config.ProjectConfig.PublicPath, "mix-manifest.json")
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("mix-manifest.json load fail: %v", err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err := dec.Decode(&manifests); err != nil {
		log.Fatalf("mix-manifest.json decode fail: %v", err)
	}

	for k, v := range manifests {
		log.Printf("%#v %#v", k, v)
	}

	result := manifests[staticFilePath]
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
