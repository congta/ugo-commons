**U**nprofessional **GO** styles, looks like Apache **COMMONS** of Java

Method Name Conventions:
* `Xxx` - normal method, returns result and error;
* `XxxQuitely` - returns void, print log if error;
* `XxxTry` - returns result or nil, print log if error;
* `XxxWildly` - returns result or nil or void, panic if error.
* `Xxx0 / Xxx1` - override methods

This is a new project, Let me known if you need more utils: `zhfchdev@gmail.com`

---

## commons-u

All basic tools.

SystemUtils:

    isWin := ucommons.IsWindows()

Handle arguments:

    map := ucommons.getArgsMap(args)

This is the default implementation, args should be `-f 1 --foo 2 --new`, args before the first `-` will be ignored, value cannot start with `-`.

Get system home dir:

    home := ucommons.Home()

## commons-lang

### ustrings - StringUtils

### unumbers - NumberUtils / Integer / ...

* `ParseXxx` - try to parse string to value, throw error if it's not a number
* `ToXxx` - convert string to value, return default value if it's not a number

## commons-io

### ufiles - FileUtils

    lines, err := ufiles.ReadLines(fileName)
    ufiles.CloseQuitely(file)
    exist := ufiles.Exist(fileName)
    ufiles.WriteLines(fileName, lines)

## commons-codec

include 3 kinds of entries.

### ucodings - for encode and decode

    dst := ucodings.EncodeBase64(src)

### udigests - calculate digest, e.g. md5 or sha

    dst := udigests.Md5(src)

### usecrets - for encrypt and decrypt

    dst, err := usecrets.AesCTREncrypt(src, key, iv)
    dst, err := usecrets.AesCTRDecrypt(src, key, iv)
    dst, err := usecrets.AesCBCEncrypt(src, key, iv)
    dst, err := usecrets.AesCBCDecrypt(src, key, iv)
