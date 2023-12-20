#!/bin/bash

#kubectl exec -it dev01-robust-78bd9d64b6-qckd9 -- bash
#kubens crm
#kubectl get pod vipbr-hml-7d95f8c76f-kxzl4 -o json | jq '.metadata.ownerReferences[].kind'

#cd scripts/

#./update-hosts

count=1
echo " Entrou no arquivo "
namespace_array=$(kubectl get namespaces -o custom-columns=:.metadata.name --no-headers)
 
for namespace in $namespace_array; do

    echo $namespace
    echo $count++

    teste=$(kubectl get pods --namespace $namespace -o custom-columns=Nome:.metadata.name --no-headers | sed 's/,$//')

    #IFS=',' # Defina a variável IFS para separar por vírgulas
    array=$teste # Transforme a string em um array

    nomes=$(echo $array | sed 's/[][]//g' | tr ' ' '\n')


    for nome in $array; do

        echo "kubectl get $nome -n $namespace"
    
        kubectl exec -it $nome -n $namespace -- bash -c 'cd scripts/ && ./update-hosts && exit'

    done

done





