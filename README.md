# Gopare
Gopare is a golang application to find duplicated files on specified directory

## How it works?

* Gopare get all paths to each file in certain directory
* Gopare compare each file using predefined size of **chunks** to avoid load the whole file on memory
* Gopare use goroutines feature to compare multiple files at the same time.

## How to use?

```
git clone https://github.com/Shinji-Mimura/gopare.git
sudo su # gopare must be run as sudo
./gopare <DIRECTORY> <THREADS NUMBER>
```
