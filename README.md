## Introduction
This repository contains the source code for the project of the "WASA: Web and Software Architecture" course.
The project required writing the API documentation , implementing it in the backend using GO, and then developing the frontend using VueJS.


## Design choices
Having developed my project over the summer, I decided to create a design reminiscent of the sea (the design was open-ended, and each student could develop it as they wished).
Feel free to develop a user interface that suits your preferences


## How to build

If you're not using the WebUI, or if you don't want to embed the WebUI into the final executable, then:

```shell
go build ./cmd/webapi/
```

If you're using the WebUI and you want to embed it into the final executable:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-embed
exit
# (outside the NPM container)
go build -tags webui ./cmd/webapi/
```

## How to run (in development mode)

You can launch the backend only using:

```shell
go run ./cmd/webapi/
```

If you want to launch the WebUI, open a new tab and launch:

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run dev
```

## How to build for production / homework delivery

```shell
./open-npm.sh
# (here you're inside the NPM container)
npm run build-prod
```

For "Web and Software Architecture" students: before committing and pushing your work for grading, please read the section below named "My build works when I use `npm run dev`, however there is a Javascript crash in production/grading"

```


## License

See [LICENSE](LICENSE).
