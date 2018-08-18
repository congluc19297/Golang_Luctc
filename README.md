# Golang_Luctc

## 3. Download the latest version (i.e 1.7.3)
```
$ wget -c https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz
```

## 2. Next, check the integrity of the tarball. Go to https://golang.org/dl/
```
$ shasum -a 256 go1.7.3.linux-amd64.tar.gz
```

## 3. Then extract the tar archive files into /usr/local directory using the command below.
```
$ sudo tar -C /usr/local -xvzf go1.7.3.linux-amd64.tar.gz
```

## 4. mkdir go

## 5. $PATH environment variable.

```
export GOPATH=$HOME/go
export GOROOT="/usr/local/go"
export PATH=$PATH:$GOROOT/bin
```
