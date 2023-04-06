# Docker 打包流程

## 本机

```powershell
docker build -t collections-backend:x.x .	# x.x 为版本号
dcoker tag collections-backend:x.x allenluuu/collections-backend:x.x
docker push allenluuu/collections-backend:x.x
docker image prune
```

## 服务器

```powershell
docker container ls
docker container stop xxx	# 删除当前运行的容器
docker container prune		# 清除不在运行的容器

docker run -dp 3001:3000 allenluuu/collections_backend:x.x

docker image ls
docker image rm xxx 		# 删除之前的镜像
docker image prune		# 删除未使用的镜像
```
