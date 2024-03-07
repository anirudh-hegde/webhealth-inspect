# 🌐️Webhealth-Inspect

## Intro
The application implemented using Golang which inspects whether the website is UP/DOWN and then opens it.
Also the application has been dockerized.

## Requirements
- Golang
- IDE (VSCode)
- Docker
  
## Build Instructions
- Clone the repository
```sh
git clone https://github.com/anirudh-hegde/webhealth-inspect.git
```

- Move to repo directory
```sh
cd webhealth-inspect
```

- Run the application
```sh
go run . -d github.com 
```

or

- Build Docker Image
```sh
docker build -t webcheck .
```

- Run the Docker Container
```sh
docker run webcheck -d github.com
```
Instead of github you can use other domain names like medium, yahoo etc.

## Conclusion
Congratulations, you have successfully run the application 🚀️.
