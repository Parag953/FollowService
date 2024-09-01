package handler

import (
	"github.com/graph-gophers/graphql-go"
)

// GraphiQL is an in-browser IDE for exploring GraphiQL APIs.
// This handler returns GraphiQL when requested.
//
// For more information, see https://github.com/graphql/graphiql.
type GraphiQL struct {
	Schema *graphql.Schema
}

var Page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.1.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.10.2/graphiql.js"></script>
		<style>
			#copied {
				color: blue
			}
		</style>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function successfulCopy() {
				if (document.getElementById("copied") != null) {
					return;
				}
				var top = document.getElementsByClassName("topBar")[0];
				var sp = document.createElement('sp');
			    	var t = document.createTextNode("copied!!");
				sp.appendChild(t)
				sp.setAttribute("id", "copied");
				top.appendChild(sp)
				setTimeout(function() {
					top.removeChild(sp);
				}, 3000)
			}

			function copyTextToClipboard(text) {
			  var textArea = document.createElement("textarea");
			  textArea.style.position = 'fixed';
			  textArea.style.top = 0;
			  textArea.style.left = 0;
			  textArea.style.width = '2em';
			  textArea.style.height = '2em';
			  textArea.style.padding = 0;
			  textArea.style.border = 'none';
			  textArea.style.outline = 'none';
			  textArea.style.boxShadow = 'none';
			  textArea.style.background = 'transparent';
			  textArea.value = text;
			  document.body.appendChild(textArea);
			  textArea.select();
			  try {
			    var successful = document.execCommand('copy');
			    if (successful) {
			    	successfulCopy()
			    }
			    var msg = successful ? 'successful' : 'unsuccessful';
			    console.log('Copying text command was ' + msg);
			  } catch (err) {
			    console.log('Oops, unable to copy');
			  }
			  document.body.removeChild(textArea);
			}

			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						window.query = graphQLParams.query
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}
			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql"),
				function() {

					var top = document.getElementsByClassName("topBar")[0];
					var a = document.createElement('a');
					a.className = "toolbar-button"
					var t = document.createTextNode("copy");
					a.appendChild(t)
					a.addEventListener( 'click', function(){
					    if (window.query !== undefined && window.query !== "") {
					    	var text = window.query.replace(/([a-z]+)(\n+)(\s*)/gi, "$1, ");
						text = text.replace(/({)(\n+)(\s*)/g, "$1 ")
						text = text.replace(/(})(\n+)(\s*)/g, "$1")
						text = text.replace(/\n/g, "")
					    	copyTextToClipboard(text)
					    }
					});
					top.appendChild(a)
				}
			);
		</script>
	</body>
</html>
`)
