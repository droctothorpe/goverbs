# GoVerbs

> The faintest ink is better than the best memory.

I made this to help me remember Golang verbs in the fmt package. Maybe it will help you too. Maybe it won't.

![IDGAF](http://giphygifs.s3.amazonaws.com/media/jrpdsvd6L12ta/giphy.gif)

### Installation
```bash
go get github.com/droctothorpe/goverbs
```

### Usage
```bash
goverbs
```
```
┌───────┬────────────────────────────────────────────┬─────────┬──────────────┐
│ VERB  │ DESCRIPTION                                │ INPUT   │ OUTPUT       │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %v    │ Default format                             │ 15      │ 15           │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %#v   │ Go-Syntax format                           │ 15      │ 15           │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %T    │ The type of value                          │ 15      │ int          │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %d    │ Base 10                                    │ 15      │ 15           │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %+d   │ Always show sign                           │ 15      │ +15          │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %4d   │ Pad with spaces (width 4, right justified) │ 15      │ ␣␣15         │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %-4d  │ Pad with spaces (width 4, left justified)  │ 15      │ 15␣␣         │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %04d  │ Pad with zeroes (width 4)                  │ 15      │ 0015         │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %b    │ Base 2                                     │ 15      │ 1111         │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %o    │ Base 8                                     │ 15      │ 17           │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %x    │ Base 16, lowercase                         │ 15      │ f            │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %X    │ Base 16, uppercase                         │ 15      │ F            │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %#x   │ Base 16, with leading 0x                   │ 15      │ 0xf          │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %c    │ Character                                  │ 65      │ A            │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %q    │ Quoted Character                           │ 65      │ 'A'          │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %U    │ Unicode                                    │ 65      │ U+0041       │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %#U   │ Unicode with character                     │ 65      │ U+0041 'A'   │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %t    │ Bolean                                     │ true    │ true         │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %e    │ Scientific notation                        │ 123.456 │ 1.234560e+02 │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %f    │ Decimal point, no exponent                 │ 123.456 │ 123.456000   │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %.2f  │ Default width, precision 2                 │ 123.456 │ 123.46       │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %8.2f │ Width 8, precision 2                       │ 123.456 │   123.46     │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %s    │ Plain string                               │ cafe    │ cafe         │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %6s   │ Width 6, right justify                     │ cafe    │   cafe       │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %-6s  │ Width 6, left justify                      │ cafe    │ cafe         │
├───────┼────────────────────────────────────────────┼─────────┼──────────────┤
│ %p    │ Pointer                                    │         │ 0xc00023f6d8 │
└───────┴────────────────────────────────────────────┴─────────┴──────────────┘
```