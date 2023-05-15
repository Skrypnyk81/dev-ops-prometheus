# Installing ArgoCD in k3d cluster with GUI interface

## Short description of ArgoCD
ArgoCD is a popular open-source continuous delivery tool for Kubernetes that automates deployment of applications, monitors their state, and provides a web-based UI
and CLI for managing and deploying applications across multiple clusters and environments.

## Get started
First you must have installed k3d https://k3d.io/v5.4.9/

and kubectl cli https://k3d.io/v5.4.9/

## Create a cluster 
Create a k3d cluster with name asciiartify-demo
```bash
k3d cluster create mycluster asciiartify-demo
```
For get all info of your cluster exsecute this command
```bash 
 kubectl get all -A
 ```
 
<img width="1086" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/a0f7ea6c-b1c5-422c-84e6-b0ab4358b0b0">

## Install ArgoCD
Run the following command for install ArgoCD
```bash
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```
## Port Forwarding¶
Kubectl port-forwarding can also be used to connect to the API server without exposing the service.
```bash
kubectl port-forward svc/argocd-server -n argocd 8080:443
```
The API server can then be accessed using https://localhost:8080

## Set password for ArgoCd
For set the password for access to argocd use this command
```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
$ 4iuiyPGTqDTWCiwF
```
Use the generate password and username _admin_ to access to webGUI of ArgoCD. 
Now you are in!

<img width="1439" alt="image" src="https://github.com/Skrypnyk81/dev-ops-prometheus/assets/24361470/0c41bd09-565d-44a4-b8cf-73106dcb5343">

- Once logged in, you will be taken to the ArgoCD dashboard, which provides an overview of all the applications managed by ArgoCD.

- From the dashboard, you can navigate to the list of applications by clicking the "Applications" link in the sidebar.

- The application list will display all of the applications that ArgoCD is managing, along with their current status. To view the details of a specific application, click on its name in the list.

- Within an application, you can view the details of its current deployment, including the Git repository it is sourced from, the target environment, and any synchronization or health status checks that are in progress.

- From the application details page, you can perform a variety of actions, including deploying the application to a new environment, synchronizing its configuration with the Git repository, and promoting a specific version to production.

- To manage the configuration for an application, navigate to its "Configuration" tab. Here you can view the raw YAML configuration files for the application, as well as any overlays that are applied to customize the configuration for different environments.

- Finally, you can access a variety of settings and configuration options for ArgoCD itself by clicking the "Settings" link in the sidebar. From here, you can manage authentication settings, configure Git repositories and webhooks, and customize the behavior of ArgoCD's automated deployment and synchronization features.
