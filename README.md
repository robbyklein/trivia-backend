### Compile protobuf for Go

```
cd protobuf && protoc --go_out=. --go_opt=paths=source_relative messages.proto && cd ../
```
