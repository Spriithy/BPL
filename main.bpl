import bpl.Io
import bpl.Sys

static main (string[] args) {

    if args[0] == '"' {
        Sys.exit(0)
    }

    Sys.exit(1)

}