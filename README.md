I create service account and its token related to this way:

```bash
kubectl create serviceaccount my-test-sa # in default ns
```
After this, I create clusterrole and clusterrolebinding related to SA and deploy my POD with SA attach to it.

