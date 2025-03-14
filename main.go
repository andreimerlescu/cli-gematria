package main

import (
	"bufio"
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/andreimerlescu/gematria"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

var NonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z]+`)

//go:embed VERSION
var VERSION embed.FS

func main() {
	showVersion := flag.Bool("v", false, "Display the version from the VERSION file")
	flag.Parse()
	if *showVersion {
		data, err := fs.ReadFile(VERSION, "VERSION")
		if err != nil {
			log.Fatalf("Failed to read embedded VERSION file: %v", err)
		}
		fmt.Println(string(data))
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n" + color.New(color.FgYellow).Sprint("Received Ctrl+C, exiting... 👋"))
		os.Exit(0)
	}()

	fmt.Println(color.New(color.FgYellow).Sprintf("Welcome to Gematria Calculator! 🎉"))
	fmt.Println("Type '?' for menu, 'exit' or 'quit' to exit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(color.New(color.FgCyan).Sprint("> "))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(color.New(color.FgRed).Sprint("Error reading input: ", err))
			continue
		}
		input = NonAlphanumericRegex.ReplaceAllString(strings.TrimSpace(input), "")
		if input == "" {
			continue
		}

		if input == "exit" || input == "quit" {
			fmt.Println(color.New(color.FgYellow).Sprint("Exiting... 👋"))
			os.Exit(0)
		} else if input == "?" {
			displayMenu(reader)
		} else {
			result := CalculateGematria(input)
			fmt.Println(color.New(color.FgGreen).Sprintf("🔢 Gematria result for '%s'\n%s", input, result))
		}
	}
}

func displayMenu(reader *bufio.Reader) {
	fmt.Println(color.New(color.FgBlue).Sprint("Available options:"))

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Option", "Description"})
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Append([]string{"1", "❓ Help"})
	table.Append([]string{"2", "🚪 Close Menu"})
	table.Append([]string{"3", "🚪 Exit App"})
	table.Render()

	fmt.Print(color.New(color.FgCyan).Sprint("Select an option (1-3): "))
	selection, _ := reader.ReadString('\n')
	selection = strings.TrimSpace(selection)

	switch selection {
	case "1":
		showHelp()
	case "2":
		return
	case "3":
		fmt.Println(color.New(color.FgYellow).Sprint("Exiting... 👋"))
		os.Exit(0)
	default:
		fmt.Println(color.New(color.FgRed).Sprint("Invalid option, please select 1-3"))
	}
}

func showHelp() {
	fmt.Println(color.New(color.FgBlue).Sprint("Help information:"))
	fmt.Println("- Type any text to calculate its gematria.")
	fmt.Println("- Type '?' to show this menu.")
	fmt.Println("- Type 'exit' or 'quit' to exit.")
}

func CalculateGematria(text string) string {
	gematrix := gematria.FromString(text)

	var buf bytes.Buffer

	table := tablewriter.NewWriter(&buf)
	table.SetHeader([]string{"Kind", "Value"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.Append([]string{"Simple", "★ " + strconv.FormatUint(gematrix.Simple, 10)})
	table.Append([]string{"Jewish", "✡ " + strconv.FormatUint(gematrix.Jewish, 10)})
	table.Append([]string{"English", "♚ " + strconv.FormatUint(gematrix.English, 10)})
	table.Append([]string{"Majestic", "♛ " + strconv.FormatUint(gematrix.Majestic, 10)})
	table.Append([]string{"Mystery", "♜ " + strconv.FormatUint(gematrix.Mystery, 10)})
	table.Append([]string{"Eights", "∞ " + strconv.FormatUint(gematrix.Eights, 10)})
	table.Render()

	return buf.String()
}
