## Connect repository to ArgoCd

In the web GUI of ArgoCD add new app 
<img width="829" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/988a38a7-c200-4167-bef9-8c620ed4fbea">

Insert Application Name and select default for Project Name
<img width="1144" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/f0fc9c56-ed4d-441f-b87d-ce099fd047bf">

In Source space insert Repository URL to track your repository, revision and path of helm chart
<img width="1145" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/1b0a0833-eeb5-4f37-b622-1ffc3aef3eb1">

In Destination choose Cluster URL by default and write name of namespace
<img width="1056" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/56a225f0-6312-423f-adbf-fb366cb3dae1">

In Sync Options check 
<img width="896" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/f1b04734-9c9c-437c-a272-2329960e14c7">

When you click on the application demo for the first time, you need to synchronize the application with the repository
<img width="1038" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/b62381fa-2a03-4caa-9f26-b837bc30ec86">

After synchronizing the application, you can see the health status of the cluster
<img width="1202" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/a70bfbca-8d5b-4bc1-a4e0-b93939d47e10">

Now, when you change the repository, the change is automatically applied to the cluster. 

## Run ambassador program in created cluster
You need run the program with command for opening port
```bash
kubectl port-forward -n demo svc/ambassador 8088:80
```
In another terminal for checking program run
```bash
$ curl localhost:8088
k8sdiy-api:599e1af%
```
Next download the image of google's logo with command
```bash
wget -O /tmp/g.png https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png 
```
```bash
--2023-05-18 21:46:16--  https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png
Risoluzione di www.google.com (www.google.com)... 142.250.180.132
Connessione a www.google.com (www.google.com)|142.250.180.132|:443... connesso.
Richiesta HTTP inviata, in attesa di risposta... 200 OK
Lunghezza: 5969 (5,8K) [image/png]
Salvataggio in: «/tmp/g.png»

/tmp/g.png                                         100%[================================================================================================================>]   5,83K  --.-KB/s    in 0,007s  

2023-05-18 21:46:16 (857 KB/s) - «/tmp/g.png» salvato [5969/5969]
```
Pass the downlod image to our service ambassador with command
```bash
curl -F 'image=@/tmp/g.png' localhost:8088/img/
```
You must see the result

<img width="425" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/6f5ac0ee-dc34-4f63-a681-638304b4c2f7">

## Conclusion
ArgoCD is a popular open-source tool for managing and deploying applications on Kubernetes. It provides a user-friendly web interface for managing the entire deployment lifecycle of applications, from defining the desired state to tracking and rolling back changes. ArgoCD offers features such as automatic synchronization with a Git repository, application diff and synchronization, application rollbacks, and the ability to manage multiple clusters from a single control plane. With its ease of use and powerful functionality, ArgoCD is a valuable tool for managing complex Kubernetes deployments and is widely used in enterprise environments.
