import http.server, ssl

server_address = ('', 5000)
httpd = http.server.HTTPServer(server_address, http.server.SimpleHTTPRequestHandler)
httpd.socket = ssl.wrap_socket(httpd.socket,
                               server_side=True,
                               certfile='./assets/certs/running_app-cert.pem',
                               keyfile='./assets/certs/running_app-key.pem',
                               ssl_version=ssl.PROTOCOL_TLS)
httpd.serve_forever()