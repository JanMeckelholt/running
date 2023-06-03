import http.server, ssl, socketserver
import sys
buffer = 1
PORT = 5000
# DIRECTORY ="/running_app/"

# class Handler(http.server.SimpleHTTPRequestHandler):
#     def __init__(self, *args, **kwargs):
#         super().__init__(*args, directory=DIRECTORY, **kwargs)

#context =  ssl.SSLContext(ssl.PROTOCOL_TLS_SERVER)
#context.load_cert_chain(certfile='/running_app/certs/running_app-cert.pem', keyfile='/running_app/certs/running_app-key.pem')
handler = http.server.SimpleHTTPRequestHandler
server_address = ('0.0.0.0', PORT)
with socketserver.TCPServer(server_address, handler) as httpd:
    #httpd.socket = context.wrap_socket(httpd.socket, server_side=True)
    sys.stderr = open('/log/server_logfile.txt', 'w', buffer)
    httpd.serve_forever()
