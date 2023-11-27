package main

import (
	"os"
	"strings"

	"github.com/ecrmnn/str"
)

type doc struct {
	filename    string
	title       string
	signature   string
	returnValue string
	description string
	examples    []string
}

var signatures = make(map[string]string)
var returnValues = make(map[string]string)
var descriptions = make(map[string]string)
var docs []doc
var currentDoc doc
var readMeTemplates []string
var linkTemplates []string

var previousLineWasActual bool = false

func main() {
	// Delete all files in api folder
	os.RemoveAll("api")

	// Extract comments from str.go and store them in @var descriptions
	strFileContent, err := os.ReadFile("../str.go")

	if err != nil {
		panic(err)
	}

	strFileContentString := string(strFileContent)
	strFileContentSlice := strings.Split(strFileContentString, "\n")

	for _, line := range strFileContentSlice {
		isSignature := strings.Contains(line, "func (s *Str)")
		isDescription := strings.Contains(line, "// ")

		if isSignature {
			functionName := str.New(line).Between("func (s *Str) ", "(").String()
			signature := line[:len(line)-2]
			signatures[functionName] = signature
			returnValues[functionName] = str.New(line).AfterLast(")").BeforeLast("{").TrimAll().String()
		}

		if isDescription {
			words := strings.Split(line, " ")
			functionName := words[1]
			functionNameCode := "`" + functionName + "`"
			descriptions[functionName] = functionNameCode + " " + strings.Join(words[2:], " ")
		}
	}

	// Extract tests to build documentation
	testFileContent, err := os.ReadFile("../str_test.go")

	if err != nil {
		panic(err)
	}

	testFileContentString := string(testFileContent)
	testFileContentSlice := strings.Split(testFileContentString, "\n")

	for _, line := range testFileContentSlice {
		isFunctionName := strings.Contains(line, "\t\tname: ")

		if isFunctionName && currentDoc.title != "" {
			docs = append(docs, currentDoc)
			currentDoc = doc{}
		}

		if isFunctionName {
			functionName := getFunctionName(line)
			currentDoc.filename = getFileName(functionName)
			currentDoc.title = functionName
			currentDoc.signature = signatures[functionName]
			currentDoc.returnValue = returnValues[functionName]
			currentDoc.description = descriptions[functionName]
			previousLineWasActual = false
		} else if strings.Contains(line, "\t\t\t\tactual:") {
			currentDoc.examples = append(currentDoc.examples, getTestActual(line))
			previousLineWasActual = true
		} else if strings.Contains(line, "\t\t\t\texpected:") {
			currentDoc.examples = append(currentDoc.examples, getTestExpected(line))
			previousLineWasActual = false
		} else if previousLineWasActual {
			l := str.New(line[4:])

			if l.EndsWith("String(),") {
				l.Replace("(),", "()")
			}

			if l.EndsWith("}(),") {
				l.Replace("}(),", "}()")
			}

			currentDoc.examples = append(currentDoc.examples, l.String())
		}
	}

	// Create api folder
	os.Mkdir("api", 0755)

	docStub, err := os.ReadFile("stubs/api.stub.txt")

	if err != nil {
		panic(err)
	}

	docStubString := string(docStub)

	readMeStub, err := os.ReadFile("stubs/short.stub.txt")

	if err != nil {
		panic(err)
	}

	readMeStubString := string(readMeStub)

	for _, d := range docs {
		m := map[string]string{
			"{{.Title}}":       d.title,
			"{{.ReturnValue}}": d.returnValue,
			"{{.Signature}}":   d.signature,
			"{{.Description}}": d.description,
			"{{.Examples}}":    strings.Join(d.examples, "\n"),
		}

		template := str.New(docStubString).Swap(m).String()

		f, _ := os.Create("api/" + d.filename)
		f.WriteString(template)
		defer f.Close()

		readMeTemplate := str.New(readMeStubString).Swap(m).String()
		readMeTemplates = append(readMeTemplates, readMeTemplate)

		linkLabel := str.New("- [{{.Title}}]").Swap(m).String()
		linkTarget := str.New("(#{{.Title}}-{{.ReturnValue}})").Swap(m).
			Swap(map[string]string{
				"*": "",
				"[": "",
				"]": "",
			}).Lower().String()

		linkTemplates = append(linkTemplates, linkLabel+linkTarget)
	}

	// Update repo README.md
	readme, err := os.ReadFile("../README.md")

	if err != nil {
		panic(err)
	}

	readmeString := string(readme)
	readmeStringWithoutOldApi :=
		str.New(readmeString).
			Before("### API").
			Append("### API\n\n").
			Append(strings.Join(linkTemplates, "\n")).
			Append("\n").
			Append(strings.Join(readMeTemplates, "\n\n")).
			String()

	err = os.Remove("../README.md")

	if err != nil {
		panic(err)
	}

	f, _ := os.Create("../README.md")
	f.WriteString(readmeStringWithoutOldApi)
	defer f.Close()
}

func getFunctionName(line string) string {
	return str.New(line).Between("name: \"", "\"").String()
}

func getFileName(functionName string) string {
	return str.New(functionName).Snake().Append(".md").String()
}

func getTestActual(line string) string {
	return strings.Trim(str.New(line).After("\t\t\t\tactual:").BeforeLast(",").String(), " ")
}

func getTestExpected(line string) string {
	return strings.Trim(str.New(line).After("\t\t\t\texpected:").BeforeLast(",").Prepend("//").Append("\n").String(), " ")
}
