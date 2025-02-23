make undeploy
make manifests
make install
make docker-build
k3d image import controller:latest -c local-k8s
make deploy
# kubectl delete -f test.yaml
# sleep 10
kubectl apply -f test.yaml
sleep 10
kubectl logs -f deployment.apps/db-backup-gen-controller-manager -n db-backup-gen-system