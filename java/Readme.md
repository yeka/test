# Learns Java

Using docker:

```
docker run -it --rm -v $(pwd):/workspaces -w /workspaces maven:3 bash
```

## Setting VSCode Development Environment Inside Container

First, you need "Dev Containers" entension installed.

Create file `.devcontainer/devcontainer.json` with following content:
```json
{
    "name": "Example Project",
    "image": "maven:3",
    // "dockerFile": "dockerfile" // if you want to custom your base image

    // used if you want to easily labeled your container
    "runArgs": ["--name", "example_project"],

    // used if you want to expose or sync some files on your host directory
    "mounts": [
        "source=${localEnv:HOME}/.m2/repository,target=/root/.m2/repository,type=bind,consistency=cached",
        "source=${localEnv:HOME}/.ssh,target=/root/.ssh,type=bind,readonly",
    ],

    // put your beloved extensinos & custom settings here
    "customizations": {
        "vscode": {
            "extensions": [
                "vscjava.vscode-java-pack"
            ],
            "settings": {
                "java.jdt.ls.java.home": "/opt/java/openjdk",
                "java.saveActions.organizeImports": true
            }
        }
    }
}
```

## Installing Java & Maven using Downloadable Archive

Get your maven binary archive from here `https://maven.apache.org/download.cgi`.

Extract on a folder of your choice, then add these lines on `~/.bashrc` or `~/.bashrc`
```text
export MAVEN_HOME=[your maven path]
export PATH=$PATH:$MAVEN_HOME/bin
```

Get you java builds from here `https://jdk.java.net/25/`
Extract on a folder of your choice, then add these lines on `~/.bashrc` or `~/.bashrc`
```text
export JAVA_HOME=[your java path]
export PATH=$PATH:$JAVA_HOME/bin
```

Then Ctrl+Shift+P, "Open User Setting (JSON)" and add this line:
```json
{
  "java.configuration.detectJdksAtStart": true,
  "java.jdt.ls.java.home": "<your java home>",
  "java.saveActions.organizeImports": true
}
```

### Using Podman

If you're using "podman" instead of docker, change/or add the following settings by using "Ctrl+Shift+P", "Preferences: Open User Settings (JSON)":

```json
{
    "dev.containers.dockerPath": "podman",
    "dev.containers.dockerComposePath": "podman compose"
}
```



## Hello World using Maven

Create a simple project using maven (adjust groupId & artifactId):

```bash
mvn archetype:generate -DgroupId=com.example.application -DartifactId=my-app -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
```

Add following build plugins to `pom.xml`:

```xml
  <build>
    <plugins>
      <!-- to make jar -->
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-jar-plugin</artifactId>
        <version>3.5.0</version>
        <configuration>
          <archive>
            <manifest>
              <addClasspath>true</addClasspath>
              <classpathPrefix>lib/</classpathPrefix>
              <mainClass>com.example.application.App</mainClass>
            </manifest>
          </archive>
        </configuration>
      </plugin>

      <!-- to copy dependencies -->
      <plugin>
        <groupId>com.hubspot.maven.plugins</groupId>
        <artifactId>slimfast-plugin</artifactId>
        <version>1.0.0</version>
        <executions>
          <execution>
            <goals>
              <goal>copy</goal>
            </goals>
            <phase>package</phase>
          </execution>
        </executions>
      </plugin>

      <!-- to execute jar using `mvn package exec:java -q` -->
      <plugin>
        <groupId>org.codehaus.mojo</groupId>
        <artifactId>exec-maven-plugin</artifactId>
        <version>3.6.3</version>
        <configuration>
            <mainClass>com.example.application.App</mainClass>
        </configuration>
      </plugin>

    </plugins>
  </build>
```

Build your project using `mvn package` and run your project using `java -jar target/my-reactive-app-1.0-SNAPSHOT.jar`.

To distribute the project, just copy the `target/my-reactive-app-1.0-SNAPSHOT.jar` & `target/lib` directory (if any) as its dependencies.


## SpringBoot

Start using Spring Initializr, no maven archetype.
Use Spring Boot version 3.5.10 or 4.0.1.

```bash
curl https://start.spring.io/starter.zip \
  -d type=maven-project \
  -d language=java \
  -d bootVersion=3.5.10 \
  -d javaVersion=25 \
  -d dependencies=web,jdbc,postgresql \
  -d groupId=com.example \
  -d artifactId=demo \
  -d name=demo \
  -o demo.zip

unzip demo.zip -d <target_directory>
cd <target_directory>
./mvnw spring-boot:run
```

Use `./mvnw package` to package it.

Adjust `src/main/resources/application.yaml`:

```yaml
spring:
  datasource:
    url: jdbc:postgresql://localhost:5432/mydb
    username: myuser
    password: secret
    driver-class-name: org.postgresql.Driver
```

### Manually adding dependencies

Add these in `pom.xml` under `proejct/dependencies`:
```xml
		<!-- Spring JDBC -->
		<dependency>
			<groupId>org.springframework.boot</groupId>
			<artifactId>spring-boot-starter-jdbc</artifactId>
		</dependency>

		<!-- PostgreSQL Driver -->
		<dependency>
			<groupId>org.postgresql</groupId>
			<artifactId>postgresql</artifactId>
			<scope>runtime</scope>
		</dependency>
```

And run `mvn dependency:resolve` to download the dependencies.

Java spring-boot heavily use annotation.