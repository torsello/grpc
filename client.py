import os
import pingpong_pb2
import pingpong_pb2_grpc
import grpc

def close(channel):
    channel.close()

def run():
    with grpc.insecure_channel("localhost:9999") as channel:
        stub=pingpong_pb2_grpc.RequestReplyServiceStub(channel)
        while True:
            print("Menu: \n 1) Traer aeropuerto (Con exepcion)\n 2) Traer usuario (Con exepcion)\n 3) Traer todos los vuelos\n 4) Traer todos los tickets\n 5) Traer cantidad de tickets\n 6) Traer usuario con mas compras\n 7) Salir")
            opt = input()
            if opt == 1:
                print("Ingrese codigo aeropuerto:")
                aeropuerto = raw_input()
                response = stub.traerAeropuerto(pingpong_pb2.Request(cadena=aeropuerto))
                print "\n"
                print response.cadena
                print "\n"
            if opt == 2:
                print("Ingrese dni o pasaporte del usuario:")
                dni = raw_input()
                response = stub.traerUsuario(pingpong_pb2.Request(cadena=dni))
                print "\n"
                print response.cadena
                print "\n"
            if opt == 3:
                print("Vuelos:")
                response = stub.traerVuelos(pingpong_pb2.Request(cadena=''))
                print "\n"
                print response.cadena
                print "\n"
            if opt == 4:
                print("Tickets:")
                response = stub.traerTickets(pingpong_pb2.Request(cadena=''))
                print "\n"
                print response.cadena
                print "\n"
            if opt == 5:
                print("Cantidad Tickets totales:")
                response = stub.cantidadTickets(pingpong_pb2.Request(cadena=''))
                print "\n"
                print response.cadena
                print "\n"
            if opt == 6:
                print("Usuario con mayor compras:")
                response = stub.usuarioMasCompras(pingpong_pb2.Request(cadena=''))
                print "\n"
                print response.cadena
                print "\n"
            if opt == 7:
                channel.unsubscribe(close)
                exit()
    


if __name__ == "__main__":
    run()