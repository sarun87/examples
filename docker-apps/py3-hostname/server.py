import http.server
import os
import socket
import struct
import time

HOST_NAME = socket.gethostname()
PORT = os.getenv("SERVER_PORT", 80)


class MyHandler(http.server.SimpleHTTPRequestHandler):
    def do_GET(s):
        """Respond to a GET request."""
        if s.path != '/':
            s.send_response(404)
            return
        s.send_response(200)
        s.send_header("Content-type", "text/html")
        s.end_headers()
        if s.path != '/':
            return
        msg = "%s : Hello World from %s" % (time.asctime(), HOST_NAME)
        tmp = bytearray()
        tmp.extend(map(ord,msg))
        s.wfile.write(tmp)


def run(server_class=http.server.HTTPServer, handler_class=http.server.BaseHTTPRequestHandler):
    print(time.asctime(), "Server Starts - %s:%s" % ("0.0.0.0", PORT))
    server_address = ('', PORT)
    httpd = server_class(server_address, MyHandler)
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        pass
    print(time.asctime(), "Server Stops - %s:%s" % ("0.0.0.0", PORT))

if __name__ == '__main__':
    run()
