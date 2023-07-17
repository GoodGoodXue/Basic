package main

// type Page struct {
// 	CurrentPage int64   //当前页码
// 	PageSize    int64   //每页条数
// 	TotalPage   int64   //总条数
// 	Nums        []int64 //分页序数
// 	PageCount   int64   //总页数
// }

// func (_ Page) CreatePage(currentPage, pageSize, total int64) Page {
// 	if currentPage < 1 {
// 		currentPage = 1
// 	}
// 	if pageSize < 1 {
// 		pageSize = 1
// 	}
// 	this := Page{}
// 	this.CurrentPage = currentPage
// 	this.PageSize = pageSize
// 	this.TotalPage = total
// 	this.PageCount = int64(math.Ceil((float64(total)) / float64(pageSize)))
// 	this.setNum()
// 	return this
// }

// // 设置序号
// func (this *Page) setNum() {
// 	this.Nums = []int64{}
// 	if this.PageCount == 0 {
// 		return
// 	}
// 	var begin int64 = 1
// 	var end int64 = 5
// 	if this.PageCount <= 5 {
// 		end = this.PageCount
// 	} else {
// 		begin = this.CurrentPage - 2
// 		end = this.CurrentPage + 2
// 		if begin <= 0 {
// 			begin = 1
// 			end = 5
// 		}
// 		if end >= this.PageCount {
// 			end = this.PageCount
// 			begin = end - 4
// 		}
// 	}
// 	for i := begin; i <= end; i++ {
// 		this.Nums = append(this.Nums, i)
// 	}
// }

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
// func Paginator(page, prepage int, nums int64) map[string]interface{} {

// 	var firstpage int //前一页地址
// 	var lastpage int  //后一页地址
// 	//根据nums总数，和prepage每页数量 生成分页总数
// 	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
// 	if page > totalpages {
// 		page = totalpages
// 	}
// 	if page <= 0 {
// 		page = 1
// 	}
// 	var pages []int
// 	switch {
// 	case page >= totalpages-5 && totalpages > 5: //最后5页
// 		start := totalpages - 5 + 1
// 		firstpage = page - 1
// 		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
// 		pages = make([]int, 5)
// 		for i, _ := range pages {
// 			pages[i] = start + i
// 		}
// 	case page >= 3 && totalpages > 5:
// 		start := page - 3 + 1
// 		pages = make([]int, 5)
// 		firstpage = page - 3
// 		for i, _ := range pages {
// 			pages[i] = start + i
// 		}
// 		firstpage = page - 1
// 		lastpage = page + 1
// 	default:
// 		pages = make([]int, int(math.Min(5, float64(totalpages))))
// 		for i, _ := range pages {
// 			pages[i] = i + 1
// 		}
// 		firstpage = int(math.Max(float64(1), float64(page-1)))
// 		lastpage = page + 1
// 		//fmt.Println(pages)
// 	}
// 	paginatorMap := make(map[string]interface{})
// 	paginatorMap["pages"] = pages
// 	paginatorMap["totalpages"] = totalpages
// 	paginatorMap["firstpage"] = firstpage
// 	paginatorMap["lastpage"] = lastpage
// 	paginatorMap["currpage"] = page
// 	return paginatorMap
// }
