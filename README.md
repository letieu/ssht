# ssht
tool  for manage SSH, no more care about password, host, port, username

## INSTALL
```bash
git clone https://github.com/letieu/ssht
cp ssht $HOME/.local/bin
mkdir $HOME/.ssht
```
## USAGE
```bash
# create new target
ssht new -n myserver -u root -a 123.44.1.1 -p=2206

# go to target
ssht go myserver

# list targets
ssht list

# help
ssht -h
```
