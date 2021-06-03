@echo off
cls
setlocal EnableDelayedExpansion

set NAME=method2go

echo getting info from Git ...
git pull
git rev-parse --short HEAD > build.txt
set /p BUILD=<build.txt
git describe --tags > version.txt
set /p VERSION=<version.txt
echo.

echo NAME    = %NAME%
echo BUILD   = %BUILD%
echo VERSION = %VERSION%
echo.

echo start building ...
go build --ldflags "-s -w -X=main.build=%BUILD% -X=main.name=%NAME% -X=main.version=%VERSION%" -o %GOPATH%\bin\%NAME%.exe clients\cli\main.go

echo cleaning up ...
del build.txt version.txt

echo.
echo done.