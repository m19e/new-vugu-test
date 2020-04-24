# Novel Formatter

This is [Vugu](https://www.vugu.org/) App. Format text for Novel following rules.

## Formatting Rules

### 1. Insert a full-width space at the beginning of the line. However, the lines that match the following are excluded.

-   Line that begins with an angle bracket (「,『).
-   Line that already have full-width spaces inserted.
-   Blank line.

### 2. Insert a full-width space after the full-width exclamation mark (!) And question mark (?). However, except when it matches with the following.

-   When an exclamation mark (!), question mark (?), or angle bracket (「,『) immediately follows the mark.
-   When full-width space is already inserted.

---

You can get started with:

```sh
git clone https://github.com/m19e/novel-formatter.git
go get -u github.com/vugu/vgrun
cd novel-formatter
vgrun devserver.go
```

Then browse to the running server: http://localhost:8844/
