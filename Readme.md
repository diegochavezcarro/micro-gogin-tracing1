Similar a https://github.com/diegochavezcarro/tracing-micros/

Pero utilizando Gin Gonic. Se intento usar librerias de opentracing que no agregaron nada.(https://github.com/opentracing-contrib/go-gin)

docker build -t diegochavezcarro/micro-gogin-tracing1:v3 .

docker push diegochavezcarro/micro-gogin-tracing1:v3

minikube start --memory 6144 -p istio4

minikube tunnel -p istio2

minikube dashboard -p istio2

export PATH="$PATH:/Users/diegochavez/tools/istio-1.9.5/bin"

istioctl install --set profile=demo -y  

kubectl label namespace default istio-injection=enabled

kubectl apply -f samples/addons

kubectl rollout status deployment/kiali -n istio-system

Primero construir los otros dos micros:

https://github.com/diegochavezcarro/micro-gogin-tracing2

https://github.com/diegochavezcarro/micro-gogin-tracing3


kubectl apply -f tracing-micros.yaml

Si se mofifican las versiones de las imagenes, con v1 no se muestra la trazabilidad total, con v2 sí, con la libreria de opentracing sin cumplir funcion, v3 anda y sin utilizar libreria opentracing

kubectl apply -f tracing-gateway.yaml

minikube -p istio4 service istio-ingressgateway -n istio-system

Usar el segundo browser que se levanto, sumando "/call" al path

Se podra ver la trazabilidad total en:

istioctl dashboard jaeger
