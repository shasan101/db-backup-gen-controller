make undeploy
make manifests
make install
make docker-build
k3d image import controller:latest -c local
make deploy
kubectl delete -f test.yaml
sleep 10
kubectl apply -f test.yaml
sleep 2
kubectl logs -f jobs/backup-operator-my-db-watcher -n db-backup-gen-system