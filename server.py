from concurrent import futures
import grpc
import pingpong_pb2
import pingpong_pb2_grpc
import time
import threading
import mysql.connector
from mysql.connector import Error, MySQLConnection


class Listener(pingpong_pb2_grpc.RequestReplyServiceServicer):
    def traerAeropuerto(self, request, context):
        try:
            connection = mysql.connector.connect(host='db4free.net',
                                            database='rappioeste',
                                            user='rapiuser',
                                            password='rapiuserpass')
            cursor = connection.cursor()
            args=(request.cadena,)
            cursor.callproc("traerAeropuerto", args)

            for result in cursor.stored_results():
                tup=result.fetchone()
            
            str1=" \n".join(map(str,tup))
            
            return pingpong_pb2.Reply(cadena=str1)
        except mysql.connector.Error as error:
            print("Failed to execute stored procedure: {}".format(error))
        finally:
            if (connection.is_connected()):
                cursor.close()
                connection.close()
                print("MySQL connection is closed")


    def traerUsuario(self, request, context):
            try:
                connection = mysql.connector.connect(host='db4free.net',
                                                database='rappioeste',
                                                user='rapiuser',
                                                password='rapiuserpass')
                cursor = connection.cursor()
                args=(request.cadena,)
                cursor.callproc("traerUsuario", args)

                for result in cursor.stored_results():
                    tup=result.fetchone()
                
                str1=" \n".join(map(str,tup))
                
                return pingpong_pb2.Reply(cadena=str1)
            except mysql.connector.Error as error:
                print("Failed to execute stored procedure: {}".format(error))
            finally:
                if (connection.is_connected()):
                    cursor.close()
                    connection.close()
                    print("MySQL connection is closed")

            
    def traerVuelos(self, request, context):
        try:
            connection = mysql.connector.connect(host='db4free.net',
                                            database='rappioeste',
                                            user='rapiuser',
                                            password='rapiuserpass')
            cursor = connection.cursor()
            cursor.callproc("traerVuelos")

            for result in cursor.stored_results():
                tup=result.fetchall()
            
            str1=", ".join(map(str,tup))
            
            return pingpong_pb2.Reply(cadena=str1)
        except mysql.connector.Error as error:
            print("Failed to execute stored procedure: {}".format(error))
        finally:
            if (connection.is_connected()):
                cursor.close()
                connection.close()
                print("MySQL connection is closed")


    def traerTickets(self, request, context):
        try:
            connection = mysql.connector.connect(host='db4free.net',
                                            database='rappioeste',
                                            user='rapiuser',
                                            password='rapiuserpass')
            cursor = connection.cursor()
            cursor.callproc("traerTickets")

            for result in cursor.stored_results():
                tup=result.fetchall()
            
            str1=", ".join(map(str,tup))
            
            return pingpong_pb2.Reply(cadena=str1)
        except mysql.connector.Error as error:
            print("Failed to execute stored procedure: {}".format(error))
        finally:
            if (connection.is_connected()):
                cursor.close()
                connection.close()
                print("MySQL connection is closed")


    def cantidadTickets(self, request, context):
        try:
            connection = mysql.connector.connect(host='db4free.net',
                                            database='rappioeste',
                                            user='rapiuser',
                                            password='rapiuserpass')
            cursor = connection.cursor()
            cursor.callproc("cantidadTickets")

            for result in cursor.stored_results():
                tup=result.fetchone()
            
            str1=" ".join(map(str,tup))
            
            return pingpong_pb2.Reply(cadena=str1)
        except mysql.connector.Error as error:
            print("Failed to execute stored procedure: {}".format(error))
        finally:
            if (connection.is_connected()):
                cursor.close()
                connection.close()
                print("MySQL connection is closed")


    def usuarioMasCompras(self, request, context):
        try:
            connection = mysql.connector.connect(host='db4free.net',
                                            database='rappioeste',
                                            user='rapiuser',
                                            password='rapiuserpass')
            cursor = connection.cursor()
            cursor.callproc("usuarioMasCompras")

            for result in cursor.stored_results():
                tup=result.fetchone()
            
            str1=' '.join(map(str,tup))
            return pingpong_pb2.Reply(cadena=str1)
        except mysql.connector.Error as error:
            print("Failed to execute stored procedure: {}".format(error))
        finally:
            if (connection.is_connected()):
                cursor.close()
                connection.close()
                print("MySQL connection is closed")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pingpong_pb2_grpc.add_RequestReplyServiceServicer_to_server(Listener(), server)
    server.add_insecure_port("[::]:9999")
    server.start()
    try:
        while True:
            print("Server on: threads %i" % (threading.active_count()))
            time.sleep(10)
    except KeyboardInterrupt:
        print("Interrumpido")
        server.stop(0)


if __name__ == "__main__":
    serve()
