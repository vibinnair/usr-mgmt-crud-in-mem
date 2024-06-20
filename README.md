# Project Setup

**Clone the github repository**

``` git@github.com:vibinnair/usr-mgmt-crud-in-mem.git ```

**Go to the usr-mgmt-crud-in-mem folder**

``` cd usr-mgmt-crud-in-mem ```

**Now build the docker image (Please make sure docker is installed on your machine)**

```docker build -t user-management-in-mem:1.0 .```

Here we are using -t to set a tag name to the image, which is ```user-management-in-mem:1.0```. We can give any tag name to the image.

**Once the image is create, run the docker container using the generated docker image**

``` docker run -d -p 8282:8282 --name user-management-in-mem user-management-in-mem:1.0 ```

Here 
 - `-p` is used to map the external port 8282 of our localhost to docker container's port `8282`. The user-management-in-mem service will listen on `8282` inside the docker container.
 - `--name` is used to give a name to our container, in this case it is `user-management-in-mem`
 - `user-management-in-mem:1.0` is the name we had given to our image during build our docker image in previous step


**Once the container is up and running, try the following curl command**

```curl -i http://localhost:8282/ping```

You should get a 200 OK response 
```
HTTP/1.1 200 OK
Date: Thu, 20 Jun 2024 14:39:23 GMT
Content-Length: 0
```
