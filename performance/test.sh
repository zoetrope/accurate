#!/bin/bash
echo "Performance Test - Controller Start/Restart Time"

rootname[1]="tplg-wide-0"
rootname[2]="tplg-full-0-0"
rootname[3]="tplg-skewer-1"

for tplg in {1..3}
do
  echo "************Loading the topology ${tplg}************"
  echo "Deleting accuate-controller-manager deployment"
  kubectl -n accurate delete deployment accurate-controller-manager
  echo "Disabling validating webhook"
  kubectl delete validatingwebhookconfigurations.admissionregistration.k8s.io/accurate-validating-webhook-configuration
  kubectl delete mutatingwebhookconfigurations.admissionregistration.k8s.io/accurate-mutating-webhook-configuration
  case $tplg in
  1)
    ./performance/load-topology-wide.sh
    ;;
  2)
    ./performance/load-topology-full.sh
    ;;
  3)
    ./performance/load-topology-skewer.sh
    ;;
  esac
  root=${rootname[tplg]}

  echo "************Starting up the controllers on topology ${tplg}************"
  ./performance/start-controllers.sh

  echo "************Restarting the controllers on topology ${tplg}************" 
  echo "Deleting accurate-controller-manager deployment"
  kubectl -n accurate delete deployment accurate-controller-manager
  ./performance/start-controllers.sh
done

# Clean up topologies
./performance/clean-up-topologies.sh
