


---



# Mint Package

- Mint Package is a lightweight Go library designed for file compression, encryption, and intelligent categorisation. It allows you to compress and encrypt files by type, ensuring secure storage and easy retrieval. Mint Package is implemented entirely in pure Go, without any C dependencies, and supports integration with other languages via a shared library (.so) if needed.

# Features

- File type classification: Automatically splits files into categories:

- SOUND for audio files (`mp3`, .`wav`, .`ogg`)

- DOCUMENT for documents (.`txt`, .`pdf`, .`docx`)

- DATA for `.data` files

- FILE for unknown types

-------

- Compression: Uses flate for efficient data compression

- Encryption: AES-GCM encryption for secure storage

- Pure Go: No external dependencies

- Optional .so build: Can be compiled as a shared library for use in other languages



---

# Installation

Ensure you have Go 1.21 or later installed on your system. You can download Go from the official source:

https://go.dev/dl/

Clone the repository to your local machine:

```
git clone https://github.com/CoolyDucks/mintpackage.git
cd mintpackage
```

Initialize Go modules:

```
go mod tidy
```

---

# Usage

Importing Mint Package

Mint Package is imported as a module:

import "mintpackage"

Example: Compressing and Encrypting a File

```Go Script 
package main

import (
	"fmt"
	"mintpackage"
)

func main() {
	key := []byte("12345678901234567890123456789012")
	content := []byte("This is an example content for text.txt")

	section, err := mintpackage.EncryptFile("text.txt", content, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("Encrypted length:", len(section.Data))

	decrypted, err := mintpackage.DecryptSection(section, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("Decrypted content:", string(decrypted))
}
```

---

# Building the Shared Library (.so)

Mint Package can be compiled as a shared library to be used in other languages such as C or Python:
```
go build -buildmode=c-shared -o libmintpackage-1.0.so libmintpackage.go
```
This generates two files:

libmintpackage-1.0.so: the shared library

libmintpackage-1.0.h: the header file



---

# How Mint Package Works

1. File Classification
When a file is passed to Mint Package, it is classified based on its extension. This allows each file to be handled according to its type (SOUND, DOCUMENT, DATA, or FILE).


2. Compression
The file data is compressed using Go's flate implementation. This reduces file size while maintaining speed and compatibility.


3. Encryption
The compressed data is encrypted using AES-GCM, which provides both confidentiality and integrity. Each encrypted section contains a unique nonce for security.


4. Decryption
When decrypting, Mint Package reverses the encryption and decompression, restoring the original file content exactly as it was.




---

File Flow Example

1. Original file: text.txt


2. Classified as DOCUMENT


3. Compressed to reduce size


4. Encrypted with AES-GCM


5. Stored as a Mint Package section (can be saved to disk as .mint)


6. Decryption and decompression restores the original file




---

# Testing Mint Package

To test Mint Package locally:

1. Create a Go file test.go in the project root.


2. Use the example code provided above.


3. Run:



go run test.go

4. Verify that encryption and decryption work as expected.




---

# Notes

Ensure your AES key is 32 bytes for security.

Mint Package can handle multiple file types in the same program.

For very large files, consider reading and processing data in chunks to reduce memory usage.



---

# License

Mint Package is open source and licensed under the "BSD 3 Clause"




---
