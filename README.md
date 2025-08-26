Here’s a clean **README.md** you can drop into your Go project 🚀

---

````markdown
# gfc – Git Find Commits

A simple CLI tool to collect all Git commits for a specific day across one or multiple repositories.  
Results are always written to a text file (no terminal output).

---

## Features

- Search commits for a given date (`dd/mm/yyyy`).
- Works in a single repository or recursively through all repos in a folder.
- Automatically writes results to a file:
  - Default: `commits-<dd>-<mm>-<yyyy>.txt`
  - Custom filename with `-f`
- Output includes repository name followed by commit list.

---

## Installation

1. Clone this project:
   ```bash
   git clone https://github.com/kaustubha-chaturvedi/find-dated-commits.git
   cd gfc
````

2. Build the binary:

   ```bash
   go build -o gfc
   ```

3. (Optional) Move it to your `$PATH`:

   ```bash
   mv gfc ~/bin/       # if ~/bin is in PATH
   # or
   sudo mv gfc /usr/local/bin/
   ```

---

## Usage

### Syntax

```bash
gfc [flags] [directory] [date]
```

### Flags

* `-r` : Search recursively in all repositories under the directory.
* `-f <filename>` : Output file name.
  Default: `commits-<dd>-<mm>-<yyyy>.txt`

### Arguments

* `directory` : Path to repository or folder of repositories. Default: `.` (current directory).
* `date` : Date in `dd/mm/yyyy` format. Default: today.

---

## Examples

* **Commits for today in current repo**

  ```bash
  gfc
  ```

  → writes `commits-26-08-2025.txt`

* **Commits for a specific date in current repo**

  ```bash
  gfc . 15/08/2025
  ```

* **Recursive search in all repos under `~/projects`**

  ```bash
  gfc -r ~/projects 15/08/2025
  ```

* **Recursive search with custom filename**

  ```bash
  gfc -r ~/projects 15/08/2025 -f my-commits.txt
  ```

---

## Output Format

Example `commits-26-08-2025.txt`:

```
📂 /home/user/projects/repo1
  a1b2c3d Gaurav 26/08/2025 Fixed bug in API
  d4e5f6g Alice 26/08/2025 Added logging

📂 /home/user/projects/repo2
  h7i8j9k Bob 26/08/2025 Refactored scheduler
```

---

## Notes

* Date must be in **`dd/mm/yyyy`** format.
* Output file will be overwritten if it already exists (use `-f` to set a unique name).
* Works with **Git ≥ 2.x** and tested on Linux/macOS with Go ≥ 1.20.