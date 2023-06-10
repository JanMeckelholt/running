import http.server, ssl, socketserver
import os
import sys
buffer = 1
PORT = 443
CERTFILE = '/volumes-data/certs/running_app-cert.pem'
KEYFILE = '/volumes-data/certs/running_app-key.pem'


# context =  ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
# context.load_cert_chain(CERTFILE, KEYFILE)
# #context.load_verify_locations('/running_app/certs/ca-cert.pem')
# context.verify_mode = ssl.CERT_REQUIRED
# handler = http.server.SimpleHTTPRequestHandler
# server_address = ('0.0.0.0', PORT)
# with socketserver.TCPServer(server_address, handler) as httpd:
#     httpd.socket = context.wrap_socket(httpd.socket, server_side=True)
#     sys.stderr = open('/log/server_logfile.txt', 'w', buffer)
#     httpd.serve_forever()




server_address = ('0.0.0.0', PORT)
httpd = http.server.HTTPServer(server_address, http.server.SimpleHTTPRequestHandler)
httpd.socket = ssl.wrap_socket(httpd.socket,
                               server_side=True,
                               certfile=CERTFILE,
                               keyfile=KEYFILE,
                               ssl_version=ssl.PROTOCOL_TLS_SERVER)
httpd.serve_forever()