#!/bin/bash
# Clean up topologies
./scripts/performance/clean-up-topologies.sh

# N is the number of children of the root. Default to 500 (501 nodes).
N=200

# O is the number of objects per namespace. Default to 1.
O=1

echo "Loading Topology Wide($N children, $O objects/node)..."

# Create all namespaces
for ((i=0;i<=N;i++))
do
  kubectl create ns tplg-wide-$i
  for ((k=1;k<=O;k++))
  do
    kubectl -n tplg-wide-$i create role role$k-$i --verb=update --resource=deployments
    kubectl -n tplg-wide-$i annotate role role$k-$i accurate.cybozu.com/propagate='update'
    kubectl -n tplg-wide-$i create rolebinding rolebinding$k-$i --role role$k-$i --serviceaccount=tplg-wide-$i:default
    kubectl -n tplg-wide-$i annotate rolebinding rolebinding$k-$i accurate.cybozu.com/propagate='update'
  done
done

# Create wide topology
kubectl accurate ns set-type tplg-wide-0 root
for ((i=1;i<=N;i++))
do
  kubectl accurate sub graft tplg-wide-$i tplg-wide-0
done
