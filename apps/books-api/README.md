# Go REST API Kubernetes Deployment Guide

This repository contains a full, end-to-end blueprint for building, containerizing, and deploying a Go REST API with Horizontal Pod Autoscaling (HPA) and Cloud Load Balancing.

## 1. Authentication & Docker Hub Push
Because remote clusters (like Linode or Azure) cannot see the local Docker images on your laptop, you must push your built image to a public registry like Docker Hub.

**1. Log into Docker Hub:**
Make sure Docker Desktop is running, then log in:
```bash
docker login -u <your-exact-docker-id>
```

**2. Build the Docker Image:**
*Note: The namespace prefix MUST exactly match your Docker Hub ID!*
```bash
docker build -t <your-docker-id>/books-api:latest .
```

**3. Push to Docker Hub:**
```bash
docker push <your-docker-id>/books-api:latest
```

### Tagging / Renaming Images
If you built the image with the wrong name, or you want to release a new version (e.g., `v2`), you don't need to rebuild it. You can just tag it:
```bash
# docker tag <old-name> <new-name>
docker tag arpanpathak/books-api:latest arpanpathak/books-api:v2

# Push the new tag
docker push arpanpathak/books-api:v2
```

---

## 2. Kubernetes Deployment
Before applying, ensure that `image: ...` in `k8s/deployment.yaml` precisely matches the image you just pushed.

**1. Apply all manifests (Deployment, HPA, Service):**
```bash
kubectl apply -f k8s/
```

**2. Get your Public IP:**
Because `service.yaml` uses `type: LoadBalancer`, your cloud provider will physically provision a Cloud Load Balancer. 
```bash
kubectl get svc books-api -w
```
Wait for `EXTERNAL-IP` to change from `<pending>` to a real IP address. Hit `Ctrl+C` to exit watch mode.

---

## 3. Interacting with the API
Once your Public IP is active, you can hit your API from anywhere in the world.

**Get all books:**
```bash
curl -s http://<PUBLIC-IP>/books
```

**Add a new book (POST):**
```bash
curl -s -X POST http://<PUBLIC-IP>/books \
  -H "Content-Type: application/json" \
  -d '{"title": "System Design Interview", "author": "Alex Xu"}'
```

---

## 4. Autoscaling (HPA) & Load Testing
The cluster is configured with a **Horizontal Pod Autoscaler (HPA)** that monitors CPU and Memory. If CPU crosses 70%, it will automatically spawn new replicas (up to 10 maximum).

To see this in action, you need 3 terminal windows.

**Terminal 1: Watch the HPA**
```bash
kubectl get hpa books-api-hpa -w
```

**Terminal 2: Watch the Pods**
```bash
kubectl get pods -w
```

**Terminal 3: Drop the Nuke (Load Generator)**
Go APIs are notoriously efficient. A standard `wget` loop won't even wake it up. To actually hurt the API and force an autoscale event, we use a modern load-testing tool called `hey`. 

Run this command to spin up a temporary pod *inside* your cluster that blasts your API with **200 concurrent worker threads** for 60 seconds straight:
```bash
kubectl run hey-load-generator --rm -i --tty --image=williamyeh/hey --restart=Never -- -z 60s -c 200 http://books-api.default.svc.cluster.local/books
```

**The Result:**
You will watch the CPU utilization in Terminal 1 spike past 300%. The HPA will instantly trigger a scale-up, and you will see Terminal 2 spawn 8 new pods to absorb the attack. Once the 60 seconds are over, Kubernetes will wait a 5-minute cooldown period and gracefully scale back down to 2 pods to save money.
