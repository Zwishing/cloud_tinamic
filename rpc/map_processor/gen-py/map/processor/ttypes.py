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
import base.ttypes

from thrift.transport import TTransport
all_structs = []


class VectorThumbnailRequest(object):
    """
    Attributes:
     - cloud_optimized_path
     - cloud_optimized_bucket_name
     - width
     - height

    """
    thrift_spec = None


    def __init__(self, cloud_optimized_path = None, cloud_optimized_bucket_name = None, width = None, height = None,):
        self.cloud_optimized_path = cloud_optimized_path
        self.cloud_optimized_bucket_name = cloud_optimized_bucket_name
        self.width = width
        self.height = height

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
                if ftype == TType.STRING:
                    self.cloud_optimized_path = iprot.readString().decode('utf-8', errors='replace') if sys.version_info[0] == 2 else iprot.readString()
                else:
                    iprot.skip(ftype)
            elif fid == 2:
                if ftype == TType.STRING:
                    self.cloud_optimized_bucket_name = iprot.readString().decode('utf-8', errors='replace') if sys.version_info[0] == 2 else iprot.readString()
                else:
                    iprot.skip(ftype)
            elif fid == 3:
                if ftype == TType.I32:
                    self.width = iprot.readI32()
                else:
                    iprot.skip(ftype)
            elif fid == 4:
                if ftype == TType.I32:
                    self.height = iprot.readI32()
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
        oprot.writeStructBegin('VectorThumbnailRequest')
        if self.cloud_optimized_path is not None:
            oprot.writeFieldBegin('cloud_optimized_path', TType.STRING, 1)
            oprot.writeString(self.cloud_optimized_path.encode('utf-8') if sys.version_info[0] == 2 else self.cloud_optimized_path)
            oprot.writeFieldEnd()
        if self.cloud_optimized_bucket_name is not None:
            oprot.writeFieldBegin('cloud_optimized_bucket_name', TType.STRING, 2)
            oprot.writeString(self.cloud_optimized_bucket_name.encode('utf-8') if sys.version_info[0] == 2 else self.cloud_optimized_bucket_name)
            oprot.writeFieldEnd()
        if self.width is not None:
            oprot.writeFieldBegin('width', TType.I32, 3)
            oprot.writeI32(self.width)
            oprot.writeFieldEnd()
        if self.height is not None:
            oprot.writeFieldBegin('height', TType.I32, 4)
            oprot.writeI32(self.height)
            oprot.writeFieldEnd()
        oprot.writeFieldStop()
        oprot.writeStructEnd()

    def validate(self):
        return

    def __repr__(self):
        L = ['%s=%r' % (key, value)
             for key, value in self.__dict__.items()]
        return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

    def __eq__(self, other):
        return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not (self == other)


class VectorThumbnailRespose(object):
    """
    Attributes:
     - base
     - thumbnail

    """
    thrift_spec = None


    def __init__(self, base = None, thumbnail = None,):
        self.base = base
        self.thumbnail = thumbnail

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
                if ftype == TType.STRUCT:
                    self.base = base.ttypes.BaseResp()
                    self.base.read(iprot)
                else:
                    iprot.skip(ftype)
            elif fid == 2:
                if ftype == TType.STRING:
                    self.thumbnail = iprot.readBinary()
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
        oprot.writeStructBegin('VectorThumbnailRespose')
        if self.base is not None:
            oprot.writeFieldBegin('base', TType.STRUCT, 1)
            self.base.write(oprot)
            oprot.writeFieldEnd()
        if self.thumbnail is not None:
            oprot.writeFieldBegin('thumbnail', TType.STRING, 2)
            oprot.writeBinary(self.thumbnail)
            oprot.writeFieldEnd()
        oprot.writeFieldStop()
        oprot.writeStructEnd()

    def validate(self):
        return

    def __repr__(self):
        L = ['%s=%r' % (key, value)
             for key, value in self.__dict__.items()]
        return '%s(%s)' % (self.__class__.__name__, ', '.join(L))

    def __eq__(self, other):
        return isinstance(other, self.__class__) and self.__dict__ == other.__dict__

    def __ne__(self, other):
        return not (self == other)
all_structs.append(VectorThumbnailRequest)
VectorThumbnailRequest.thrift_spec = (
    None,  # 0
    (1, TType.STRING, 'cloud_optimized_path', 'UTF8', None, ),  # 1
    (2, TType.STRING, 'cloud_optimized_bucket_name', 'UTF8', None, ),  # 2
    (3, TType.I32, 'width', None, None, ),  # 3
    (4, TType.I32, 'height', None, None, ),  # 4
)
all_structs.append(VectorThumbnailRespose)
VectorThumbnailRespose.thrift_spec = (
    None,  # 0
    (1, TType.STRUCT, 'base', [base.ttypes.BaseResp, None], None, ),  # 1
    (2, TType.STRING, 'thumbnail', 'BINARY', None, ),  # 2
)
fix_spec(all_structs)
del all_structs
