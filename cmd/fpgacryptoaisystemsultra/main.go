// cmd/fpgacryptoaisystemsultra/main.go
package main

import (
"flag"
"log"
"os"

"fpgacryptoaisystemsultra/internal/fpgacryptoaisystemsultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := fpgacryptoaisystemsultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
