package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func replace_and_or(commend string) string {
	end_replace := strings.ReplaceAll(commend, "; and ", "&&")
	or_replace := strings.ReplaceAll(end_replace, "; or ", "||")
	return or_replace
}

func check_fish_history_file() string {
	user_home_path, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	path_fish_history := user_home_path + "/.local/share/fish/fish_history"
	if _, err := os.Stat(path_fish_history); err != nil {
		log.Fatalf("File %s does not exist", path_fish_history)
	}
	return path_fish_history
}

func linesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func convert_history(file_data []string) {
	user_home_path, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	path_zsh_history := user_home_path + "/.zsh_history.convert"
	zsh_history_file, err := os.Create(path_zsh_history)
	if err != nil {
		log.Fatal(err)
	}
	for index, line := range file_data {
		if strings.HasPrefix(line, "- cmd:") {
			command := strings.Join(strings.Split(line, "- cmd:"), " ")
			if index+1 <= len(file_data)-1 {
				when := strings.Join(strings.Split(file_data[index+1], "when: "), " ")
				fmt.Fprintf(zsh_history_file, ": %s:0;%s\n", when, command)
			}
		}
	}
}

func main() {
	fish_history_path := check_fish_history_file()
	file_array := linesInFile(fish_history_path)
	convert_history(file_array)
}
