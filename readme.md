## A golang web app using Docker
### Pre-requisites:
* Golang installed locally
* Docker (This tutorial is created using Mac)
* Heroku account
### Let's get started
> My root folder for the project name is *goapp* so I assuming that as prefix for docker images/containers. feel free to use your own custom names.
Get/clone code from git
```
git clone https://github.com/balindersingh/<appname>
cd <appname>
```
Let's test locally if it works
```
go build -o bin/goapp
```
Now visit localhost:5000 . You should see webpage with simple hello message. 
## Let's Build docker container for your app using *alpine* image with golang (See Dockerfile in the project)
```
docker build . -t goapp-image
```
Let's create the container from the image we created and expose it to local machine on 3030 port
```
docker run -p 3030:3001 --rm --name goapp-container goapp-image
```
Let's see if it is working.
Visit localhost:3030 . You should see webpage with simple hello message. 
Some handly docker commands:
```
// list current docker images
docker images
// Remove docker image by IMAGE ID
 docker image rm e5294acd187b
// List all docker container (including inactive ones using -a)
docker container ls -a
// Remove docker container by CONTAINER ID
 docker conatiner rm 75fba963b8ca
```
## Now let's see if we can push this app to Heroku
Let's create a heroku app. (Create a heroku account if you don't have one. It has forever free version- which we will be using.)
```
heroku create
```
It will give you back your app name somthing like this : *shrouded-castle-12345* which we will be using in next steps.
Login to heroku container (assuming you are already login to heroku login, it should auto login you right away)
```
heroku container:login
```
Push your app codebase to heroku container. (you don't need to commit the code)
```
heroku container:push -a <your_heroku_app_name> web
```
Now you just need to release it to make it live.
```
heroku container:release -a <your_heroku_app_name> web
```
## References:
* https://medium.com/@rrgarciach/bootstrapping-a-go-application-with-docker-47f1d9071a2a
* https://www.cloudreach.com/blog/containerize-this-golang-dockerfiles/