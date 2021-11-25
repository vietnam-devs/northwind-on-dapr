# Dockerized

1. Before you run `tye`, you need to package your code as 

  - Use Buildpacks

```
> mvn spring-boot:build-image
```

2. Manual build Dockerfile (not recommendation)

  - Create `Dockerfile`

```
FROM openjdk:17-jdk-alpine
EXPOSE 8080
ARG JAR_FILE=target/*.jar
ADD ${JAR_FILE} app.jar
ENTRYPOINT ["java","-jar","/app.jar"]
```

  - Then build a `shipping-0.0.1-SNAPSHOT.jar`

```
> ./mvnw package
> ./mvnw package && java -jar target/shipping-0.0.1-SNAPSHOT.jar # or run the package to test it
```

  - Finally, we run `tye run` on root of this project.

# Install Java and Maven

- JDK: https://developers.redhat.com/products/openjdk/download
  - java-17-openjdk-17.0.1.0.12-1.win.x86_64
- Maven
  - https://howtodoinjava.com/maven/how-to-install-maven-on-windows/