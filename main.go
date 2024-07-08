package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

)

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
				err := os.Rename(path, newPath)
				if err != nil {
					fmt.Println("Error renaming file:", err)
				} else {
					fmt.Printf("Renamed: %s -> %s\n", path, newPath)
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
}
