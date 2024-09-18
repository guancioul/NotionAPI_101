window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">

  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  
  let url = document.location.href.split('#')[0];
  if(url.endsWith('/index.html')){
    url = url.substring(0, url.length - 'index.html'.length) + 'docs/swagger.json'
  } else if(url.endsWith('/')){
    url = url + 'docs/swagger.json'
  } else {
    url = url + '/docs/swagger.json'
  }
  window.ui = SwaggerUIBundle({
    url: url,
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  });

  //</editor-fold>
};
