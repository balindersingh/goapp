## A golang app using Docker
```
docker build . -t goapp-image
```
Let's create the container from the image we created and expose it to local machine on 3030 port
```
docker run -p 3030:3001 --rm --name goapp-container goapp-image
```
## References:
* https://medium.com/@rrgarciach/bootstrapping-a-go-application-with-docker-47f1d9071a2a
* https://www.cloudreach.com/blog/containerize-this-golang-dockerfiles/