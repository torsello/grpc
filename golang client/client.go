/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"fmt"
	"os"
	"context"
	"log"
	"google.golang.org/grpc"
	pb "pingpong"
)

const (
	address     = "localhost:9999"
	defaultName = "world"
)

func main() {
	var opt int
	var cadena string
	
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRequestReplyServiceClient(conn)
	for{
		fmt.Println("Menu: \n 1) Traer aeropuerto (Con exepcion)\n 2) Traer usuario (Con exepcion)\n 3) Traer todos los vuelos\n 4) Traer todos los tickets\n 5) Traer cantidad de tickets\n 6) Traer usuario con mas compras\n 7) Salir")
		fmt.Scan(&opt)
		fmt.Println("\n")
		if opt==1{
			fmt.Println("Ingrese codigo de aeropuerto:")
			fmt.Scan(&cadena)
			response, err := c.TraerAeropuerto(context.Background(), &pb.Request{Cadena: cadena})
			fmt.Println("\n")
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Printf("Aeropuerto: %s", response.Cadena)
			fmt.Println("\n")
		}
		if opt==2{
			fmt.Println("Ingrese dni/pasaporte del usuario:")
			fmt.Scan(&cadena)
			response, err := c.TraerUsuario(context.Background(), &pb.Request{Cadena: cadena})
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Printf("Usuario: %s", response.Cadena)
			fmt.Println("\n")
		}
		if opt==3{
			response, err := c.TraerVuelos(context.Background(), &pb.Request{Cadena: ""})
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Printf("Cantidad de vuelos totales: %s", response.Cadena)
			fmt.Println("\n")
		}
		if opt==4{
			response, err := c.TraerTickets(context.Background(), &pb.Request{Cadena: ""})
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Printf("Tickets: %s", response.Cadena)
			fmt.Println("\n")
		}
		if opt==5{
			response, err := c.CantidadTickets(context.Background(), &pb.Request{Cadena: ""})
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Printf("Cantidad de tickets emitidos: %s", response.Cadena)
			fmt.Println("\n")
		}
		if opt==6{
			response, err := c.UsuarioMasCompras(context.Background(), &pb.Request{Cadena: ""})
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Printf("Usuario con mas compras: %s", response.Cadena)
			fmt.Println("\n")
		}
		if opt==7{
			conn.Close()
			os.Exit(0)
		}
	
	}
}
