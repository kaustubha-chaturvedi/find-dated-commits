package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func isGitRepo(path string) bool {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	return err == nil
}

func runGitLog(repo, searchDate string, out *os.File) {
	fmt.Fprintf(out, "📂 %s\n", repo)

	cmd := exec.Command("git", "-C", repo, "log",
		"--since="+searchDate+" 00:00",
		"--until="+searchDate+" 23:59",
		"--pretty=  %h %an %ad %s",
		"--date=format:%d/%m/%Y",
	)
	cmd.Stdout = out
	cmd.Stderr = out
	_ = cmd.Run()
	fmt.Fprintln(out)
}

func main() {
	recursive := flag.Bool("r", false, "Search recursively in all repos")
	outFile := flag.String("f", "", "Output filename (default: commits-<date>.txt)")
	flag.Parse()

	args := flag.Args()
	dir := "."
	dateStr := time.Now().Format("02/01/2006") 

	if len(args) > 0 {
		dir = args[0]
	}
	if len(args) > 1 {
		dateStr = args[1]
	}

	
	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		fmt.Fprintln(os.Stderr, "Invalid date format, use dd/mm/yyyy")
		os.Exit(1)
	}
	searchDate := fmt.Sprintf("%s-%s-%s", parts[2], parts[1], parts[0])

	
	if *outFile == "" {
		*outFile = fmt.Sprintf("commits-%s-%s-%s.txt", parts[0], parts[1], parts[2])
	}

	
	foundRepo := false
	if *recursive {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err == nil && info.IsDir() && info.Name() == ".git" {
				foundRepo = true
				return filepath.SkipDir
			}
			return nil
		})
	} else {
		if isGitRepo(dir) {
			foundRepo = true
		}
	}

	if !foundRepo {
		fmt.Fprintln(os.Stderr, "Error: no git repository found in the given path")
		os.Exit(1)
	}

	
	out, err := os.Create(*outFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	if *recursive {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err == nil && info.IsDir() && info.Name() == ".git" {
				runGitLog(filepath.Dir(path), searchDate, out)
				return filepath.SkipDir
			}
			return nil
		})
	} else {
		runGitLog(dir, searchDate, out)
	}
}
