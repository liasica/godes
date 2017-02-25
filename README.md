Golang DES encrypt or decrypt
======

### PHP Javascript Perl version
http://www.tero.co.uk/des/code.php

### Useage
```golang
key := "this is a 24 byte key !!"
msg := "This is a test message"
enret := godes.Des(key, msg, 1, 0, "", 0)
deret := godes.Des(key, enret, 0, 0, "", 0)
```
