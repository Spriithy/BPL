import bpl.Io // Foo bar baz
import bpl.Sys

static main (string[] args) {

    if args[0] == '\n' {
        Sys.exit(1)
    }

    const x = 1000.3
    string str = "Foo\t"

    const report = Sys::exit

    report(1)
}