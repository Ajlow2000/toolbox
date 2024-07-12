outpath := "./bin/toolbox"

goFlags := "-ldflags '-linkmode external -extldflags -static' -o {{ outpath }}"

run: prebuild
    @go run main.go

build: prebuild
    @go build {{ goFlags }} main.go
    @echo -e "Final binary built: {{ outpath }}"

prebuild:
    @echo "Running prebuild steps:"
    zig build-lib -femit-bin=./app/misc/printEnviron/main.bin -femit-h=./app/misc/printEnviron/main.h ./app/misc/printEnviron/main.zig
    @echo ""
    @echo "================================================="
    @echo ""
