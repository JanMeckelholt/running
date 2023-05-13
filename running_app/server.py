import http.server, ssl, socketserver

context =  ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
context.load_cert_chain(certfile='./assets/certs/running_app-cert.pem', keyfile='./assets/certs/running_app-key.pem')
handler = http.server.BaseHTTPRequestHandler
server_address = ('', 5000)
with socketserver.TCPServer(server_address, handler) as httpd:
    httpd.socket = context.wrap_socket(httpd.socket, server_side=True)
    httpd.serve_forever()


# httpd = http.server.HTTPServer(server_address, http.server.BaseHTTPRequestHandler)
# httpd.socket = ssl.wrap_socket(httpd.socket,
#                                server_side=True,
#                                certfile='./assets/certs/running_app-cert.pem',
#                                keyfile='./assets/certs/running_app-key.pem',
#                                ssl_version=ssl.PROTOCOL_TLS)
# httpd.serve_forever()