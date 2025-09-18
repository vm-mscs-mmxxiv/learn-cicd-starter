# 🐳 Docker CLI Cheatsheet

## 🔹 Images

```bash
# Builds container out from Dockerfile, path to file is '.' current directory.
# -t set name of contaner part after : set tag of  container
docker build -t my-app:latest .

# Search for an image on Docker Hub
docker search <name>

# Download (pull) an image
docker pull <image>[:tag]

# List local images
docker images
docker image ls

# Remove image
docker rmi <image_id>
```

## 🔹 Containers

```bash
# Run a container (interactive + remove after exit)
docker run -it --rm <image> /bin/bash

# Run in background (detached)
docker run -d --name myapp <image>

# Map ports
docker run -d -p 8080:80 nginx

# List running containers
docker ps

# List all containers (including stopped)
docker ps -a

# Stop container
docker stop <container_id>

# Start container
docker start <container_id>

# Remove container
docker rm <container_id>

# View logs
docker logs -f <container_id>

# Execute command inside container
docker exec -it <container_id> /bin/bash
```

## 🔹 Volumes & Data

```bash
# List volumes
docker volume ls

# Create a volume
docker volume create mydata

# Mount volume into container
docker run -v mydata:/data -it ubuntu
```

## 🔹 Networks

```bash
# List networks
docker network ls

# Create a network
docker network create mynet

# Connect a container to a network
docker network connect mynet <container_id>
```

## 🔹 System & Cleanup

```bash
# Show resource usage
docker stats

# Disk usage
docker system df

# Remove stopped containers, unused networks, and dangling images
docker system prune

# Remove everything unused (CAREFUL!)
docker system prune -a
```

## 🔹 Save & Load Images

```bash
# Save image to tar file
docker save -o myimage.tar <image>:tag

# Load image from tar file
docker load -i myimage.tar
```
