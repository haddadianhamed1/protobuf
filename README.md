Step1) Install VSCode protobuf extensions
```
vscode-proto3  & Clang-Format
```

Step2) Install clang-format
```
MacOSX: brew install clang-format 
```

step3 ) Install protoc
```
In order to perform code generation, need to install protoc
brew install protobuf
```

step4) Generate
```
python
protoc -I=proto --python_out=python proto/*.proto
```