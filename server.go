package main

import (
	"fmt";
	"net";
	"io";
	"os";
	"strings"
)

func main() {
	var (
		host = "127.0.0.1";
		port = "9998";
		remote = host + ":" + port;
		data = make([]byte, 1024);
	)
	DB := make(map[string] string);
	fmt.Println("Initiating server... (Ctrl-C to stop)");

	lis, error := net.Listen("tcp", remote);
	defer lis.Close();
	if error != nil { 
		fmt.Printf("Error creating listener: %s\n", error ); 
		os.Exit(1); 
	}
	for {
		var query string;
		var read = true;
		con, error := lis.Accept();
		if error != nil { fmt.Printf("Error: Accepting data: %s\n", error); os.Exit(2); }
		fmt.Printf("=== New Connection received from: %s \n", con.RemoteAddr()); 
		for read {
			n, error := con.Read(data);
			switch error { 
			case io.EOF:
				fmt.Printf("Warning: End of data reached: %s \n", error); 
				read = false;
			case nil:
				//fmt.Println(string(data[0:n])); // Debug
				//response = response + string(data[0:n]);
				query = string(data[0:n]);
				fmt.Println("Query : ",query); 
				var temp []string=strings.Split(query,"|");
				if temp[0]=="insert" || temp[0]=="update"{
					DB[temp[1]]=temp[2];
				}else if temp[0]=="delete"{
					delete(DB,temp[1]);
				}else if temp[0]=="view"{
					if DB[temp[1]]!=""{
						in, error := con.Write([]byte(DB[temp[1]]));
						if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
					}else{
						in, error := con.Write([]byte("null"));
						if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
					}
				}
			default:
				fmt.Printf("Error: Reading data : %s \n", error); 
				read = false;
			}
		}
		con.Close();
	}
}
