# web

protoc --proto_path=protos --js_out=import_style=commonjs,binary:web/src/grpc/gen/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:web/src/grpc/gen/ protos/project.proto

## Project setup

```
npm install
```

### Compiles and hot-reloads for development

```
npm run serve
```

### Compiles and minifies for production

```
npm run build
```

### Lints and fixes files

```
npm run lint
```

### Customize configuration

See [Configuration Reference](https://cli.vuejs.org/config/).
