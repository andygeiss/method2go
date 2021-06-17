@echo off
cls
setlocal EnableDelayedExpansion

set NAME={{ .Name }}

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
go get github.com/andygeiss/utilities/logging
go get github.com/andygeiss/utilities/tracing
go build --ldflags "-s -w -X=main.build=%BUILD% -X=main.name=%NAME% -X=main.version=%VERSION%" -o %GOPATH%\bin\%NAME%_api.exe clients\api\main.go
go build --ldflags "-s -w -X=main.build=%BUILD% -X=main.name=%NAME% -X=main.version=%VERSION%" -o %GOPATH%\bin\%NAME%_cli.exe clients\cli\main.go
go build --ldflags "-s -w -X=main.build=%BUILD% -X=main.name=%NAME% -X=main.version=%VERSION%" -o %GOPATH%\bin\%NAME%_web.exe clients\web\main.go

echo cleaning up ...
del build.txt version.txt

echo.
echo done.