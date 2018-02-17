#!/bin/bash

kubectl  --username=admin --password=EJHmxMkMxe0CSdeC   delete clusterrolebinding cassandra-operator
kubectl --username=admin --password=EJHmxMkMxe0CSdeC   delete clusterrole cassandra-operator
