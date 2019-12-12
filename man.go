package main

//cat用法
func cat(){

}
/*
func main() {
	f,err:=os.Create("E:/mnt/a.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println(f)
	for i:=0;i<10;i++{
		n, errs := f.WriteString("hello-world'\n'")//写数据
		if errs != nil {
			fmt.Println(errs)
			fmt.Println(n)
		}
	}
	defer f.Close()

}



func main(){
	file,err:=os.OpenFile("E:/mnt/site.txt",os.O_RDONLY,0766)
	if err != nil{
		fmt.Println(err)
		return
}
	//fmt.Println(file)
    defer file.Close()
	reader:=bufio.NewReader(file)
	for {
		 line,err:=reader.ReadString('\n')

		 if err == io.EOF{
		 	fmt.Println("已结束")
		 	break
		 }


		 if err !=nil{
		 	fmt.Println(err)
		 	return
		}
		fmt.Println(line)
	}




	output,err:=ioutil.ReadAll(file)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
	context:="golange000" //写入会替代
	for i:=0;i<10;i++{
		if ioutil.WriteFile ("E:/mnt/a.txt",[]byte(context),0644) !=nil{
			fmt.Println(err)
		}else{
			fmt.Println("success")
		}
	}

}
//读文件写入文件
func main() {
	lines,err:=os.OpenFile("E:/mnt/site1.txt",os.O_RDWR|os.O_CREATE,0666)
	if err !=nil {
		fmt.Println(err)
		return
	}
    defer  lines.Close()

	line,err:=os.OpenFile("E:/mnt/site.txt",os.O_RDONLY,0666)
	if err !=nil {
		fmt.Println(err)
		return
	}
	defer line.Close()

	reader:=bufio.NewReader(line)//读文件
	writes:=bufio.NewWriter(lines)//写文件
	for {
		str,err:=reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("已结束")
			break
		}
		if err!=nil{
			fmt.Println(err)
			return
		}

		writes.WriteString(str)
	}

	writes.Flush()
}



//复制文件copy用法
func main(){
	_,err:=copyfile("E:/mnt/site4.txt","E:/mnt/site.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}

	fmt.Println("copygone")

}

func copyfile(destfile,srcfile string) (written int64,err error){
	src,err:=os.Open(srcfile)
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer src.Close()
	des,err:=os.OpenFile(destfile,os.O_CREATE|os.O_RDWR,0664)
	if err!=nil{
		fmt.Println(err)
		return
	}
    defer des.Close()
	return io.Copy(des,src)

}

 */