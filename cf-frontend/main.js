/* gRPC leverages client definitions in order to interact with gRPC services.
*  This makes it really easy to call, and handle responses from gRPC services.
* Since the service is encapsulated into a pre-generated client, you can simply
* use NEW to create a new client, and then call the supported rpc calls. */

// These files were generated by using proto-gen-grpc-web.exe using the .proto files
// defined in the proto/ folder.
// proto-gen-grpc-web can be found here https://github.com/grpc/grpc-web/releases
const { CatFactRequest } = require('./proto/catfact_pb.js');
const { CatFactServiceClient } = require('./proto/catfact_grpc_web_pb.js');

var client = new CatFactServiceClient('http://localhost:31380');
var factPlaceholder = document.getElementById('factPlaceholder');

window.getFact = function() {

  var request = new CatFactRequest();

  // Using the pre-generated client object provided by proto-gen-grpc-web, call the .get() method
  // which ultimately calls the Get() method on our gRPC service.
  // this will return a CatFactResponse object.
  client.get(request, {}, (err, response) => {
    factPlaceholder.innerText = response.getFact();
  });
}
