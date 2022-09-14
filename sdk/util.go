package sdk

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func ReadFileToSlice(input string) ([][]string, error) {
	dataSlice := make([][]string, 0)
	if input == "" {
		return nil, errors.New("empty input")
	}
	file, err := os.OpenFile(input, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer func() { _ = file.Close() }()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		arr := strings.Split(sc.Text(), ",")
		dataSlice = append(dataSlice, arr)
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return dataSlice, nil
}

func ReadFileToChan(filename string) (dataChannel chan string, err error) {

	file, err := os.OpenFile(filename, os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	buf := bufio.NewReader(file)
	dataSlice := make([]string, 0)
	for true {
		line, err := buf.ReadString('\n')
		if err != nil {
			dataSlice = append(dataSlice, line)
			break
		} else {
			line = strings.TrimSpace(line)
			dataSlice = append(dataSlice, line)
		}
	}

	dataChannel = make(chan string, len(dataSlice))
	for _, v := range dataSlice {
		dataChannel <- v
	}

	return dataChannel, nil

}

// GetDataFromChan 防止管道数据读完后造成阻塞
func GetDataFromChan(ch chan string) (string, error) {
	select {
	case x := <-ch:
		return x, nil
	default:
		return "0", errors.New("channel has no data")
	}
}
