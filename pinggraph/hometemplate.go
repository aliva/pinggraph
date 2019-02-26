package pinggraph

var homeTemplate = `
<!DOCTYPE html>
<html lang="en">

<head>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/sigma.js/1.2.1/sigma.min.js"></script>
  <script
    src="https://cdnjs.cloudflare.com/ajax/libs/sigma.js/1.2.1/plugins/sigma.renderers.parallelEdges.min.js"></script>
  <script
    src="https://cdnjs.cloudflare.com/ajax/libs/sigma.js/1.2.1/plugins/sigma.renderers.edgeLabels.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/sigma.js/1.2.1/plugins/sigma.plugins.dragNodes.min.js"></script>
  <script type="text/javascript">
    var nodes = []
    var edgesData = {}
    var edges = []

    function addNode(name) {
      if (nodes.includes(name)) {
        return;
      }

      nodes.push(name);

      s.graph.addNode({
        "id": name,
        "label": name,
        "x": Math.random(),
        "y": Math.random(),
        "size": 3,
      })
      s.refresh()
    }

    function addEdge(data) {
      let id = data.HostName + "-" + data.RemoteName;
      let revId = data.RemoteName + "-" + data.HostName;
      edgesData[id] = {
        "success": data.Success,
        "counter": data.Counter,
      }
      let success = edgesData[id].success
      let counter = edgesData[id].counter
      if (revId in edgesData) {
        success += edgesData[revId].success
        counter += edgesData[revId].counter
      }
      if (edges.includes(id) == false) {
        id = revId;
      }
      if (edges.includes(id) == false) {
        edges.push(id)
      }

      if (s.graph.edges(id)) {
        s.graph.dropEdge(id)
      }
      s.graph.addEdge({
        "id": id,
        "size": 100,
        "source": data.HostName,
        "target": data.RemoteName,
        // "type": "curvedArrow",
        "label": data.Result + "ms (" + success + "/" + counter + ")",
      })
      s.refresh()
    }

    window.onload = function () {
      var conn;
      var msg = document.getElementById("msg");
      var log = document.getElementById("log");

      if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
          console.log("Connection closed");
        };
        conn.onmessage = function (evt) {
          var message = evt.data;
          message = JSON.parse(message)
          addNode(message.HostName)
          addNode(message.RemoteName)
          addEdge(message)
        };
      } else {
        console.log("Your browser does not support WebSockets.");
      }
    };

  </script>
  <style type="text/css">
    #container {
      top: 0;
      bottom: 0;
      left: 0;
      right: 0;
      position: absolute;
    }
  </style>
</head>

<body>
  <div id="container"></div>
</body>
<script>
  s = new sigma({
    container: 'container',
    renderer: {
      container: document.getElementById('container'),
      type: 'canvas'
    },
    settings: {
      defaultEdgeLabelSize: 15,
      edgeLabelSize: 'proportional'
    }
  });
  sigma.plugins.dragNodes(s, s.renderers[0]);
</script>

</html>
`
