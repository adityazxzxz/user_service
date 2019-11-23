package api

import (
	"fmt"
	"log"
	"os"
)

func Serve(){
	
}

func Initialize(Dbdriver,Dbuser,Dbpassword,Dbport,Dbhost,DbName string){
	if Dbdriver == "mysql"{
		DBURL := fmt.Sprintf("%s:$s@tcp(%s:%s)/%s?charset=utf&parseTime=True&loc=Local",Dbuser,Dbpassword,Dbhost,Dbport,Dbname)
	}
}