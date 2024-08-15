package main

import (
	"bufio"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/image/webp"
)

func convertAndOverwrite(inputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	var img image.Image
	ext := filepath.Ext(inputPath)

	switch ext {
	case ".webp":
		img, err = webp.Decode(file)
		if err != nil {
			return err
		}
	case ".jpeg", ".jpg":
		img, _, err = image.Decode(file)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported file format: %s", ext)
	}

	outFile, err := os.Create(inputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return jpeg.Encode(outFile, img, &jpeg.Options{Quality: 90})
}

func isValidJPEG(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = jpeg.Decode(file)
	if err == nil {
		return true
	}

	file.Seek(0, 0)
	_, err = png.Decode(file)
	return err == nil
}

func renameJPEGToPNG(filePath string) error {
	newPath := filePath[:len(filePath)-len(filepath.Ext(filePath))] + ".png"
	return os.Rename(filePath, newPath)
}

func runAIRenamer() {
	cmd := exec.Command("ai-renamer", ".", "--regex", "v[012]")

	// Create pipes for stdout and stderr
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating stdout pipe: %s\n", err)
		return
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("Error creating stderr pipe: %s\n", err)
		return
	}

	// Start the command
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Error starting command: %s\n", err)
		return
	}

	// Create a scanner to read stdout
	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			fmt.Println("STDOUT:", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading stdout: %s\n", err)
		}
	}()

	// Create a scanner to read stderr
	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println("STDERR:", scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading stderr: %s\n", err)
		}
	}()

	// Wait for the command to complete
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Command finished with error: %s\n", err)
	}
}

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	files, err := os.ReadDir(currentDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			path := filepath.Join(currentDir, file.Name())
			ext := filepath.Ext(path)

			switch ext {
			case ".webp":
				newPath := path[:len(path)-len(ext)] + ".jpeg"
				err := convertAndOverwrite(path)
				if err != nil {
					fmt.Println("Error converting file:", err)
				} else {
					fmt.Printf("Converted: %s -> %s\n", path, newPath)
					err := os.Remove(path)
					if err != nil {
						fmt.Println("Error deleting original file:", err)
					} else {
						fmt.Printf("Deleted original: %s\n", path)
					}
				}

			case ".jpeg", ".jpg":
				if !isValidJPEG(path) {
					err := renameJPEGToPNG(path)
					if err != nil {
						fmt.Println("Error renaming JPEG to PNG:", err)
					} else {
						fmt.Printf("Renamed JPEG to PNG: %s\n", path)
					}
				}

			case ".mp4":
				newPath := path[:len(path)-len(ext)] + ".gif"
				cmd := exec.Command("ffmpeg", "-i", path, "-vf", "fps=10,scale=320:-1:flags=lanczos", "-c:v", "gif", newPath)
				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Printf("Failed to convert %s\n", path)
					fmt.Printf("ffmpeg error: %s\n", string(output))
				} else {
					fmt.Printf("Converted: %s -> %s\n", path, newPath)
					err := os.Remove(path)
					if err != nil {
						fmt.Println("Error deleting original file:", err)
					} else {
						fmt.Printf("Deleted original: %s\n", path)
					}
				}
			default:
				continue
			}
		}
	}

	runAIRenamer()
}
