import sys
import os
# 获取当前文件的父目录
parent_dir = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
# 构造目标目录路径，这里假设目标目录名为 'gen-py'
target_dir = os.path.join(parent_dir, 'gen-py')
# 将目标目录添加到 sys.path
sys.path.append(target_dir)

from thrift.transport import TSocket, TTransport
from thrift.protocol import TBinaryProtocol
from thrift.server import TServer

from handler import MapProcessorServiceHandler

from map.processor import MapProcessorService

if __name__ == '__main__':
    handler = MapProcessorServiceHandler()
    processor = MapProcessorService.Processor(handler)
    transport = TSocket.TServerSocket(host='0.0.0.0', port=9090)
    tfactory = TTransport.TBufferedTransportFactory()
    pfactory = TBinaryProtocol.TBinaryProtocolFactory()

    server = TServer.TSimpleServer(processor, transport, tfactory, pfactory)
    print("Starting MapProcessor Server...")
    server.serve()

