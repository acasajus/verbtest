# Backend tech stack selected
******

  - [Go](https://golang.org/) has been selected as the implementation language. It's main benefits are:
    - It is a compiled language. 
      - Generates a statically linked binary so it is extremely easy to deploy in the customer computing center if needed, or in verbio's center.
      - Can take advantage of all the CPU cores in the host
      - Dinamic languages are typically very difficult to test since only code that is executed can be tested. Whereas in compiled languagues the compiler can catch a lot of errors beforehand.
      - Testing is intregrated in the toolkit
      - Indentation rules are predefined by the toolkit making understanding code from other contributors easier
    - Compiles into a static binary so it is extremely easy to deploy
  - [Consul](https://www.consul.io/) as the configuration service and service registrar for service discovery
    - Each service on startup should register in consul and deregegister on stop. This way anyone that needs to call a service can discover all instances of it and load-balance by using round-robin selection or more advanced selection methods if needed
    - All configuration is in one place
    - It is also a single statically linked binary that also makes it easy to deploy
    - Can be run on a developers laptop to create a whole instance of the full service for development/debug purposes
    - If services are deployed via kubernetes the service discovery part can be replaced by kubernetes dns
  - gRPC as the transport protocol
    - Service endpoints are defined in a separate file from the code so any change in the interface requires a chande in this file. This makes trivial seeing if there is a change in the interface and prevents errors
    - It is a binary protocol leveraging HTTP/2 muxing so although it is more difficult to debug by hand, it allows bidirectional data flow, has less overheat and is faster serializind and deserializing than JSON.

# Architecture 
******

  - The API gw is a stateless service that forwards the requests to the appropriate backend service. In order to make this stateless it shoult only connect to consul to discover where the services are and to the AuthNZ service that validates the requests come from the valid users and what their privileges are. If everything is OK then it forwards the request to the appropiate service.
  - The frontend connects with the gw via a websocket to easy the bidirectional dataflow required in a chat application. Once HTTP/2 trailers are standardised websockets can be replaced directly with gRPC connections.
  - In the mockup developed the transport between clients and the apigw can be done with websockets althouth official standards doesn't allow to send any header (like Authorization: header). So a workaround would be needed.

  - There is a helper webpage included made in VueJS. It sends the requests to the apigw and prints the results. 

# How to run
*******

 Running make docker should install dependencies, build and start docker-compose. The apigw listens in port 8000 so make sure you don't have anything running there. Once the containers have been build you can run them again just by runnind `docker-compose up`. An example request can be executed with curl by doing `curl http://localhost:8000/dialog -X POST -d '{"text": "nlu: hello"}' --header "Authorization: Bearer aaa"`. Each service publishes itself in consul, so at any time you may connect to http://localhost:8500 and see in realtime what services are connected and how many. The webpage runs in port 8080. So once everything has been built and is up-and-running load http://localhost:8080 in your browser to run the webpage.

 
