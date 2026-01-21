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
                "java.jdt.ls.java.home": "/opt/java/openjdk"
            }
        }
    }
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
    </plugins>
  </build>
```

Build your project using `mvn package` and run your project using `java -jar target/my-reactive-app-1.0-SNAPSHOT.jar`.

To distribute the project, just copy the `target/my-reactive-app-1.0-SNAPSHOT.jar` & `target/lib` directory (if any) as its dependencies.
