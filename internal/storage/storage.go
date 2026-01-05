package storage

import (
	"bufio"
	"os"
	"strings"
	"todo/internal/tasks"
)

func LoadTasksFromFile() {

	file, err := os.Open("data/tasks.txt") // пытаемся открыть файл "tasks.txt" для чтения

	if err != nil { // если файла нет или произошла ошибка при открытии — просто выходим из функции
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
			continue
		}

		done := parts[0] == "1"                        // превратили "1"/"0" в bool ("1" → true, "0" → false)
		task := tasks.Task{Text: parts[1], Done: done} // создаём новую задачу Task с текстом и статусом
		tasks.List = append(tasks.List, task)          // добавляем эту задачу в глобальный список List

	}
}

func SaveTasksToFile() {
	file, err := os.Create("data/tasks.txt") // создаём новый файл tasks.txt или перезаписываем старый

	if err != nil {
		return // если ошибка — просто выходим
	}
	defer file.Close() // гарантируем, что файл закроется в конце функции

	for _, task := range tasks.List {
		status := "0" // по умолчанию считаем задачу невыполненной
		if task.Done {
			status = "1" //если задача выполнена, ставим "1"
		}
		line := status + "|" + task.Text // формируем строку для файла
		file.WriteString(line + "\n")    // записываем строку в файл с переносом на новую строку
	}
}
