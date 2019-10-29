# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import pingpong_pb2 as pingpong__pb2


class RequestReplyServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.traerAeropuerto = channel.unary_unary(
        '/RequestReplyService/traerAeropuerto',
        request_serializer=pingpong__pb2.Request.SerializeToString,
        response_deserializer=pingpong__pb2.Reply.FromString,
        )
    self.traerUsuario = channel.unary_unary(
        '/RequestReplyService/traerUsuario',
        request_serializer=pingpong__pb2.Request.SerializeToString,
        response_deserializer=pingpong__pb2.Reply.FromString,
        )
    self.traerVuelos = channel.unary_unary(
        '/RequestReplyService/traerVuelos',
        request_serializer=pingpong__pb2.Request.SerializeToString,
        response_deserializer=pingpong__pb2.Reply.FromString,
        )
    self.traerTickets = channel.unary_unary(
        '/RequestReplyService/traerTickets',
        request_serializer=pingpong__pb2.Request.SerializeToString,
        response_deserializer=pingpong__pb2.Reply.FromString,
        )
    self.cantidadTickets = channel.unary_unary(
        '/RequestReplyService/cantidadTickets',
        request_serializer=pingpong__pb2.Request.SerializeToString,
        response_deserializer=pingpong__pb2.Reply.FromString,
        )
    self.usuarioMasCompras = channel.unary_unary(
        '/RequestReplyService/usuarioMasCompras',
        request_serializer=pingpong__pb2.Request.SerializeToString,
        response_deserializer=pingpong__pb2.Reply.FromString,
        )


class RequestReplyServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def traerAeropuerto(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def traerUsuario(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def traerVuelos(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def traerTickets(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def cantidadTickets(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def usuarioMasCompras(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RequestReplyServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'traerAeropuerto': grpc.unary_unary_rpc_method_handler(
          servicer.traerAeropuerto,
          request_deserializer=pingpong__pb2.Request.FromString,
          response_serializer=pingpong__pb2.Reply.SerializeToString,
      ),
      'traerUsuario': grpc.unary_unary_rpc_method_handler(
          servicer.traerUsuario,
          request_deserializer=pingpong__pb2.Request.FromString,
          response_serializer=pingpong__pb2.Reply.SerializeToString,
      ),
      'traerVuelos': grpc.unary_unary_rpc_method_handler(
          servicer.traerVuelos,
          request_deserializer=pingpong__pb2.Request.FromString,
          response_serializer=pingpong__pb2.Reply.SerializeToString,
      ),
      'traerTickets': grpc.unary_unary_rpc_method_handler(
          servicer.traerTickets,
          request_deserializer=pingpong__pb2.Request.FromString,
          response_serializer=pingpong__pb2.Reply.SerializeToString,
      ),
      'cantidadTickets': grpc.unary_unary_rpc_method_handler(
          servicer.cantidadTickets,
          request_deserializer=pingpong__pb2.Request.FromString,
          response_serializer=pingpong__pb2.Reply.SerializeToString,
      ),
      'usuarioMasCompras': grpc.unary_unary_rpc_method_handler(
          servicer.usuarioMasCompras,
          request_deserializer=pingpong__pb2.Request.FromString,
          response_serializer=pingpong__pb2.Reply.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'RequestReplyService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))