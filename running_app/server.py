import http.server, ssl, socketserver
import sys
buffer = 1

context =  ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
context.load_cert_chain(certfile='/running_app/certs/running_app-cert.pem', keyfile='/running_app/certs/running_app-key.pem')
handler = http.server.SimpleHTTPRequestHandler
server_address = ('0.0.0.0', 5000)
with socketserver.TCPServer(server_address, handler) as httpd:
    httpd.socket = context.wrap_socket(httpd.socket, server_side=True)
    sys.stderr = open('/log/server_logfile.txt', 'w', buffer)
    httpd.serve_forever()
