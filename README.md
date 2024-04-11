# Tutorial

## install

git clone
go mod tity

## run

go run main.go

OR

go build .
./sdv-m2-renf-backend-go-cobra.exe

you can also name the executable as you wish with

go build -o tournoi-cli.exe
./tournoi-cli.exe

for the following, I will be using ./tournoi-cli.exe

When the command is ran, it will list the commands you can use.
Commands can also contain subcommands (like auth)
So if you run
./tournoi-cli.exe

it will list the subcommands you can use (login, signup)

## To initialize the database

./tournoi-cli init

## problems: go wont accept a package named tournament so I called it tournament\_

# Front: cobra

# Back

## Tutorials followed:

[JWT Authentication in Go (Gin/Gorm)](https://www.youtube.com/watch?v=ma7rUS_vW9M)

[How to write beautiful Golang CLI](https://www.youtube.com/watch?v=SSRIn5DAmyw)

# libs used (no need to install)

## Cobra

go install github.com/spf13/cobra-cli@latest

### promptui

go get github.com/manifoldco/promptui

## Gorm

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite

## Bcrypt

go get -u golang.org/x/crypto/bcrypt

## Jwt

go get -u github.com/golang-jwt/jwt/v5

## Godotenv

go get github.com/joho/godotenv

### jwt token valid for 1 day
