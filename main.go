package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	log.Printf("currenttime-resource")

	err := run(os.Args, time.Now().Round(time.Second).UTC(), os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}

func run(args []string, now time.Time, stdout io.Writer) error {
	switch filepath.Base(args[0]) {
	case "check":
		return check(now, stdout)
	case "in", "get":
		return in(args, now, stdout)
	case "out", "put":
		return out(now, stdout)
	default:
		return in(args, now, stdout)
	}
}

func check(now time.Time, stdout io.Writer) error {
	fmt.Fprintf(stdout, `[{"time":"%s"}]`+"\n", now.Format(time.RFC3339))

	return nil
}

func in(args []string, now time.Time, stdout io.Writer) error {
	if len(args) < 2 {
		args = append(args, ".")
	}

	err := writeCurrentTime(now, args[1])
	if err != nil {
		return err
	}

	fmt.Fprintf(stdout, `{"version":{"time":"%s"}}`+"\n", now.Format(time.RFC3339))

	return nil
}

func out(now time.Time, stdout io.Writer) error {
	fmt.Fprintf(stdout, `{"version":{"time":"%s"}}`+"\n", now.Format(time.RFC3339))

	return nil
}

func writeCurrentTime(now time.Time, directory string) error {
	if _, err := os.Stat(directory); err != nil {
		return fmt.Errorf("%s: %w", directory, err)
	}

	for filename, value := range map[string]string{
		"time":      now.Format(time.RFC3339),
		"rfc3339":   now.Format(time.RFC3339),
		"unixmilli": fmt.Sprintf("%d", now.UnixMilli()),
	} {
		fullFilename := filepath.Join(directory, filename)

		if err := os.WriteFile(fullFilename, []byte(value+"\n"), 0o644); err != nil {
			return fmt.Errorf("%s: %w", fullFilename, err)
		}
	}

	return nil
}
