# Login to dockerhub
# docker login

# Build the go App into a container
docker build -t arpanpathak/books-api:latest .

# Push the container to the cloud registry
docker push arpanpathak/books-api:latest

