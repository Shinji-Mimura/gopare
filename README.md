# Gopare

Gopare is a Go application designed to identify duplicated files within a specified directory.

## Overview

- **Purpose**: Efficiently find duplicate files in a directory.
- **Implementation**: Uses Go's goroutines for concurrent file comparison.
- **Comparison Method**: Compares files in chunks to avoid loading the entire file into memory.

## Features

- **Chunked Comparison**: Instead of loading the entire file into memory, Gopare reads files in chunks of size 64,000 bytes for efficient comparison.
- **Concurrency**: Utilizes Go's goroutines to compare multiple files simultaneously.
- **Progress Bar**: Provides a progress bar to track the comparison process.

## Prerequisites

- Gopare requires sudo privileges to run.

## Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/Shinji-Mimura/gopare.git
   ```

2. Navigate to the cloned directory and run the application with sudo privileges:
   ```bash
   sudo su
   ./gopare <DIRECTORY> <THREADS NUMBER>
   ```

   - `<DIRECTORY>`: The directory where you want to search for duplicated files.
   - `<THREADS NUMBER>`: Number of threads (goroutines) to use for file comparison. A recommended value is 10.

## License

This project is licensed under the terms of the [MIT License](LICENSE).

---

Feel free to use this improved README for your repository or make further modifications as needed.
