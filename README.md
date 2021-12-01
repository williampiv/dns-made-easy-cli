# DNS Made Easy CLI

A cli for use against the DNS Made Easy api [https://api-docs.dnsmadeeasy.com/](https://api-docs.dnsmadeeasy.com)

```
user@thing $ dns-made-easy-cli
Easy DNS CLI

>>> domains
Domain: manchego.com
ID:123456

Domain: cheddar.com
ID:654321

>>> records 123456
Name: www
Value: 3.111.222.2
Type: A

Name: mail
Value: 3.111.222.3
Type: A

>>> quit
Bye!
```

## Quickstart

1. Build `make` OR grab your correct binary from one of the release tags
2. Copy `secrets.ini.dist` to `secrets.ini` and fill out with actual API secrets
3. Run `build/dns-made-easy-cli`

NOTE: if you wish to name your secrets file something else, it can be specified at runtime, i.e. `build/dns-made-easy-cli nifty-file.ini`
