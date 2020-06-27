

下载代码，编译镜像
operator-sdk build huiwq1990/mongondb-operator
需要改为自己的仓库
docker push  huiwq1990/mongondb-operator
docker push  huiwq1990/mongondb-operator

删除资源
kubectl delete -f deploy/cr.yaml
kubectl delete -f deploy/bundle.yaml

kubectl delete pvc mongod-data-my-cluster-name-rs0-0
kubectl delete pvc mongod-data-my-cluster-name-rs0-1
kubectl delete pvc mongod-data-my-cluster-name-rs0-2

新建资源

kubectl apply -f deploy/bundle.yaml
kubectl apply -f deploy/secrets.yaml
kubectl apply -f deploy/cr.yaml

