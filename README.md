<p style="text-align:center;" align="center">
  <img src="https://github.com/anirudh-hegde/webhealth-inspect/assets/105560839/766f096f-0f83-4c29-a4eb-dd2ab6c5898d"width="700px" height="400px">
</p>


## Intro
The application implemented using Golang which inspects whether the website is UP/DOWN and then opens it.
Also the application has been dockerized.

## Demo
[Screencast from 2024-03-08 09-33-43.webm](https://github.com/anirudh-hegde/webhealth-inspect/assets/105560839/0c91f921-3562-4a7d-991a-7ed2f46d3340)

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
