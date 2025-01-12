![containerd banner light mode](./docs//logo/convers1.png#gh-light-mode-only)
![containerd banner dark mode](./docs//logo/convers2.png##gh-dark-mode-only)

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0)

# Convērs
Start your conversation threads securely across distributed systems.

## Getting Started

### Requirement
- Docker engine
- Go 1.23.3
- Protocol Buffer

### 🛠️ Local development
Run command bellow to clone all project's source code

```
git clone --recurse-submodules -j8 https://github.com/convrz/convers.git $GOPATH/src/github.com/convrz/convers
```


##### With SSH Key
```
git clone --recurse-submodules -j8 git@github.com:convrz/convers.git $GOPATH/src/github.com/convrz/convers
```

Your project is located in
```
cd $GOPATH/src/github.com/convrz/convers
```

## Submodules

### 📂 Protobuf Third Party
- ```googleapis``` at https://github.com/googleapis/googleapis.git standard rule for REST API
- ```grpc-gateway``` at https://github.com/grpc-ecosystem/grpc-gateway.git plugin protocol buffer to generate reverse proxy
- ```protovalidate``` at https://github.com/bufbuild/protovalidate.git library to validate structure.

#
Made in 🇻🇳 🚀