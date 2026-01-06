package storage

import (
	"bufio"
	"log"
	"os"
	"strings"
	"todo/internal/tasks"
)

func LoadTasksFromFile() {

	tasks.List = nil // очистить перед загрузкой
	// ... существующий код загрузки

	err := os.MkdirAll("data", 0755) // создаём папку "data", если её нет
	if err != nil {
		log.Println("[ERROR] cannot create data folder")
		return
	}

	file, err := os.OpenFile("data/tasks.txt", os.O_RDWR|os.O_CREATE, 0644) // открываем файл "tasks.txt" для чтения/записи, если файла нет — создаём
	if err != nil {
		log.Println("[ERROR] cannot create or open tasks file")
		return // это нормально, например, при первом запуске программы
	}
	defer file.Close() // гарантируем, что файл закроется автоматически, когда функция завершится

	scanner := bufio.NewScanner(file) // создаём Scanner — объект, который умеет читать файл построчно

	for scanner.Scan() { // пока есть следующая строка в файле, делаем тело цикла, пока действует эта функция в виде условия,
		// она возвращает тру фолс, это стандартный способ читать файл построчно в Go Scan() возвращает true, пока есть строки
		line := scanner.Text()            // получаем текущую строку из файла в виде текста(возвращает текст текущей строки без символов переноса строки)
		parts := strings.Split(line, "|") // разбиваем строку по символу "|" на две части:
		// parts[0] — "1" или "0" (Done)
		// parts[1] — текст задачи
		if len(parts) < 2 { // если строка некорректная — пропускаем
			log.Println("[WARN] corrupted task line skipped:", line)
			continue
		}

		done := parts[0] == "1"                        // превратили "1"/"0" в bool ("1" → true, "0" → false)
		task := tasks.Task{Text: parts[1], Done: done} // создаём новую задачу Task с текстом и статусом
		tasks.List = append(tasks.List, task)          // добавляем эту задачу в глобальный список List

	}

	if err := scanner.Err(); err != nil {
		log.Println("[ERROR] reading tasks file:", err)
	}

	log.Printf("[INFO] Loaded %d tasks from file", len(tasks.List))
}

func SaveTasksToFile() {

	file, err := os.Create("data/tasks.txt") // создаём новый файл tasks.txt или перезаписываем старый

	if err != nil {
		log.Println("[ERROR] failed to create tasks file:", err)
		return // если ошибка — просто выходим
	}
	defer file.Close() // гарантируем, что файл закроется в конце функции

	for _, task := range tasks.List {
		status := "0" // по умолчанию считаем задачу невыполненной
		if task.Done {
			status = "1" //если задача выполнена, ставим "1"
		}
		line := status + "|" + task.Text        // формируем строку для файла
		_, err := file.WriteString(line + "\n") // записываем строку в файл с переносом на новую строку
		if err != nil {
			log.Println("[ERROR] failed to write task to file:", err)
			return
		}
	}
	log.Printf("[INFO] Saved %d tasks to file", len(tasks.List))
}
