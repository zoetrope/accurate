#!/bin/bash
# Clean up topologies
./scripts/performance/clean-up-topologies.sh

# N is the number of children per node(namespace). Default to 20 (421 nodes).
N=10

# O is the number of objects per namespace. Default to 1.
O=1

echo "Loading Topology Full($N children/node, $O objects/node)..."

# Create root namespace
kubectl create ns tplg-full-0-0
for ((k=1;k<=O;k++))
do
  kubectl -n tplg-full-0-0 create role role$k-0-0 --verb=update --resource=deployments
  kubectl -n tplg-full-0-0 annotate role role$k-0-0 accurate.cybozu.com/propagate='update'
  kubectl -n tplg-full-0-0 create rolebinding rolebinding$k-0-0 --role role$k-0-0 --serviceaccount=tplg-full-0-0:default
  kubectl -n tplg-full-0-0 annotate rolebinding rolebinding$k-0-0 accurate.cybozu.com/propagate='update'
done

# Create all namespaces with tree depth 1
for ((i=1;i<=N;i++))
do
  echo "Creating namespace tplg-full-0-$i"
  kubectl create ns tplg-full-0-$i
  for ((k=1;k<=O;k++))
  do
    kubectl -n tplg-full-0-$i create role role$k-0-$i --verb=update --resource=deployments
    kubectl -n tplg-full-0-$i annotate role role$k-0-$i accurate.cybozu.com/propagate='update'
    kubectl -n tplg-full-0-$i create rolebinding rolebinding$k-0-$i --role role$k-0-$i --serviceaccount=tplg-full-0-$i:default
    kubectl -n tplg-full-0-$i annotate rolebinding rolebinding$k-0-$i accurate.cybozu.com/propagate='update'
  done
  # Create all namespaces with tree depth 2
  for ((j=1;j<=N;j++))
  do
    echo "Creating namespace tplg-full-$i-$j"
    kubectl create ns tplg-full-$i-$j
    for ((k=1;k<=O;k++))
    do
      kubectl -n tplg-full-$i-$j create role role$k-$i-$j --verb=update --resource=deployments
      kubectl -n tplg-full-$i-$j annotate role role$k-$i-$j accurate.cybozu.com/propagate='update'
      kubectl -n tplg-full-$i-$j create rolebinding rolebinding$k-$i-$j --role role$k-$i-$j --serviceaccount=tplg-full-$i-$j:default
      kubectl -n tplg-full-$i-$j annotate rolebinding rolebinding$k-$i-$j accurate.cybozu.com/propagate='update'
    done
  done
done

# Create full topology.
kubectl accurate ns set-type tplg-full-0-0 root
for ((i=1;i<=N;i++))
do
  kubectl accurate sub graft tplg-full-0-$i tplg-full-0-0
  for ((j=1;j<=N;j++))
  do
    kubectl accurate sub graft tplg-full-$i-$j tplg-full-0-$i
  done
done
