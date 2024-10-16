#
# Autogenerated by Thrift Compiler (0.21.0)
#
# DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
#
#  options string: py
#

from thrift.Thrift import TType, TMessageType, TFrozenDict, TException, TApplicationException
from thrift.protocol.TProtocol import TProtocolException
from thrift.TRecursive import fix_spec
from uuid import UUID

import sys

from thrift.transport import TTransport
all_structs = []


class Code(object):
    SUCCESS = 0
    INVALID_REQUEST = 1
    NOT_FOUND = 2
    SERVER_ERROR = 3
    UNAUTHORIZED = 4
    FAIL = 5

    _VALUES_TO_NAMES = {
        0: "SUCCESS",
        1: "INVALID_REQUEST",
        2: "NOT_FOUND",
        3: "SERVER_ERROR",
        4: "UNAUTHORIZED",
        5: "FAIL",
    }

    _NAMES_TO_VALUES = {
        "SUCCESS": 0,
        "INVALID_REQUEST": 1,
        "NOT_FOUND": 2,
        "SERVER_ERROR": 3,
        "UNAUTHORIZED": 4,
        "FAIL": 5,
    }


class BaseResp(object):
    """
    Attributes:
     - code
     - msg

    """
    thrift_spec = None


    def __init__(self, code = None, msg = None,):
        self.code = code
        self.msg = msg

    def read(self, iprot):
        if iprot._fast_decode is not None and isinstance(iprot.trans, TTransport.CReadableTransport) and self.thrift_spec is not None:
            iprot._fast_decode(self, iprot, [self.__class__, self.thrift_spec])
            return
        iprot.readStructBegin()
        while True:
            (fname, ftype, fid) = iprot.readFieldBegin()
            if ftype == TType.STOP:
                break
            if fid == 1:
                if ftype == TType.I32:
                    self.code = iprot.readI32()
                else:
                    iprot.skip(ftype)
            elif fid == 2:
                if ftype == TType.STRING:
                    self.msg = iprot.readString().decode('utf-8', errors='replace') if sys.version_info[0] == 2 else iprot.readString()
                else:
                    iprot.skip(ftype)
            else:
                iprot.skip(ftype)
            iprot.readFieldEnd()
        iprot.readStructEnd()

    def write(self, oprot):
        self.validate()
        if oprot._fast_encode is not None and self.thrift_spec is not None:
            oprot.trans.write(oprot._fast_encode(self, [self.__class__, self.thrift_spec]))
            return
        oprot.writeStructBegin('BaseResp')
        if self.code is not None:
            oprot.writeFieldBegin('code', TType.I32, 1)
            oprot.writeI32(self.code)
            oprot.writeFieldEnd()
        if self.msg is not None:
            oprot.writeFieldBegin('msg', TType.STRING, 2)
            oprot.writeString(self.msg.encode('utf-8') if sys.version_info[0] == 2 else self.msg)
            oprot.writeFieldEnd()
        oprot.writeFieldStop()
        oprot.writeStructEnd()

    def validate(self):
        if self.code is None:
            raise TProtocolException(message='Required field code is unset!')
        if self.msg is None:
            raise TProtocolException(message='Required field msg is unset!')
        return

    def __repr__(self):
        L = ['%s=%r' % (key, value)
             for key, value in self.__dict__.items()]
        return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

    def __eq__(self, other):
        return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not (self == other)
all_structs.append(BaseResp)
BaseResp.thrift_spec = (
    None,  # 0
    (1, TType.I32, 'code', None, None, ),  # 1
    (2, TType.STRING, 'msg', 'UTF8', None, ),  # 2
)
fix_spec(all_structs)
del all_structs
