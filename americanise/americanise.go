// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var britishAmerican = "C:/Users/VVV/IdeaProjects/ProgInGo/americanise/british-american.txt"

func init() {
	//dir, _ := filepath.Split(os.Args[0])
	//britishAmerican = filepath.Join(dir, britishAmerican)
	//fmt.Println(dir)
	//fmt.Println(britishAmerican)
}

func main() {
	inFilename, outFilename, err := filenamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inFile, outFile := os.Stdin, os.Stdout
	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}

	if outFilename != "" {
		if outFile, err = os.Create(outFilename); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}

	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)
	}

}

func filenamesFromCommandLine() (inFilename, outFilename string, err error) {

	args := os.Args
	l := len(args)
	if (l > 1 && (args[1] == "-h" || args[1] == "--help")) || l == 1  {
		err = fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt", filepath.Base(args[0]))
		return "", "", err
	}

	if l > 1 {
		inFilename = args[1]
		if l > 2 {
			outFilename = args[2]
		}
	}

	if inFilename != "" && inFilename == outFilename {
		log.Fatal("won't overwrite the infile")
	}

	return inFilename, outFilename, nil

}

func americanise(inFile io.Reader, outFile io.Writer) (err error) {

	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	var replacer func(string) string
	if replacer, err = makeReplacerFunction(britishAmerican); err != nil {
		return err
	}

	wordRx := regexp.MustCompile("[A-Za-z]+")
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil  // в действительности признак io.EOF не является ошибкой
			eof = true // это вызовет прекращение цикла в следующей итерации
		} else if err != nil {
			return err // в случае настоящей ошибки выйти немедленно
		}
		line = wordRx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}

func makeReplacerFunction(file string) (func(string) string, error) {
	rawBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	text := string(rawBytes)

	usForBritish := make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}

	return func(word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}
		return word
	}, nil
}
