# evmkey

> ATTENTION. evmkey is now a simple key util, Please backup your mnemonic phrase yourself

usage:

- `go install -ldflags="-w -s" github.com/crustio/evmkey@latest`
- `evmkey account new`

example output:

```bash
evmkey account new
Enter Password:
Confirm Password:
2024/05/17 11:42:41 Keystore saved:  keystore/0xF81Ed07908875F2bE0e2EB8524dC31e772b58FB8.keystore
2024/05/17 11:42:41 Mnemonic saved:  keystore/0xF81Ed07908875F2bE0e2EB8524dC31e772b58FB8.mnemonic
```