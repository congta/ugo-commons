**U**nprofessional **GO** styles, looks like Apache **COMMONS** of Java

Method Name Conventions:
* `Xxx` - normal method, returns result and error;
* `XxxQuitely` - returns void, print log if error;
* `XxxTry` - returns result or nil, print log if error;
* `XxxWildly` - returns result or nil or void, panic if error.
* `Xxx0 / Xxx1` - override methods

This is a new project, Let me known if you need more utils: `zhfchdev@gmail.com`

---

## commons-lang

### ustrings - StringUtils

## commons-io

### ufiles - FileUtils

    lines, err := ufiles.ReadLines(fileName)
    ufiles.CloseQuitely(file)

## commons-codec

### ucodings - for encode and decode

    dst := ucodings.EncodeBase64(src)

### udigests - digest by md5 or sha

    dst := udigests.Md5(src)

### usecrets - for encrypt and decrypt