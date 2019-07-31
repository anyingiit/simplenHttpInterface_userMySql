package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"time"
)

func main() {
	WebserBase()
}
func WebserBase() {
	fmt.Println("This webserver running!")
	http.HandleFunc("/user_info",user_infoTask)
	err:=http.ListenAndServe(":8000",nil)
	if err!=nil {
		println("ListenServer Error!"+err.Error())
	}
}
func user_infoTask(w http.ResponseWriter, req *http.Request) {
	fmt.Println("user_infoTask running!")
	time.Sleep(time.Second/2)
	req.ParseForm()
	param_limit,existLimit:=req.Form["limit"]
	param_page,existPage:=req.Form["page"]
	if !(existLimit && existPage) {
		fmt.Fprintf(w,"请勿非法访问")
		return
	}
	baseJsonBean,data:=newBaseJsonBean()
	page,err:=strconv.Atoi(param_page[0])
	if err != nil {
		println("获取page失败!",err.Error())
	}
	limit,err:=strconv.Atoi(param_limit[0])
	if err != nil {
		println("获取limit失败!",err.Error())

	}
	fmt.Println("page:",page,"limit:",limit)//打印获取到的数据
	data.Data = conAndGetDatas(limit,page)
	data.Limit = limit
	data.Page = page
	baseJsonBean.Code = 0
	baseJsonBean.Message = "成功"
	baseJsonBean.Data = data

	//返回数据
	bytes,err:=json.Marshal(baseJsonBean)
	if err != nil {
		println("转换为Json格式失败!",err.Error())
	}
	fmt.Fprint(w,string(bytes))
}


func conAndGetDatas(limit, page int) [][]string {
	db :=connSql()
	trueSqlLen:=getSqlLen(db)
	//trueMaxPageL:=getTrueMaxPage(trueSqlLen,limit)
	firstElement,lastElement:=getFirstElementAndLastElement(limit,page,trueSqlLen)
	return getSqlDataForIndex(db,firstElement,lastElement)
}

func getSqlDataForIndex(db sql.DB,start,end int) ([][]string) {
	var datas [][]string = make([][]string,0)
	var sql string = "SELECT * FROM `user_info` LIMIT"+" "+strconv.Itoa(start)+","+strconv.Itoa(end-start)
	rows,err:=db.Query(sql)
	if err != nil {
		println("获取数组长度时发生错误!查询或获取长度失败.",err.Error())
	}
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址


	for rows.Next() {
		var id int
		var name string
		var telephone string
		var age int
		err = rows.Scan(&id,&name,&telephone,&age)
		var tmpArry []string
		tmpArry = append(tmpArry,strconv.Itoa(id),name,telephone,strconv.Itoa(age))
		datas = append(datas,tmpArry)
	}

	for i := 0; i < len(datas); i++ {//本地输出获取到的数据
		for j := 0; j < len(datas[i]); j++ {
			fmt.Print(datas[i][j]+" ")
		}
		fmt.Println()
	}
	return datas
}

func getFirstElementAndLastElement(limit,page,dbLenght int) (FirstElement,LastElement int){
	var trueMaxPage int = getTrueMaxPage(dbLenght,limit)
	var trueLastElement,trueFirstElement int
	if (page<=trueMaxPage&&page>0){//检查limit和page是否非法
		if (page==trueMaxPage&&dbLenght%limit!=0){//如果用户想获取最后一页数据并且数据不能被正好显示完
			trueLastElement = dbLenght;
			trueFirstElement = (page-1)*limit;
			//trueFirstElement = trueLastElement - (trueLastElement-(page-1)*limit); //最后一个元素位置-(最后一个与元素位置-((页数-1)*每页显示个数)  最后一个元素位置-(最后一个元素位置-上一页最后元素位置)
		}else {//否则正常输出
			trueLastElement = page*limit;
			trueFirstElement = trueLastElement-limit;
		}
		return trueFirstElement,trueLastElement //返回数据开始和结束角标
	}else {
		println("获取数据开始和结束角标失败!输入的limit或者page过大或不是正整数.")
	}
	return -1,-1
}
func getTrueMaxPage(dbLenght, limit int) (trueMaxPage int) { //获取表单数据真实长度
	if dbLenght%limit ==0 {
		trueMaxPage=dbLenght/ limit
	}else {
		trueMaxPage=dbLenght/limit +1
	}
	return trueMaxPage
}
func connSql() sql.DB{
	db,err:=sql.Open("mysql","root:nHHfkW1i1zfwGqAy@tcp(192.168.222.3)/simplenpage?charset=utf8")
	if err != nil {
		println("ConnSql Error!",err.Error())
	}
	return *db
}
func getSqlLen(db sql.DB) int {
	var dbLenght int
	err:=db.QueryRow("select count(1) from `user_info`").Scan(&dbLenght)
	if err != nil {
		println("获取数组长度时发生错误,查询或获取长度失败!",err.Error())
	}
	return dbLenght
}

type BaseJsonBean struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data *Data `json:"data"`
}

type Data struct {
	Limit int `json:"limit"`
	Page int `json:"page"`
	Data interface{} `json:"data"`
}




/*func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}*/

func newBaseJsonBean() (*BaseJsonBean,*Data) {
	return &BaseJsonBean{},&Data{}
}