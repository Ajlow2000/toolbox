outpath := "./bin/toolbox"

run: prebuild
    @go run main.go

build: prebuild
    @go build -o {{ outpath }} main.go
    @echo -e "Final binary built: {{ outpath }}"

prebuild:
    @echo -e "Running prebuild steps\n"
