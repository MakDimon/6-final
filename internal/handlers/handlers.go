package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	// Читаем содержимое файла
	fileContent, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "error during file reading", http.StatusInternalServerError)
		return
	}
	// Устанавливаем заголовок Content-Type
	w.Header().Set("Content-Type", "text/html")
	// Записываем содержимое файла в ответ
	w.Write(fileContent)
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1) //1Mb
	// получаем файл из формы
	file, _, err := r.FormFile("myFile") //_handler_ if needed
	if err != nil {
		http.Error(w, "error during parsing script", http.StatusInternalServerError)
		return
	}
	// закрываем файл
	defer file.Close()
	// Чтение файла
	content, err := io.ReadAll(file) //тут сидит содержимое файла
	if err != nil {
		http.Error(w, "file reading error", http.StatusInternalServerError)
		return
	}
	converted := service.Service(string(content)) //перевели из/в морзе
	fn := time.Now().UTC().Add(3 * time.Hour).Format("2006-01-02_15-04-05")
	err = os.WriteFile(fn, []byte(converted), 0755)
	if err != nil {
		http.Error(w, "file writing error", http.StatusInternalServerError)
		return
	}

	// Выводим результат в браузер
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(content))
	s := fmt.Sprint(" результат: ", converted, " имя файла: ", fn)
	w.Write([]byte(s))
}
