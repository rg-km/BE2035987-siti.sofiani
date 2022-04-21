<<<<<<< HEAD
# BE2035987-siti.sofiani
=======
# Ruangguru Playground

[![GitHub Super-Linter](https://github.com/ruang-guru/playground/workflows/Lint%20Code%20Base/badge.svg)](https://github.com/marketplace/actions/super-linter)

![banner](banner.png)

---

## To generate repo without answer

- `go run cli/main.go answer remove -r backend` or `go run cli/main.go answer remove -r frontend`
- `rsync -vhra . ../playground --include='**.gitignore' --exclude='/.git' --filter=':- .gitignore' --delete-after`
- `cd ../playground`
- `git add .`
- `git commit -nm "(sync)"`
- `git push`
<<<<<<< HEAD
>>>>>>> 3180bbb8cad16953700c234b831b3f2e1f414397
=======

## Adding or Modifying Assignment
- Create or modify your assignment on your folder
- Update file [assignments.json](./assignments.json)
  - Assignment format should be:
    ```json
    {
        "course": "<your course>",
        "path": "path/to/your/assignment/directory",
        "weight": 50.0
    }
    ```
- Create pull request
>>>>>>> 841cc8058b2af3b230b3f3261ca1d9451b8e55f0
