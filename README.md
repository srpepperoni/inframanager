Instrucciones para despliegue en cluster

Despliegue en el cluster

creacion de namespace mis-pruebas
kubectl create ns mis-pruebas

despliegue de las plantillas del deployment y servicio

kubectl apply -f deployment.yml -n mis-pruebas
kubectl apply -f service.yml -n mis-pruebas

portforward al pod

kubectl port-forward <nombre_del_pod> mipuerto:puertopod -n mis-pruebas

https://kubernetes.io/docs/reference/access-authn-authz/authentication/#service-account-tokens

tokenFile  = "/var/run/secrets/kubernetes.io/serviceaccount/token"
rootCAFile = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"