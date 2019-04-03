const { CatFactRequest } = require('./protos/catfact_pb.js');
const { CatFactServiceClient } = require('./protos/catfact_grpc_web_pb.js');

var client = new CatFactServiceClient('http://localhost:8000');
var factPlaceholder = document.getElementById('factPlaceholder');

window.getFact = function() {

  var request = new CatFactRequest();

  client.get(request, {}, (err, response) => {
    factPlaceholder.innerText = response.getFact();
  });
}
