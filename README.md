# IHazMemes
## Introduction

IHazMemes is a simple Go program created out of frustration with manually renaming meme files and converting video files to GIFs. This tool automates the process, making it easier to manage and format your meme collection.

## Features

- Renames `.webp` files to `.jpeg`.
- Renames `.jpeg` files to `.png` if incorrectly named.
- Converts `.mp4` video files to `.gif` using `ffmpeg`.

## Requirements

- [Go](https://golang.org/dl/) (1.16 or higher)
- [FFmpeg](https://ffmpeg.org/download.html)

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/mishalhossin/IHazMemes.git
    cd IHazMemes
    ```

2. **Build the program for your platform**:

    To build for your current platform:
    ```sh
    go build -o IHazMemes
    ```
## Usage

1. **Place IHazMemes in the directory with your meme files**.

2. **Run the program**:

    For Windows:
    ```sh
    .\IHazMemes.exe
    ```

    For Linux/Mac:
    ```sh
    ./IHazMemes
    ```

3. **The program will automatically rename and convert the files** as specified.

## Supported Platforms

- Windows (amd64, 386, arm, arm64)
- Linux (amd64, 386, arm, arm64)
- macOS (amd64, arm64)

## Contributing

Feel free to fork this repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

This project was created because I was tired of renaming meme files and turning videos into GIFs manually.
