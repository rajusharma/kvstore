package main

import (
	"fmt";
	"net";
	"os";
	"io"
	"strings";
)

func main() {
	var (
		host = "127.0.0.1";
		port = "9998";
		remote = host + ":" + port;
		msg string;
		action string;
		send = true;
	)
	con, error := net.Dial("tcp",remote);
	if error != nil { fmt.Printf("Host not found: %s\n", error ); os.Exit(1); }
	fmt.Println("Actions : insert,update,delete,view,exit");
	
	
	for send{
		fmt.Println("Please enter the valid action :");
		fmt.Scanf("%s", &action);
		action=strings.ToLower(action);//taking input from the user for valid action
		
		switch action { 
			case "insert":
				var key string
				var val string
				fmt.Printf("Enter Key: ");
				fmt.Scanf("%s", &key);
				fmt.Printf("Enter Value: ");
				fmt.Scanf("%s", &val);
				msg=action+"|"+key+"|"+val;
				in, error := con.Write([]byte(msg));
				if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
				fmt.Println("stored");
			case "update":
				var key string
				var val string
				fmt.Printf("Enter Key: ");
				fmt.Scanf("%s", &key);
				fmt.Printf("Enter NewValue: ");
				fmt.Scanf("%s", &val);
				msg=action+"|"+key+"|"+val;
				in, error := con.Write([]byte(msg));
				if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
				fmt.Println("Updated");
			case "delete":
				var key string
				fmt.Printf("Enter Key to be deleted: ");
				fmt.Scanf("%s", &key);
				msg=action+"|"+key;
				in, error := con.Write([]byte(msg));
				if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
				fmt.Println("Deleted");
			case "view":
				var key string
				var output string
				var data = make([]byte, 1024);
				fmt.Printf("Enter the key: ");
				fmt.Scanf("%s", &key);
				msg=action+"|"+key;
				in, error := con.Write([]byte(msg));
				if error != nil { fmt.Printf("Error sending data: %s, in: %d\n", error, in ); os.Exit(2); }
				n, error := con.Read(data);
				switch error { 
					case io.EOF:
						fmt.Printf("Warning: End of data reached: %s \n", error); 
					case nil:
						output = string(data[0:n]);
					default:
						fmt.Printf("Error: Reading data : %s \n", error); 
				}
				if output=="null"{
					fmt.Println("Key Doesn't exists !!!!");
				}else{
					fmt.Println(key+" == "+output);
				}
			case "exit":
				send=false;
			default:
				fmt.Println("Invalid Action"); 
				
		}
	}
	
	con.Close();

}

