package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите главу, которую хотите найти: ")
	reader := bufio.NewReader(os.Stdin)
	textInput, _ := reader.ReadString('\n')
	textInput = strings.TrimSpace(textInput)

	file, err := os.Open("witcher.txt")
	if err != nil {
		log.Fatal("Не удалось открыть файл: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	found := false
	collecting := false

	for scanner.Scan() { // Основной цикл добавления строк, пока нужная нам глава
		line := scanner.Text()
		var lowerLine = strings.ToLower(line)

		if strings.Contains(lowerLine, "глава") && strings.HasPrefix(lowerLine, strings.ToLower(textInput)) {
			found = true
			collecting = true
		}

		if strings.Contains(lowerLine, "глава") && collecting && !strings.HasPrefix(lowerLine, strings.ToLower(textInput)) {
			break
		}

		if collecting { // Собираем текст главы
			lines = append(lines, line)
		}
	}

	if !found {
		log.Fatal("Введенная " + textInput + " не была найдена.")
	}

	if len(lines) == 0 {
		log.Fatal("Нет данных для записи.")
	}

	outName := strings.ReplaceAll(textInput, " ", "_") + ".txt" // Формируем название нового файла с главой пользователя
	outFile, err := os.Create(outName)

	if err != nil {
		log.Fatal("Ошибка при создании файла: ", err)
	}

	defer outFile.Close() // Отложенный вызов
	writer := bufio.NewWriter(outFile)

	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	if err := writer.Flush(); err != nil {
		log.Fatal("Ошибка при сохранении файла: ", err)
	}
}
