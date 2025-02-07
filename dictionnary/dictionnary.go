package dictionnary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

func Load(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		words = append(words, scan.Text())
	}
	if err_s := scan.Err(); err_s != nil {
		return err_s

	}
	return nil
}

func PickWord() string {
	rand.Seed(time.Now().Unix()) // I Know it's deprecated
	i := rand.Intn(len(words))
	return words[i]

}
