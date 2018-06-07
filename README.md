# Astify

[![Build Status](https://travis-ci.org/cristaloleg/astify.svg?branch=master)](https://travis-ci.org/cristaloleg/astify)

### An easy Go AST explorer.

```
a, err := astify.Parse("~/go/src/github.com/cristaloleg/astify")
if err != nil {
    panic(err)
}

a.Walk(func(f *astify.GoFile, n astify.Node) error {
    switch {
    case astify.IsStruct(n):
        strct := astify.AsStruct(n)
        fmt.Printf("# %v : %v \n", strct.Name(), len(strct.Fields()))

    case astify.IsInterface(n):
        iface := astify.AsInterface(n)
        for _, m := range iface.Methods() {
            fmt.Printf("# %v : %v \n", m.Name(), len(m.Params()))
        }

    case astify.IsFunction(n):
        fn := astify.AsFunction(n)
        if fn.IsMethod() && !fn.IsExported() {
            fmt.Printf("$ %v : %v \n", fn.Name(), len(fn.Results()))
        }
    }

    return nil
})
```