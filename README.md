# Nasa Probe

![Probe](https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/160/htc/37/rocket_1f680.png)

### Installation

Nasa Probe requires [Golang](https://golang.org/dl/) to run.

Install the dependencies and devDependencies and start the server.

```sh
$ cd nasa-sonda
$ go run main.go
```

For production environments...

```sh
$ go build main.go
$ ./main
```

### Development

Want to contribute? Great!

Dillinger uses Gulp + Webpack for fast developing.
Make a change in your file and instantaneously see your updates!

Open your favorite Terminal and run these commands.

First Tab:
```sh
$ go run main.go
```

(optional) Second Tab:
```sh
$ go test -v ./... -cover
```

#### Building for source
For production release:
```sh
$ go build main.go
```

### Docker
Nasa Probe is very easy to install and deploy in a Docker container.

By default, the Docker will expose port 8080, so change this within the Dockerfile if necessary. 
When ready, simply use the Dockerfile to build the image.

```sh
$ cd nasa-app
$ docker build -t nasa-probe .
```

Once done, run the Docker image and map the port to whatever you wish on your host. In this example, we simply map port 8000 of the host to port 8080 of the Docker (or whatever port was exposed in the Dockerfile):

```sh
$ docker container run -p 8889:8889 sonda
```

Rodar com docker-compose

```sh
$ docker-compose up --build
```

Verify the deployment by navigating to your server address in your preferred browser.

```sh
127.0.0.1:8089
```

#### Heroku

See [Heroku](https://devcenter.heroku.com/)


```sh
$ heroku login
$ heroku create
$git push heroku main

```
