#Implementing Distributed Job Scheduler  using Hashicorp consul

Used Consul so as to have kind of locking mechanism through which one node among various nodes decides to run the job.
Consul CAS(check and set) api uses modify Index attribute to determine only one write succeeds on the given index, thus only one node receives the successful resposne.

The Makefile has few delete instruction which user need to be aware of before executing make commands so that it didn't conflict with any existing container names and docker network.

#Environment config

The root directory has envconfig file in which various parameters can be changed as per requirement.

#Commands:

make run := This will pull consul and run the container exposing 8500 port.This also builds and runs the application. This deletes/cleans up the resources that were earlier started by this command.  
