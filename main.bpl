struct Lexer {

    string path
    *Token tokens

    private (
        int at, lno
        byte cur, nxt
        *Token tail
    )

    new (string path) {
        this.path = path

        input = Io.fopen(path, "r").read()
        tokens = null // TODO
        tail = null
    }

    private init () {
        cur, nxt = null, null
        at, lno = 0, 1
    }

    public lex() {
        init()
        while next() { process() }
        tail.next = EOF(-1)
    }

    private process() {
        Io.printf("%-5dc:%c\tn:%c\n", at, cur, nxt)
        switch cur {
        match '\\': break;
        any:
            break;
        }
    }

}