# Honker-go-starter

## 1. Clone the repository and build the project.
```bash
cd ~
git clone git@github.com:qiuzhanghua/honker-go-starter.git -o gh
cd honker-go-starter
```

# Windows

## 2. copy `honker_ext.dll` to the project root directory.
```powershell
cd ~
git clone git@github.com:russellromney/honker.git -o gh 
cd honker
cargo build --release

cd ~
cd honker-go-starter

cp ~\space\rust\honker\target\release\honker_ext.dll .
```

## 3. Build and run the project.
```powershell
# Install mingw-w64 if you haven't already, and add it to your PATH environment variable.
# gcc is required to build the project with CGO enabled.

# Enable CGO and build the project.
$Env:CGO_ENABLED=1
go build -o honker-go-starter.exe main.go
./honker-go-starter.exe
```