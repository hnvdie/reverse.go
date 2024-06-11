1. **Compile the Program**: Pastikan untuk memberikan instruksi kepada pengguna untuk menjalankan perintah ini dari direktori proyek yang benar, di mana file `reverse.go` berada.

2. **Run the Program**: Tambahkan beberapa informasi tambahan tentang apa yang harus dilakukan pengguna jika mereka ingin menggunakan opsi `-r` atau `-l`, misalnya, tentang apa yang dilakukan opsi tersebut.

3. **Example**: Tambahkan output yang diharapkan setelah menjalankan perintah `gorev`, sehingga pengguna dapat melihat contoh hasil yang diharapkan.

4. **License**: Pastikan untuk mengganti `[year]` dan `[fullname]` dengan tahun dan nama Anda dalam teks lisensi.

Berikut adalah versi revisi dari README.md Anda:

```markdown
# GoRev - Reverse DNS Lookup Tool

GoRev is a simple command-line tool written in Go for performing reverse DNS lookup on a list of IP addresses specified in a text file. It recursively resolves each IP address to its corresponding domain name and saves the results to an output file.

## Features

- Reverse DNS lookup for multiple IP addresses
- Recursive resolution to find all associated domain names
- Option to print live results to the console
- Saves results to an output file

## Usage

1. **Compile the Program**:

   Before using the tool, compile the Go source code into an executable binary using the following command from the project directory:

   ```bash
   go build -o bin/gorev reverse.go && sudo mv bin/gorev /bin
   ```

2. **Run the Program**:

   Execute the compiled binary `gorev` followed by the name of the input file containing the list of IP addresses. Optionally, use the `-r` flag for recursive lookup and the `-l` flag to print live results.

   ```bash
   $ gorev hosts.txt -r -l
   ```

   Use `gorev -h` for more options.

3. **View the Results**:

   The results will be saved to an output file named `output.txt` in the current directory.

## Example

Suppose you have a file named `hosts.txt` containing a list of IP addresses:

```
192.0.2.1
198.51.100.5
203.0.113.12
```

Running the command:

```bash
$ gorev hosts.txt -r -l
```

will perform a recursive reverse DNS lookup on each IP address, printing live results to the console, and saving the final results to `output.txt`.

Expected output:

```
example.com
example.net
example.org
```

## License

This project is licensed under the [MIT License](LICENSE).
```

