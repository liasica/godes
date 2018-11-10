Golang DES encrypt or decrypt
======

`⌘` + `V`

Usage
=========================================
`go get -u -v github.com/liasica/godes`

```golang
key := "this is a 24 byte key !!"
msg := "This is a test message"
enret, err := godes.Des(key, msg, 1, 0, "", 0)
if err != nil {
    fmt.Printf("encrypt has error: %v", err)
} else {
    deret, err := godes.Des(key, enret, 0, 0, "", 0)
    if err != nil {
        fmt.Printf("decrypt has error: %v", err)
    }
    ret := ""
    for _, l := range []byte(deret) {
        if l != 0 {
            ret += string([]byte{l})
        }
    }
    fmt.Println(ret)
}
```

How Encryption-decryption affects content
=========================================

```
e0) INPUT as entered (any Unicode character)
      |
  (browser)
      |
e1) input as displayed (any supported Unicode character)
            |
      ----------------------------------
      |                                |
    (C2E)                           (as is)
      |                                |
    text with entities (ASCII)         |
      |                                |
   (filter does not affect)        (filter)
      |                                |
    whole content (**)   'censored' content
      |                                |
      ----------------------------------
            |
e2) filtered input (ASCII)
      |
  (ENCRYPT)
      |
e3) encrypted string (0-255 8 bit)
         |
  (output encoding)
         |
   ------------------------------------
   |           |              |       |
 plain   plain + entities    hex    base64
   |           |              |       |
 8 bit      (ASCII)        0-9 A-F   0-9
 0-255        (*)          (ASCII)   A-Z
  (?)          |              |      a-z
   |           |              |      + /
   |           |              |    (ASCII)
   |           |              |       |
   ------------------------------------
       |
e4) OUTPUT
      ||
d1) input to be decrypted
      |
  (DECRYPT)
      |
d2) the same as in e2) (ASCII)
                |
      ----------------------------------
      |                                |
    whole content        'censored' content
    as in e2)             as in e2)    |
      |                                |
    original text                      |
    with entities (ASCII)           (ASCII)
           |                           |
     --------------                    |
     |            |                    |
  (use as is)   (E2C)             (use as is)
     |            |                    |
```
